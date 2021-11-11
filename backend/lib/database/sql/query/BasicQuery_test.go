package query

import (
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Id       int    `sqldb:"pkey"`
	Name     string `sqldb:"nnull,unique"`
	Nickname string `sqldb:"unique,vchar-10"`
}

func makeUser() User {
	return User{Id: 0, Name: "bob", Nickname: "marley"}
}

func logQueryStr(q RowQuery) {
	logrus.Info(q.ToString(), " ", q.GetValues())
}

func TestCreateTableQuery(t *testing.T) {
	user := makeUser()
	q := MakeCreateTableQuery(user)
	logrus.Info(q.Qstr)
}

func TestDropTableQuery(t *testing.T) {
	q := MakeDropTableQuery(User{})
	qexpect := "DROP TABLE users CASCADE"
	assert.Equal(t, strings.ToLower(qexpect), strings.ToLower(q.Qstr))
	logrus.Info(q.Qstr)
}

func TestSelectQuery(t *testing.T) {
	var sqc SelectQueryChain
	user := makeUser()

	sqc = MakeSelectQueryChain(user)
	logQueryStr(&sqc)

	sqc = MakeSelectQueryChain(user).WhereKey("id", 10)
	logQueryStr(&sqc)

	sqc = MakeSelectQueryChain(user).WhereModel(user, "id", "nickname")
	logQueryStr(&sqc)

	sqc = MakeSelectQueryChain(user).
		Select("name").
		Where("name = ? OR nickname = ?", "bro", "bruh")
	logQueryStr(&sqc)

	sqc = MakeSelectQueryChain(user).Select("name", "nickname").WhereModel(user, "name")
	logQueryStr(&sqc)
}

func TestInsertQuery(t *testing.T) {
	var iqc InsertQueryChain
	user := makeUser()
	iqc = MakeInsertQueryChain(user).InsertModel(user)
	logQueryStr(&iqc)

	iqc = MakeInsertQueryChain(user).InsertManyModel(user, user, user)
	logQueryStr(&iqc)
}

func TestUpdateQuery(t *testing.T) {
	var uqc UpdateQueryChain
	user := makeUser()
	uqc = MakeUpdateQueryChain(user).Set("name", "Robert").Where("id = ?", user.Id)
	logQueryStr(&uqc)

	uqc = MakeUpdateQueryChain(user).SetModel(user, "name", "nickname").WhereModel(user)
	logQueryStr(&uqc)

	uqc = MakeUpdateQueryChain(user).SetModel(user).WhereKey("id", user.Id)
	logQueryStr(&uqc)
}

func TestDeleteQuery(t *testing.T) {
	var dqc DeleteQueryChain

	user := makeUser()
	dqc = MakeDeleteQueryChain(user).WhereKey("id", user.Id)
	logQueryStr(&dqc)

	dqc = MakeDeleteQueryChain(user).WhereModel(user, "name")
	logQueryStr(&dqc)
}
