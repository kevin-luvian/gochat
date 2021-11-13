package dao

import (
	libsql "database/sql"
	"gochat/database"
	"gochat/env"
	"gochat/model"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	env.LoadDotEnvForTest()
}

func makeTestUserDB(t *testing.T) *libsql.DB {
	logrus.Info("Creating MYSQLDB for testing")

	testDB := database.GetMYSQLDBTest()

	// clean the database
	testDB.DropCreateTables(model.User{})

	return testDB.GetDatabase()
}

func checkUserEqual(userA model.User, userB model.User) bool {
	return userA.Equal(userB)
}

func TestFindAllUser(t *testing.T) {
	db := makeTestUserDB(t)

	userDao := MakeUserDAO(db)

	// on table creation table users should be empty
	if ok, users := userDao.FindAll(); !ok || len(users) != 0 {
		t.Fatalf(`users is not empty. ok: %t, users: %v`, ok, users)
	} else {
		logrus.Info(`table user is empty`)
	}
}

func TestCreateUser(t *testing.T) {
	db := makeTestUserDB(t)

	userDao := MakeUserDAO(db)

	// create new user
	user := model.User{Name: "Rick", UserID: "Sanchez"}
	if ok, _ := userDao.Create(user); !ok {
		t.Fatalf(`can't create new user. ok: %t, user: %v`, ok, user)
	} else {
		logrus.Info("new user created")
	}

	// get all user, check if new user added
	ok, users := userDao.FindAll()
	if !ok || len(users) != 1 {
		t.Fatalf(`database users doesn't change. ok: %t, users: %v`, ok, users)
	} else {
		logrus.Info("list user database increased")
	}

	// check if user is the same
	if ok := checkUserEqual(user, users[0]); !ok {
		t.Fatalf(`user is not the same. user:%v users: %v`, user, users[0])
	} else {
		logrus.Info("added user matched")
	}
}

func TestFindUserById(t *testing.T) {
	db := makeTestUserDB(t)

	userDao := MakeUserDAO(db)

	// create new user
	modelUser := model.User{Name: "Rick", UserID: "Sanchez"}
	_, id := userDao.Create(modelUser)
	logrus.Info("new user created")

	// find user
	if ok, userFound := userDao.FindById(id); !ok {
		t.Fatalf(`findById operation failed. ok: %t, userFound: %v`, ok, userFound)
	} else if !checkUserEqual(modelUser, userFound) {
		t.Fatalf(`user is not the same. user:%v users: %v`, modelUser, userFound)
	} else {
		logrus.Info("matching user found in database")
	}

	// user with id 1000 shouldn't exist
	if ok, userFound := userDao.FindById(1000); ok {
		t.Fatalf(`user with non existent id found. ok:%t users: %v`, ok, userFound)
	} else {
		logrus.Info("non existing user not found in database")
	}
}

func TestDeleteUserById(t *testing.T) {
	db := makeTestUserDB(t)

	userDao := MakeUserDAO(db)

	// create new user
	modelUser := model.User{Name: "Rick", UserID: "Sanchez"}
	ok, id := userDao.Create(modelUser)
	logrus.Info("new user created")

	// check create user and check if users list increased to 1
	if !ok {
		t.Fatalf(`can't create new user.`)
	} else if ok, users := userDao.FindAll(); !ok || len(users) != 1 {
		t.Fatalf(`database users doesn't increase. ok: %t, users: %v`, ok, users)
	} else {
		logrus.Info("list user in database increased")
	}

	// delete user and check if users list decreased.
	if ok := userDao.DeleteById(id); !ok {
		t.Fatalf(`delete by id operation failed. user id: %d`, id)
	} else {
		logrus.Info("user deleted")
	}

	if ok, users := userDao.FindAll(); !ok || len(users) != 0 {
		t.Fatalf(`database users doesn't decrease. ok: %t, users: %v`, ok, users)
	} else {
		logrus.Info("list user in database decreased")
	}

	// delete a nonexistent user id '1000'
	if ok := userDao.DeleteById(1000); ok {
		t.Fatalf(`user with nonexistent id deleted.`)
	} else {
		logrus.Info("user with non exist id not deleted")
	}
}

func TestUpdateUserById(t *testing.T) {
	db := makeTestUserDB(t)

	userDao := MakeUserDAO(db)

	// create new user. assume user create is functional and passed previous test
	modelUser := model.User{Name: "Rick", UserID: "Sanchez"}
	_, id := userDao.Create(modelUser)
	logrus.Info("new user created")

	modelUserToUpdate := model.User{Name: "Rock", UserID: "Dwayne"}
	ok, updatedUser := userDao.UpdateById(id, modelUserToUpdate)
	if !ok {
		t.Fatalf(`user update failed`)
	} else {
		logrus.Info("user updated")
	}

	if !checkUserEqual(modelUserToUpdate, updatedUser) {
		t.Fatalf(`updated user with update values doesn't match`)
	} else {
		logrus.Info("updated user with update values match")
	}

	if ok, userFound := userDao.FindById(id); !ok {
		t.Fatalf(`user not found`)
	} else if !checkUserEqual(userFound, updatedUser) {
		t.Fatalf(`user in database and updated user doesn't match`)
	} else {
		logrus.Info("updated user in database match")
	}
}
