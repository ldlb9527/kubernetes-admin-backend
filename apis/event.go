package apis

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
)

func ListEvent(c *gin.Context) {
	clusterName := c.Param("clusterName")
	namespace := c.DefaultQuery("namespace", "")
	eventList := service.ListEvent(clusterName, namespace)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, eventList, "查询成功"))
}
