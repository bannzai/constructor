package raw

// Definition is each elemnt for yaml
type Definition struct {
	Package           string `yaml:"package"`
	SourcePaths       []Path `yaml:"sourcePaths"`
	IgnoredPaths      []Path `yaml:"ignoredPaths"`
	TemplateFilePaths []Path `yaml:"templatePaths"`
	DestinationPath   Path   `yaml:"destinationPath"`
}

// Yaml is file data for format of .yaml
type Yaml struct {
	Definitions []Definition `yaml:"definitions"`
}
