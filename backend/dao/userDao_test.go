package dao

import (
	libsql "database/sql"
	"gochat/lib/database/sql"
	"gochat/model"
	"testing"

	"github.com/sirupsen/logrus"
)

func makeTestUserDB(t *testing.T) (db *libsql.DB, cleanup func() error) {
	logrus.Info("Creating MYSQLDB for testing")
	defer logrus.Info("MYSQLDB created")

	testDB := sql.MakeTestSQLDB()
	testDB.Connect()
	testDB.DropCreateTables(model.User{})

	return testDB.GetDatabase(), testDB.GetDatabase().Close
}

func checkUserEqual(userA model.User, userB model.User) bool {
	return userA.Name == userB.Name &&
		userA.Nickname == userB.Nickname
}

func TestFindAllUser(t *testing.T) {
	db, close := makeTestUserDB(t)
	defer close()

	userDao := MakeUserDAO(db)

	// on table creation table users should be empty
	if ok, users := userDao.FindAll(); !ok || len(users) != 0 {
		t.Fatalf(`users is not empty. ok: %t, users: %v`, ok, users)
	}
}

func TestCreateUser(t *testing.T) {
	db, close := makeTestUserDB(t)
	defer close()

	userDao := MakeUserDAO(db)

	// create new user
	user := model.User{Name: "Rick", Nickname: "Sanchez"}
	if ok, _ := userDao.Create(user); !ok {
		t.Fatalf(`can't create new user. ok: %t, user: %v`, ok, user)
	}

	// get all user, check if new user added
	ok, users := userDao.FindAll()
	if !ok || len(users) != 1 {
		t.Fatalf(`database users doesn't change. ok: %t, users: %v`, ok, users)
	}

	// check if user is the same
	if ok := checkUserEqual(user, users[0]); !ok {
		t.Fatalf(`user is not the same. user:%v users: %v`, user, users[0])
	}
}

func TestFindUserById(t *testing.T) {
	db, close := makeTestUserDB(t)
	defer close()

	userDao := MakeUserDAO(db)

	// create new user
	modelUser := model.User{Name: "Rick", Nickname: "Sanchez"}
	_, id := userDao.Create(modelUser)

	// find user
	if ok, userFound := userDao.FindById(id); !ok {
		t.Fatalf(`findById operation failed. ok: %t, userFound: %v`, ok, userFound)
	} else if !checkUserEqual(modelUser, userFound) {
		t.Fatalf(`user is not the same. user:%v users: %v`, modelUser, userFound)
	}

	// user with id 1000 shouldn't exist
	if ok, userFound := userDao.FindById(1000); ok {
		t.Fatalf(`user with non existent id found. ok:%t users: %v`, ok, userFound)
	}
}

func TestDeleteUserById(t *testing.T) {
	db, close := makeTestUserDB(t)
	defer close()

	userDao := MakeUserDAO(db)

	// create new user
	modelUser := model.User{Name: "Rick", Nickname: "Sanchez"}
	ok, id := userDao.Create(modelUser)

	// check create user and check if users list increased to 1
	if !ok {
		t.Fatalf(`can't create new user.`)
	} else if ok, users := userDao.FindAll(); !ok || len(users) != 1 {
		t.Fatalf(`database users doesn't increase. ok: %t, users: %v`, ok, users)
	}

	// delete user and check if users list decreased.
	if ok := userDao.DeleteById(id); !ok {
		t.Fatalf(`delete by id operation failed. user id: %d`, id)
	} else if ok, users := userDao.FindAll(); !ok || len(users) != 0 {
		t.Fatalf(`database users doesn't decrease. ok: %t, users: %v`, ok, users)
	}

	// delete a nonexistent user id '1000'
	if ok := userDao.DeleteById(1000); ok {
		t.Fatalf(`user with nonexistent id deleted.`)
	}
}

func TestUpdateUserById(t *testing.T) {
	db, close := makeTestUserDB(t)
	defer close()

	userDao := MakeUserDAO(db)

	// create new user. assume user create is functional and passed previous test
	modelUser := model.User{Name: "Rick", Nickname: "Sanchez"}
	_, id := userDao.Create(modelUser)

	modelUserToUpdate := model.User{Name: "Rock", Nickname: "Dwayne"}
	userDao.UpdateById(id, modelUserToUpdate)
}
