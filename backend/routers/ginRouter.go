package routers

import (
	"gochat/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	apiGr := r.Group("/api")
	api.AuthRoutes(apiGr.Group("/auth"))
	return r
}
