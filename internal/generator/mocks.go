package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func GenerateMock(spec InterfaceSpec, structSpec StructSpec, common CommonSpec) error {
	const mockTemplate = `package {{ .Importer }}

	import "github.com/jackclarke/GoStubGen/generated/{{ .Package }}"

// {{ .MockConfigName }} stores mock flags and responses
type {{ .MockConfigName }} struct {
{{ range .Methods }}
	{{ .Name }} methodConfig[func({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }}]
{{- end }}
}

// {{ .MockName }} embeds a concrete {{ .Interface }} and its mocks
type {{ .MockName }} struct {
	real   {{ .Package }}.{{ .Interface }}
	mocked {{ .MockConfigName }}
}

// {{ .MockFactory }} returns a new mock
func {{ .MockFactory }}(v {{ .Package }}.{{ .Interface }}) *{{ .MockName }} {
	return &{{ .MockName }}{
		real:   v,
		mocked: {{ .MockConfigName }}{},
	}
}

{{- range .Methods }}
/* -------------------------- {{ .Name }} Mock Helpers --------------------------- */

// {{ .Name }} overrides the method to return the mock response
func (m *{{ $.MockName }}) {{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }} {
	if m.mocked.{{ title .Name }}.enabled {
		{{- if gt (len .Outputs) 0 }}
		return m.mocked.{{ .Name }}.response({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
		{{- else }}
		m.mocked.{{ .Name }}.response({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
		{{- end }}
	}
	{{- if gt (len .Outputs) 0 }}
	return m.real.{{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
	{{- else }}
	m.real.{{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }}{{ end }})
	{{- end }}
}

// set{{ title .Name }}Func sets the function for {{ .Name }}
func (m *{{ $.MockName }}) set{{ title .Name }}Func(f func({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }}) {
	m.mocked.{{ .Name }}.response = f
}

{{ if gt (len .Outputs) 0 }}
// set{{ title .Name }}Response sets the response for {{ .Name }}
func (m *{{ $.MockName }}) set{{ title .Name }}Response{{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}output{{ $index }} {{ $param.Type }}{{ end }}){{ end }} {
	m.set{{ .Name }}Func(func({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }} {
		return {{ if gt (len .Outputs) 0 }}{{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}output{{ $index }}{{ end }}{{ end }}
	})
}
{{ end }}

// enable{{ title .Name }}Mock turns the mock on
func (m *{{ $.MockName }}) enable{{ title .Name }}Mock() {
	m.mocked.{{ title .Name }}.enabled = true
}

// disable{{ title .Name }}Mock turns the mock off
func (m *{{ $.MockName }}) disable{{ title .Name }}Mock() {
	m.mocked.{{ title .Name }}.enabled = false
}
{{- end }}
`

	if err := os.MkdirAll("generated/importer", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	filePath := fmt.Sprintf("generated/importer/%s_mock_test.go", strings.ToLower(spec.Name))
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	tmpl, err := template.New("mock").Funcs(template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
	}).Parse(mockTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	return tmpl.Execute(file, struct {
		Interface      string
		Concrete       string
		MockName       string
		MockConfigName string
		MockFactory    string
		Methods        []Method
		Package        string
		Importer       string
	}{
		Interface:      spec.Name,
		Concrete:       structSpec.Name,
		MockName:       fmt.Sprintf("mock%s", spec.Name),
		MockConfigName: fmt.Sprintf("mock%sConfig", spec.Name),
		MockFactory:    fmt.Sprintf("%sMock", strings.ToLower(spec.Name[:1])+spec.Name[1:]),
		Methods:        spec.Methods,
		Package:        common.Package,
		Importer:       common.Importer,
	})
}
