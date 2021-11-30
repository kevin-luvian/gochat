package docs

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRapidoc(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	Rapidoc("/docs")(c)

	rOpts := RapiDocOpts{}
	rOpts.ToDefault("/docs")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<title>%s</title>", rOpts.Title))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<link rel='icon' type='image/png' href='%s' sizes='16x16' />", rOpts.FavIconURL))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("<script type='module' src='%s'></script>", rOpts.RapiDocURL))
	assert.Contains(t, w.Body.String(), fmt.Sprintf("spec-url='%s'", rOpts.SpecURL))
}
