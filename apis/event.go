package apis

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
)

func ListEvent(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "")
	eventList := service.ListEvent(namespace)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, eventList, "查询成功"))
}
