package auth

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.GET("/temp", Temp)
	r.GET("/login", Temp)
	r.POST("/login/google", LoginGoogle)
	r.POST("/signup", Signup)
}
