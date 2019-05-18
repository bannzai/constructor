package output

import (
	"os"
)

type Template struct{}

const templateFileName = "constructor.tpl"
const defaultTemplate = `

{{- $dot := . -}}

{{range $i, $field := .Structs -}}

{{- $suffix := $dot.Package | upperCamelCase -}}
{{- $structureName := $dot.Name -}}

// New{{$structureName}}{{.Suffix}} insitanciate {{$dot.Name}}
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
	file, err := os.Create(templateFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(defaultTemplate))
}
