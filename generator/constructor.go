package generator

import (
	"context"
	"html/template"
	"path/filepath"

	"github.com/constructor/file"
	"github.com/constructor/reader"
	"github.com/constructor/structure"
)

type Constructor struct {
	YamlReader       YamlReader
	TemplateReader   reader.Template
	SourceCodeReader reader.Code
}

func (impl Constructor) Generate(ctx context.Context) {
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
		file.WriteFile(component.DestinationPath, component.Content())
		file.GoImports(component.DestinationPath)
		file.GoFormat(component.DestinationPath)
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
