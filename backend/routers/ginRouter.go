package routers

import (
	"gochat/pkg/docs"
	"gochat/routers/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(makeCORS())

	{
		rdocs := r.Group("/docs")
		rdocs.Static("/static", "./docs")
		rdocs.GET("/rapidoc", docs.Rapidoc("/docs/static"))
		rdocs.GET("/redoc", docs.Redoc("/docs/static"))
		rdocs.GET("/swagger", docs.Swagger("/docs/static"))
	}
	api.Routes(r.Group("/api"))
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
