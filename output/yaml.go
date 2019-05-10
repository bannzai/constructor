package output

type Yaml interface {
	Setup() error
}
type YamlImpl struct{}

func (YamlImpl) Setup() error {
	return nil
}
