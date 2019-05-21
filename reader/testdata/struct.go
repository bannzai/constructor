package testdata

type Alias string
type Struct struct {
	P string
	F func(aaa int, bbb bool) string
	A Alias
}
