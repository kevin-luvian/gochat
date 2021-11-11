package query

import (
	"gochat/lib/database/sql/query/internal"
	"strings"
)

func MakeSelectQueryChain(o interface{}) SelectQueryChain {
	return SelectQueryChain{
		tablename:  internal.GetModelTablename(o),
		fieldnames: []string{},
		qwhere:     internal.MakeWhereQuery(),
	}
}

type SelectQueryChain struct {
	tablename  string
	fieldnames []string
	qwhere     internal.WhereQuery
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
	if iqc.qwhere.HasAny() {
		b.WriteString(" WHERE ")
		b.WriteString(iqc.qwhere.GetQuery())
	}
	return b.String()
}

func (sqc *SelectQueryChain) GetValues() []interface{} {
	return sqc.qwhere.GetValues()
}

func (iqc SelectQueryChain) Select(fieldnames ...string) SelectQueryChain {
	iqc.fieldnames = fieldnames
	return iqc
}

func (iqc SelectQueryChain) Where(w string, vals ...interface{}) SelectQueryChain {
	iqc.qwhere = iqc.qwhere.Args(w, vals)
	return iqc
}

func (iqc SelectQueryChain) WhereKey(k string, v interface{}) SelectQueryChain {
	iqc.qwhere = iqc.qwhere.Key(k, v)
	return iqc
}

func (iqc SelectQueryChain) WhereModel(o interface{}, fields ...string) SelectQueryChain {
	iqc.qwhere = iqc.qwhere.Model(o, fields)
	return iqc
}
