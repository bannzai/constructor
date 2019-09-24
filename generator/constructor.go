package generator

import (
	"bytes"
)

type Constructor struct {
	TemplateReader
	SourceCodeReader
	FileWriter
}

func (generator Constructor) Generate(sourcePath, destinationPath string) {
	templateExecutor := generator.TemplateReader.Read(sourcePath)
	sourceCode := generator.SourceCodeReader.Read(sourcePath)

	buf := &bytes.Buffer{}
	if err := templateExecutor.Execute(buf, map[string]interface{}{
		"Structs": sourceCode.Structs,
	}); err != nil {
		panic(err)
	}

	content := buf.String()
	generator.FileWriter.Write(destinationPath, content)
}
