package testdata

import "io"

// referenced playground: https://play.golang.org/p/NROZdl65DlM
type XXX interface{}
type Alias string
type Struct struct {
	P    string
	F    func(aaa int, bbb bool) string
	A    Alias
	O    io.Writer
	L    []int
	LO   []io.Writer
	M    map[string]bool
	MKO  map[io.Writer]bool
	MVO  map[string]io.Writer
	MKVO map[io.Writer]io.Writer
	I    string
	XXX
}
