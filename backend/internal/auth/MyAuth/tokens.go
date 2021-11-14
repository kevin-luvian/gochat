package MyAuth

import (
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
	jwt.StandardClaims
}

func (t tokenData) Valid() error {
	if t.ExpiresAt < time.Now().Unix() {
		return errors.New("token has expired")
	}
	return nil
}

func generateToken(claims tokenData, secret string) (string, error) {
	claims.Issuer = "gochat"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signkey := []byte(secret)
	signedToken, err := token.SignedString(signkey)
	if err != nil {
		logrus.Error("Something Went Wrong: ", err.Error())
		return "", err
	}

	return signedToken, nil
}

func parseToken(token string, secret string) (data tokenData, ok bool) {
	signkey := []byte(secret)
	var claims tokenData

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("there was an error in parsing")
			}
			return signkey, nil
		},
	)

	if err != nil || !parsedToken.Valid || !claims.VerifyIssuer("gochat", true) {
		logrus.Warn("failed to parse token. ", err.Error())
		return tokenData{}, false
	}

	return claims, true
}

func GenerateAccessToken(userid string) (string, error) {
	exp := time.Now().Add(time.Minute * 15) // 15 minutes expiry
	data := tokenData{
		Authorized: false,
		UserId:     userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
	return generateToken(data, AccessTokenSecret)
}

func GenerateRefreshToken(userid string) (string, error) {
	exp := time.Now().Add(time.Hour * 24 * 30) // 30 days expiry
	data := tokenData{
		Authorized: true,
		UserId:     userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}
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
