package input

import (
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/constructor/raw"
)

type Code interface {
	Read(filePath raw.Path) raw.Code
}
type CodeImpl struct{}

func (impl CodeImpl) Read(filePath raw.Path) raw.Code {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, filePath, buf, 0)
	if err != nil {
		panic(err)
	}
	return raw.Code{
		FilePath: filePath,
		ASTFile:  *astFile,
	}
}
