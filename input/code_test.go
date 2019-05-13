package input

import (
	"fmt"
	"go/ast"
	"reflect"
	"testing"

	"github.com/constructor/raw"
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

func Test_parseASTFile(t *testing.T) {
	type args struct {
		filePath raw.Path
	}
	tests := []struct {
		name string
		args args
		want *ast.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTFile(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseASTFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseASTStructs(t *testing.T) {
	type args struct {
		file *ast.File
	}
	tests := []struct {
		name        string
		args        args
		wantStructs []ast.StructType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStructs := parseASTStructs(tt.args.file); !reflect.DeepEqual(gotStructs, tt.wantStructs) {
				t.Errorf("parseASTStructs() = %v, want %v", gotStructs, tt.wantStructs)
			}
		})
	}
}
