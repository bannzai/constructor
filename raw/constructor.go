// DO NOT EDIT THIS FILE.
// File generated by constructor.
// https://github.com/bannzai/constructor
package raw

// NewArgumentRaw insitanciate Argument
func NewArgumentRaw(
	yamlPath Path,
) Argument {
	return Argument{
		YamlPath: yamlPath,
	}
}

// NewCodeRaw insitanciate Code
func NewCodeRaw(
	filePath Path,
) Code {
	return Code{
		FilePath: filePath,
	}
}

// NewStructRaw insitanciate Struct
func NewStructRaw(
	name string,
) Struct {
	return Struct{
		Name: name,
	}
}

// NewFieldRaw insitanciate Field
func NewFieldRaw(
	name string,
	_type string,
) Field {
	return Field{
		Name: name,
		Type: _type,
	}
}

// NewYamlRaw insitanciate Yaml
func NewYamlRaw() Yaml {
	return Yaml{}
}

// NewDefinitionRaw insitanciate Definition
func NewDefinitionRaw(
	_package string,
	sourcePath Path,
	destinationPath Path,
) Definition {
	return Definition{
		Package:         _package,
		SourcePath:      sourcePath,
		DestinationPath: destinationPath,
	}
}
