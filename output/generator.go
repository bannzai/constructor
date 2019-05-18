package generator

import (
	"html/template"
	"io"
	"strings"

	"github.com/constructor/raw"
)

// Generator is inter struct for generate constructor.
type Generator struct {
	Package         string
	Template        *template.Template
	SourceCode      raw.Code
	DestinationPath raw.Path

	Writer io.Writer
}

var functions = template.FuncMap{
	"upperCamelCase": upperCamelCase,
	"paramterCase":   lowerCamelCase,
	"argumentCase":   lowerCamelCase,
}

func lowerCamelCase(target string) string {
	firstString := strings.ToLower(target[:1])
	dropedFirstString := target[1:]
	return firstString + dropedFirstString
}

func upperCamelCase(target string) string {
	firstString := strings.ToUpper(target[:1])
	dropedFirstString := target[1:]
	return firstString + dropedFirstString
}

func (t Generator) Generate() {
	if err := t.Template.Funcs(functions).Execute(t.Writer, map[string]interface{}{
		"SourceFilePath": t.SourceCode.FilePath,
		"Structs":        t.SourceCode.Structs,
	}); err != nil {
		panic(err)
	}
}
