package reader

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/bannzai/constructor/structure"
)

func parseASTFile(filePath structure.Path) *ast.File {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, filePath, buf, 0)
	if err != nil {
		panic(err)
	}
	return astFile
}

func parseASTStructs(file *ast.File) (typeNameAndStruct map[string]*ast.StructType) {
	typeNameAndStruct = map[string]*ast.StructType{}
	ast.Inspect(file, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return false
		}

		typeSpec, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		name := typeSpec.Name.Name
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		typeNameAndStruct[name] = structType
		return true
	})
	return
}
