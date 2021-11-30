package docs

import (
	"github.com/gin-gonic/gin"
)

type RapiDocOpts struct {
	Title      string
	SpecURL    string
	RapiDocURL string
	FavIconURL string
}

func (r *RapiDocOpts) ToDefault(basePath string) {
	r.Title = "GoChat API Documentation"
	r.SpecURL = basePath + "/swagger.json"
	r.RapiDocURL = basePath + "/rapidoc-min.js"
	r.FavIconURL = basePath + "/favicon-16x16.png"
}

func Rapidoc(staticFilePath string) gin.HandlerFunc {
	opts := RapiDocOpts{}
	opts.ToDefault(staticFilePath)
	b := ReadTemplate(rapidocTemplate, opts)

	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.Write(b)
		c.Writer.Flush()
	}
}

const rapidocTemplate = `<!doctype html>
<html>
<head>
  <title>{{ .Title }}</title>
  <meta charset="utf-8"> <!-- Important: rapi-doc uses utf8 charecters -->
  <link rel='icon' type='image/png' href='{{.FavIconURL}}' sizes='16x16' />
  <script type='module' src='{{ .RapiDocURL }}'></script>
</head>
<body>
  <rapi-doc 
  	spec-url='{{ .SpecURL }}'
	theme='dark'
	route-prefix='#' />
</body>
</html>
`
