package GOAuth

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.Info("Starting credentials test")
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
