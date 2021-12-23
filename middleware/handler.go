package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/AndersDJ/FMC/models/fileOperations"
	"github.com/AndersDJ/FMC/pkg/errCode"
	"github.com/AndersDJ/FMC/pkg/utils"
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("test")
		JsonRspOk(c, "test", true)
	}
}

func ListDirHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		bb, err := base64.StdEncoding.DecodeString(c.Request.Header.Get("basePath"))
		if err != nil {
			JsonRspErr(c, errCode.CMDErr, err)
			return
		}
		b := string(bb)
		log.Infof("basePath: %s", b)
		// if b == "/" {
		// 	b = viper.GetString("control_path")
		// }
		d, _ := fileOperations.BP.Open(b)
		data, err := fileOperations.ListDir(d, b)
		if err != nil {
			JsonRspErr(c, errCode.AuthErr, err)
		} else {
			JsonRspOk(c, data, false)
		}
	}
}

func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer log.Flush()
		// log.Info("base path: " + c.Request.Header.Get("basePath"))
		form, err := c.MultipartForm()
		// log.Info(form)
		files := form.File["file"]
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		for _, f := range files {
			filePath := viper.GetString("control_path") + c.Request.Header.Get("basePath") + "/" + f.Filename
			b, err := utils.Exists(filePath)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			if b {
				i := 0
				for {
					i++
					fp := filePath + "(" + fmt.Sprint(i) + ")"
					c, _ := utils.Exists(fp)
					if !c {
						filePath = fp
						break
					}
				}
			}
			log.Infof(filePath)
			c.SaveUploadedFile(f, filePath)
		}
		c.JSON(http.StatusOK, gin.H{})
	}
}
