package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors 跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method != "OPTIONS" {
			c.Next()
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT,PATCH,DELETE")
			c.Header("Access-Control-Allow-Headers", "content-type,origin, authorization, accept")
			c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(200)
		}
	}
}
