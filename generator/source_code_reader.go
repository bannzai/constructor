//go:generate mockgen -source=$GOFILE -destination=$GOPATH/src/github.com/bannzai/constructor/generator/source_code_reader_interface_mock_test.go -package=$GOPACKAGE
package generator

import "github.com/bannzai/constructor/structure"

type SourceCodeReader interface {
	Read(filePath structure.Path, ignoreFieldNames []string) structure.Code
	ReadWithType(filePath structure.Path, generatedTypeName string, ignoreFieldNames []string) (code structure.Code)
}
