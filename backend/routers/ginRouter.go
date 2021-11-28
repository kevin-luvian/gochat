package routers

import (
	"gochat/routers/api"
	"gochat/routers/swagger"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	swagger.SwaggerRoute(r.Group("/swagger"))

	{
		apis := r.Group("/api")
		api.AuthRoutes(apis.Group("/auth"))
	}
	return r
}
