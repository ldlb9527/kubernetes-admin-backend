package apis

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/service"
	"net/http"
)

// Version 查询集群的版本信息
func Version(c *gin.Context) {
	version, _ := service.Version()
	c.JSON(http.StatusOK, version)
}

// Nodes 查询节点列表详情
func Nodes(c *gin.Context) {
	nodeList, _ := service.ListClusters()
	c.JSON(http.StatusOK, nodeList)
}

// ExtraClusterInfo 统计就绪节点、cpu使用、内存使用占比
func ExtraClusterInfo(c *gin.Context) {
	extraClusterInfo := service.ExtraClusterInfo()
	c.JSON(http.StatusOK, extraClusterInfo)
}

// ApiGroups 查询api-resources列表
func ApiGroups(c *gin.Context) {
	groups := service.QueryApiGroups()
	c.JSON(http.StatusOK, groups)
}
