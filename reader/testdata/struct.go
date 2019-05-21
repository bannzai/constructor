package testdata

import "io"

type Alias string
type Struct struct {
	P string
	F func(aaa int, bbb bool) string
	A Alias
	T io.Writer
}
