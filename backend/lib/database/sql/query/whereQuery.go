package query

import (
	"gochat/lib/util"
	"strings"

	"github.com/sirupsen/logrus"
)

type whereQuery struct {
	wheres    string
	wherevals []interface{}
}

func makeWhereQuery() whereQuery {
	return whereQuery{"", []interface{}{}}
}

func (w *whereQuery) hasAny() bool {
	return len(w.wherevals) > 0
}

func (w *whereQuery) where(wstr string, vals []interface{}) {
	wqlen := len(strings.Split(wstr, "?"))

	if wqlen-1 != len(vals) {
		logrus.Panic("given arguments doesn't match query")
	}

	w.wheres = wstr
	w.wherevals = vals
}

func (w *whereQuery) whereKey(k string, v interface{}) {
	w.wheres = k + " = ?"
	w.wherevals = []interface{}{v}
}

func (w *whereQuery) whereModel(o interface{}, fields []string) {
	wherestring := make([]string, 0, len(fields))
	wherevals := make([]interface{}, 0, len(fields))
	mmeta := MakeModelMetadata(o)
	for _, field := range mmeta.Fields {
		if len(fields) == 0 || util.ArrStringContains(fields, field.name) {
			wherestring = append(wherestring, field.name+" = ?")
			wherevals = append(wherevals, util.AnyToStringQuery(field.value))
		}
	}
	w.wheres = strings.Join(wherestring, " AND ")
	w.wherevals = wherevals
}
