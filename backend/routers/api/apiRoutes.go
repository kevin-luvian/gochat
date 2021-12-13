package api

import (
	"gochat/internal/auth"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	auth.Routes(r.Group("auth"))
}
