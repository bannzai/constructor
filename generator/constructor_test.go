package generator

import (
	template "html/template"
	"testing"

	structure "github.com/bannzai/constructor/structure"
	"github.com/golang/mock/gomock"
)

const testTemplate = "package {{.Package}}\n" +
	"\n" +
	"{{range $i, $struct := .Structs -}}\n" +
	"struct {{$struct.Name}} {\n" +
	"	{{range $i, $field := $struct.Fields -}}\n" +
	"	{{$field.Name}} {{$field.Type -}}\n" +
	"	{{end}}\n" +
	"}\n" +
	"{{end -}}\n"

func TestConstructor_Generate(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		TemplateReader
		SourceCodeReader
		FileWriter
	}
	type args struct {
		sourcePath       string
		destinationPath  string
		templatePath     string
		typeName         string
		ignoreFieldNames []string
		packageName      string
	}
	tests := []struct {
		name   string
		fields fields
		args
	}{
		{
			name: "Successfully generate constructor function when not specify type name",
			fields: fields{
				TemplateReader: func() TemplateReader {
					mock := NewTemplateReaderMock(ctrl)
					mock.EXPECT().Read("template.tpl").Return(
						template.Must(template.New("template.tpl").Parse(testTemplate)),
					)
					return mock
				}(),
				SourceCodeReader: func() SourceCodeReader {
					mock := NewMockSourceCodeReader(ctrl)
					mock.EXPECT().Read("source_code.go", []string{}).Return(
						structure.Code{
							FilePath: "source_code.go",
							Structs: []structure.Struct{
								structure.Struct{
									Name: "X",
									Fields: []structure.Field{
										structure.Field{
											Name: "Field",
											Type: "int",
										},
									},
								},
								structure.Struct{
									Name: "Y",
									Fields: []structure.Field{
										structure.Field{
											Name: "Field",
											Type: "string",
										},
									},
								},
							},
						},
					)
					return mock
				}(),
				FileWriter: func() FileWriter {
					expect := "package abcd\n" +
						"\n" +
						"struct X {\n" +
						"	Field int\n" +
						"}\n" +
						"struct Y {\n" +
						"	Field string\n" +
						"}\n"

					mock := NewWriterMock(ctrl)
					mock.EXPECT().Write("destination.go", expect)
					return mock
				}(),
			},
			args: args{
				sourcePath:       "source_code.go",
				destinationPath:  "destination.go",
				templatePath:     "template.tpl",
				typeName:         "",
				ignoreFieldNames: []string{},
				packageName:      "abcd",
			},
		},
		{
			name: "Successfully generate constructor function when specify type name",
			fields: fields{
				TemplateReader: func() TemplateReader {
					mock := NewTemplateReaderMock(ctrl)
					mock.EXPECT().Read("template.tpl").Return(
						template.Must(template.New("template.tpl").Parse(testTemplate)),
					)
					return mock
				}(),
				SourceCodeReader: func() SourceCodeReader {
					mock := NewMockSourceCodeReader(ctrl)
					mock.EXPECT().ReadWithType("source_code.go", "X", []string{}).Return(
						structure.Code{
							FilePath: "source_code.go",
							Structs: []structure.Struct{
								structure.Struct{
									Name: "X",
									Fields: []structure.Field{
										structure.Field{
											Name: "Field",
											Type: "int",
										},
									},
								},
							},
						},
					)
					return mock
				}(),
				FileWriter: func() FileWriter {
					expect := "package abcd\n" +
						"\n" +
						"struct X {\n" +
						"	Field int\n" +
						"}\n"

					mock := NewWriterMock(ctrl)
					mock.EXPECT().Write("destination.go", expect)
					return mock
				}(),
			},
			args: args{
				sourcePath:       "source_code.go",
				destinationPath:  "destination.go",
				templatePath:     "template.tpl",
				typeName:         "X",
				ignoreFieldNames: []string{},
				packageName:      "abcd",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := Constructor{
				TemplateReader:   tt.fields.TemplateReader,
				SourceCodeReader: tt.fields.SourceCodeReader,
				FileWriter:       tt.fields.FileWriter,
			}
			impl.Generate(tt.args.sourcePath, tt.args.destinationPath, tt.args.templatePath, tt.args.typeName, tt.args.ignoreFieldNames, tt.args.packageName)
		})
	}
}
