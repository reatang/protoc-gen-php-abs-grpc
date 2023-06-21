package demo

import (
	_ "embed"
	"text/template"
)

//go:embed template_demo.tmpl
var demoTmpl string

func Template() *template.Template {
	t := template.New("template_demo")

	// 引入一些模板函数
	// t = t.Funcs(template.FuncMap{
	//  	"some_func": nil,
	// })

	t = template.Must(t.Parse(demoTmpl))
	return t
}
