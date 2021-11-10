package query

import (
	"strings"
)

func MakeDeleteQueryChain(o interface{}) DeleteQueryChain {
	return DeleteQueryChain{
		tablename: getModelTablename(o),
		qwhere:    makeWhereQuery(),
	}
}

type DeleteQueryChain struct {
	tablename string
	qwhere    whereQuery
}

func (dqc *DeleteQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("DELETE FROM ")
	b.WriteString(dqc.tablename)
	b.WriteString(" WHERE ")
	b.WriteString(dqc.qwhere.wheres)
	return b.String()
}

func (dqc *DeleteQueryChain) GetValues() []interface{} {
	return dqc.qwhere.wherevals
}

func (dqc DeleteQueryChain) Where(w string, vals ...interface{}) DeleteQueryChain {
	dqc.qwhere.where(w, vals)
	return dqc
}

func (dqc DeleteQueryChain) WhereKey(k string, v interface{}) DeleteQueryChain {
	dqc.qwhere.whereKey(k, v)
	return dqc
}

func (dqc DeleteQueryChain) WhereModel(o interface{}, fields ...string) DeleteQueryChain {
	dqc.qwhere.whereModel(o, fields)
	return dqc
}
