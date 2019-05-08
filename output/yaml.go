package output

type Yaml interface {
	Output() error
}
type YamlImpl struct{}

func (YamlImpl) Output() error {

}
