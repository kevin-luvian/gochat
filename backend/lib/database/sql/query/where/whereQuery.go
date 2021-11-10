package where

import (
	"gochat/lib/database/sql/query/metadata"
	"gochat/lib/util"
	"strings"

	"github.com/sirupsen/logrus"
)

type WhereQuery struct {
	Wheres    string
	Wherevals []interface{}
}

func (w *WhereQuery) GetWheresString() string {
	return w.Wheres
}

func (w *WhereQuery) GetWheresValues() []interface{} {
	return w.Wherevals
}

func MakeWhereQuery() WhereQuery {
	return WhereQuery{"", []interface{}{}}
}

func HasAny(w *WhereQuery) bool {
	return len(w.Wherevals) > 0
}

func Where(w WhereQuery, wstr string, vals []interface{}) WhereQuery {
	wqlen := len(strings.Split(wstr, "?"))

	if wqlen-1 != len(vals) {
		logrus.Panic("given arguments doesn't match query")
	}

	w.Wheres = wstr
	w.Wherevals = vals
	return w
}

func WhereKey(w WhereQuery, k string, v interface{}) WhereQuery {
	w.Wheres = k + " = ?"
	w.Wherevals = []interface{}{v}
	return w
}

func WhereModel(w WhereQuery, o interface{}, fields []string) WhereQuery {
	wherestring := make([]string, 0, len(fields))
	wherevals := make([]interface{}, 0, len(fields))
	mmeta := metadata.MakeModelMetadata(o)
	for _, field := range mmeta.Fields {
		if len(fields) == 0 || util.ArrStringContains(fields, field.Name) {
			wherestring = append(wherestring, field.Name+" = ?")
			wherevals = append(wherevals, util.AnyToStringQuery(field.Value))
		}
	}
	w.Wheres = strings.Join(wherestring, " AND ")
	w.Wherevals = wherevals
	return w
}
