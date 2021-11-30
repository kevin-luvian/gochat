package docs

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRedoc(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Redoc("/docs")(c)

	rOpts := RedocOpts{}
	rOpts.ToDefault("/docs")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<title>%s</title>", rOpts.Title))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<link href='%s' rel='stylesheet'>", rOpts.FontURL))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<redoc spec-url='%s'></redoc>", rOpts.SpecURL))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<script src='%s'> </script>", rOpts.RedocURL))
}
