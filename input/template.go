package input

import "github.com/constructor/value"

type Template interface {
	read() value.Template
}
type TemplateImpl struct{}

func (TemplateImpl) read() value.Template {
	return value.Template{}
}
