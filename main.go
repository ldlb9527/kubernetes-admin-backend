package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/klog"
	"kubernetes-admin-backend/config"
	"kubernetes-admin-backend/middleware"
	"kubernetes-admin-backend/router"
	"kubernetes-admin-backend/service"
	"net/http"
)

func main() {
	go service.Informers()  //todo informers监控
	engine := gin.Default() //返回默认的路由引擎
	gin.SetMode(gin.DebugMode)
	engine.Use(middleware.Cors()) //解决跨域问题
	engine.GET("/metrics", func(handler http.Handler) gin.HandlerFunc {
		return func(c *gin.Context) {
			handler.ServeHTTP(c.Writer, c.Request)
		}
	}(promhttp.Handler()))
	router.CollectRoute(engine)
	err := engine.Run(fmt.Sprintf("%s:%d", config.GetString(config.ServerHost), config.GetInt(config.ServerPort)))
	if err != nil {
		klog.Fatal(err)
		return
	}
}

//engine.GET("/metrics", promHandler(promhttp.Handler()))
func promHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
