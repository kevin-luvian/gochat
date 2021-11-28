package swagger

import (
	_ "gochat/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func SwaggerRoute(r *gin.RouterGroup) {
	r.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
