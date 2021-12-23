package router

import (
	"net/http"

	"github.com/AndersDJ/FMC/middleware"

	"github.com/gin-gonic/gin"
)

var notFound = &gin.H{"result": http.StatusNotFound, "msg": "Not Found"}

func New() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, notFound)
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, notFound)
	})

	r1 := r.Group("/test", middleware.LogHeader())
	r1.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	r1.GET("/test", middleware.ListDirHandler())

	r2 := r.Group("/service", middleware.LogHeader())
	r2.GET("/fileList", middleware.ListDirHandler())
	r2.POST("/uploadFiles", middleware.UploadHandler())

	return r
}
