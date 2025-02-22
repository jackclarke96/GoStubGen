package generator

type PackageSpec struct {
	// TODO
}

// InterfaceSpec represents an interface definition
type InterfaceSpec struct {
	Name     string   `yaml:"name"`
	Methods  []Method `yaml:"methods"`
	Concrete string   `yaml:"concrete"`
	Package  string   `yaml:"package"`
}

// Method represents a method signature
type Method struct {
	Name    string  `yaml:"name"`
	Inputs  []Param `yaml:"inputs"`
	Outputs []Param `yaml:"outputs"`
}

// Param represents a function parameter (input or output)
type Param struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type"`
}

// StructSpec represents a custom struct definition
type StructSpec struct {
	Name   string  `yaml:"name"`
	Fields []Param `yaml:"fields"`
}
