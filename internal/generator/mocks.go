package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func GenerateMock(spec InterfaceSpec, structSpec StructSpec, common CommonSpec) error {
	const mockTemplate = `package {{ .Package }}

// {{ .MockConfigName }} stores mock flags and responses
type {{ .MockConfigName }} struct {
{{ range .Methods }}
	// {{ .Name }} flag and mock response
	mock{{ title .Name }}     bool
	{{ .Name }}Response func({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }}
{{ end }}
}

// {{ .MockName }} embeds {{ .Concrete }} and its mocks
type {{ .MockName }} struct {
	{{ .ConcreteVar }}   {{ .Concrete }}
	mocks {{ .MockConfigName }}
}

// {{ .MockFactory }} returns a new mock
func {{ .MockFactory }}() {{ .Interface }} {
	{{ .ConcreteVar }} := New{{ .Concrete }}()
	return {{ .MockName }}{
		{{ .ConcreteVar }}:   *{{ .ConcreteVar }},
		mocks: {{ .MockConfigName }}{},
	}
}

{{- range .Methods }}
/* -------------------------- {{ .Name }} Mock Helpers --------------------------- */

// {{ .Name }} overrides the method to return the mock response
func (m {{ $.MockName }}) {{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }} {
	if m.mocks.mock{{ title .Name }} {
		{{- if gt (len .Outputs) 0 }}
		return m.mocks.{{ .Name }}Response({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
		{{- else }}
		m.mocks.{{ .Name }}Response({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
		{{- end }}
	}
	{{- if gt (len .Outputs) 0 }}
	return m.{{ $.ConcreteVar }}.{{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
	{{- else }}
	m.{{ $.ConcreteVar }}.{{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
	{{- end }}
}

// set{{ title .Name }}Response sets the response for {{ .Name }}
func (m {{ $.MockName }}) set{{ title .Name }}Response(resp func({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }}) {
	m.mocks.{{ .Name }}Response = resp
}

// enable{{ title .Name }}Response turns the mock on
func (m {{ $.MockName }}) enable{{ title .Name }}Response() {
	m.mocks.mock{{ title .Name }} = true
}

// disable{{ title .Name }}Response turns the mock off
func (m {{ $.MockName }}) disable{{ title .Name }}Response() {
	m.mocks.mock{{ title .Name }} = false
}
{{- end }}
`

	// Ensure "generated/" directory exists
	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	filePath := fmt.Sprintf("generated/%s_mock.go", structSpec.Name)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Compute a variable name for the concrete type by lower-casing the first letter.
	concreteVar := ""
	if len(structSpec.Name) > 0 {
		concreteVar = strings.ToLower(string(structSpec.Name[0])) + structSpec.Name[1:]
	}

	tmpl, err := template.New("mock").Funcs(template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
	}).Parse(mockTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Execute template with mock struct details, including the dynamic concrete variable.
	return tmpl.Execute(file, struct {
		Interface      string
		Concrete       string
		ConcreteVar    string
		MockName       string
		MockConfigName string
		MockFactory    string
		Methods        []Method
		Package        string
	}{
		Interface:      spec.Name,
		Concrete:       structSpec.Name,
		ConcreteVar:    concreteVar,
		MockName:       fmt.Sprintf("mock%s", spec.Name),
		MockConfigName: fmt.Sprintf("mock%sConfig", spec.Name),
		MockFactory:    fmt.Sprintf("%sMock", strings.ToLower(spec.Name)),
		Methods:        spec.Methods,
		Package:        common.Package,
	})
}
