package GOAuth

import (
	"gochat/database"
	"gochat/env"
	"gochat/lib/database/redis"
	"log"
	"testing"
)

func init() {
	env.LoadDotEnvForTest()
	rpool := database.GetRedisConnection()
	redis.FLUSH(rpool)
}

func TestCheckCredentialsData(t *testing.T) {
	log.Println("My Creds ID: ", GetGOAuthConf().ClientID)
	log.Println("My Creds Secret: ", GetGOAuthConf().ClientSecret)

	log.Println("Google OAuth Config: ", GetGOAuthConf())
}

func TestCheckLoginCredential(t *testing.T) {
	config := GetGOAuthConf()
	token, url := MakeLoginURLCredential(config)
	log.Println("Token: ", token)
	log.Println("Login Url: ", url)
}
