package generator

import (
	"github.com/constructor/file"
	"github.com/constructor/structure"
)

type Writer interface {
	Write(destinationPath structure.Path, content []byte)
}

type FileWriter struct{}

func (FileWriter) Write(destinationPath structure.Path, content []byte) {
	file.WriteFile(destinationPath, content)
	file.GoImports(destinationPath)
	file.GoFormat(destinationPath)
}
