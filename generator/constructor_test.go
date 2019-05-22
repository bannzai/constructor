package generator

import (
	"testing"

	structure "github.com/constructor/structure"
	"github.com/golang/mock/gomock"
)

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
									Package:           "generator",
									SourcePath:        "",
									IgnoredPaths:      []structure.Path{},
									TemplateFilePaths: []structure.Path{},
									DestinationPath:   "",
								},
							},
						},
					)
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
