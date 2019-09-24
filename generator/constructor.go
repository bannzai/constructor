package generator

import (
	"bytes"
	"html/template"

	"github.com/bannzai/constructor/structure"
)

type Constructor struct {
	TemplateReader
	SourceCodeReader
	FileWriter
}

func (generator Constructor) Generate() {

}

func (generator Constructor) content(
	packageName string,
	template *template.Template,
	sourceCodes []structure.Code,
	destinationPath structure.Path,
) string {
	structs := []structure.Struct{}
	for _, sourceCode := range sourceCodes {
		structs = append(structs, sourceCode.Structs...)
	}

	buf := &bytes.Buffer{}
	if err := template.Execute(buf, map[string]interface{}{
		"Package": packageName,
		"Structs": structs,
	}); err != nil {
		panic(err)
	}
	return buf.String()
}
