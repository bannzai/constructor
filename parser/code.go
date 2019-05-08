package parser

import "github.com/constructor/input"

type CodeParsedResult string
type Code interface {
	parse() CodeParsedResult
}
type CodeImpl struct {
	Input input.Code
}

func (impl CodeImpl) parse() CodeParsedResult {
	return ""
}
