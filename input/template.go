package input

import "github.com/constructor/raw"

type Template interface {
	read() raw.Template
}
type TemplateImpl struct{}

func (TemplateImpl) read() raw.Template {
	return raw.Template{}
}
