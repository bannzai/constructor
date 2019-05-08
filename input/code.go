package input

import (
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/constructor/raw"
)

type Code interface {
	Read() raw.Code
}
type CodeImpl struct {
	FilePath string
}

func (impl CodeImpl) Read() raw.Code {
	buf, err := ioutil.ReadFile(impl.FilePath)
	if err != nil {
		panic(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, impl.FilePath, buf, 0)
	if err != nil {
		panic(err)
	}
	return raw.Code{
		FilePath: impl.FilePath,
		ASTFile:  astFile,
	}
}
