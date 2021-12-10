package auth

type LoginGoogleReq struct {
	RedirectUrl string `json:"redirect_url" validate:"validurl" example:"http://localhost:8000/auth/google"`
}

type LoginGoogleRes struct {
	OAuthUrl string `json:"oauth_url" example:"https://accounts.google.com/o/oauth2/auth?..."`
	State    string `json:"state" example:"GoogleAuthCredential_12345"`
}

type SignupReq struct {
	Username string `json:"username" example:"rick" validate:"nestr"`
	Email    string `json:"email" example:"abc@def.gh" validate:"email"`
	Password string `json:"password" example:"pass1234" validate:"nestr"`
}

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
