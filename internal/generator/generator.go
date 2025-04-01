package generator

type CommonSpec struct {
	Package  string `yaml:"package"`
	Importer string `yaml:"importer"`
}

// InterfaceSpec represents an interface definition
type InterfaceSpec struct {
	Name     string   `yaml:"name"`
	Embedded []string `yaml:"embedded"`
	Methods  []Method `yaml:"methods"`
}

// The top level yaml entry containing all interfaces
type Spec struct {
	Interfaces []InterfaceSpec `yaml:"interfaces"`
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
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type"`
}

// fields represents the field of a struct
type field struct {
	Name        string `yaml:"name,omitempty"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
}

// StructSpec represents a struct definition
type StructSpec struct {
	Name          string   `yaml:"name"`
	Embedded      []string `yaml:"embedded"`
	Implements    []string `yaml:"implements"`
	Fields        []field  `yaml:"fields"`
	Description   string   `yaml:"description,omitempty"`
	Methods       []Method
	FlattenEmbeds *bool `yaml:"flatten_embeds,omitempty"` // optional override

}

// CustomTypeSpec represents a custom (non-struct) type definition
type CustomTypesSpec struct {
	Name        string `yaml:"name"`
	Definition  string `yaml:"definition"`
	Description string `yaml:"description,omitempty"`
}

func (s StructSpec) getName() string {
	return s.Name
}

func (i InterfaceSpec) getName() string {
	return i.Name
}
