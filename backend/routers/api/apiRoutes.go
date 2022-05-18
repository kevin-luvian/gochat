package api

import (
	"gochat/internal/auth"
	"gochat/internal/user"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	auth.Routes(r.Group("auth"))
	user.Routes(r.Group("user"))
}
