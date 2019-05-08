package input

import "github.com/constructor/raw"

type Yaml interface {
	read() raw.Yaml
}

type YamlImpl struct{}

func (YamlImpl) read() raw.Yaml {
	return raw.Yaml{}
}
