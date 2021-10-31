package dao

import "database/sql"

type DAO struct {
	tablename string
	db        *sql.DB
}
