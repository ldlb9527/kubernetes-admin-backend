package router

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/apis"
)

func CollectRoute(engine *gin.Engine) {

	engine.GET("/terminal", apis.VisitorWebsocketServer)

	commonGroup := engine.Group("/common")
	commonGroup.GET("/:group/:version/:resource/:namespace/:name", apis.GetYaml)
	commonGroup.POST("/apply", apis.ApplyYaml)
	commonGroup.GET("/find/gvr", apis.FindGVRs)

	clusterGroup := engine.Group("/cluster")
	clusterGroup.GET("/version", apis.Version)
	clusterGroup.GET("/nodes", apis.Nodes)
	clusterGroup.GET("/extra/info", apis.ExtraClusterInfo)
	clusterGroup.GET("/groups", apis.ApiGroups)

	namespaceGroup := engine.Group("/namespace")
	namespaceGroup.GET("/get", apis.GetNamespaces)
	namespaceGroup.POST("/create", apis.CreateNamespace)
	namespaceGroup.POST("/delete/:name", apis.DeleteNamespace)
	namespaceGroup.POST("/update", apis.UpdateNamespace)

	podGroup := engine.Group("/pod")
	podGroup.GET("/list", apis.GetAllPods)

	svcGroup := engine.Group("/svc")
	svcGroup.GET("/list", apis.ListSvc)

	configMapsGroup := engine.Group("/configmaps")
	configMapsGroup.GET("/list", apis.ListConfigMap)
	configMapsGroup.GET("/get/:namespace/:name", apis.GetConfigMapByName)
}
