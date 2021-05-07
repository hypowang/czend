package main

import (
	"czend/bootstrap"
	"czend/pkg/conf"
	"czend/pkg/setting"
	"czend/pkg/util"
	"czend/web/routers"
	"flag"
	"fmt"
	"net/http"
)

var (
	confPath      string
	isShowVersion bool
)

//程序初始化，会先于main()自动执行
func init() {

	//解析并定义命令行参数
	flag.StringVar(&confPath, "c", "configs/app.ini", "config file path")
	//flag.IntVar(&portNumber, "p", 8000, "bind port number")
	flag.BoolVar(&isShowVersion, "v", false, "check version")
	flag.Parse()

	bootstrap.Init(confPath)

}

//主函数，默认入口
func main() {

	if isShowVersion {
		fmt.Println("version:1.0")
		return
	}

	setting.InitSetting(confPath)

	util.Log().Info("config: %v", util.RelativePath(confPath))

	util.Log().Info("Listen: %v", conf.SystemConfig.Listen)
	util.Log().Info("Listen: %v", setting.HTTPPort)

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
