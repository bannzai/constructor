package input

import (
	"fmt"
	"testing"
)

const thisFileName = "code_test.go"

func TestCodeImpl_Read(t *testing.T) {
	type args struct {
		FilePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successfully read go file.",
			args: args{
				FilePath: thisFileName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := CodeImpl{}
			got := impl.Read(tt.args.FilePath)
			fmt.Printf("Successfully got: %v ", got)
		})
	}
}
