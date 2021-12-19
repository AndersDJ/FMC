package middleware

import (
	"encoding/base64"

	"github.com/AndersDJ/FMC/models/fileOperations"
	"github.com/AndersDJ/FMC/pkg/errCode"
	log "github.com/cihub/seelog"

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
