package generator

import (
	"github.com/bannzai/constructor/structure"
)

func definitions(yaml structure.Yaml) []structure.Definition {
	return yaml.Definitions
}

func templateFilePaths(definition structure.Definition) []structure.Path {
	return definition.TemplateFilePaths
}
