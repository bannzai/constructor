package input

import (
	"fmt"
	"testing"
)

const thisFileName = "code_test.go"

func TestCodeImpl_Read(t *testing.T) {
	type fields struct {
		FilePath string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Successfully read go file.",
			fields: fields{
				FilePath: thisFileName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := CodeImpl{
				FilePath: tt.fields.FilePath,
			}
			got := impl.Read()
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
