package main

import (
	"gochat/env"
	"gochat/pkg/setting"
	"gochat/routers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	env.LoadMainDotEnv()
	env.CheckAllVars()

	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	logrus.Info("starting server in ", env.GetStr(env.ENV), " mode...")

	router := routers.InitRouter()
	endPoint := setting.ServerSetting.EndPoint

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}
	logrus.Info("start http server listening ", endPoint)
	server.ListenAndServe()
}
