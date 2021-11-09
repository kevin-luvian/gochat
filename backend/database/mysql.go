package database

import (
	"gochat/lib/database/sql"
	"os"
	"sync"
)

var mysqldbonce sync.Once

var mysqldbInstance *sql.SQLDB

func GetMYSQLDB() *sql.SQLDB {
	if mysqldbInstance == nil {
		mysqldbonce.Do(makeMYSQLDB)
	}
	return mysqldbInstance
}

func makeMYSQLDB() {
	sqldb := sql.MakeSQLDB(
		"mysql",
		os.Getenv("DB_MYSQL_HOST"),
		os.Getenv("DB_MYSQL_PORT"),
		os.Getenv("DB_MYSQL_USER"),
		os.Getenv("DB_MYSQL_PASSWORD"),
		os.Getenv("DB_MYSQL_DATABASE"))
	mysqldbInstance = &sqldb
	mysqldbInstance.Connect()
}
