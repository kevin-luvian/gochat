package query

import "gochat/lib/database/sql/query/metadata"

func MakeDropTableQuery(o interface{}) string {
	tablename := metadata.GetModelTablename(o)
	return "DROP TABLE " + tablename + " CASCADE"
}
