package api

import (
	"gochat/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.GET("/temp", auth.Temp)
}
