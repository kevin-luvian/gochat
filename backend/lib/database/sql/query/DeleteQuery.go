package query

import (
	"gochat/lib/database/sql/query/internal"
	"strings"
)

func MakeDeleteQueryChain(o interface{}) DeleteQueryChain {
	return DeleteQueryChain{
		tablename: internal.GetModelTablename(o),
		qwhere:    internal.MakeWhereQuery(),
	}
}

type DeleteQueryChain struct {
	tablename string
	qwhere    internal.WhereQuery
}

func (dqc *DeleteQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("DELETE FROM ")
	b.WriteString(dqc.tablename)
	b.WriteString(" WHERE ")
	b.WriteString(dqc.qwhere.GetQuery())
	return b.String()
}

func (dqc *DeleteQueryChain) GetValues() []interface{} {
	return dqc.qwhere.GetValues()
}

func (dqc DeleteQueryChain) Where(w string, vals ...interface{}) DeleteQueryChain {
	dqc.qwhere = dqc.qwhere.Args(w, vals)
	return dqc
}

func (dqc DeleteQueryChain) WhereKey(k string, v interface{}) DeleteQueryChain {
	dqc.qwhere = dqc.qwhere.Key(k, v)
	return dqc
}

func (dqc DeleteQueryChain) WhereModel(o interface{}, fields ...string) DeleteQueryChain {
	dqc.qwhere = dqc.qwhere.Model(o, fields)
	return dqc
}
