package generator

import "github.com/bannzai/constructor/structure"

type SourceCodeReader interface {
	Read(filePath structure.Path) (code structure.Code)
	ReadWithType(filePath structure.Path, typeName string) (code structure.Code)
}
