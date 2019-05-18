package reader

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

		switch types := field.Type.(type) {
		case *ast.Ident:
			fieldTypeName := types.Name
			for _, nameIdentifier := range field.Names {
				name := nameIdentifier.Name
				typeAndNames[fieldTypeName] = append(typeAndNames[fieldTypeName], name)
			}
		case *ast.FuncType:
			statement := "func("
			for i, param := range types.Params.List {
				parameterType := param.Type.(*ast.Ident).Name
				parameterNames := param.Names[0:len(param.Names)]
				if i != 0 {
					statement += ", "
				}
				for i, parameterName := range parameterNames {
					if i == 0 {
						statement += parameterName.Name
					} else {
						statement += ", " + parameterName.Name
					}
				}
				statement += " " + parameterType
			}
			statement += ")"

			results := types.Results.List
			if len(results) > 1 {
				statement += "("
			}
			for i, result := range types.Results.List {
				resultType := result.Type.(*ast.Ident).Name
				resultNames := result.Names[0:len(result.Names)]
				if i != 0 {
					statement += ", "
				}
				for i, resultName := range resultNames {
					if i == 0 {
						statement += resultName.Name
					} else {
						statement += ", " + resultName.Name
					}
				}
				statement += " " + resultType
			}
			if len(results) > 1 {
				statement += ")"
			}
			fieldName := field.Names[0].Name
			typeAndNames[statement] = append(typeAndNames[statement], fieldName)
			// No continue next child
			// Because ast.FuncType has *ast.Field node.
			// It will duplicate call function
			return false
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