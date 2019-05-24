package generator

import (
	"path/filepath"

	"github.com/constructor/structure"
)

type FilePathFetcher interface {
	sourceFilePaths(definition structure.Definition) []structure.Path
}

type Glob struct{}

func (Glob) sourceFilePaths(definition structure.Definition) []structure.Path {
	filePaths, err := filepath.Glob(definition.SourcePath)
	if err != nil {
		panic(err)
	}
	return filePaths
}
