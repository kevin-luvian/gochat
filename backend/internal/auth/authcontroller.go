package auth

import (
	"gochat/database"
	"gochat/helper"
	httpHelper "gochat/helper/http"
	"gochat/helper/util"
	"gochat/internal/auth/GOAuth"
	"gochat/internal/auth/MyAuth"
	"net/http"

	"github.com/sirupsen/logrus"
)

func temp(w http.ResponseWriter, r *http.Request) {
	helper.SuccessJSON(w, "healthy", "empty")
}

func loginGoogle(w http.ResponseWriter, r *http.Request) {
	state, url := GOAuth.MakeLoginURLCredential()
	database.GetRedis().SETEX(state, util.SecHour(1), "")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func loginGoogleRedirect(w http.ResponseWriter, r *http.Request) {
	body, err := httpHelper.RequestBody(r)
	if err != nil {
		helper.FailedJSON(w, 400, `cant read the request body`, nil)
		return
	}

	redirectUrl, ok := body["redirect_url"].(string)
	if !ok {
		helper.FailedJSON(w, 400, `redirect url must be provided`, nil)
		return
	}

	state, url := GOAuth.MakeLoginRedirect(redirectUrl)
	database.GetRedis().SETEX(state, util.SecHour(1), "")

	res := struct {
		OAuthUrl string `json:"oauth_url"`
		State    string `json:"state"`
	}{url, state}

	helper.SuccessJSON(w, `google oauth url created`, res)
}

func authGoogle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	state := query.Get("state")
	code := query.Get("code")

	if !database.GetRedis().EXIST(state) {
		helper.FailedJSON(w, http.StatusBadRequest, "state is not valid", nil)
		return
	}

	tok, err := GOAuth.MakeToken(code)
	if err != nil {
		helper.FailedJSON(w, http.StatusBadRequest, "failed to make token exchange", nil)
		return
	}

	userInfo, err := GOAuth.GetUserInfo(tok)
	if err != nil {
		logrus.Error(err.Error())
		helper.FailedJSON(w, http.StatusBadRequest, "failed to get user info", nil)
		return
	}

	gentok, _ := MyAuth.GenerateAccessToken("12345")
	result := struct {
		GOAuth.UserInfo
		AccessToken string `json:"access_token"`
		GToken      string `json:"generated_token"`
	}{userInfo, tok.AccessToken, gentok}

	helper.SuccessJSON(w, "login using gmail success", result)
}

func refresh(w http.ResponseWriter, r *http.Request) {
	body, err := httpHelper.RequestBody(r)
	if err != nil {
		helper.FailedJSON(w, 400, `cant read the request body`, nil)
		return
	}

	rToken := body["refresh_token"].(string)
	userId, ok := MyAuth.ParseAccessToken(rToken)
	if !ok {
		helper.FailedJSON(w, 400, `refresh token is invalid`, nil)
		return
	}

	newAToken, natErr := MyAuth.GenerateAccessToken(userId)
	newRToken, nrtErr := MyAuth.GenerateRefreshToken(userId)
	if natErr != nil || nrtErr != nil {
		helper.FailedJSON(w, 400, `failed to generate token, please try again.`, nil)
		return
	}

	result := struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}{newAToken, newRToken}

	helper.SuccessJSON(w, `success generating new tokens`, result)
}

func parseToken(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	token := query.Get("token")

	ptok, ok := MyAuth.ParseAccessToken(token)
	if !ok {
		helper.FailedJSON(w, 400, `token is invalid`, nil)
		return
	}

	helper.SuccessJSON(w, "parsed token", ptok)
}
