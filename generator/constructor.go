package generator

import (
	"context"
	"html/template"
	"path/filepath"

	"github.com/constructor/file"
	"github.com/constructor/raw"
	"github.com/constructor/reader"
)

type Constructor interface {
	Generate(ctx context.Context)
}
type ConstructorImpl struct {
	YamlReader       reader.Yaml
	TemplateReader   reader.Template
	SourceCodeReader reader.Code
}

func (impl ConstructorImpl) Generate(ctx context.Context) {
	yaml := impl.YamlReader.Read()
	generateSources := []generateComponent{}

	for _, definition := range definitions(yaml) {
		templates := []*template.Template{}
		for _, path := range templateFilePaths(definition) {
			templates = append(templates, impl.TemplateReader.Read(path))
		}

		codes := []raw.Code{}
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
	}
}
func definitions(yaml raw.Yaml) []raw.Definition {
	return yaml.Definitions
}

func templateFilePaths(definition raw.Definition) []raw.Path {
	return definition.TemplateFilePaths
}

func sourceFilePaths(definition raw.Definition) []raw.Path {
	filePaths, err := filepath.Glob(definition.SourcePath)
	if err != nil {
		panic(err)
	}
	return filePaths
}
