package input

import (
	"html/template"

	"github.com/constructor/raw"
)

type Template interface {
	Read() raw.Template
}
type TemplateImpl struct {
	FilePath raw.Path
}

func (impl TemplateImpl) Read() raw.Template {
	tpl := template.Must(template.New("input").Parse(impl.FilePath))
	return raw.Template{
		Template: *tpl,
	}
}
