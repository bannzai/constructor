package input

import (
	"reflect"
	"testing"

	"github.com/constructor/raw"
)

func TestCodeImpl_Read(t *testing.T) {
	type fields struct {
		FilePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   raw.Code
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := CodeImpl{
				FilePath: tt.fields.FilePath,
			}
			if got := impl.Read(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CodeImpl.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
