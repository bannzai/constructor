package output

type Template interface {
	Output() error
}
type TemplateImpl struct{}

func (impl TemplateImpl) output() error {

}
