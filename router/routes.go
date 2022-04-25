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
	engine.GET("/terminal/pod/:podName/:namespace/:container", apis.TerminalPod)

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
