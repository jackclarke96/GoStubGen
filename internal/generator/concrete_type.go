package generator

import (
	"fmt"
	"os"
	"text/template"
)

// GenerateStruct generates a Go struct file with stub methods
func GenerateConcreteType(spec InterfaceSpec, structSpec StructSpec, common CommonSpec) error {
	const structTemplate = `package {{ .Common.Package }}

// {{ .Struct.Name }} is the concrete implementation of {{ .Interface.Name }}
type {{ .Struct.Name }} struct {
{{- range .Struct.Fields }}
    {{ .Name }} {{ .Type }}
{{- end }}
}

// New{{ .Struct.Name }} creates a new instance of {{ .Struct.Name }} with default values
func New{{ .Struct.Name }}() *{{ .Struct.Name }} {
	return &{{ .Struct.Name }}{
		{{- range .Struct.Fields }}
		{{ .Name }}: {{ getDefaultReturnValue .Type }},
		{{- end }}
	}
}

{{- range .Interface.Methods }}
func (s *{{ $.Struct.Name }}) {{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}) ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}) {
	// TODO: Implement
	{{- if gt (len .Outputs) 0 }}
	return {{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ getDefaultReturnValue .Type }}{{ end }}
	{{- end }}
}
{{- end }}
`

	// Ensure "generated/" directory exists
	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	filePath := fmt.Sprintf("generated/%s.go", structSpec.Name)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	tmpl, err := template.New("struct").Funcs(template.FuncMap{
		"getDefaultReturnValue": getDefaultReturnValue,
	}).Parse(structTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template with combined struct and interface data
	return tmpl.Execute(file, struct {
		Interface InterfaceSpec
		Struct    StructSpec
		Common    CommonSpec
	}{
		Interface: spec,
		Struct:    structSpec,
		Common:    common,
	})
}

// GenerateConstructor generates a basic constructor for the given struct
func GenerateConstructor(structSpec StructSpec) error {
	const constructorTemplate = `package generated

// New{{ .Name }} creates a new instance of {{ .Name }} with default values
func New{{ .Name }}() *{{ .Name }} {
	return &{{ .Name }}{
		{{- range .Fields }}
		{{ .Name }}: {{ getDefaultReturnValue .Type }},
		{{- end }}
	}
}
`

	// Ensure "generated/" directory exists
	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	file, err := os.Create(fmt.Sprintf("generated/%s.go", structSpec.Name))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	tmpl, err := template.New("constructor").Funcs(template.FuncMap{
		"getDefaultReturnValue": getDefaultReturnValue,
	}).Parse(constructorTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	return tmpl.Execute(file, structSpec)
}
