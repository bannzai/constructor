package output

type Constructor interface {
	Generate() error
}
type ConstructorImpl struct{}

func (ConstructorImpl) Generate() error {
	return nil
}
