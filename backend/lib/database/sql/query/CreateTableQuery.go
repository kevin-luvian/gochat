package query

import (
	"strings"
)

func MakeCreateTableQuery(o interface{}) string {
	mmeta := MakeModelMetadata(o)
	b := strings.Builder{}
	b.WriteString("CREATE TABLE ")
	b.WriteString(mmeta.Tablename)
	b.WriteString("( ")
	nametags := make([]string, len(mmeta.Fields))
	for i, field := range mmeta.Fields {
		nametags[i] = field.name + " " + field.getSQLType() + " " + field.getTagConstraints()
	}
	b.WriteString(strings.Join(nametags, ", "))
	b.WriteString(")")
	return b.String()
}
