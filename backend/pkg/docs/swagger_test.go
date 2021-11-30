package docs

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSwagger(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Swagger("/docs")(c)

	config := SwaggerConfig{}
	config.ToDefault("/docs")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<title>%s</title>", config.Title))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<link href='%s' rel='stylesheet' />", config.FontsCSS))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<link rel='stylesheet' type='text/css' href='%s' />", config.CSS))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<script src='%s'> </script>", config.BundleJS))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<script src='%s'> </script>", config.StandaloneJS))
}
