package GOAuth

import (
	"encoding/json"
	"gochat/dao"
	"gochat/database"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var authStateDAO dao.AuthStateDAO
var MyCredentials Credentials
var GOAuthConf oauth2.Config

// Credentials which stores google ids.
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

func init() {
	path, _ := os.Getwd()
	file, err := ioutil.ReadFile(path + "/internal/auth/GOAuth/creds.json")
	if err != nil {
		logrus.Fatal("File error: ", err)
	}

	err = json.Unmarshal(file, &MyCredentials)
	if err != nil {
		logrus.Fatal("Unmarshal error: ", err)
	}

	GOAuthConf = oauth2.Config{
		ClientID:     MyCredentials.Cid,
		ClientSecret: MyCredentials.Csecret,
		RedirectURL:  "http://localhost:8000/auth/google",
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	authStateDAO = dao.MakeAuthStateDAO(database.MYSQLDB.GetDatabase())
}

func MakeLoginURLCredential() (string, string) {
	_, state := authStateDAO.Create()
	return state, GOAuthConf.AuthCodeURL(state)
}
