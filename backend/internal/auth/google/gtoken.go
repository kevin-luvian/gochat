package google

import (
	"gochat/env"
	"gochat/helper/util"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func makeGOAuthConf(redirectUrl string) oauth2.Config {
	return oauth2.Config{
		ClientID:     env.GetStr(env.GOOGLE_CID),
		ClientSecret: env.GetStr(env.GOOGLE_CSECRET),
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

func MakeLoginRedirect(redirectUrl string) (state, url string) {
	stateID := "GoogleAuthCredential_" + util.MakeUUIDNoDash()
	conf := makeGOAuthConf(redirectUrl)
	return stateID, conf.AuthCodeURL(stateID)
}
