package MyAuth

import (
	"gochat/env"
	"gochat/helper/util"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	env.LoadDotEnvForTest()
}

func TestGenerateTokens(t *testing.T) {
	userid := "12345"

	if tok, err := GenerateAccessToken(userid); err != nil {
		t.Fatal("cant generate access token. ", err.Error())
	} else {
		logrus.Info("Access Token: ", tok)
	}

	if _, err := GenerateRefreshToken(userid); err != nil {
		t.Fatal("cant generate refresh token. ", err.Error())
	}

	rand := util.RandPrime(200)
	logrus.Info("Rand ", rand)
}

func TestParseAccessToken(t *testing.T) {
	userid := "12345"
	aToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjpmYWxzZSwiZXhwIjoxNjM2Nzg3MTc1LCJ1c2VyX2lkIjoiMTIzNDUifQ.O4FrGXe5CV_hn1MIB3Ur60rQTGLsnVU9hIiL4NY-oH4"
	if data, ok := ParseAccessToken(aToken); ok {
		t.Fatal("expired access token is still valid! ", data)
	}

	aToken, _ = GenerateAccessToken(userid)
	if _, ok := ParseAccessToken(aToken); !ok {
		t.Fatal("cant parse token: ", aToken)
	}

	if data, ok := ParseRefreshToken(aToken); ok {
		t.Fatal(`access token is accessible using refresh token parser `, data)
	}
}

func TestParseRefreshToken(t *testing.T) {
	userid := "12345"
	rToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzY3ODcwOTQsInVzZXJfaWQiOiIxMjM0NSJ9.fkvD-dY4XC02RdL4bUhopx0tnnna1Q1gZS8cAHeu75A"
	if data, ok := ParseRefreshToken(rToken); ok {
		t.Fatal("expired access token is still valid! ", data)
	}

	rToken, _ = GenerateRefreshToken(userid)
	if _, ok := ParseRefreshToken(rToken); !ok {
		t.Fatal("cant parse token: ", rToken)
	}
	if _, ok := ParseAccessToken(rToken); !ok {
		t.Fatal("cant parse token: ", rToken)
	}
}
