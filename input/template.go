package input

import "github.com/constructor/raw"

type Template interface {
	Read() raw.Template
}
type TemplateImpl struct{}

func (TemplateImpl) Read() raw.Template {
	return raw.Template{}
}
