package generator

import "github.com/bannzai/constructor/structure"

type SourceCodeReader interface {
	Read(filePath structure.Path, ignoreFieldNames []string) structure.Code
	ReadWithType(filePath structure.Path, generatedTypeName string, ignoreFieldNames []string) (code structure.Code)
}
