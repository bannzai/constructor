package input

import (
	"reflect"
	"testing"

	"github.com/constructor/raw"
)

func TestYamlImpl_Read(t *testing.T) {
	type fields struct {
		argument raw.Argument
	}
	tests := []struct {
		name   string
		fields fields
		want   raw.Yaml
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := YamlImpl{
				argument: tt.fields.argument,
			}
			if got := impl.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YamlImpl.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
