package input

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/constructor/raw"
)

const testdataPath = "testdata/"
const testdataStructPath = testdataPath + "struct.go"

func TestCodeImpl_Read(t *testing.T) {
	type args struct {
		FilePath string
	}
	tests := []struct {
		name string
		args args
		want raw.Code
	}{
		{
			name: "Successfully read go file.",
			args: args{
				FilePath: testdataStructPath,
			},
			want: raw.Code{
				FilePath: testdataStructPath,
				Structs: []raw.Struct{
					raw.Struct{
						Name: "Struct",
						Fields: []raw.Field{
							raw.Field{
								Name: "P",
								Type: "string",
							},
							raw.Field{
								Name: "F",
								Type: "func(aaa int, bbb bool) string",
							},
							raw.Field{
								Name: "A",
								Type: "Alias",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := CodeImpl{}
			got := impl.Read(tt.args.FilePath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
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
				filePath: testdataStructPath,
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
				file: parseASTFile(testdataStructPath),
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
