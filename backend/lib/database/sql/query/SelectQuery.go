package query

import (
	"strings"
)

func MakeSelectQueryChain(o interface{}) SelectQueryChain {
	return SelectQueryChain{
		tablename:  getModelTablename(o),
		fieldnames: []string{},
		qwhere:     makeWhereQuery(),
	}
}

type SelectQueryChain struct {
	tablename  string
	fieldnames []string
	qwhere     whereQuery
}

func (iqc *SelectQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("SELECT ")
	if len(iqc.fieldnames) == 0 {
		b.WriteString("*")
	} else {
		b.WriteString(strings.Join(iqc.fieldnames, ", "))
	}
	b.WriteString(" FROM ")
	b.WriteString(iqc.tablename)
	if iqc.qwhere.hasAny() {
		b.WriteString(" WHERE ")
		b.WriteString(iqc.qwhere.wheres)
	}
	return b.String()
}

func (sqc *SelectQueryChain) GetValues() []interface{} {
	return sqc.qwhere.wherevals
}

func (iqc SelectQueryChain) Select(fieldnames ...string) SelectQueryChain {
	iqc.fieldnames = fieldnames
	return iqc
}

func (iqc SelectQueryChain) Where(w string, vals ...interface{}) SelectQueryChain {
	iqc.qwhere.where(w, vals)
	return iqc
}

func (iqc SelectQueryChain) WhereKey(k string, v interface{}) SelectQueryChain {
	iqc.qwhere.whereKey(k, v)
	return iqc
}

func (iqc SelectQueryChain) WhereModel(o interface{}, fields ...string) SelectQueryChain {
	iqc.qwhere.whereModel(o, fields)
	return iqc
}
