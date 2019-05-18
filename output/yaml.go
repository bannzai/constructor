package output

import (
	"fmt"

	"github.com/constructor/raw"
	yaml "gopkg.in/yaml.v2"
)

type Yaml struct{}

const yamlFilePathName = "constructor.yaml"

func (Yaml) Setup() {
	if fileExists(yamlFilePathName) {
		fmt.Println(yamlFilePathName + " is already exists. Not generate " + yamlFilePathName)
		return
	}
	y := raw.Yaml{
		Definitions: []raw.Definition{
			raw.Definition{
				Package:           "",
				PackagePath:       "",
				IgnoredPaths:      []raw.Path{},
				TemplateFilePaths: []raw.Path{},
				DestinationPath:   "",
			},
		},
	}

	output, err := yaml.Marshal(y)
	if err != nil {
		panic(err)
	}

	writeFile(yamlFilePathName, output)
}
