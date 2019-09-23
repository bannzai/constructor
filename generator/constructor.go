package generator

type Constructor struct {
	TemplateReader   TemplateReader
	SourceCodeReader SourceCodeReader
	FileWriter       Writer
}

func (impl Constructor) Generate() {
}
