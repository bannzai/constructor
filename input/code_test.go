package input

import (
	"fmt"
	"go/ast"
	"testing"

	"github.com/constructor/raw"
)

const thisFileName = "code_test.go"
const testdataPath = "testdata/"

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
		name    string
		args    args
		wantNil bool
	}{
		{
			name: "Successfully parse AST file.",
			args: args{
				filePath: testdataPath + "struct.go",
			},
			wantNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseASTFile(tt.args.filePath); (got == nil) != tt.wantNil {
				t.Errorf("parseASTFile() = %v, want %v", got, tt.wantNil)
			}
		})
	}
}

func Test_parseASTStructs(t *testing.T) {
	type args struct {
		file *ast.File
	}
	tests := []struct {
		name             string
		args             args
		wantStructsCount int
	}{
		// reference: https://play.golang.org/p/BMcvmVmSgtM
		{
			name: "Successfully parse AST file.",
			args: args{
				file: parseASTFile(testdataPath + "struct.go"),
			},
			wantStructsCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotStructs := parseASTStructs(tt.args.file); len(gotStructs) != tt.wantStructsCount {
				t.Errorf("parseASTStructs() = %v, wantStructsCount %v", gotStructs, tt.wantStructsCount)
			}
		})
	}
}
