package generator

import "testing"

func TestConstructor_Generate(t *testing.T) {
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
		// TODO: Add test cases.
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
