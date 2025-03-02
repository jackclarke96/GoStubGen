package generator

type CommonSpec struct {
	Package  string `yaml:"package"`
	Importer string `yaml:"importer"`
}

// InterfaceSpec represents an interface definition
type InterfaceSpec struct {
	Name    string   `yaml:"name"`
	Methods []Method `yaml:"methods"`
}

// Method represents a method signature
type Method struct {
	Name        string  `yaml:"name"`
	Inputs      []Param `yaml:"inputs"`
	Outputs     []Param `yaml:"outputs"`
	Description string  `yaml:"description,omitempty"`
}

// Param represents a function parameter (input or output)
type Param struct {
	Name        string `yaml:"name,omitempty"`
	Type        string `yaml:"type"`
	Description string `yaml:"description,omitempty"`
}

// StructSpec represents a custom struct definition
type StructSpec struct {
	Name        string  `yaml:"name"`
	Fields      []Param `yaml:"fields"`
	Description string  `yaml:"description,omitempty"`
}

type CustomTypesSpec struct {
	Name        string `yaml:"name"`
	Definition  string `yaml:"definition"`
	Description string `yaml:"description,omitempty"`
}
