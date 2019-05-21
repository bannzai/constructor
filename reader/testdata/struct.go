package testdata

import "io"

type Alias string
type Struct struct {
	P string
	F func(aaa int, bbb bool) string
	A Alias
	O io.Writer
	L []int
	M map[string]bool
}
