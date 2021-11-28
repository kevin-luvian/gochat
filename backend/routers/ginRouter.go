package routers

import (
	"gochat/routers/api"
	"gochat/routers/swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(makeCORS())

	swagger.SwaggerRoute(r.Group("/swagger"))

	{
		apis := r.Group("/api")
		api.AuthRoutes(apis.Group("/auth"))
	}
	return r
}

func makeCORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("authorization")
	// corsConfig.AddAllowMethods("OPTIONS")
	return cors.New(corsConfig)
}
