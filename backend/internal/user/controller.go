package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/user

// Temp function user
// @Summary Temporary Handler
// @Description test REST on User
// @Tags user
// @Produce json
// @Success 200 {object} nil
// @Router /api/user/temp [get]
func Temp(c *gin.Context) {
	c.Status(http.StatusOK)
}
