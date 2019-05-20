package generator

import (
	"fmt"

	"github.com/constructor/file"
	"github.com/constructor/raw"
	yaml "gopkg.in/yaml.v2"
)

type Yaml struct{}

func (Yaml) Setup() {
	if file.FileExists(raw.YamlFilePathName) {
		fmt.Println(raw.YamlFilePathName + " is already exists. Not generate " + raw.YamlFilePathName)
		return
	}
	y := raw.Yaml{
		Definitions: []raw.Definition{
			raw.Definition{
				Package:           "",
				SourcePath:        "",
				IgnoredPaths:      []raw.Path{},
				TemplateFilePaths: []raw.Path{},
				DestinationPath:   "",
			},
		},
	}

	generator, err := yaml.Marshal(y)
	if err != nil {
		panic(err)
	}

	file.WriteFile(raw.YamlFilePathName, generator)
}
