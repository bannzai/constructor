package reader

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/constructor/structure"
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
		want structure.Code
	}{
		{
			name: "Successfully read go file.",
			args: args{
				FilePath: testdataStructPath,
			},
			want: structure.Code{
				FilePath: testdataStructPath,
				Structs: []structure.Struct{
					structure.Struct{
						Name: "Struct",
						Fields: []structure.Field{
							structure.Field{
								Name: "A",
								Type: "Alias",
							},
							structure.Field{
								Name: "F",
								Type: "func(aaa int, bbb bool) string",
							},
							structure.Field{
								Name: "L",
								Type: "[]int",
							},
							structure.Field{
								Name: "LO",
								Type: "[]io.Writer",
							},
							structure.Field{
								Name: "M",
								Type: "map[string]bool",
							},
							structure.Field{
								Name: "MKO",
								Type: "map[io.Writer]bool",
							},
							structure.Field{
								Name: "MKVO",
								Type: "map[io.Writer]io.Writer",
							},
							structure.Field{
								Name: "MVO",
								Type: "map[string]io.Writer",
							},
							structure.Field{
								Name: "O",
								Type: "io.Writer",
							},
							structure.Field{
								Name: "P",
								Type: "string",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := Code{}
			got := impl.Read(tt.args.FilePath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v,\n want %v", got, tt.want)
			}
		})
	}
}

func Test_parseASTFile(t *testing.T) {
	type args struct {
		filePath structure.Path
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
