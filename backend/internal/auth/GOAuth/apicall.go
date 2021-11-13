package GOAuth

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func makeClient(token *oauth2.Token) *http.Client {
	conf := makeDefaultConf()
	return conf.Client(context.Background(), token)
}

func MakeToken(code string) (*oauth2.Token, error) {
	conf := makeDefaultConf()
	return conf.Exchange(context.Background(), code)
}

func GetUserInfo(token *oauth2.Token) (UserInfo, error) {
	userInfo := UserInfo{}

	url := "https://www.googleapis.com/oauth2/v3/userinfo"
	client := makeClient(token)
	resp, err := client.Get(url)
	if err != nil {
		return userInfo, err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err := json.Unmarshal(data, &userInfo); err != nil {
		logrus.Error("Failed to unmarshal ", err)
		return userInfo, err
	}
	return userInfo, nil
}
