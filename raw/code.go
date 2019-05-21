package raw

type (
	// Code is presentation of .go file content.
	Code struct {
		FilePath Path `ignore_constructor:true`
		Structs  []Struct
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
