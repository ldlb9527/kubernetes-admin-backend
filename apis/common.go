package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/restmapper"
	"kubernetes-admin-backend/client"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
	"sigs.k8s.io/yaml"
)

// GetYaml 通过gvr查询资源的yaml
func GetYaml(c *gin.Context) {
	unstructured := service.GetYaml(c.Param("group"), c.Param("version"), c.Param("resource"),
		c.Param("namespace"), c.Param("name"))
	bytes, _ := yaml.Marshal(unstructured)

	fmt.Println(string(bytes))
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, bytes, "查询yaml成功"))
}

// ApplyYaml 通过yaml更新资源 如果不存在
func ApplyYaml(c *gin.Context) {
	u, err := getYamlData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, "更新失败,"+err.Error()))
		return
	}

	unStructured := service.ApplyYaml(u)
	bytes, _ := yaml.Marshal(unStructured)
	fmt.Println(string(bytes))
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, bytes, "更新yaml成功"))
}

// FindGVRs 通过yaml更新资源 如果不存在
func FindGVRs(c *gin.Context) {
	config, _ := client.GetRestConfig()
	dc, _ := discovery.NewDiscoveryClientForConfig(config)
	mapper := restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))
	gvk := schema.GroupVersionKind{Group: c.Query("group"), Version: c.Query("version"), Kind: c.Query("kind")}

	mapping, err := mapper.RESTMappings(gvk.GroupKind(), gvk.Version)
	fmt.Println(err)
	for _, restMapping := range mapping {
		fmt.Println(restMapping.Resource)
	}
}

func getYamlData(c *gin.Context) (*unstructured.Unstructured, error) {
	var body map[string][]byte
	if err := c.ShouldBind(&body); err != nil {
		return nil, err
	}

	var u *unstructured.Unstructured
	if err := yaml.Unmarshal(body["data"], &u); err != nil {
		return nil, err
	}
	// managedFields字段不能patch 也不建议update 如果存在则删除
	delete(u.Object["metadata"].(map[string]interface{}), "managedFields")
	return u, nil
}
