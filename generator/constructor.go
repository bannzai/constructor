package generator

import (
	"html/template"
	"path/filepath"

	"github.com/constructor/structure"
)

type Constructor struct {
	YamlReader       YamlReader
	TemplateReader   TemplateReader
	SourceCodeReader SourceCodeReader
	FileWriter       Writer
}

func (impl Constructor) Generate() {
	yaml := impl.YamlReader.Read()
	generateSources := []generateComponent{}

	for _, definition := range definitions(yaml) {
		templates := []*template.Template{}
		for _, path := range templateFilePaths(definition) {
			templates = append(templates, impl.TemplateReader.Read(path))
		}

		codes := []structure.Code{}
		for _, filePath := range sourceFilePaths(definition) {
			codes = append(codes, impl.SourceCodeReader.Read(filePath))
		}

		for _, template := range templates {
			generateSources = append(generateSources, generateComponent{
				Package:         definition.Package,
				Template:        template,
				SourceCodes:     codes,
				DestinationPath: definition.DestinationPath,
			})
		}
	}

	for _, component := range generateSources {
		impl.FileWriter.Write(component.DestinationPath, component.Content())
	}
}
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
