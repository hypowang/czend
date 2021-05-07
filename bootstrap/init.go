package bootstrap

import (
	"czend/pkg/conf"
	"czend/pkg/setting"

	"github.com/gin-gonic/gin"
)

func Init(path string) {

	InitApp()

	//read config
	conf.Init(path)
	setting.InitSetting(path)

	//setup debug
	if !conf.SystemConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
