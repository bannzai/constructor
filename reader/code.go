package reader

import (
	"github.com/constructor/structure"
)

type Code struct{}

func (impl Code) Read(filePath structure.Path) (code structure.Code) {
	code.FilePath = filePath
	for typeName, structure := range parseASTStructs(parseASTFile(code.FilePath)) {
		code.Structs = append(code.Structs, convert(typeName, structure))
	}
	code.Structs = sortedStructs(code.Structs)
	return
}
