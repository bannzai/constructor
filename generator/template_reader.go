package generator

import (
	"html/template"

	"github.com/bannzai/constructor/structure"
)

type TemplateReader interface {
	Read(filePath structure.Path) *template.Template
}
