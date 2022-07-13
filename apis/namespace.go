package apis

import (
	"github.com/gin-gonic/gin"
	"kubernetes-admin-backend/proto"
	"kubernetes-admin-backend/service"
	"net/http"
)

// GetNamespaces 查询所有命名空间
func GetNamespaces(c *gin.Context) {
	clusterName := c.Param("clusterName")
	namespaces, err := service.GetNamespaces(clusterName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
		return
	}

	nsList := make([]proto.NameSpace, 0, len(namespaces))
	for _, namespace := range namespaces {
		var ns proto.NameSpace
		ns.Name = namespace.Name
		ns.Labels = namespace.Labels
		ns.Annotations = namespace.Annotations
		ns.CreationTimestamp = namespace.CreationTimestamp.Time
		ns.Status = string(namespace.Status.Phase)
		nsList = append(nsList, ns)
	}
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nsList, "查询成功"))
}

// CreateNamespace 创建命名空间
func CreateNamespace(c *gin.Context) {
	clusterName := c.Param("clusterName")
	var nameSpace proto.NameSpace
	if err := c.ShouldBind(&nameSpace); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
		return
	}

	namespace, err := service.CreateNamespace(clusterName, nameSpace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
		return
	}

	ns := proto.NameSpace{Name: namespace.Name, Labels: namespace.Labels, Annotations: namespace.Annotations,
		CreationTimestamp: namespace.CreationTimestamp.Time, Status: string(namespace.Status.Phase)}
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, ns, "创建成功"))
}

// DeleteNamespace 根据名称删除命名空间
func DeleteNamespace(c *gin.Context) {
	clusterName := c.Param("clusterName")
	if err := service.DeleteNamespace(clusterName, c.Param("name")); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nil, "删除成功"))
}

// UpdateNamespace 修改命名空间
func UpdateNamespace(c *gin.Context) {
	clusterName := c.Param("clusterName")
	var nameSpace proto.NameSpace
	if err := c.ShouldBind(&nameSpace); err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
		return
	}

	namespace, err := service.UpdateNamespace(clusterName, nameSpace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, (&proto.Result{}).Error(500, nil, err.Error()))
		return
	}

	nameSpace.Status = string(namespace.Status.Phase)
	nameSpace.CreationTimestamp = namespace.CreationTimestamp.Time
	c.JSON(http.StatusOK, (&proto.Result{}).Ok(200, nameSpace, "修改成功"))
}
