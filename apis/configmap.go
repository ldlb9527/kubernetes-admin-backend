package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
	"sigs.k8s.io/yaml"
)

// ListConfigMap 根据命名空间查询configmap,命名空间为空字符串时查询所有
func ListConfigMap(c *gin.Context) {
	clusterName := c.Param("clusterName")
	namespace := c.DefaultQuery("namespace", "")
	configMaps := service.ListConfigMap(clusterName, namespace)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, configMaps, "查询成功"))
}

// GetConfigMapByName 根据名称查询configmap的yaml信息
func GetConfigMapByName(c *gin.Context) {
	clusterName := c.Param("clusterName")
	namespace := c.Param("namespace")
	name := c.Param("name")
	unstructured := service.GetConfigMapByName(clusterName, namespace, name)
	//bytes, _ := yaml.Marshal(unstructObj)
	bytes, _ := yaml.Marshal(unstructured)
	fmt.Println(string(bytes))
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, bytes, "查询成功"))
}
