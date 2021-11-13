package main

import (
	"gochat/database"
	"gochat/env"
	"gochat/internal/auth"
	"gochat/internal/check"
	"gochat/internal/sample"
	"gochat/model"
	"gochat/router"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	env.LoadMainDotEnv()
	logrus.Info("ENV: ", os.Getenv("ENV"))
}

func main() {
	logrus.Info("Starting Server...")

	db := database.GetMYSQLDB()
	db.CreateTables(model.User{})

	router := router.MakeMyRouter()
	router.Handle("/check", &check.Routes)
	router.Handle("/sample", &sample.Routes)
	router.Handle(auth.Endpoint, &auth.Routes)

	router.Serve(8000)
}
