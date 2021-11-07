package dao

import (
	libsql "database/sql"
	"gochat/lib/database/sql"
	"gochat/lib/database/sql/query"
	"gochat/model"

	"github.com/sirupsen/logrus"
)

type AuthStateDAO struct {
	db *libsql.DB
}

func MakeAuthStateDAO(db *libsql.DB) AuthStateDAO {
	return AuthStateDAO{db}
}

func (dao *AuthStateDAO) Create() (bool, string) {
	aState := model.MakeAuthState()
	qchain := query.MakeInsertQueryChain(aState).InsertModel(aState)
	_, err := sql.InsertRowQuery(dao.db, qchain.ToString(), qchain.GetValues())
	if err != nil {
		logrus.Warn("Failed to create auth state ", err)
		return false, ""
	}
	return true, aState.State
}

func (dao *AuthStateDAO) Exist(state string) bool {
	ok, aState := dao.findByState(state)
	if !ok {
		return false
	}
	if ok := dao.deleteById(aState.Id); !ok {
		return false
	}
	return true
}

func (dao *AuthStateDAO) findByState(state string) (bool, model.AuthState) {
	aState := model.AuthState{}
	qchain := query.MakeSelectQueryChain(aState).WhereKey("state", state)
	if err := sql.FindRowQuery(
		dao.db,
		qchain.ToString(),
		qchain.GetValues(),
		aState.GetAddresses(),
	); err != nil {
		return false, aState
	}
	return true, aState
}

func (dao *AuthStateDAO) deleteById(id int64) bool {
	qchain := query.MakeDeleteQueryChain(model.AuthState{}).
		WhereKey("id", id)
	if raff, err := sql.DeleteRowQuery(dao.db, qchain.ToString(), qchain.GetValues()); err != nil {
		return false
	} else if raff != 1 {
		return false
	}
	return true
}
