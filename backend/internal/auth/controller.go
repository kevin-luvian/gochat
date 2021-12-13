package auth

import (
	"gochat/internal/auth/MyAuth"
	"gochat/internal/auth/google"
	"gochat/pkg/app"
	"gochat/pkg/db"
	"gochat/pkg/errc"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

// @Tags auth
// @Summary Create new user account
// @Param data body SignupReq true "create new user"
// @Success 200 {object} TokenRes
// @Failure 400 {object} app.ValidationError
// @Failure 500 {object} app.ErrResponse
// @Router /api/auth/signup [post]
func Signup(c *gin.Context) {
	gapp := app.Gin{C: c}
	var form SignupReq
	logrus.Info("Signing up")

	if errCode := gapp.BindAndValid(&form); errCode != errc.Success {
		return
	}

	// username := form.Username
	// password := form.Password
	// email := form.Email
	// state, url := google.MakeLoginRedirect(redirectUrl)
	accTok, erra := MyAuth.GenerateAccessToken("abc")
	refTok, errf := MyAuth.GenerateRefreshToken("abc")
	if erra != nil || errf != nil {
		gapp.AppErrResponse("failed to generate token")
		return
	}

	gapp.OkResponse(TokenRes{AccessToken: accTok, RefreshToken: refTok})
}
