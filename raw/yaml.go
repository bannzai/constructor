package raw

type (
	// Yaml is file data for format of .yaml
	Yaml struct {
		Definitions []Definition `yaml:"definitions"`
	}
	// Definition is each elemnt for yaml
	Definition struct {
		Package           string `yaml:"package"`
		SourcePaths       []Path `yaml:"sourcePaths"`
		IgnoredPaths      []Path `yaml:"ignoredPaths"`
		TemplateFilePaths []Path `yaml:"templatePaths"`
		DestinationPath   Path   `yaml:"destinationPath"`
	}
)
