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

// @title GoChat API documentation
// @description This is the core server to manage user accounts and contacts in
// GOChat application.
// @version 1.0.0
// @host localhost:8000
// @BasePath /api
// @termsOfService http://swagger.io/terms/

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	logrus.Info("starting server in ", env.GetStr(env.ENV), " mode...")

	router := routers.InitRouter()
	endPoint := setting.ServerSetting.EndPoint

	server := &http.Server{
		Addr:    endPoint,
		Handler: router,
	}

	logrus.Info("server listening on ", endPoint)
	server.ListenAndServe()
}
