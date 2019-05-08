package input

import (
	"io/ioutil"

	"github.com/constructor/raw"
	"gopkg.in/yaml.v2"
)

// Yaml is interface for read setting yaml file
type Yaml interface {
	Read() raw.Yaml
}

// YamlImpl for implement for Yaml
type YamlImpl struct {
	argument raw.Argument
}

func (impl YamlImpl) Read() raw.Yaml {
	buf, err := ioutil.ReadFile(impl.argument.YamlPath)
	if err != nil {
		panic(err)
	}

	var d raw.Yaml
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}

	return d
}
