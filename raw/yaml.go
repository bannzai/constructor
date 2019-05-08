package raw

// Path for `constructor` about yaml configuration
type Path = string

// Yaml is file data for format of .yaml
type Yaml struct {
	Paths        []Path `yaml:"paths"`
	IgnoredPaths []Path `yaml:"ignoredPaths"`
}
