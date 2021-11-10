package query

import (
	"gochat/lib/database/sql/query/metadata"
	"strings"
)

func MakeCreateTableQuery(o interface{}) string {
	mmeta := metadata.MakeModelMetadata(o)
	b := strings.Builder{}
	b.WriteString("CREATE TABLE ")
	b.WriteString(mmeta.Tablename)
	b.WriteString("( ")
	nametags := make([]string, len(mmeta.Fields))
	for i, field := range mmeta.Fields {
		nametags[i] = field.Name + " " + field.GetSQLType() + " " + field.GetTagConstraints()
	}
	b.WriteString(strings.Join(nametags, ", "))
	b.WriteString(")")
	return b.String()
}
