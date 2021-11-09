package main

import (
	"gochat/database"
	"gochat/internal/auth"
	"gochat/internal/check"
	"gochat/internal/sample"
	"gochat/model"
	"gochat/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logrus.Info("ENV: ", os.Getenv("ENV"))
}

func main() {
	logrus.Info("Starting Server...")

	db := database.GetMYSQLDB()
	db.CreateTables(model.User{})

	router := router.MakeMyRouter()
	router.Handle("/check", &check.Routes)
	router.Handle("/sample", &sample.Routes)
	router.Handle("/auth", &auth.Routes)

	router.Serve(8000)
}
