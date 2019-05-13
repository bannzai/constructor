package input

import (
	"html/template"

	"github.com/constructor/raw"
)

type Template interface {
	Read(filePath raw.Path) raw.Template
}
type TemplateImpl struct{}

func (impl TemplateImpl) Read(filePath raw.Path) raw.Template {
	tpl := template.Must(template.New("input").Parse(filePath))
	return raw.Template{
		Template: *tpl,
	}
}
