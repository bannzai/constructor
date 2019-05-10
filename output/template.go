package output

type Template interface {
	Setup() error
}
type TemplateImpl struct{}

func (impl TemplateImpl) Setup() error {
	return nil
}
