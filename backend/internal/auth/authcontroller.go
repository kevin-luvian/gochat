package auth

import (
	"context"
	"encoding/json"
	"gochat/database"
	"gochat/helper"
	"gochat/helper/util"
	"gochat/internal/auth/GOAuth"
	"gochat/lib/database/redis"
	"io/ioutil"
	"net/http"
)

func temp(w http.ResponseWriter, r *http.Request) {
	helper.SuccessJSON(w, "healthy", "empty")
}

func loginGoogle(w http.ResponseWriter, r *http.Request) {
	config := GOAuth.GetGOAuthConf()
	state, url := GOAuth.MakeLoginURLCredential(config)

	rpool := database.GetRedisConnection()
	redis.SETEX(rpool, state, util.SecHour(1), "")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func authGoogle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	state := query.Get("state")
	code := query.Get("code")

	rpool := database.GetRedisConnection()

	if !redis.EXIST(rpool, state) {
		helper.FailedJSON(w, http.StatusBadRequest, "state is not valid", nil)
		return
	}

	// Handle the exchange code to initiate a transport.
	tok, err := GOAuth.GetGOAuthConf().Exchange(context.Background(), code)
	if err != nil {
		helper.FailedJSON(w, http.StatusBadRequest, "failed to token make exchange", nil)
		return
	}

	// Construct the client.
	client := GOAuth.GetGOAuthConf().Client(context.Background(), tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		helper.FailedJSON(w, http.StatusBadRequest, "failed to get client", nil)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	m := make(map[string]string)
	json.Unmarshal(data, &m)
	m["state"] = state

	helper.SuccessJSON(w, "login using gmail success", m)
}
