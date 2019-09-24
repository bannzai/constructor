package generator

import (
	"github.com/bannzai/constructor/file"
	"github.com/bannzai/constructor/structure"
)

type FileWriter interface {
	Write(destinationPath structure.Path, content string)
}

type FileWriterImpl struct{}

func (FileWriterImpl) Write(destinationPath structure.Path, content string) {
	file.WriteFile(destinationPath, content)
	file.GoImports(destinationPath)
}
