package MyAuth

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

var AccessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
var RefreshTokenSecret = os.Getenv("REFRESH_TOKEN_SECRET")

type tokenData struct {
	Authorized bool   `json:"authorized"`
	UserId     string `json:"user_id"`
	Exp        int64  `json:"exp"`
}

func generateToken(tdata tokenData, secret string) (string, error) {
	signkey := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	dataByte, _ := json.Marshal(tdata)
	json.Unmarshal(dataByte, &claims)

	tokenString, err := token.SignedString(signkey)
	if err != nil {
		logrus.Error("Something Went Wrong: ", err.Error())
		return "", err
	}

	return tokenString, nil
}

func parseToken(token string, secret string) (data tokenData, ok bool) {
	res := tokenData{}
	signkey := []byte(secret)
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return signkey, nil
	})

	if err != nil {
		logrus.Warn("failed to parse token. ", err.Error())
		return res, false
	}

	if m, ok := parsedtoken.Claims.(jwt.MapClaims); ok && parsedtoken.Valid {
		dataByte, _ := json.Marshal(m)
		if err := json.Unmarshal(dataByte, &res); err != nil {
			logrus.Warn("failed to unmarshal, token field is invalid. ", err.Error())
			return res, false
		}
		return res, true
	}
	return res, false
}

func GenerateAccessToken(userid string) (string, error) {
	exp := time.Now().Add(time.Minute * 15) // 15 minutes expiry
	data := tokenData{false, userid, exp.Unix()}
	return generateToken(data, AccessTokenSecret)
}

func GenerateRefreshToken(userid string) (string, error) {
	exp := time.Now().Add(time.Hour * 24 * 30) // 30 days expiry
	data := tokenData{true, userid, exp.Unix()}
	return generateToken(data, RefreshTokenSecret)
}

func ParseAccessToken(token string) (userid string, ok bool) {
	data, ok := parseToken(token, AccessTokenSecret)
	if !ok {
		return "", false
	}
	return data.UserId, true
}

func ParseRefreshToken(token string) (userid string, ok bool) {
	data, ok := parseToken(token, RefreshTokenSecret)
	if !ok || !data.Authorized {
		return "", false
	}
	return data.UserId, ok
}
