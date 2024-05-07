package php

import (
	_ "embed"
	"text/template"
)

//go:embed template_gateway.gotmpl
var grpcGatewayTmpl string

// GrpcGatewayTemplate gets the templates to for the typescript file
func GrpcGatewayTemplate() *template.Template {
	t := template.New("template_gateway")

	t = t.Funcs(template.FuncMap{
		"scalaType": mapScalaType,
	})

	t = template.Must(t.Parse(grpcGatewayTmpl))
	return t
}
