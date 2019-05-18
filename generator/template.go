package generator

import (
	"fmt"

	"github.com/constructor/file"
)

type Template struct{}

const templateFileName = "constructor.tpl"
const defaultTemplate = ` // DO NOT EDIT THIS FILE.
// File generated by constructor.
// https://github.com/bannzai/constructor

{{- $dot := . -}}
package {{.Package}}

{{range $i, $field := .Structs -}}

{{- $suffix := $dot.Package | upperCamelCase -}}
{{- $structureName := $dot.Name -}}

// New{{$structureName}}{{$suffix}} insitanciate {{$dot.Name}}
func New{{$structureName}}{{$suffix}}(
	{{range $i, $field := $dot.Fields -}}
		{{parameterCase $field}} {{$field}},
	{{end}}
) {{$structureName}} {
	return {{$structureName}}{{$suffix}}{
		{{range $i, $field := $dot.Fields -}}
			{{$field}}: {{argumentCase $field}},
		{{end}}
	}
}
{{end}}
`

func (impl Template) Setup() {
	if file.FileExists(templateFileName) {
		fmt.Println(templateFileName + " is already exists. Not generate " + templateFileName)
		return
	}
	file.WriteFile(templateFileName, []byte(defaultTemplate))
}
