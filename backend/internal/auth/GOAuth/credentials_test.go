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
	log.Println("My Creds ID: ", GetGOAuthConf().ClientID)
	log.Println("My Creds Secret: ", GetGOAuthConf().ClientSecret)

	log.Println("Google OAuth Config: ", GetGOAuthConf())
}

func TestCheckLoginCredential(t *testing.T) {
	token, url := MakeLoginURLCredential()
	log.Println("Token: ", token)
	log.Println("Login Url: ", url)
}
