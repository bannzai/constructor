package output

import (
	"context"

	"github.com/constructor/input"
	"github.com/constructor/model"
	"github.com/constructor/raw"
)

type Constructor interface {
	Generate(ctx context.Context)
}
type ConstructorImpl struct {
	YamlReader       input.Yaml
	TemplateReader   input.Template
	SourceCodeReader input.Code
}

func (impl ConstructorImpl) Generate(ctx context.Context) {
	yaml := impl.YamlReader.Read()

	generateSources := []model.GenerateElementEachPackage{}
	for _, definition := range definitions(yaml) {
		templates := []raw.Template{}
		for _, path := range templateFilePaths(definition) {
			templates = append(templates, impl.TemplateReader.Read(path))
		}

		code := input.CodeImpl{}.Read(packagePath(definition))

		for _, template := range templates {
			generateSources = append(generateSources, model.GenerateElementEachPackage{
				Package:         definition.Package,
				Template:        template,
				SourceCode:      code,
				DestinationPath: definition.DestinationPath,
			})
		}
	}

	for _, source := range generateSources {
		content := source.Content()
		writeFile(source.DestinationPath, content)
	}
}
func definitions(yaml raw.Yaml) []raw.Definition {
	return yaml.Definitions
}

func templateFilePaths(definition raw.Definition) []raw.Path {
	return definition.TemplateFilePaths
}

func packagePath(definition raw.Definition) raw.Path {
	return definition.PackagePath
}
