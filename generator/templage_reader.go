package generator

import (
	"html/template"

	"github.com/constructor/structure"
)

type TemplateReader interface {
	Read(filePath structure.Path) *template.Template
}
