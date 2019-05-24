package reader

import (
	"fmt"
	"testing"

	"github.com/constructor/structure"
)

func TestTemplateImpl_Read(t *testing.T) {
	type args struct {
		FilePath structure.Path
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successfully read template file.",
			args: args{
				FilePath: "../constructor.tpl",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := Template{}
			got := impl.Read(tt.args.FilePath)
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
