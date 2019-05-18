package model

import (
	"html/template"
	"io"
	"strings"

	"github.com/constructor/raw"
)

// GenerateElementEachPackage is inter struct for generate constructor.
type GenerateElementEachPackage struct {
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

func (t GenerateElementEachPackage) Content() []byte {
	if err := t.Template.Funcs(functions).Execute(t.Writer, map[string]interface{}{
		"SourceFilePath": t.SourceCode.FilePath,
		"Structs":        t.SourceCode.Structs,
	}); err != nil {
		panic(err)
	}

	return []byte{}
}
