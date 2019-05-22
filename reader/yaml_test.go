package reader

import (
	"fmt"
	"testing"

	"github.com/constructor/structure"
)

func TestYamlImpl_Read(t *testing.T) {
	type fields struct {
		argument structure.Argument
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Successfully read yaml file.",
			fields: fields{
				argument: structure.Argument{
					YamlPath: "../constructor.yaml",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := YamlImpl{
				Argument: tt.fields.argument,
			}
			got := impl.Read()
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
