package main

import (
	db "gochat/database"
	"gochat/model"

	"github.com/sirupsen/logrus"
)

func main() {
	// mdef.Test()

	db.MYSQL.Connect()
	// db.MYSQL.DropCreateTables([]interface{}{model.User{}})
	db.MYSQL.CreateTables(model.User{})

	logrus.Info("server started")

	// router := router.MakeMyRouter()
	// router.Handle("/check", &check.Routes)
	// router.Handle("/sample", &sample.Routes)

	// router.Serve(8000)
}
