package output

type Yaml interface {
	output() string
}
type YamlImpl struct{}
