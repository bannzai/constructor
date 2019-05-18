package output

import (
	"os"

	"github.com/constructor/raw"
	yaml "gopkg.in/yaml.v2"
)

type Yaml struct{}

const yamlFilePathName = "constructor.yaml"

func (Yaml) Setup() {
	y := raw.Yaml{
		Definitions: []raw.Definition{
			raw.Definition{
				Package:           "",
				SourcePaths:       []raw.Path{},
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

	file, err := os.Create(yamlFilePathName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(output)
}
