package sql

import (
	"database/sql"
	"strings"

	"github.com/sirupsen/logrus"
)

// TODO: DuplicateERR

func FindRowsQuery(
	db *sql.DB,
	qstr string,
	qvals []interface{},
	addrs []interface{},
	callback func()) error {

	rows, err := db.Query(qstr, qvals...)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(addrs...); err != nil {
			if err == sql.ErrNoRows {
				return err
			}
			logrus.Panic("Cant query ", qstr, "\n", err)
		}
		callback()
	}

	return nil
}

func FindRowQuery(
	db *sql.DB, qstr string, qvals []interface{}, addrs []interface{}) error {
	// Query for a value based on a single row.
	if err := db.
		QueryRow(qstr, qvals...).
		Scan(addrs...); err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		logrus.Panic("Cant query ", qstr, "\n", err)
	}
	return nil
}

func InsertRowQuery(db *sql.DB, qstr string, qvals []interface{}) (int64, error) {
	qresult, err := db.Exec(qstr, qvals...)
	if err != nil {
		switch {
		case strings.Contains(strings.ToLower(err.Error()), "duplicate entry"):
			return -1, err
		default:
			logrus.Panic("UNKNOWN ERROR OCCURED", err)
		}
	}
	rId, err := qresult.LastInsertId()
	if err != nil {
		logrus.Panic("Driver doesnt support LID ", err)
	}
	return rId, nil
}

func UpdateRowQuery(db *sql.DB, qstr string, qvals []interface{}) (int64, error) {
	return execRowQuery(db, qstr, qvals)
}

func DeleteRowQuery(db *sql.DB, qstr string, qvals []interface{}) (int64, error) {
	return execRowQuery(db, qstr, qvals)
}

func execRowQuery(db *sql.DB, qstr string, qvals []interface{}) (int64, error) {
	qresult, err := db.Exec(qstr, qvals...)
	if err != nil {
		logrus.Panic("UNKNOWN ERROR OCCURED", err)
	}

	raff, err := qresult.RowsAffected()
	if err != nil {
		logrus.Panic("UNKNOWN ERROR OCCURED", err)
	}

	return raff, nil
}
