package parser

import (
	"github.com/constructor/raw"
)

type CodeParsedResult string
type Code interface {
	Parse(code raw.Code) CodeParsedResult
}
type CodeImpl struct{}

func (impl CodeImpl) Parse(code raw.Code) CodeParsedResult {
	return ""
}
