package model

type User struct {
	Id       int    `json:"id" sqldb:"pkey"`
	Name     string `json:"name" sqldb:"nnull"`
	Nickname string `json:"nickname" sqldb:"unique"`
}

func (usr *User) GetAddresses() []interface{} {
	return []interface{}{
		&usr.Id,
		&usr.Name,
		&usr.Nickname,
	}
}
