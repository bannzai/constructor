package parser

import "github.com/constructor/input"

type CodeParsedResult string
type Code interface {
	Parse() CodeParsedResult
}
type CodeImpl struct {
	Input input.Code
}

func (impl CodeImpl) Parse() CodeParsedResult {
	return ""
}
