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
	MYDOCCSS   string
}

func (r *RedocOpts) ToDefault(basePath string) {
	r.Title = "GoChat API Documentation"
	r.SpecURL = basePath + "/swagger.json"
	r.RedocURL = basePath + "/redoc.standalone.js"
	r.FontURL = basePath + "/fonts.css"
	r.FavIconURL = basePath + "/favicon-16x16.png"
	r.MYDOCCSS = basePath + "/my-doc-ui.css"
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
		<link rel='stylesheet' type='text/css' href='{{.MYDOCCSS}}' />
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
	<script type="text/javascript">
	function redirectTo(endpoint) {
	  const baseUrl = window.location.protocol+"//"+window.location.host;
	  const currPath = window.location.pathname.split('/');
	  const docPath = currPath.slice(0, currPath.length - 1).join('/');
	  const targetUrl = baseUrl+docPath+"/"+endpoint; 
	  window.location.replace(targetUrl);
	}
	</script>
  </head>
  <body>
	<div class="doc-header">
		<div class="doc-button unselectable" onclick="redirectTo('swagger')">swagger</div>
		<div class="doc-button unselectable">redoc</div>
		<div class="doc-button unselectable" onclick="redirectTo('rapidoc')">rapidoc</div>
	</div>
    <redoc spec-url='{{ .SpecURL }}'></redoc>
    <script src='{{ .RedocURL }}'> </script>
  </body>
</html>
`
