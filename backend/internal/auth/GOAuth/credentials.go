package GOAuth

import (
	"gochat/dao"
	"gochat/database"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var authStateDAO dao.AuthStateDAO

var once sync.Once

var confInstance *oauth2.Config

func GetGOAuthConf() *oauth2.Config {
	if confInstance == nil {
		once.Do(makeGOAuthConf)
	}
	return confInstance
}

func makeGOAuthConf() {
	confInstance = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CID"),
		ClientSecret: os.Getenv("GOOGLE_CSECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func init() {
	authStateDAO = dao.MakeAuthStateDAO(database.MYSQLDB.GetDatabase())
}

func MakeLoginURLCredential() (string, string) {
	_, state := authStateDAO.Create()
	return state, GetGOAuthConf().AuthCodeURL(state)
}
