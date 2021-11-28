package auth

import (
	"gochat/pkg/app"
	"gochat/pkg/err"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/auth

// Temp function godoc
// @Summary Temporary Handler
// @Description check for server connection
// @Tags auth
// @Produce json
// @Success 200 {object} app.Response
// @Router /auth/temp [get]
func Temp(c *gin.Context) {
	ag := app.Gin{C: c}
	ag.Response(http.StatusOK, err.SUCCESS, nil)
}
