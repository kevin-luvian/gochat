package database

import (
	"gochat/lib/database/sql"
)

var MYSQLDB sql.SQLDB

func init() {
	MYSQLDB = sql.MakeDefaultSQLDB()
	MYSQLDB.Connect()
}
