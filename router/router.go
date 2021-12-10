package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r1 := r.Group("/test")
	r1.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	return r
}
