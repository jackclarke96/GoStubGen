package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func GenerateInterface(spec InterfaceSpec, common CommonSpec) error {
	const interfaceTemplate = `package {{ .Common.Package}}

// {{ .Interface.Name }} defines the interface
type {{ .Interface.Name }} interface {
{{- range .Interface.Methods }}
	{{ if .Description }}// {{ .Description }} {{- end }}
	{{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}) ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }})
{{- end }}
}
`

	tmpl, err := template.New("interface").Parse(interfaceTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	if err := os.MkdirAll(fmt.Sprintf("generated/%s", common.Package), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	file, err := os.Create(fmt.Sprintf("generated/%s/%s.go", common.Package, strings.ToLower(spec.Name)))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer file.Close()

	return tmpl.Execute(file, struct {
		Interface InterfaceSpec
		Common    CommonSpec
	}{
		Interface: spec,
		Common:    common,
	})
}
