package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/klog"
	"kubernetes-admin-backend/config"
	"kubernetes-admin-backend/middleware"
	"kubernetes-admin-backend/router"
	"kubernetes-admin-backend/service"
)

func main() {
	go service.Informers()  //todo informers监控
	engine := gin.Default() //返回默认的路由引擎
	gin.SetMode(gin.DebugMode)
	engine.Use(middleware.Cors()) //解决跨域问题
	router.CollectRoute(engine)
	err := engine.Run(fmt.Sprintf("%s:%d", config.GetString(config.ServerHost), config.GetInt(config.ServerPort)))
	if err != nil {
		klog.Fatal(err)
		return
	}
}
