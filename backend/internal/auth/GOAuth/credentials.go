package GOAuth

import (
	"gochat/helper/util"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func makeGOAuthConf(redirectUrl string) oauth2.Config {
	return oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CID"),
		ClientSecret: os.Getenv("GOOGLE_CSECRET"),
		RedirectURL:  redirectUrl,
		// You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
}

func makeDefaultConf() oauth2.Config {
	return makeGOAuthConf(os.Getenv("GOOGLE_REDIRECT_URL"))
}

func MakeLoginURLCredential() (string, string) {
	state := "GoogleAuthCredential_" + util.MakeUUIDNoDash()
	conf := makeDefaultConf()
	return state, conf.AuthCodeURL(state)
}

func MakeLoginRedirect(redirectUrl string) (string, string) {
	state := "GoogleAuthCredential_" + util.MakeUUIDNoDash()
	conf := makeGOAuthConf(redirectUrl)
	return state, conf.AuthCodeURL(state)
}
