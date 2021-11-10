package query

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Id       int    `json:"id" sqldb:"pkey"`
	Name     string `json:"name" sqldb:"nnull,unique"`
	Nickname string `json:"nickname" sqldb:"unique,vchar-10"`
}

func makeUser() User {
	return User{Id: 0, Name: "bob", Nickname: "marley"}
}

func TestCreateTableQuery(t *testing.T) {
	user := makeUser()
	q := MakeCreateTableQuery(user)
	logrus.Info(q)
}

func TestDropTableQuery(t *testing.T) {
	q := MakeDropTableQuery(User{})
	qexpect := "DROP TABLE users CASCADE"
	assert.Equal(t, strings.ToLower(qexpect), strings.ToLower(q))
}

func TestSelectQuery(t *testing.T) {
	var sqc SelectQueryChain
	user := makeUser()

	sqc = MakeSelectQueryChain(user)
	logrus.Info(sqc.ToString(), sqc.GetValues())

	sqc = MakeSelectQueryChain(user).WhereKey("id", 10)
	logrus.Info(sqc.ToString(), sqc.GetValues())

	sqc = MakeSelectQueryChain(user).WhereModel(user, "id", "nickname")
	logrus.Info(sqc.ToString(), sqc.GetValues())

	sqc = MakeSelectQueryChain(user).
		Select("name").
		Where("name = ? OR nickname = ?", "bro", "bruh")
	logrus.Info(sqc.ToString(), sqc.GetValues())

	sqc = MakeSelectQueryChain(user).Select("name", "nickname").WhereModel(user, "name")
	logrus.Info(sqc.ToString(), sqc.GetValues())
}

func TestInsertQuery(t *testing.T) {
	var iqc InsertQueryChain
	user := makeUser()
	iqc = MakeInsertQueryChain(user).InsertModel(user)
	logrus.Info(iqc.ToString(), iqc.GetValues())

	iqc = MakeInsertQueryChain(user).InsertManyModel(user, user, user)
	logrus.Info(iqc.ToString(), iqc.GetValues())
}

func TestUpdateQuery(t *testing.T) {
	var uqc UpdateQueryChain
	user := makeUser()
	uqc = MakeUpdateQueryChain(user).Set("name", "Robert").Where("id = ?", user.Id)
	logrus.Info(uqc.ToString(), uqc.GetValues())

	uqc = MakeUpdateQueryChain(user).SetModel(user, "name", "nickname").WhereModel(user)
	logrus.Info(uqc.ToString(), uqc.GetValues())

	uqc = MakeUpdateQueryChain(user).SetModel(user).WhereKey("id", user.Id)
	logrus.Info(uqc.ToString(), uqc.GetValues())
}

func TestDeleteQuery(t *testing.T) {
	var dqc DeleteQueryChain

	user := makeUser()
	dqc = MakeDeleteQueryChain(user).WhereKey("id", user.Id)
	logrus.Info(dqc.ToString(), dqc.GetValues())

	dqc = MakeDeleteQueryChain(user).WhereModel(user, "name")
	logrus.Info(dqc.ToString(), dqc.GetValues())
}
