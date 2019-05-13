package output

import (
	"context"
	"io/ioutil"

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

		codes := []raw.Code{}
		for _, path := range sourceCodeFilePaths(definition) {
			codes = append(codes, input.CodeImpl{}.Read(path))
		}

		for _, template := range templates {
			generateSources = append(generateSources, model.GenerateElementEachPackage{
				Package:         definition.Package,
				Template:        template,
				Codes:           codes,
				DestinationPath: definition.DestinationPath,
			})
		}
	}

	for _, source := range generateSources {
		content := source.Content()
		if err := ioutil.WriteFile(source.DestinationPath, content, 0644); err != nil {
			panic(err)
		}
	}
}
func definitions(yaml raw.Yaml) []raw.Definition {
	return []raw.Definition{}
}

func templateFilePaths(defition raw.Definition) []raw.Path {
	return []raw.Path{}
}

func sourceCodeFilePaths(defition raw.Definition) []raw.Path {
	return []raw.Path{}
}
