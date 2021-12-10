package logger

import (
	"fmt"

	log "github.com/cihub/seelog"
)

func Init() {
	fmt.Println("seelog init")
	// 记录日志后需要 flush
	// defer log.Flush()

	logger, err := log.LoggerFromConfigAsFile("config/log.xml")
	// logger, err := log.LoggerFromConfigAsFile("/log.xml")
	if err != nil {
		fmt.Println("parse log.xml error")
		panic(err)
	}

	if err = log.ReplaceLogger(logger); err != nil {
		fmt.Println("log.ReplaceLogger: ", err)
		panic(err)
	}
}

//TestLog 测试日志输出
func TestLog() {
	// 从上往下，级别依次升高
	log.Trace("trace: xxx")
	log.Debug("debug: xxx")
	log.Info("info: xxx")
	log.Warn("warn: xxx")
	log.Error("error: xxx")
	log.Critical("critical: xxx")
	log.Tracef("%s%s", "aa", "bb")
}
