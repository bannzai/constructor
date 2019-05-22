package generator

import (
	"bytes"
	"html/template"

	"github.com/constructor/structure"
)

type generateComponent struct {
	Package         string
	Template        *template.Template
	SourceCodes     []structure.Code
	DestinationPath structure.Path
}

func (g generateComponent) Content() string {
	structs := []structure.Struct{}
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
	return buf.String()
}
