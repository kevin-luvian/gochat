package query

import (
	"gochat/lib/database/sql/query/metadata"
	"gochat/lib/database/sql/query/where"
	"strings"
)

func MakeSelectQuery(o interface{}) SelectQuery {
	return SelectQuery{
		tablename:  metadata.GetModelTablename(o),
		fieldnames: []string{},
		qwhere:     where.MakeWhereQuery(),
	}
}

type SelectQuery struct {
	tablename  string
	fieldnames []string
	qwhere     where.WhereQuery
}

func (iqc *SelectQuery) ToString() string {
	b := strings.Builder{}
	b.WriteString("SELECT ")
	if len(iqc.fieldnames) == 0 {
		b.WriteString("*")
	} else {
		b.WriteString(strings.Join(iqc.fieldnames, ", "))
	}
	b.WriteString(" FROM ")
	b.WriteString(iqc.tablename)
	if where.HasAny(&iqc.qwhere) {
		b.WriteString(" WHERE ")
		b.WriteString(iqc.qwhere.GetWheresString())
	}
	return b.String()
}

func (sqc *SelectQuery) GetValues() []interface{} {
	return sqc.qwhere.GetWheresValues()
}

func Select(iqc SelectQuery, fieldnames ...string) SelectQuery {
	iqc.fieldnames = fieldnames
	return iqc
}

func Where(iqc SelectQuery, w string, vals ...interface{}) SelectQuery {
	iqc.qwhere = where.Where(iqc.qwhere, w, vals)
	return iqc
}

func WhereKey(iqc SelectQuery, k string, v interface{}) SelectQuery {
	iqc.qwhere = where.WhereKey(iqc.qwhere, k, v)
	return iqc
}

func WhereModel(iqc SelectQuery, o interface{}, fields ...string) SelectQuery {
	iqc.qwhere = where.WhereModel(iqc.qwhere, o, fields)
	return iqc
}
