package api

import (
	"gochat/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/temp", auth.Temp)
	r.GET("/login", auth.Temp)
	r.POST("/login/google", auth.LoginGoogle)
	r.GET("/signup", auth.Temp)
}
