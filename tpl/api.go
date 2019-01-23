package tpl

import (
	"bytes"
	"io"
)

type templateExecuter interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type api struct {
	t templateExecuter
}

func (p *api) include(file string, data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := p.t.ExecuteTemplate(&buf, file, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
