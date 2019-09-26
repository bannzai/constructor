// DO NOT EDIT THIS FILE.
// File generated by constructor.
// https://github.com/bannzai/constructor
package structure

// NewCodeStructure insitanciate Code
func NewCodeStructure(
	filePath Path,
	structs []Struct,
) Code {
	return Code{
		FilePath: filePath,
		Structs:  structs,
	}
}

// NewFieldStructure insitanciate Field
func NewFieldStructure(
	name string,
	_type string,
) Field {
	return Field{
		Name: name,
		Type: _type,
	}
}

// NewStructStructure insitanciate Struct
func NewStructStructure(
	fields []Field,
	name string,
) Struct {
	return Struct{
		Fields: fields,
		Name:   name,
	}
}