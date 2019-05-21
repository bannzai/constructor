package raw

type (
	// Code is presentation of .go file content.
	Code struct {
		FilePath Path
		Structs  []Struct `ignore_constructor:true`
	}
	Struct struct {
		Name   string
		Fields []Field
	}
	Field struct {
		Name string
		Type string
	}
)
