package reader

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/constructor/structure"
)

type TypeAndNames = map[string][]string

func convert(typeName string, astStruct *ast.StructType) structure.Struct {
	typeAndNames := TypeAndNames{}
	ast.Inspect(astStruct, func(node ast.Node) bool {
		lastChildNode := node == nil
		if lastChildNode {
			return false
		}

		field, ok := node.(*ast.Field)
		if !ok {
			return true
		}

		if hasIgnoreTag(field) {
			return true
		}

		switch types := field.Type.(type) {
		case *ast.Ident:
			fieldTypeName := types.Name
			for _, nameIdentifier := range field.Names {
				name := nameIdentifier.Name
				typeAndNames[fieldTypeName] = append(typeAndNames[fieldTypeName], name)
			}
		case *ast.ArrayType:
			var fieldTypeName string
			if selector, ok := types.Elt.(*ast.SelectorExpr); ok {
				x, sel := parseSelectorExpr(selector)
				fieldTypeName = "[]" + x + "." + sel
			}
			if ident, ok := types.Elt.(*ast.Ident); ok {
				fieldTypeName = "[]" + ident.Name
			}
			if len(fieldTypeName) == 0 {
				panic(fmt.Errorf("Unknown pattern when ast.ArrayType.Elt receive %v", reflect.TypeOf(types.Elt)))
			}
			for _, nameIdentifier := range field.Names {
				name := nameIdentifier.Name
				typeAndNames[fieldTypeName] = append(typeAndNames[fieldTypeName], name)
			}
		case *ast.MapType:
			var key string
			var value string
			if k, ok := types.Key.(*ast.Ident); ok {
				key = k.Name
			}
			if v, ok := types.Value.(*ast.Ident); ok {
				value = v.Name
			}
			if k, ok := types.Key.(*ast.SelectorExpr); ok {
				x, sel := parseSelectorExpr(k)
				key = x + "." + sel
			}
			if v, ok := types.Value.(*ast.SelectorExpr); ok {
				x, sel := parseSelectorExpr(v)
				value = x + "." + sel
			}

			if len(key) == 0 {
				panic(fmt.Errorf("Unknown pattern when ast.MapType.Key receive %v", reflect.TypeOf(types.Key)))
			}
			if len(value) == 0 {
				panic(fmt.Errorf("Unknown pattern when ast.MapType.Value receive %v", reflect.TypeOf(types.Value)))
			}

			fieldTypeName := "map[" + key + "]" + value
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
		case *ast.SelectorExpr:
			x, sel := parseSelectorExpr(types)
			fieldTypeName := x + "." + sel
			for _, nameIdentifier := range field.Names {
				name := nameIdentifier.Name
				typeAndNames[fieldTypeName] = append(typeAndNames[fieldTypeName], name)
			}
		}
		return true
	})

	fields := []structure.Field{}
	for fieldType, names := range typeAndNames {
		for _, name := range names {
			fields = append(fields, structure.Field{
				Name: name,
				Type: fieldType,
			})
		}
	}

	fields = sortedFields(fields)

	return structure.Struct{
		Name:   typeName,
		Fields: fields,
	}
}

func hasIgnoreTag(field *ast.Field) bool {
	if field.Tag == nil {
		return false
	}

	separator := ":"
	annotation := structure.IgnoreCaseKeyword + separator
	if !strings.Contains(field.Tag.Value, annotation) {
		return false
	}

	head := "`" + annotation
	return field.Tag.Value[len(head):len(head)+len("true")] == "true" // FIXME: Good code
}

func parseSelectorExpr(types *ast.SelectorExpr) (string, string) {
	x, ok := types.X.(*ast.Ident)
	if !ok {
		panic(fmt.Errorf("Unknown pattern when ast.SelectorExpr.X receive %v", reflect.TypeOf(types.X)))
	}
	return x.Name, types.Sel.Name
}
