package query

import "gochat/lib/database/sql/query/internal"

func MakeDropTableQuery(o interface{}) TableQuery {
	mmeta := internal.MakeModelMetadata(o)
	qstr := "DROP TABLE " + mmeta.Tablename + " CASCADE"
	return TableQuery{mmeta.Tablename, qstr}
}
