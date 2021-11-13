package GOAuth

import (
	"gochat/database"
	"gochat/env"
	"log"
	"testing"
)

func init() {
	env.LoadDotEnvForTest()
	database.GetRedis().FLUSH()
}

func TestCheckCredentialsData(t *testing.T) {
	conf := makeDefaultConf()
	log.Println("My Creds ID: ", conf.ClientID)
	log.Println("My Creds Secret: ", conf.ClientSecret)

	log.Println("Google OAuth Config: ", conf)
}

func TestCheckLoginCredential(t *testing.T) {
	token, url := MakeLoginURLCredential()
	log.Println("Token: ", token)
	log.Println("Login Url: ", url)
}
