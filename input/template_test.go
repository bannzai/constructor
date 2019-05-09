package input

import (
	"reflect"
	"testing"

	"github.com/constructor/raw"
)

func TestTemplateImpl_Read(t *testing.T) {
	type fields struct {
		FilePath raw.Path
	}
	tests := []struct {
		name   string
		fields fields
		want   raw.Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := TemplateImpl{
				FilePath: tt.fields.FilePath,
			}
			if got := impl.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TemplateImpl.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
