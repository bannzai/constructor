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
	for typeName, structure := range parseASTStructs(parseASTFile(code.FilePath)) {
		code.Structs = append(code.Structs, convert(typeName, structure))
	}
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

func parseASTStructs(file *ast.File) (typeNameAndStruct map[string]*ast.StructType) {
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

func convert(typeName string, astStruct *ast.StructType) raw.Struct {
	typeAndNames := map[string][]string{}
	ast.Inspect(astStruct, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return false
		}

		field, ok := node.(*ast.Field)
		if !ok {
			return true
		}

		identifier, ok := field.Type.(*ast.Ident)
		if !ok {
			return true
		}
		fieldTypeName := identifier.Name
		if typeAndNames[fieldTypeName] == nil {
			typeAndNames[fieldTypeName] = []string{}
		}
		for _, nameIdentifier := range field.Names {
			name := nameIdentifier.Name
			typeAndNames[fieldTypeName] = append(typeAndNames[fieldTypeName], name)
		}
		return true
	})

	fields := []raw.Field{}
	for fieldType, names := range typeAndNames {
		for _, name := range names {
			fields = append(fields, raw.Field{
				Name: name,
				Type: fieldType,
			})
		}
	}

	return raw.Struct{
		Name:   typeName,
		Fields: fields,
	}
}
