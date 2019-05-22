package generator

import (
	"path/filepath"

	"github.com/constructor/structure"
)

func definitions(yaml structure.Yaml) []structure.Definition {
	return yaml.Definitions
}

func templateFilePaths(definition structure.Definition) []structure.Path {
	return definition.TemplateFilePaths
}

func sourceFilePaths(definition structure.Definition) []structure.Path {
	filePaths, err := filepath.Glob(definition.SourcePath)
	if err != nil {
		panic(err)
	}
	return filePaths
}
