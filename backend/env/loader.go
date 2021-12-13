package env

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func loadEnv(path string) {
	err := godotenv.Load(path)

	if err != nil {
		logrus.Fatal("Error loading .env file. ", err.Error())
	}
}

func LoadMainDotEnv() {
	path, _ := os.Getwd()
	loadEnv(path + "/.env")
}

func LoadDotEnvForTest() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		logrus.Fatal("No caller information")
	}

	filePath := strings.TrimRight(path.Dir(filename), "/env")

	loadEnv(filePath + "/.env")
	os.Setenv(ENV, "test")
}
