package input

import (
	"fmt"
	"testing"

	"github.com/constructor/raw"
)

func TestTemplateImpl_Read(t *testing.T) {
	type args struct {
		FilePath raw.Path
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successfully read template file.",
			args: args{
				FilePath: "../template_test.go",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := TemplateImpl{}
			got := impl.Read(tt.args.FilePath)
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
