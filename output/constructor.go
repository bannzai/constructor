package output

import (
	"io/ioutil"

	"github.com/constructor/model"
)

type Constructor interface {
	Generate()
}
type ConstructorImpl struct {
	Source model.GenerateElementEachPackage
}

func (impl ConstructorImpl) Generate() {
	source := impl.Source
	content := source.Content()
	if err := ioutil.WriteFile(source.DestinationPath, content, 0644); err != nil {
		panic(err)
	}
}
