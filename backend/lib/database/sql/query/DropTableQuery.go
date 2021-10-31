package query

func MakeDropTableQuery(o interface{}) string {
	mmeta := MakeModelMetadata(o)
	return "DROP TABLE " + mmeta.Tablename + " CASCADE"
}
