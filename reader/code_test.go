package reader

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/bannzai/constructor/structure"
)

const testdataPath = "testdata/"
const testdataStructPath = testdataPath + "struct.go"
const testdataForMultipleStructPath = testdataPath + "struct_for_multiple_type.go"

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

func TestCode_Read(t *testing.T) {
	type args struct {
		filePath         string
		ignoreFieldNames []string
	}
	tests := []struct {
		name string
		args args
		want structure.Code
	}{
		{
			name: "Successfully read go file.",
			args: args{
				filePath:         testdataStructPath,
				ignoreFieldNames: []string{"I"},
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
							structure.Field{
								Name: "XXX",
								Type: "XXX",
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
			got := impl.Read(tt.args.filePath, tt.args.ignoreFieldNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v,\n want %v", got, tt.want)
			}
		})
	}
}

func TestCode_ReadWithType(t *testing.T) {
	type args struct {
		filePath          structure.Path
		generatedTypeName string
		ignoreFieldNames  []string
	}
	tests := []struct {
		name string
		impl Code
		args args
		want structure.Code
	}{
		{
			name: "Successfully read go file with specify type X",
			args: args{
				filePath:          testdataForMultipleStructPath,
				generatedTypeName: "X",
				ignoreFieldNames:  []string{},
			},
			want: structure.Code{
				FilePath: testdataForMultipleStructPath,
				Structs: []structure.Struct{
					structure.Struct{
						Name: "X",
						Fields: []structure.Field{
							structure.Field{
								Name: "A",
								Type: "int",
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
			if gotCode := impl.ReadWithType(tt.args.filePath, tt.args.generatedTypeName, tt.args.ignoreFieldNames); !reflect.DeepEqual(gotCode, tt.want) {
				t.Errorf("Code.ReadWithType() = %v, want %v", gotCode, tt.want)
			}
		})
	}
}
