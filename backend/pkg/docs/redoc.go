package docs

import (
	"github.com/gin-gonic/gin"
)

type RedocOpts struct {
	Title      string
	SpecURL    string
	RedocURL   string
	FontURL    string
	FavIconURL string
}

func (r *RedocOpts) ToDefault(basePath string) {
	r.Title = "GoChat API Documentation"
	r.SpecURL = basePath + "/swagger.json"
	r.RedocURL = basePath + "/redoc.standalone.js"
	r.FontURL = basePath + "/fonts.css"
	r.FavIconURL = basePath + "/favicon-16x16.png"
}

func Redoc(staticFilePath string) gin.HandlerFunc {
	opts := RedocOpts{}
	opts.ToDefault(staticFilePath)
	b := ReadTemplate(redocTemplate, opts)

	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.Write(b)
		c.Writer.Flush()
	}
}

const redocTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>{{ .Title }}</title>
		<!-- needed for adaptive design -->
		<meta charset="utf-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link href='{{ .FontURL }}' rel='stylesheet'>
		<link rel="icon" type="image/png" href='{{.FavIconURL}}' sizes="16x16" />
    <!--
    ReDoc doesn't change outer page styles
    -->
    <style>
      body {
        margin: 0;
        padding: 0;
      }
    </style>
  </head>
  <body>
  	<h1> my new top bar </h1>
    <redoc spec-url='{{ .SpecURL }}'></redoc>
    <script src='{{ .RedocURL }}'> </script>
  </body>
</html>
`
