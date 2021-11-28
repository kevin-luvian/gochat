package env

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	ENV = "ENV"

	ACCESS_TOKEN_SECRET  = "ACCESS_TOKEN_SECRET"
	REFRESH_TOKEN_SECRET = "REFRESH_TOKEN_SECRET"

	SERVER_ENDPOINT      = "SERVER_ENDPOINT"
	SERVER_READ_TIMEOUT  = "SERVER_READ_TIMEOUT"
	SERVER_WRITE_TIMEOUT = "SERVER_WRITE_TIMEOUT"

	GOOGLE_CID          = "GOOGLE_CID"
	GOOGLE_CSECRET      = "GOOGLE_CSECRET"
	GOOGLE_REDIRECT_URL = "GOOGLE_REDIRECT_URL"

	DB_MYSQL_HOST     = "DB_MYSQL_HOST"
	DB_MYSQL_PORT     = "DB_MYSQL_PORT"
	DB_MYSQL_DATABASE = "DB_MYSQL_DATABASE"
	DB_MYSQL_USER     = "DB_MYSQL_USER"
	DB_MYSQL_PASSWORD = "DB_MYSQL_PASSWORD"

	DB_MYSQL_TEST_HOST     = "DB_MYSQL_TEST_HOST"
	DB_MYSQL_TEST_PORT     = "DB_MYSQL_TEST_PORT"
	DB_MYSQL_TEST_DATABASE = "DB_MYSQL_TEST_DATABASE"
	DB_MYSQL_TEST_USER     = "DB_MYSQL_TEST_USER"
	DB_MYSQL_TEST_PASSWORD = "DB_MYSQL_TEST_PASSWORD"

	DB_REDIS_DEV_HOST = "DB_REDIS_DEV_HOST"
	DB_REDIS_DEV_PORT = "DB_REDIS_DEV_PORT"
	DB_REDIS_URL      = "DB_REDIS_URL"
)

func IsDevEnv() bool {
	return GetStr(ENV) == "development"
}

func GetStr(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	logrus.Panic("Cant read env ", key)
	return ""
}

func GetInt(key string) int {
	if val, ok := os.LookupEnv(key); ok {
		i, _ := strconv.ParseInt(val, 10, 0)
		return int(i)
	}
	logrus.Panic("Cant read env ", key)
	return 0
}

func CheckAllVars() {
	allVars := []string{
		ENV,
		ACCESS_TOKEN_SECRET,
		REFRESH_TOKEN_SECRET,
		SERVER_ENDPOINT,
		SERVER_READ_TIMEOUT,
		SERVER_WRITE_TIMEOUT,
		GOOGLE_CID,
		GOOGLE_CSECRET,
		GOOGLE_REDIRECT_URL,
		DB_MYSQL_HOST,
		DB_MYSQL_PORT,
		DB_MYSQL_DATABASE,
		DB_MYSQL_USER,
		DB_MYSQL_PASSWORD,
		DB_MYSQL_TEST_HOST,
		DB_MYSQL_TEST_PORT,
		DB_MYSQL_TEST_DATABASE,
		DB_MYSQL_TEST_USER,
		DB_MYSQL_TEST_PASSWORD,
		DB_REDIS_DEV_HOST,
		DB_REDIS_DEV_PORT,
		DB_REDIS_URL,
	}
	for _, v := range allVars {
		if _, ok := os.LookupEnv(v); !ok {
			logrus.Panic("ENV variable for ", v, " is not defined")
		}
	}
	logrus.Info("all ENV variables are defined")
}
