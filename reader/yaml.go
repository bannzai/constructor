package reader

import (
	"io/ioutil"

	"github.com/constructor/structure"
	"gopkg.in/yaml.v2"
)

// Yaml is interface for read setting yaml file
type Yaml interface {
	Read() structure.Yaml
}

// YamlImpl for implement for Yaml
type YamlImpl struct {
	Argument structure.Argument
}

// Read for yaml file
func (impl YamlImpl) Read() structure.Yaml {
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
