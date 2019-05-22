package generator

import "github.com/constructor/structure"

type YamlReader interface {
	Read() structure.Yaml
}
