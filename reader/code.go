package reader

import (
	"github.com/bannzai/constructor/structure"
)

type Code struct{}

func (impl Code) Read(filePath structure.Path) structure.Code {
	return impl.ReadWithType(filePath, "")
}

func (impl Code) ReadWithType(filePath structure.Path, generatedTypeName string) (code structure.Code) {
	code.FilePath = filePath
	isNotSpecifyType := 0 == len(generatedTypeName)
	for typeName, structure := range parseASTStructs(parseASTFile(code.FilePath)) {
		if isNotSpecifyType || typeName != generatedTypeName {
			continue
		}
		code.Structs = append(code.Structs, convert(typeName, structure))
	}
	code.Structs = sortedStructs(code.Structs)
	return
}
