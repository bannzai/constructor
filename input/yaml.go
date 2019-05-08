package input

import "github.com/constructor/value"

type Yaml interface {
	read() value.Yaml
}

type YamlImpl struct{}

func (YamlImpl) read() value.Yaml {
	return value.Yaml{}
}
