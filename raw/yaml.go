package raw

const YamlFilePathName = "constructor.yaml"

type (
	// Yaml is file data for format of .yaml
	Yaml struct {
		Definitions []Definition `yaml:"definitions"`
	}
	// Definition is each elemnt for yaml
	Definition struct {
		Package           string `yaml:"package"`
		SourcePath        Path   `yaml:"sourcePath"`
		IgnoredPaths      []Path `yaml:"ignoredPaths"`
		TemplateFilePaths []Path `yaml:"templatePaths"`
		DestinationPath   Path   `yaml:"destinationPath"`
	}
)
