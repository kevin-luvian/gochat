package query

import (
	"gochat/lib/util"
	"strings"
)

func MakeUpdateQueryChain(o interface{}) UpdateQueryChain {
	return UpdateQueryChain{
		tablename: getModelTablename(o),
		sets:      []string{},
		setvals:   []interface{}{},
		qwhere:    makeWhereQuery(),
	}
}

type UpdateQueryChain struct {
	tablename string
	sets      []string
	setvals   []interface{}
	qwhere    whereQuery
}

func (uqc *UpdateQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("UPDATE ")
	b.WriteString(uqc.tablename)
	b.WriteString(" SET ")
	b.WriteString(strings.Join(uqc.sets, ", "))
	b.WriteString(" WHERE ")
	b.WriteString(uqc.qwhere.wheres)
	return b.String()
}

func (uqc UpdateQueryChain) GetValues() []interface{} {
	allvals := make([]interface{}, len(uqc.setvals)+len(uqc.qwhere.wherevals))
	for i := range uqc.setvals {
		allvals[i] = uqc.setvals[i]
	}
	for i := range uqc.qwhere.wherevals {
		allvals[i+len(uqc.setvals)] = uqc.qwhere.wherevals[i]
	}
	return allvals
}

func (uqc UpdateQueryChain) SetModel(o interface{}, fields ...string) UpdateQueryChain {
	mmeta := MakeModelMetadata(o).removePrimary()
	uqc.sets = make([]string, 0, len(mmeta.Fields))
	uqc.setvals = make([]interface{}, 0, len(mmeta.Fields))
	for _, field := range mmeta.Fields {
		if len(fields) == 0 || util.ArrStringContains(fields, field.name) {
			uqc.sets = append(uqc.sets, field.name+" = ?")
			uqc.setvals = append(uqc.setvals, field.value)
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
	uqc.qwhere.where(w, vals)
	return uqc
}

func (uqc UpdateQueryChain) WhereKey(k string, v interface{}) UpdateQueryChain {
	uqc.qwhere.whereKey(k, v)
	return uqc
}

func (uqc UpdateQueryChain) WhereModel(o interface{}, fields ...string) UpdateQueryChain {
	uqc.qwhere.whereModel(o, fields)
	return uqc
}
