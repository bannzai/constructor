package reader

import (
	"github.com/bannzai/constructor/structure"
)

type Code struct{}

func (impl Code) Read(filePath structure.Path, ignoreFieldNames []string) structure.Code {
	return impl.ReadWithType(filePath, "", ignoreFieldNames)
}

func (impl Code) ReadWithType(filePath structure.Path, generatedTypeName string, ignoreFieldNames []string) (code structure.Code) {
	code.FilePath = filePath
	for typeName, structure := range parseASTStructs(parseASTFile(code.FilePath)) {
		code.Structs = append(code.Structs, convert(typeName, ignoreFieldNames, structure))
	}
	code.Structs = sortedStructs(code.Structs)
	return
}
