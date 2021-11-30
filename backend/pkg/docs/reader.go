package docs

import (
	"bytes"
	"html/template"
)

func ReadTemplate(tmplstr string, data interface{}) []byte {
	tmpl := template.Must(template.New("template").Parse(tmplstr))

	buf := bytes.NewBuffer(nil)
	_ = tmpl.Execute(buf, data)
	return buf.Bytes()
}
