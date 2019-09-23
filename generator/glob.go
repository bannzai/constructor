package generator

import (
	"path/filepath"

	"github.com/bannzai/constructor/structure"
)

type FilePathFetcher interface {
	sourceFilePaths(definition structure.Definition) []structure.Path
}

type Glob struct{}

func (Glob) sourceFilePaths(definition structure.Definition) []structure.Path {
	sourceFilePaths, err := filepath.Glob(definition.SourcePath)
	if err != nil {
		panic(err)
	}

	ignoreFilePaths := []string{}

	for _, ignorePath := range definition.IgnoredPaths {
		ignores, err := filepath.Glob(ignorePath)
		if err != nil {
			panic(err)
		}

		ignoreFilePaths = append(ignoreFilePaths, ignores...)
	}

	paths := []structure.Path{}
	for _, sourceFilePath := range sourceFilePaths {
	inner:
		for _, ignoreFilePath := range ignoreFilePaths {
			if sourceFilePath == ignoreFilePath {
				continue inner
			}
			paths = append(paths, sourceFilePath)
		}
	}

	return paths
}
