package query

import (
	"gochat/lib/database/sql/query/internal"
	"gochat/lib/util"
	"strings"
)

func MakeUpdateQueryChain(o interface{}) UpdateQueryChain {
	return UpdateQueryChain{
		tablename: internal.GetModelTablename(o),
		sets:      []string{},
		setvals:   []interface{}{},
		qwhere:    internal.MakeWhereQuery(),
	}
}

type UpdateQueryChain struct {
	tablename string
	sets      []string
	setvals   []interface{}
	qwhere    internal.WhereQuery
}

func (uqc *UpdateQueryChain) ToString() string {
	b := strings.Builder{}
	b.WriteString("UPDATE ")
	b.WriteString(uqc.tablename)
	b.WriteString(" SET ")
	b.WriteString(strings.Join(uqc.sets, ", "))
	b.WriteString(" WHERE ")
	b.WriteString(uqc.qwhere.GetQuery())
	return b.String()
}

func (uqc *UpdateQueryChain) GetValues() []interface{} {
	qwValues := uqc.qwhere.GetValues()
	svLen := len(uqc.setvals)
	qwLen := len(qwValues)

	allvals := make([]interface{}, svLen+qwLen)
	for i := 0; i < svLen; i++ {
		allvals[i] = uqc.setvals[i]
	}
	for i := 0; i < qwLen; i++ {
		allvals[svLen+i] = qwValues[i]
	}
	return allvals
}

func (uqc UpdateQueryChain) SetModel(o interface{}, fields ...string) UpdateQueryChain {
	mmeta := internal.MakeModelMetadata(o)
	mmeta = internal.RmvMModelPK(mmeta)

	uqc.sets = make([]string, 0, mmeta.FLen())
	uqc.setvals = make([]interface{}, 0, mmeta.FLen())

	fnames := mmeta.GetNames()
	fvalues := mmeta.GetValues()
	for i := 0; i < mmeta.FLen(); i++ {
		if len(fields) == 0 || util.ArrStringContains(fields, fnames[i]) {
			uqc.sets = append(uqc.sets, fnames[i]+" = ?")
			uqc.setvals = append(uqc.setvals, fvalues[i])
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
	uqc.qwhere = uqc.qwhere.Args(w, vals)
	return uqc
}

func (uqc UpdateQueryChain) WhereKey(k string, v interface{}) UpdateQueryChain {
	uqc.qwhere = uqc.qwhere.Key(k, v)
	return uqc
}

func (uqc UpdateQueryChain) WhereModel(o interface{}, fields ...string) UpdateQueryChain {
	uqc.qwhere = uqc.qwhere.Model(o, fields)
	return uqc
}
