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

func (generator Constructor) Generate(templatePath, sourcePath, destinationPath string, typeName string, ignoreFieldNames []string) {
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
	}); err != nil {
		panic(err)
	}

	content := buf.String()
	generator.FileWriter.Write(destinationPath, content)
}
