package generator

import (
	"html/template"

	"github.com/bannzai/constructor/structure"
)

type Constructor struct {
	YamlReader       YamlReader
	TemplateReader   TemplateReader
	SourceCodeReader SourceCodeReader
	FileWriter       Writer
	FilePathFetcher
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
		for _, filePath := range impl.FilePathFetcher.sourceFilePaths(definition) {
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
