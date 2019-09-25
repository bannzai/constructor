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

func (generator Constructor) Generate(sourcePath, destinationPath string, typeName string) {
	templateExecutor := generator.TemplateReader.Read(sourcePath)
	var sourceCode structure.Code
	if len(typeName) > 0 {
		sourceCode = generator.SourceCodeReader.ReadWithType(sourcePath, typeName)
	} else {
		sourceCode = generator.SourceCodeReader.Read(sourcePath)
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
