package input

import (
	"fmt"
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
	}{
		{
			name: "Successfully read template file.",
			fields: fields{
				FilePath: "../template_test.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := TemplateImpl{
				FilePath: tt.fields.FilePath,
			}
			got := impl.Read()
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
