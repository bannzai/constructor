package generator

import (
	"bytes"
	"html/template"

	"github.com/constructor/raw"
)

type generateComponent struct {
	Package         string
	Template        *template.Template
	SourceCode      raw.Code
	DestinationPath raw.Path
}

var functions = template.FuncMap{
	"upperCamelCase": upperCamelCase,
	"paramterCase":   lowerCamelCase,
	"argumentCase":   lowerCamelCase,
}

func (t generateComponent) Content() []byte {
	buf := &bytes.Buffer{}
	if err := t.Template.Funcs(functions).Execute(buf, map[string]interface{}{
		"Package":        t.Package,
		"SourceFilePath": t.SourceCode.SourcePath,
		"Structs":        t.SourceCode.Structs,
	}); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
