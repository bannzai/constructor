package parser

import "github.com/constructor/input"

type YamlParsedResult string
type Yaml interface {
	parse() YamlParsedResult
}
type YamlImpl struct {
	Input input.Yaml
}

func (impl CodeImpl) parse() YamlParsedResult {
	return ""
}
