package auth

import (
	"gochat/controllers/auth/google"
	"gochat/pkg/app"
	"gochat/pkg/db"
	"gochat/pkg/errc"
	"strings"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/auth

// Temp function godoc
// @Summary Temporary Handler
// @Description check for server connection
// @Tags auth
// @Produce json
// @Success 200 {object} nil
// @Router /auth/temp [get]
func Temp(c *gin.Context) {
	ag := app.Gin{C: c}
	ag.OkResponse(nil)
}

type LoginGoogleReq struct {
	RedirectUrl string `json:"redirect_url" validate:"validurl" example:"http://localhost:3000/redirect/google"`
}

type LoginGoogleRes struct {
	OAuthUrl string `json:"oauth_url"`
	State    string `json:"state"`
}

// @Tags auth
// @Produce json
// @Summary Create google oauth redirect
// @Param data body LoginGoogleReq true "login google request"
// @Success 200 {object} LoginGoogleRes
// @Failure 400 {object} app.ValidationError
// @Router /auth/login/google [post]
func LoginGoogle(c *gin.Context) {
	gapp := app.Gin{C: c}
	var form LoginGoogleReq

	if errCode := gapp.BindAndValid(&form); errCode != errc.Success {
		return
	}

	redirectUrl := strings.TrimSpace(form.RedirectUrl)
	state, url := google.MakeLoginRedirect(redirectUrl)
	tenMinutes := 60 * 10
	db.GetRedis().SETEX(state, tenMinutes, "")
	gapp.OkResponse(LoginGoogleRes{url, state})
}
