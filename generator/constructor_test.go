package generator

import (
	template "html/template"
	"testing"

	structure "github.com/constructor/structure"
	"github.com/golang/mock/gomock"
)

const testTemplate = `
package {{.Package}}

struct A{
{{range $i, $struct := .Structs -}}
	{{$struct.Name}}{{$i}}
}
{{end}}
`

func TestConstructor_Generate(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		YamlReader       YamlReader
		TemplateReader   TemplateReader
		SourceCodeReader SourceCodeReader
		FileWriter       Writer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Successfully generate constructor function",
			fields: fields{
				YamlReader: func() YamlReader {
					mock := NewYamlReaderMock(ctrl)
					mock.EXPECT().Read().Return(
						structure.Yaml{
							Definitions: []structure.Definition{
								structure.Definition{
									Package:           "abcd",
									SourcePath:        "source_code.go",
									IgnoredPaths:      []structure.Path{},
									TemplateFilePaths: []structure.Path{"template.tpl"},
									DestinationPath:   "destination.go",
								},
							},
						},
					)
					return mock
				}(),
				TemplateReader: func() TemplateReader {
					mock := NewTemplateReaderMock(ctrl)
					mock.EXPECT().Read("template.tpl").Return(
						template.Must(template.New("template.tpl").Parse(testTemplate)),
					)
					return mock
				}(),
				SourceCodeReader: func() SourceCodeReader {
					mock := NewSourceCodeReaderMock(ctrl)
					mock.EXPECT().Read("source_code.go").Return(
						structure.Code{
							FilePath: "source_code.go",
							Structs: []structure.Struct{
								structure.Struct{
									Name: "X",
								},
								structure.Struct{
									Name: "Y",
								},
							},
						},
					)
					return mock
				}(),
				FileWriter: func() Writer {
					expect := []byte(`
package abcd
struct A{
	X1
	Y2
}
						`)
					mock := NewWriterMock(ctrl)
					mock.EXPECT().Write("destination.go", expect).Return()
					return mock
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := Constructor{
				YamlReader:       tt.fields.YamlReader,
				TemplateReader:   tt.fields.TemplateReader,
				SourceCodeReader: tt.fields.SourceCodeReader,
				FileWriter:       tt.fields.FileWriter,
			}
			impl.Generate()
		})
	}
}
