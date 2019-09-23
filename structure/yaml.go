package structure

type (
	// Definition is each elemnt for yaml
	Definition struct {
		Package           string `yaml:"package"`
		SourcePath        Path   `yaml:"sourcePath"`
		IgnoredPaths      []Path `yaml:"ignoredPaths"`
		TemplateFilePaths []Path `yaml:"templatePaths"`
		DestinationPath   Path   `yaml:"destinationPath"`
	}
)
