package bootstrap

import (
	"czend/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Init(path string) {

	InitApp()

	//read config
	conf.Init(path)

	//setup debug
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
