package internal

import (
	"gochat/lib/util"
	"strings"

	"github.com/sirupsen/logrus"
)

func MakeWhereQuery() WhereQuery {
	return WhereQuery{"", []interface{}{}}
}

type WhereQuery struct {
	wheres    string
	wherevals []interface{}
}

func (w *WhereQuery) GetQuery() string {
	return w.wheres
}

func (w *WhereQuery) GetValues() []interface{} {
	return w.wherevals
}

func (w *WhereQuery) HasAny() bool {
	return len(w.wherevals) > 0
}

func (w WhereQuery) Args(wstr string, vals []interface{}) WhereQuery {
	wqlen := len(strings.Split(wstr, "?"))

	if wqlen-1 != len(vals) {
		logrus.Panic("given arguments doesn't match query")
	}

	w.wheres = wstr
	w.wherevals = vals
	return w
}

func (w WhereQuery) Key(k string, v interface{}) WhereQuery {
	w.wheres = k + " = ?"
	w.wherevals = []interface{}{v}
	return w
}

func (w WhereQuery) Model(o interface{}, fields []string) WhereQuery {
	wherestring := make([]string, 0, len(fields))
	wherevals := make([]interface{}, 0, len(fields))
	mmeta := MakeModelMetadata(o)

	fnames := mmeta.GetNames()
	fvalues := mmeta.GetValues()

	for i := 0; i < mmeta.FLen(); i++ {
		if len(fields) == 0 || util.ArrStringContains(fields, fnames[i]) {
			wherestring = append(wherestring, fnames[i]+" = ?")
			wherevals = append(wherevals, fvalues[i])
		}
	}

	w.wheres = strings.Join(wherestring, " AND ")
	w.wherevals = wherevals
	return w
}
