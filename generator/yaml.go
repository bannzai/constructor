package generator

import (
	"fmt"

	"github.com/constructor/file"
	"github.com/constructor/structure"
	yaml "gopkg.in/yaml.v2"
)

type Yaml struct{}

func (Yaml) Setup() {
	if file.FileExists(structure.YamlFilePathName) {
		fmt.Println(structure.YamlFilePathName + " is already exists. Not generate " + structure.YamlFilePathName)
		return
	}
	y := structure.Yaml{
		Definitions: []structure.Definition{
			structure.Definition{
				Package:           "",
				SourcePath:        "",
				IgnoredPaths:      []structure.Path{},
				TemplateFilePaths: []structure.Path{},
				DestinationPath:   "",
			},
		},
	}

	content, err := yaml.Marshal(y)
	if err != nil {
		panic(err)
	}

	file.WriteFile(structure.YamlFilePathName, string(content))
}
