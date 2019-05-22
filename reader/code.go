package reader

import (
	"sort"

	"github.com/constructor/structure"
)

type Code struct{}

func sortedStructs(structs []structure.Struct) []structure.Struct {
	sort.SliceStable(structs, func(l, r int) bool {
		return sort.StringsAreSorted(
			[]string{
				structs[l].Name,
				structs[r].Name,
			},
		)
	})
	return structs
}

func sortedFields(fields []structure.Field) []structure.Field {
	sort.SliceStable(fields, func(l, r int) bool {
		return sort.StringsAreSorted(
			[]string{
				fields[l].Name,
				fields[r].Name,
			},
		)
	})
	return fields
}

func (impl Code) Read(filePath structure.Path) (code structure.Code) {
	code.FilePath = filePath
	for typeName, structure := range parseASTStructs(parseASTFile(code.FilePath)) {
		code.Structs = append(code.Structs, convert(typeName, structure))
	}
	code.Structs = sortedStructs(code.Structs)
	return
}
