package auth

import (
	"gochat/controllers/auth/google"
	"gochat/pkg/app"
	"gochat/pkg/errc"
	"net/http"

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
	RedirectUrl string `json:"redirect_url" validate:"nestr" example:"http://localhost:3000/redirect/google"`
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
// @Failure 400 {array} app.VErr
// @Failure 500 {object} app.ErrResponse
// @Router /auth/login/google [post]
func LoginGoogle(c *gin.Context) {
	gapp := app.Gin{C: c}
	var form LoginGoogleReq

	errCode, vErr := gapp.BindAndValid(&form)
	if errCode == errc.FailedValidation {
		gapp.Response(http.StatusBadRequest, vErr)
		return
	} else if errCode != errc.Success {
		gapp.ErrResponse(http.StatusInternalServerError, errCode)
		return
	}

	state, url := google.MakeLoginRedirect(form.RedirectUrl)
	gapp.OkResponse(LoginGoogleRes{url, state})
}
