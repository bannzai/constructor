package generator

import (
	"bytes"
	"html/template"

	"github.com/constructor/raw"
)

type generateComponent struct {
	Package         string
	Template        *template.Template
	SourceCodes     []raw.Code
	DestinationPath raw.Path
}

var functions = template.FuncMap{
	"upperCamelCase": upperCamelCase,
	"paramterCase":   lowerCamelCase,
	"argumentCase":   lowerCamelCase,
}

func (g generateComponent) Content() []byte {
	buf := &bytes.Buffer{}
	g.Template = g.Template.Funcs(functions)

	structs := []raw.Struct{}
	for _, sourceCode := range g.SourceCodes {
		structs = append(structs, sourceCode.Structs...)
	}

	if err := g.Template.ExecuteTemplate(buf, "constructor", map[string]interface{}{
		"Package": g.Package,
		"Structs": structs,
	}); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
