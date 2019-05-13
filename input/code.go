package input

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/constructor/raw"
)

type Code interface {
	Read(filePath raw.Path) raw.Code
}
type CodeImpl struct{}

func (impl CodeImpl) Read(filePath raw.Path) (code raw.Code) {
	code.FilePath = filePath
	code.ASTFile = parseASTFile(code.FilePath)
	code.Structs = parseASTStructs(code.ASTFile)
	return
}

func parseASTFile(filePath raw.Path) *ast.File {
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

func parseASTStructs(file *ast.File) (structs []ast.StructType) {
	ast.Inspect(file, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return false
		}
		if structType, ok := node.(*ast.StructType); ok {
			structs = append(structs, *structType)
		}
		return true
	})
	return
}
