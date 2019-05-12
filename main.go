package main

import (
	"fmt"

	"github.com/constructor/cmd"
	"github.com/constructor/config"
	"github.com/constructor/input"
	"github.com/constructor/model"
	"github.com/constructor/output"
	"github.com/constructor/raw"
)

const YamlFilePathName = "constructor.yaml"

func main() {
	config.Configuration.YamlFilePath = YamlFilePathName
	yaml := input.YamlImpl{
		Argument: raw.Argument{
			YamlPath: configurationFilePath(),
		},
	}.Read()

	generateSources := []model.GenerateElementEachPackage{}
	for _, definition := range definitions(yaml) {
		templates := []raw.Template{}
		for _, path := range templateFilePaths(definition) {
			templates = append(templates, input.TemplateImpl{
				FilePath: path,
			}.Read())
		}

		codes := []raw.Code{}
		for _, path := range sourceCodeFilePaths(definition) {
			codes = append(codes, input.CodeImpl{
				FilePath: path,
			}.Read())
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
		output.ConstructorImpl{source}.Generate()
	}

	fmt.Printf("yaml = %+v\n", yaml)
	cmd.Execute()
}

func definitions(yaml raw.Yaml) []raw.Definition {
	return []raw.Definition{}
}

func configurationFilePath() raw.Path {
	return "./" + YamlFilePathName
}

func templateFilePaths(defition raw.Definition) []raw.Path {
	return []raw.Path{}
}

func sourceCodeFilePaths(defition raw.Definition) []raw.Path {
	return []raw.Path{}
}
