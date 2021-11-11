package dao

import (
	libsql "database/sql"
	"gochat/lib/database/sql"
	"gochat/lib/database/sql/query"
	"gochat/model"

	"github.com/sirupsen/logrus"
)

type UserDAO struct {
	db *libsql.DB
}

func MakeUserDAO(db *libsql.DB) UserDAO {
	return UserDAO{db}
}

func (dao *UserDAO) Create(user model.User) (bool, int64) {
	qchain := query.MakeInsertQueryChain(user).InsertModel(user)
	id, err := sql.InsertRowQuery(dao.db, &qchain)
	if err != nil {
		logrus.Warn("Failed to create user ", err)
		return false, -1
	}
	return true, id
}

func (dao *UserDAO) FindById(id int64) (bool, model.User) {
	user := model.User{}
	qchain := query.MakeSelectQueryChain(user).WhereKey("id", id)
	if err := sql.FindRowQuery(
		dao.db,
		&qchain,
		user.GetAddresses(),
	); err != nil {
		return false, user
	}
	return true, user
}

func (dao *UserDAO) FindAll() (bool, []model.User) {
	user := model.User{}
	users := []model.User{}
	onNext := func() {
		users = append(users, user)
	}
	qchain := query.MakeSelectQueryChain(user)
	qchain.GetValues()
	if err := sql.FindRowsQuery(
		dao.db,
		&qchain,
		user.GetAddresses(),
		onNext,
	); err != nil {
		return false, users
	}
	return true, users
}

func (dao *UserDAO) DeleteById(id int64) bool {
	qchain := query.MakeDeleteQueryChain(model.User{}).
		WhereKey("id", id)
	if raff, err := sql.DeleteRowQuery(dao.db, &qchain); err != nil {
		return false
	} else if raff != 1 {
		return false
	}
	return true
}

func (dao *UserDAO) UpdateById(id int64, user model.User) (bool, model.User) {
	qchain := query.MakeUpdateQueryChain(user).
		SetModel(user).
		WhereKey("id", id)
	if raff, err := sql.UpdateRowQuery(dao.db, &qchain); err != nil {
		return false, user
	} else if raff != 1 {
		return false, user
	}
	return dao.FindById(id)
}
