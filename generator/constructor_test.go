package generator

import (
	"reflect"
	"testing"

	"github.com/constructor/structure"
)

func TestConstructor_Generate(t *testing.T) {
	type fields struct {
		YamlReader       YamlReader
		TemplateReader   TemplateReader
		SourceCodeReader SourceCodeReader
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
			}
			impl.Generate()
		})
	}
}

func Test_definitions(t *testing.T) {
	type args struct {
		yaml structure.Yaml
	}
	tests := []struct {
		name string
		args args
		want []structure.Definition
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := definitions(tt.args.yaml); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("definitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_templateFilePaths(t *testing.T) {
	type args struct {
		definition structure.Definition
	}
	tests := []struct {
		name string
		args args
		want []structure.Path
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := templateFilePaths(tt.args.definition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateFilePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sourceFilePaths(t *testing.T) {
	type args struct {
		definition structure.Definition
	}
	tests := []struct {
		name string
		args args
		want []structure.Path
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sourceFilePaths(tt.args.definition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sourceFilePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
