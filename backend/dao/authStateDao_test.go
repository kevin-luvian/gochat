package dao

import (
	libsql "database/sql"
	"gochat/lib/database/sql"
	"gochat/model"
	"testing"

	"github.com/sirupsen/logrus"
)

func makeTestAuthStateDB(t *testing.T) (db *libsql.DB, cleanup func() error) {
	logrus.Info("Creating MYSQLDB for testing")
	defer logrus.Info("MYSQLDB created")

	testDB := sql.MakeTestSQLDB()
	testDB.Connect()
	testDB.DropCreateTables(model.AuthState{})

	return testDB.GetDatabase(), testDB.GetDatabase().Close
}

func TestAuthStateDAO(t *testing.T) {
	db, close := makeTestAuthStateDB(t)
	defer close()

	authStateDao := MakeAuthStateDAO(db)

	ok, state := authStateDao.Create()
	if !ok {
		logrus.Fatal("Failed to create auth state")
	}

	if isExist := authStateDao.Exist(state); !isExist {
		logrus.Fatal("State doesnt exist ", state)
	}

	if isExist := authStateDao.Exist(state); isExist {
		logrus.Fatal("State still exist ", state)
	}

	if isExist := authStateDao.Exist("asdkas"); isExist {
		logrus.Fatal("non state exist")
	}
}
