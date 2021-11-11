package query

import (
	"gochat/lib/database/sql/query/internal"
	"strings"
)

func MakeCreateTableQuery(o interface{}) TableQuery {
	mmeta := internal.MakeModelMetadata(o)
	b := strings.Builder{}
	b.WriteString("CREATE TABLE ")
	b.WriteString(mmeta.Tablename)
	b.WriteString("( ")
	b.WriteString(strings.Join(mmeta.GetFieldSQLTags(), ", "))
	b.WriteString(")")
	return TableQuery{mmeta.Tablename, b.String()}
}
