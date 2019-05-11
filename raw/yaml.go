package raw

// Yaml is file data for format of .yaml
type Yaml struct {
	Package      string `yaml:"package"`
	Paths        []Path `yaml:"paths"`
	IgnoredPaths []Path `yaml:"ignoredPaths"`
}
