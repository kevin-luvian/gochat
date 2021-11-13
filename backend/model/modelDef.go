package model

type Model interface {
	GetAddresses() []interface{}
	Equal(interface{}) bool
}
