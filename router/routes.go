package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kubernetes-admin-backend/apis"
	"net/http"
)

func CollectRoute(engine *gin.Engine) {
	// prometheus 监控
	engine.GET("/metrics", func(handler http.Handler) gin.HandlerFunc {
		return func(c *gin.Context) {
			handler.ServeHTTP(c.Writer, c.Request)
		}
	}(promhttp.Handler()))

	engine.GET("/terminal/ssh", apis.VisitorWebsocketServer)
	engine.GET("/terminal/pod/:clusterName/:namespace/:podName/:container", apis.TerminalPod)

	commonGroup := engine.Group("/common")
	commonGroup.GET("/:clusterName/:group/:version/:resource/:namespace/:name", apis.GetYaml)
	commonGroup.POST("/apply/:clusterName", apis.ApplyYaml)
	commonGroup.GET("/find/gvr/:clusterName/:group/:version/:kind", apis.FindGVRs)

	clusterGroup := engine.Group("/cluster")
	clusterGroup.GET("/version/:clusterName", apis.Version)
	clusterGroup.GET("/nodes/:clusterName", apis.Nodes)
	clusterGroup.GET("/extra/info/:clusterName", apis.ExtraClusterInfo)
	clusterGroup.GET("/groups/:clusterName", apis.ApiGroups)
	clusterGroup.POST("/add/:clusterName", apis.AddCluster)
	clusterGroup.GET("/list", apis.GetClusters)
	clusterGroup.GET("/delete/:clusterName", apis.DeleteCluster)

	namespaceGroup := engine.Group("/namespace")
	namespaceGroup.GET("/get/:clusterName", apis.GetNamespaces)
	namespaceGroup.POST("/create/:clusterName", apis.CreateNamespace)
	namespaceGroup.POST("/delete/:clusterName/:name", apis.DeleteNamespace)
	namespaceGroup.POST("/update/:clusterName", apis.UpdateNamespace)

	podGroup := engine.Group("/pod")
	podGroup.GET("/list/:clusterName", apis.GetAllPods)

	svcGroup := engine.Group("/svc")
	svcGroup.GET("/list/:clusterName", apis.ListSvc)

	eventGroup := engine.Group("/event")
	eventGroup.GET("/list/:clusterName", apis.ListEvent)

	configMapsGroup := engine.Group("/configmaps")
	configMapsGroup.GET("/list/:clusterName", apis.ListConfigMap)
	configMapsGroup.GET("/get/:clusterName/:namespace/:name", apis.GetConfigMapByName)
}
