package generator

import "github.com/bannzai/constructor/structure"

type YamlReader interface {
	Read() structure.Yaml
}
