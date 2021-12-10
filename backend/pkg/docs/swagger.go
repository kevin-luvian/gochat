package docs

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Swagger(staticFilePath string) gin.HandlerFunc {
	config := SwaggerConfig{}
	config.ToDefault(staticFilePath)
	b := ReadTemplate(swagger_index_templ, config)

	return func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.Writer.Write(b)
		c.Writer.Flush()
	}
}

type SwaggerStaticFiles struct {
	DOC          string
	CSS          string
	MYDOCCSS     string
	FontsCSS     string
	BundleJS     string
	StandaloneJS string
	FavIcon16    string
}

func (s *SwaggerStaticFiles) ToDefault(basePath string) {
	s.DOC = basePath + "/swagger.json"
	s.CSS = basePath + "/swagger-ui.css"
	s.MYDOCCSS = basePath + "/my-doc-ui.css"
	s.FontsCSS = basePath + "/fonts.css"
	s.BundleJS = basePath + "/swagger-ui-bundle.js"
	s.StandaloneJS = basePath + "/swagger-ui-standalone-preset.js"
	s.FavIcon16 = basePath + "/favicon-16x16.png"
}

type SwaggerConfig struct {
	Title                    string
	DeepLinking              bool
	DocExpansion             string
	DefaultModelsExpandDepth int
	Oauth2RedirectURL        template.JS
	SwaggerStaticFiles
}

func (s *SwaggerConfig) ToDefault(basePath string) {
	s.Title = "GoChat Interactive API Documentation"
	s.DeepLinking = true
	s.DocExpansion = "list" // list, full, none
	s.DefaultModelsExpandDepth = 1
	s.Oauth2RedirectURL = template.JS(
		"`${window.location.protocol}//${window.location.host}$" +
			"{window.location.pathname.split('/').slice(0, window.location.pathname.split('/').length - 1).join('/')}" +
			"/oauth2-redirect.html`",
	)
	s.SwaggerStaticFiles.ToDefault(basePath)
}

const swagger_index_templ = `<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>{{.Title}}</title>
  <link href='{{.FontsCSS}}' rel='stylesheet' />
  <link rel='stylesheet' type='text/css' href='{{.MYDOCCSS}}' />
  <link rel='stylesheet' type='text/css' href='{{.CSS}}' />
  <link rel="icon" type="image/png" href="{{.FavIcon16}}" sizes="16x16" />
  <style>
    html
    {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
    }
    *,
    *:before,
    *:after
    {
        box-sizing: inherit;
    }

    body {
      margin:0;
      background: #fafafa;
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
  <div class="doc-button unselectable">swagger</div>
  <div class="doc-button unselectable" onclick="redirectTo('redoc')">redoc</div>
  <div class="doc-button unselectable" onclick="redirectTo('rapidoc')">rapidoc</div>
</div>
<div id="swagger-ui"></div>

<script src='{{.BundleJS}}'> </script>
<script src='{{.StandaloneJS}}'> </script>
<script>
window.onload = function() {
  // Build a system
  const ui = SwaggerUIBundle({
    url: '{{ .DOC }}',
    dom_id: '#swagger-ui',
    validatorUrl: null,
    oauth2RedirectUrl: {{.Oauth2RedirectURL}},
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
	layout: "StandaloneLayout",
    docExpansion: "{{.DocExpansion}}",
	deepLinking: {{.DeepLinking}},
	defaultModelsExpandDepth: {{.DefaultModelsExpandDepth}}
  })

  window.ui = ui
}
</script>
</body>

</html>
`
