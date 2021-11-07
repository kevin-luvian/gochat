package model

import "gochat/helper/util"

type AuthState struct {
	Id    int64  `sqldb:"pkey"`
	State string `sqldb:"nnull,unique"`
}

func MakeAuthState() AuthState {
	return AuthState{State: util.RandString(30)}
}

func (as *AuthState) GetAddresses() []interface{} {
	return []interface{}{&as.Id, &as.State}
}
