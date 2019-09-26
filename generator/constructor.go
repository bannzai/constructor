//go:generate constructor generate --source=$GOFILE --destination=$GOPATH/src/github.com/bannzai/constructor/generator/constructor.constructor.go --package=$GOPACKAGE --template=$GOPATH/src/github.com/bannzai/constructor/constructor.tpl --type=Constructor --ignoreFields=TemplateReader,SourceCodeReader
package generator

import (
	"bytes"

	"github.com/bannzai/constructor/structure"
)

type Constructor struct {
	TemplateReader
	SourceCodeReader
	FileWriter
}

func (generator Constructor) Generate(sourcePath, destinationPath, templatePath string, typeName string, ignoreFieldNames []string, packageName string) {
	templateExecutor := generator.TemplateReader.Read(templatePath)
	var sourceCode structure.Code
	if len(typeName) > 0 {
		sourceCode = generator.SourceCodeReader.ReadWithType(sourcePath, typeName, ignoreFieldNames)
	} else {
		sourceCode = generator.SourceCodeReader.Read(sourcePath, ignoreFieldNames)
	}

	buf := &bytes.Buffer{}
	if err := templateExecutor.Execute(buf, map[string]interface{}{
		"Structs": sourceCode.Structs,
		"Package": packageName,
	}); err != nil {
		panic(err)
	}

	content := buf.String()
	generator.FileWriter.Write(destinationPath, content)
}
