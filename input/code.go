package input

import "github.com/constructor/raw"

type Code interface {
	read() raw.Code
}
type CodeImpl struct{}

func (CodeImpl) read() raw.Ast {
	return raw.Code{}
}
