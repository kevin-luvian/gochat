package query

type RowQuery interface {
	ToString() string
	GetValues() []interface{}
}
