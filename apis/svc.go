package apis

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
)

// ListSvc 根据命名空间和标签查询svc,为空字符串时该条件为所有
func ListSvc(c *gin.Context) {
	namespace := c.DefaultQuery("namespace", "")
	label := c.DefaultQuery("label", "")
	svcs := service.ListSvc(namespace, label)
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, svcs, "查询成功"))
}
