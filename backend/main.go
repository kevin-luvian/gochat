package main

import (
	db "gochat/database"
	"gochat/model"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Starting Server...")

	db.MYSQL.Connect()
	db.MYSQL.CreateTables(model.User{})

	// router := router.MakeMyRouter()
	// router.Handle("/check", &check.Routes)
	// router.Handle("/sample", &sample.Routes)

	// router.Serve(8000)
}
