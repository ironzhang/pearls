package tpl

import (
	"text/template"
)

func ParseFiles(filenames ...string) (*template.Template, error) {
	t := template.New("")
	t.Funcs(makeFuncMap(t))
	return t.ParseFiles(filenames...)
}

func ParseGlob(pattern string) (*template.Template, error) {
	t := template.New("")
	t.Funcs(makeFuncMap(t))
	return t.ParseGlob(pattern)
}

func makeFuncMap(t *template.Template) template.FuncMap {
	a := api{t: t}
	m := template.FuncMap{
		"include": a.include,
	}
	return m
}
