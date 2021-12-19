package middleware

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"

	"github.com/gin-gonic/gin"
)

func LogHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Infof("auth url= %s", c.Request.URL.Path)
		h := fmt.Sprintf("From=%s", c.Request.RemoteAddr)
		c.Set(LogHeaderKey, h)
		c.Get("Access-Control-Allow-Origin")
		log.Debug(h)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}