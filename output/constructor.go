package output

type Constructor interface {
	Output() error
}
type ConstructorImpl struct{}

func (ConstructorImpl) Output() error {
	return nil
}
