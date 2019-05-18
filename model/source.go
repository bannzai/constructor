package model

import "github.com/constructor/raw"

// GenerateElementEachPackage is inter struct for generate constructor.
type GenerateElementEachPackage struct {
	Package         string
	Template        raw.Template
	SourceCode      raw.Code
	DestinationPath raw.Path
}

func (t GenerateElementEachPackage) Content() []byte {
	return []byte{}
}
