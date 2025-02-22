package generator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateInterface(spec InterfaceSpec) error {
	const interfaceTemplate = `package {{ .Package}}

// {{ .Name }} defines the interface
type {{ .Name }} interface {
{{- range .Methods }}
	{{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}) ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }})
{{- end }}
}
`

	tmpl, err := template.New("interface").Parse(interfaceTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	// Ensure "generated/" directory exists
	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	file, err := os.Create(fmt.Sprintf("generated/%s.go", spec.Name))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	return tmpl.Execute(file, spec)
}
