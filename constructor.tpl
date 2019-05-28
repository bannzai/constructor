// DO NOT EDIT THIS FILE.
// File generated by constructor.
// https://github.com/bannzai/constructor

{{- $dot := . }}
package {{.Package}}

{{range $i, $struct := .Structs -}}
{{- $suffix := upperCamelCase $dot.Package -}}
{{- $structureName := .Name -}}
// New{{$structureName}}{{$suffix}} insitanciate {{.Name}}
func New{{$structureName}}{{$suffix}}(
	{{range $i, $field := .Fields -}}
		{{parameterCase $field.Name}} {{$field.Type}},
	{{end -}}
) {{$structureName}} {
	return {{$structureName}}{
		{{range $i, $field := .Fields -}}
			{{$field.Name}}: {{argumentCase $field.Name}},
		{{end -}}
	}
}
{{end}}
