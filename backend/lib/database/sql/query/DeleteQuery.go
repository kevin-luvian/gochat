package query

import (
	"gochat/lib/database/sql/query/metadata"
	"gochat/lib/database/sql/query/where"
	"strings"
)

func MakeDeleteQueryChain(o interface{}) DeleteQueryChain {
	return DeleteQueryChain{
		tablename: metadata.GetModelTablename(o),
		qwhere:    where.MakeWhereQuery(),
	}
}

type DeleteQueryChain struct {
	tablename string
	qwhere    where.WhereQuery
}

func (dqc *DeleteQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("DELETE FROM ")
	b.WriteString(dqc.tablename)
	b.WriteString(" WHERE ")
	b.WriteString(dqc.qwhere.GetWheresString())
	return b.String()
}

func (dqc *DeleteQueryChain) GetValues() []interface{} {
	return dqc.qwhere.GetWheresValues()
}

func (dqc DeleteQueryChain) Where(w string, vals ...interface{}) DeleteQueryChain {
	dqc.qwhere = where.Where(dqc.qwhere, w, vals)
	return dqc
}

func (dqc DeleteQueryChain) WhereKey(k string, v interface{}) DeleteQueryChain {
	dqc.qwhere = where.WhereKey(dqc.qwhere, k, v)
	return dqc
}

func (dqc DeleteQueryChain) WhereModel(o interface{}, fields ...string) DeleteQueryChain {
	dqc.qwhere = where.WhereModel(dqc.qwhere, o, fields)
	return dqc
}
