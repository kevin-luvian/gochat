package database

import (
	"gochat/lib/database/sql"
	"os"
	"sync"
)

var mysqldbtestonce sync.Once

var mysqldbtestInstance *sql.SQLDB

func GetMYSQLDBTest() *sql.SQLDB {
	if mysqldbtestInstance == nil {
		mysqldbtestonce.Do(makeMYSQLDBTest)
	}
	return mysqldbtestInstance
}

func makeMYSQLDBTest() {
	sqldb := sql.MakeSQLDB(
		"mysql",
		os.Getenv("DB_MYSQL_TEST_HOST"),
		os.Getenv("DB_MYSQL_TEST_PORT"),
		os.Getenv("DB_MYSQL_TEST_USER"),
		os.Getenv("DB_MYSQL_TEST_PASSWORD"),
		os.Getenv("DB_MYSQL_TEST_DATABASE"))
	mysqldbtestInstance = &sqldb
	mysqldbtestInstance.Connect()
}
