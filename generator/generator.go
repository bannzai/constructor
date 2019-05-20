package generator

import (
	"bytes"
	"html/template"
	"sort"

	"github.com/constructor/raw"
)

type generateComponent struct {
	Package         string
	Template        *template.Template
	SourceCodes     []raw.Code
	DestinationPath raw.Path
}

func sorted(structs []raw.Struct) []raw.Struct {
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

func (g generateComponent) Content() []byte {
	structs := []raw.Struct{}
	for _, sourceCode := range g.SourceCodes {
		structs = append(structs, sourceCode.Structs...)
	}

	buf := &bytes.Buffer{}
	if err := g.Template.Execute(buf, map[string]interface{}{
		"Package": g.Package,
		"Structs": sorted(structs),
	}); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
