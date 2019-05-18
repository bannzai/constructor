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
	t := g.Template.Funcs(functions)

	for _, sourceCode := range g.SourceCodes {
		if err := t.Execute(buf, map[string]interface{}{
			"Package":        g.Package,
			"SourceFilePath": sourceCode.FilePath,
			"Structs":        sourceCode.Structs,
		}); err != nil {
			panic(err)
		}
	}
	return buf.Bytes()
}
