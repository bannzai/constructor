package input

import "github.com/constructor/raw"

type Yaml interface {
	Read() raw.Yaml
}

type YamlImpl struct{}

func (YamlImpl) Read() raw.Yaml {
	return raw.Yaml{}
}
