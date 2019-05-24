package generator

import "github.com/constructor/structure"

type SourceCodeReader interface {
	Read(filePath structure.Path) (code structure.Code)
}
