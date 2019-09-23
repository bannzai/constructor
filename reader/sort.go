package reader

import (
	"sort"

	"github.com/bannzai/constructor/structure"
)

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
