package auth

import (
	"gochat/pkg/app"
	"gochat/pkg/err"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Temp(c *gin.Context) {
	ag := app.Gin{C: c}
	ag.Response(http.StatusOK, err.SUCCESS, nil)
}
