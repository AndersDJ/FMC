package middleware

import (
	"encoding/json"

	"github.com/AndersDJ/FMC/pkg/errCode"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func JsonRspOk(c *gin.Context, data interface{}, printData bool) {
	code := errCode.Success
	respons(c, code, data, printData)
}

func JsonRspErr(c *gin.Context, errCode *errCode.ErrCode, err error) {
	log.Error(err.Error())
	respons(c, errCode, err.Error(), false)
}

func respons(c *gin.Context, code *errCode.ErrCode, data interface{}, printData bool) {
	h, _ := c.Get(LogHeaderKey)
	rspData := gin.H{"result": code.Code, "msg": code.Msg, "data": data}
	s, _ := json.Marshal(rspData)
	if printData {
		log.Infof("%s JSONRsp status=200, result=%d, msg=%s, data=%s", h, code.Code, code.Msg, s)
	} else {
		log.Infof("%s JSONRsp status=200, result=%d, msg=%s, len(data)=%d", h, code.Code, code.Msg, len(s))
	}
	c.AbortWithStatusJSON(200, rspData)
}
