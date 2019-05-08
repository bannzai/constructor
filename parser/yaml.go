package parser

import "github.com/constructor/input"

type YamlParsedResult string
type Yaml interface {
	Parse() YamlParsedResult
}
type YamlImpl struct {
	Input input.Yaml
}

func (impl YamlImpl) Parse() YamlParsedResult {
	return ""
}
