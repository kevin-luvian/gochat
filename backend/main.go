package main

import (
	db "gochat/database"
	"gochat/internal/auth"
	"gochat/internal/check"
	"gochat/internal/sample"
	"gochat/model"
	"gochat/router"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting Server...")

	db.MYSQLDB.CreateTables(model.User{}, model.AuthState{})

	router := router.MakeMyRouter()
	router.Handle("/check", &check.Routes)
	router.Handle("/sample", &sample.Routes)
	router.Handle("/auth", &auth.Routes)

	router.Serve(8000)
}
