package input

import "github.com/constructor/raw"

type Code interface {
	Read() raw.Code
}
type CodeImpl struct{}

func (CodeImpl) Read() raw.Code {
	return raw.Code{}
}
