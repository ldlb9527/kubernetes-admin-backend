package apis

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
)

// GetAllPods 查询所有pods
func GetAllPods(c *gin.Context) {
	pods := service.GetPods()
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, pods, "查询成功"))
}

// todo 查询某一命名空间下的pod
