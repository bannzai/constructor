package output

import (
	"os"

	"github.com/pkg/errors"
)

type Template interface {
	Setup() error
}
type TemplateImpl struct{}

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

func (impl TemplateImpl) Setup() error {
	file, err := os.Create(templateFileName)
	if err != nil {
		return errors.Wrap(err, "Can not create "+templateFileName)
	}
	defer file.Close()
	file.Write([]byte(defaultTemplate))
	return nil
}
