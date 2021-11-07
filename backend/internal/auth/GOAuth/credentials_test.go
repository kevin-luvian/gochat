package GOAuth

import (
	"log"
	"testing"
)

func TestCheckCredentialsData(t *testing.T) {
	log.Println("My Creds ID: ", MyCredentials.Cid)
	log.Println("My Creds Secret: ", MyCredentials.Csecret)

	log.Println("Google OAuth Config: ", GOAuthConf)
}

func TestCheckLoginCredential(t *testing.T) {
	token, url := MakeLoginURLCredential()
	log.Println("Token: ", token)
	log.Println("Login Url: ", url)
}
