package model

type User struct {
	Id         int    `json:"id" sqldb:"pkey"`
	UserID     string `json:"user_id" sqldb:"unique"`
	Name       string `json:"name" sqldb:"nnull"`
	ProfileUrl string `json:"profile_url"`
	Password   string `json:"-"`
}

func (usr *User) GetAddresses() []interface{} {
	return []interface{}{
		&usr.Id,
		&usr.UserID,
		&usr.Name,
		&usr.ProfileUrl,
		&usr.Password,
	}
}

func (usr *User) Equal(b interface{}) bool {
	usrB, ok := b.(User)
	return ok &&
		usr.Name == usrB.Name &&
		usr.ProfileUrl == usrB.ProfileUrl &&
		usr.UserID == usrB.UserID &&
		usr.Password == usrB.Password
}
