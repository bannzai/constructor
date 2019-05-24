package reader

import (
	"io/ioutil"

	"github.com/constructor/structure"
	"gopkg.in/yaml.v2"
)

// Yaml for implement for Yaml
type Yaml struct {
	Argument structure.Argument
}

// Read for yaml file
func (impl Yaml) Read() structure.Yaml {
	buf, err := ioutil.ReadFile(impl.Argument.YamlPath)
	if err != nil {
		panic(err)
	}

	var d structure.Yaml
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	return d
}
