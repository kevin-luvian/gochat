package main

import (
	"gochat/env"
	"gochat/pkg/db"
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
	db.Setup()
}

// @title GoChat API documentation
// @description This is the core server for GoChat to manage accounts and contacts.
// @version 1.0.0
// @host localhost:8000
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
