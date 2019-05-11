package model

import "github.com/constructor/raw"

// GenerateElementEachPackage is inter struct for generate constructor.
type GenerateElementEachPackage struct {
	Package  string
	Template raw.Template
	CodeList []raw.Code
}

func (t *GenerateElementEachPackage) Add(code raw.Code) {
	t.CodeList = append(t.CodeList, code)
}

func (t GenerateElementEachPackage) GenerateContent() string {
	return ""
}
