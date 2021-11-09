package auth

import (
	"context"
	"gochat/helper"
	"gochat/internal/auth/GOAuth"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func temp(w http.ResponseWriter, r *http.Request) {
	helper.SuccessJSON(w, "healthy", "empty")
}

func authGoogle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	state := query.Get("state")
	code := query.Get("code")

	isStateExist := true
	if !isStateExist {
		helper.FailedJSON(w, http.StatusBadRequest, "state is not valid", nil)
		return
	}

	// Handle the exchange code to initiate a transport.
	tok, err := GOAuth.GetGOAuthConf().Exchange(context.Background(), code)
	if err != nil {
		helper.FailedJSON(w, http.StatusBadRequest, "failed to make exchange", nil)
	}
	logrus.Info("Token ", tok)

	// Construct the client.
	client := GOAuth.GetGOAuthConf().Client(context.Background(), tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		helper.FailedJSON(w, http.StatusBadRequest, "failed to get client", nil)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	log.Println("Resp body: ", string(data))

	helper.SuccessJSON(w, state, struct {
		State string `json:"state"`
	}{state})
}
