package generator

import (
	"github.com/constructor/file"
	"github.com/constructor/structure"
)

type Writer interface {
	Write(destinationPath structure.Path, content string)
}

type FileWriter struct{}

func (FileWriter) Write(destinationPath structure.Path, content string) {
	file.WriteFile(destinationPath, content)
	file.GoImports(destinationPath)
	file.GoFormat(destinationPath)
}
