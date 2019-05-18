package input

import (
	"html/template"

	"github.com/constructor/raw"
)

type Template interface {
	Read(filePath raw.Path) *template.Template
}
type TemplateImpl struct{}

func (impl TemplateImpl) Read(filePath raw.Path) *template.Template {
	return template.Must(template.New("input").Parse(filePath))
}
