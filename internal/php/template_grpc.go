package php

import (
	_ "embed"
	"text/template"
)

//go:embed template_grpc.gotmpl
var serviceTmpl string

// GrpcTemplate gets the templates to for the typescript file
func GrpcTemplate() *template.Template {
	t := template.New("template_grpc")

	t = t.Funcs(template.FuncMap{
		"scalaType": mapScalaType,
	})

	t = template.Must(t.Parse(serviceTmpl))
	return t
}
