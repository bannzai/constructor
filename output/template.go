package output

type Template interface {
	Setup() error
}
type TemplateImpl struct{}

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
	return nil
}
