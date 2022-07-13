package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
	"sigs.k8s.io/yaml"
)

// GetYaml 通过gvr查询资源的yaml
func GetYaml(c *gin.Context) {
	unstructured := service.GetYaml(c.Param("clusterName"), c.Param("group"), c.Param("version"), c.Param("resource"),
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

	clusterName := c.Param("clusterName")
	unStructured := service.ApplyYaml(clusterName, u)
	bytes, _ := yaml.Marshal(unStructured)
	fmt.Println(string(bytes))
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, bytes, "更新yaml成功"))
}

// FindGVRs 查询gvr
func FindGVRs(c *gin.Context) {
	clusterName := c.Param("clusterName")
	group := c.Param("group")
	version := c.Param("version")
	kind := c.Param("kind")

	service.FindGVRs(clusterName, group, version, kind)
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
