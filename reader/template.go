package reader

import (
	"html/template"
	"strings"

	"github.com/constructor/structure"
)

type Template interface {
	Read(filePath structure.Path) *template.Template
}
type TemplateImpl struct{}

var functions = template.FuncMap{
	"upperCamelCase":     upperCamelCase,
	"parameterCase":      lowerCamelCase,
	"argumentCase":       lowerCamelCase,
	"escapeReservedWord": escapeReservedWord,
}

func escapeReservedWord(target string) string {
	for _, reservedWord := range structure.ReservedWords {
		if reservedWord == target {
			return "_" + target
		}
	}

	return target
}

func lowerCamelCase(target string) string {
	firstString := strings.ToLower(target[:1])
	dropedFirstString := target[1:]
	return escapeReservedWord(firstString + dropedFirstString)
}

func upperCamelCase(target string) string {
	firstString := strings.ToUpper(target[:1])
	dropedFirstString := target[1:]
	return escapeReservedWord(firstString + dropedFirstString)
}

func (impl TemplateImpl) Read(filePath structure.Path) *template.Template {
	return template.Must(template.New(filePath).Funcs(functions).ParseFiles(filePath))
}
