package input

import (
	"fmt"
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
	}{
		{
			name: "Successfully read yaml file.",
			fields: fields{
				argument: raw.Argument{
					YamlPath: "../constructor.yaml",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := YamlImpl{
				argument: tt.fields.argument,
			}
			got := impl.Read()
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
