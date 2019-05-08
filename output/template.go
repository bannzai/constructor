package output

type Template interface {
	output() string
}
type TemplateImpl struct{}
