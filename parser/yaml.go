package parser

import (
	"github.com/constructor/raw"
)

type YamlParsedResult string
type Yaml interface {
	Parse(yaml raw.Yaml) YamlParsedResult
}
type YamlImpl struct{}

func (impl YamlImpl) Parse(yaml raw.Yaml) YamlParsedResult {
	return ""
}
