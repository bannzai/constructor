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

func (g generateComponent) Content() []byte {
	structs := []raw.Struct{}
	for _, sourceCode := range g.SourceCodes {
		structs = append(structs, sourceCode.Structs...)
	}

	buf := &bytes.Buffer{}
	if err := g.Template.Execute(buf, map[string]interface{}{
		"Package": g.Package,
		"Structs": structs,
	}); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
