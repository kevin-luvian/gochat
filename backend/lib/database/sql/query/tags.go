package query

var TAG_UNIQUE = "unique"
var TAG_NOT_NULL = "nnull"
var TAG_DEFAULT = "default"
var TAG_PRIMARY_KEY = "pkey"
var TAG_FOREIGN_KEY = "fkey"

var MOD_TAG_VARCHAR = "vchar"
var MOD_TAG_CHAR = "char"

var tagDef = map[string]string{
	TAG_UNIQUE:      "UNIQUE",
	TAG_NOT_NULL:    "NOT NULL",
	TAG_PRIMARY_KEY: "PRIMARY KEY",
	TAG_DEFAULT:     "DEFAULT",
	TAG_FOREIGN_KEY: "REFERENCES",
}
