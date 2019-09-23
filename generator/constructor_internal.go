package generator

import (
	"github.com/bannzai/constructor/structure"
)

func templateFilePaths(definition structure.Definition) []structure.Path {
	return definition.TemplateFilePaths
}
