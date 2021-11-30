package auth

import (
	"gochat/controllers/auth/google"
	"gochat/pkg/app"
	"gochat/pkg/db"
	"gochat/pkg/errc"
	"net/http"
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
// @Router /api/auth/temp [get]
func Temp(c *gin.Context) {
	c.Status(http.StatusOK)
}

type LoginGoogleReq struct {
	RedirectUrl string `json:"redirect_url" validate:"validurl" example:"http://localhost:8000/auth/google"`
}

type LoginGoogleRes struct {
	OAuthUrl string `json:"oauth_url" example:"https://accounts.google.com/o/oauth2/auth?..."`
	State    string `json:"state" example:"GoogleAuthCredential_12345"`
}

// @Tags auth
// @Produce json
// @Summary Create google oauth redirect
// @Param data body LoginGoogleReq true "login google request"
// @Success 200 {object} LoginGoogleRes
// @Failure 400 {object} app.ValidationError
// @Router /api/auth/login/google [post]
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
