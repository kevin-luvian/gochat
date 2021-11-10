package query

import (
	"gochat/lib/database/sql/query/metadata"
	"gochat/lib/database/sql/query/where"
	"gochat/lib/util"
	"strings"
)

func MakeUpdateQueryChain(o interface{}) UpdateQueryChain {
	return UpdateQueryChain{
		tablename: metadata.GetModelTablename(o),
		sets:      []string{},
		setvals:   []interface{}{},
		qwhere:    where.MakeWhereQuery(),
	}
}

type UpdateQueryChain struct {
	tablename string
	sets      []string
	setvals   []interface{}
	qwhere    where.WhereQuery
}

func (uqc *UpdateQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("UPDATE ")
	b.WriteString(uqc.tablename)
	b.WriteString(" SET ")
	b.WriteString(strings.Join(uqc.sets, ", "))
	b.WriteString(" WHERE ")
	b.WriteString(uqc.qwhere.Wheres)
	return b.String()
}

func (uqc UpdateQueryChain) GetValues() []interface{} {
	allvals := make([]interface{}, len(uqc.setvals)+len(uqc.qwhere.Wherevals))
	for i := range uqc.setvals {
		allvals[i] = uqc.setvals[i]
	}
	for i := range uqc.qwhere.Wherevals {
		allvals[i+len(uqc.setvals)] = uqc.qwhere.Wherevals[i]
	}
	return allvals
}

func (uqc UpdateQueryChain) SetModel(o interface{}, fields ...string) UpdateQueryChain {
	mmeta := metadata.MakeModelMetadata(o)
	mmeta = metadata.RemovePrimary(mmeta)
	uqc.sets = make([]string, 0, len(mmeta.Fields))
	uqc.setvals = make([]interface{}, 0, len(mmeta.Fields))
	for _, field := range mmeta.Fields {
		if len(fields) == 0 || util.ArrStringContains(fields, field.Name) {
			uqc.sets = append(uqc.sets, field.Name+" = ?")
			uqc.setvals = append(uqc.setvals, field.Value)
		}
	}
	return uqc
}

func (uqc UpdateQueryChain) Set(k string, v interface{}) UpdateQueryChain {
	uqc.sets = append(uqc.sets, k+" = ?")
	uqc.setvals = append(uqc.setvals, v)
	return uqc
}

func (uqc UpdateQueryChain) Where(w string, vals ...interface{}) UpdateQueryChain {
	uqc.qwhere = where.Where(uqc.qwhere, w, vals)
	return uqc
}

func (uqc UpdateQueryChain) WhereKey(k string, v interface{}) UpdateQueryChain {
	uqc.qwhere = where.WhereKey(uqc.qwhere, k, v)
	return uqc
}

func (uqc UpdateQueryChain) WhereModel(o interface{}, fields ...string) UpdateQueryChain {
	uqc.qwhere = where.WhereModel(uqc.qwhere, o, fields)
	return uqc
}
