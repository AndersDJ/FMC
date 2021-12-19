package middleware

import (
	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		JsonRspOk(c, "test", true)
	}
}
