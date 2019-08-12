package main

import (
	"fmt"
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-blog/pkg/gredis"
	"gin-blog/pkg/util"
	/*
		    "fmt"
		    "log"
		    "syscall"

		    "github.com/fvbock/endless"

		    "gin-blog/routers"
			"gin-blog/pkg/setting"
	*/)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	/*
		endless.DefaultReadTimeOut = setting.ReadTimeout
		endless.DefaultWriteTimeOut = setting.WriteTimeout
		endless.DefaultMaxHeaderBytes = 1 << 20
		endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

		//返回一个初始化的 endlessServer 对象，在 BeforeBegin 时输出当前进程的 pid，调用 ListenAndServe 将实际“启动”服务
		server := endless.NewServer(endPoint, routers.InitRouter())
		server.BeforeBegin = func(add string) {
			log.Printf("Actual pid is %d", syscall.Getpid())
		}

		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}

	*/
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	routersInit.Static("/statics", "./statics")

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
