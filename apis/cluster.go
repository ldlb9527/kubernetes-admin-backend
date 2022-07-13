package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
)

// AddCluster 导入一个集群
func AddCluster(c *gin.Context) {
	name := c.Param("clusterName")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, fmt.Sprintf("get form err: %s", err.Error())))
		return
	}

	if err = service.AddCluster(name, file); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, fmt.Sprintf("AddCluster err: %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nil, "导入集群成功"))
}

// GetClusters 查询所有集群
func GetClusters(c *gin.Context) {
	clusters := service.GetClusters()
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, clusters, "查询集群成功"))
}

func DeleteCluster(c *gin.Context) {
	name := c.Param("clusterName")
	if err := service.DeleteCluster(name); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, fmt.Sprintf("删除集群失败: %s", err.Error())))
		return
	}
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nil, "删除集群成功"))
}

// Version 查询集群的版本信息
func Version(c *gin.Context) {
	clusterName := c.Param("clusterName")
	version, _ := service.Version(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, version, "查询成功"))
}

// Nodes 查询节点列表详情
func Nodes(c *gin.Context) {
	clusterName := c.Param("clusterName")
	nodeList, _ := service.ListClusters(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nodeList, "查询成功"))
}

// ExtraClusterInfo 统计就绪节点、cpu使用、内存使用占比
func ExtraClusterInfo(c *gin.Context) {
	clusterName := c.Param("clusterName")
	extraClusterInfo := service.ExtraClusterInfo(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, extraClusterInfo, "查询成功"))
}

// ApiGroups 查询api-resources列表
func ApiGroups(c *gin.Context) {
	clusterName := c.Param("clusterName")
	groups := service.QueryApiGroups(clusterName)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, groups, "查询成功"))
}
