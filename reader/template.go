package reader

import (
	"html/template"
	"strings"

	"github.com/constructor/raw"
)

type Template interface {
	Read(filePath raw.Path) *template.Template
}
type TemplateImpl struct{}

var functions = template.FuncMap{
	"upperCamelCase":     upperCamelCase,
	"parameterCase":      lowerCamelCase,
	"argumentCase":       lowerCamelCase,
	"escapeReservedWord": escapeReservedWord,
}

func escapeReservedWord(target string) string {
	for _, reservedWord := range raw.ReservedWords {
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

func (impl TemplateImpl) Read(filePath raw.Path) *template.Template {
	return template.Must(template.New(filePath).Funcs(functions).ParseFiles(filePath))
}
