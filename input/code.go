package input

import "github.com/constructor/value"

type Code interface {
	read() value.Ast
}
type CodeImpl struct{}

func (CodeImpl) read() value.Ast {
	return value.Ast{}
}
