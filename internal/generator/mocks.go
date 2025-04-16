package generator

//todo add spy. add queue with locking
import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func generateMethodConfig() string {
	return `// {{ .MockConfigName }} stores mock flags and responses
type {{ .MockConfigName }} struct {
{{ range .Methods }}
	{{ .Name }} methodConfig[func({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}){{ end }}]
{{- end }}
}`
}

func generateMockStruct() string {
	return `// {{ .MockName }} embeds a concrete {{ .Interface }} and its mocks
type {{ .MockName }} struct {
	real   {{ .Package }}.{{ .Interface }}
	mocked {{ .MockConfigName }}
}`
}

func generateFactoryFunc() string {
	return `// {{ .MockFactory }} returns a new mock
func {{ .MockFactory }}(v {{ .Package }}.{{ .Interface }}) *{{ .MockName }} {
	return &{{ .MockName }}{
		real:   v,
		mocked: {{ .MockConfigName }}{},
	}
}`
}

const methodDividerTemplate = `
/* -------------------------- {{ .Name }} Mock Helpers --------------------------- */
`

const methodOverrideTemplate = `
// {{ .Name }} overrides the method to return the mock response
func (m *{{ .MockName }}) {{ .Name }}({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }} {{ $p.Type }}{{ end }}){{ if gt (len .Outputs) 0 }} ({{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}{{ $o.Type }}{{ end }}){{ end }} {
	if m.mocked.{{ title .Name }}.enabled {
		{{- if gt (len .Outputs) 0 }}
		return m.mocked.{{ .Name }}.NextResponse(func({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }} {{ $p.Type }}{{ end }}) ({{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}{{ $o.Type }}{{ end }}) {
			return m.real.{{ .Name }}({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }})
		})({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }})
		{{- else }}
		m.mocked.{{ .Name }}.NextResponse(func({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }} {{ $p.Type }}{{ end }}) {
			m.real.{{ .Name }}({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }})
		})({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }})
		return
		{{- end }}
	}
	{{- if gt (len .Outputs) 0 }}
	return m.real.{{ .Name }}({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }})
	{{- else }}
	m.real.{{ .Name }}({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Name }}{{ end }})
	{{- end }}
}`

const setFuncTemplate = `
// set{{ title .Name }}Func sets the function for {{ .Name }}
func (m *{{ .MockName }}) set{{ title .Name }}Func(f {{ responseSignature .Inputs .Outputs }}) {
	m.mocked.{{ .Name }}.fallback = f
}`

const setResponseTemplate = `
// set{{ title .Name }}Response sets the response for {{ .Name }}
func (m *{{ .MockName }}) set{{ title .Name }}Response({{ range $i, $p := .Outputs }}{{ if $i }}, {{ end }}output{{ $i }} {{ $p.Type }}{{ end }}) {
	m.set{{ title .Name }}Func(func({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Type }}{{ end }}) ({{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}{{ $o.Type }}{{ end }}) {
		return {{ range $i, $p := .Outputs }}{{ if $i }}, {{ end }}output{{ $i }}{{ end }}
	})
}`

const enableTemplate = `
// enable{{ title .Name }}Mock turns the mock on
func (m *{{ .MockName }}) enable{{ title .Name }}Mock() {
	m.mocked.{{ title .Name }}.enabled = true
}`

const disableTemplate = `
// disable{{ title .Name }}Mock turns the mock off
func (m *{{ .MockName }}) disable{{ title .Name }}Mock() {
	m.mocked.{{ title .Name }}.enabled = false
}`

const enqueueFuncTemplate = `
// enqueue{{ title .Name }}ResponseFunc enqueues a function response for {{ .Name }}
func (m *{{ .MockName }}) enqueue{{ title .Name }}ResponseFunc(f {{ responseSignature .Inputs .Outputs }}) {
	m.mocked.{{ .Name }}.EnqueueWithDelay(f, 0)
}`

const enqueueFuncWithDelayTemplate = `
// enqueue{{ title .Name }}ResponseFuncWithDelay enqueues a function response with delay for {{ .Name }}
func (m *{{ .MockName }}) enqueue{{ title .Name }}ResponseFuncWithDelay(f {{ responseSignature .Inputs .Outputs }}, d time.Duration) {
	m.mocked.{{ .Name }}.EnqueueWithDelay(f, d)
}`

const enqueueStaticTemplate = `
// enqueue{{ title .Name }}Response enqueues a static response for {{ .Name }}
func (m *{{ .MockName }}) enqueue{{ title .Name }}Response({{ range $i, $p := .Outputs }}{{ if $i }}, {{ end }}output{{ $i }} {{ $p.Type }}{{ end }}) {
	m.mocked.{{ .Name }}.EnqueueWithDelay(func({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Type }}{{ end }}) ({{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}{{ $o.Type }}{{ end }}) {
		return {{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}output{{ $i }}{{ end }}
	}, 0)
}`

const enqueueStaticWithDelayTemplate = `
// enqueue{{ title .Name }}ResponseWithDelay enqueues a static response with delay for {{ .Name }}
func (m *{{ .MockName }}) enqueue{{ title .Name }}ResponseWithDelay({{ range $i, $p := .Outputs }}{{ if $i }}, {{ end }}output{{ $i }} {{ $p.Type }}{{ end }}, d time.Duration) {
	m.mocked.{{ .Name }}.EnqueueWithDelay(func({{ range $i, $p := .Inputs }}{{ if $i }}, {{ end }}{{ $p.Type }}{{ end }}) ({{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}{{ $o.Type }}{{ end }}) {
		time.Sleep(d)
		return {{ range $i, $o := .Outputs }}{{ if $i }}, {{ end }}output{{ $i }}{{ end }}
	}, d)
}`

func writeTemplate(w io.Writer, tmplStr string, data any, funcs template.FuncMap) error {
	tmpl, err := template.New("").Funcs(funcs).Parse(tmplStr)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

func GenerateMock(spec InterfaceSpec, structSpec StructSpec, common CommonSpec) error {
	if err := os.MkdirAll("generated/importer", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	filePath := fmt.Sprintf("generated/importer/%s_mock_test.go", strings.ToLower(spec.Name))
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	funcs := template.FuncMap{
		"title": func(s string) string {
			return strings.Title(s)
		},
		"responseSignature": func(inputs, outputs []Param) string {
			var b strings.Builder
			b.WriteString("func(")
			for i, in := range inputs {
				if i > 0 {
					b.WriteString(", ")
				}
				b.WriteString(in.Type)
			}
			b.WriteString(")")
			if len(outputs) > 0 {
				b.WriteString(" (")
				for i, out := range outputs {
					if i > 0 {
						b.WriteString(", ")
					}
					b.WriteString(out.Type)
				}
				b.WriteString(")")
			}
			return b.String()
		},
	}

	headerTemplate := `package {{ .Importer }}

import "github.com/jackclarke/GoStubGen/generated/{{ .Package }}"

` + generateMethodConfig() + "\n\n" + generateMockStruct() + "\n\n" + generateFactoryFunc() + "\n"

	// Write the header section
	tmpl, err := template.New("header").Funcs(funcs).Parse(headerTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse header template: %w", err)
	}

	err = tmpl.Execute(file, struct {
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
		MockFactory:    fmt.Sprintf("new%sMock", strings.ToUpper(spec.Name[:1])+spec.Name[1:]),
		Methods:        spec.Methods,
		Package:        common.Package,
		Importer:       common.Importer,
	})
	if err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write method-specific helper functions
	for _, method := range spec.Methods {
		data := struct {
			MockName string
			Name     string
			Inputs   []Param
			Outputs  []Param
		}{
			MockName: fmt.Sprintf("mock%s", spec.Name),
			Name:     method.Name,
			Inputs:   method.Inputs,
			Outputs:  method.Outputs,
		}

		if err := writeTemplate(file, methodDividerTemplate, data, funcs); err != nil {
			return err
		}
		// Always generate core + function enqueue templates
		for _, tmplStr := range []string{
			methodOverrideTemplate,
			setFuncTemplate,
			enableTemplate,
			disableTemplate,
			enqueueFuncTemplate,
			enqueueFuncWithDelayTemplate,
		} {
			if err := writeTemplate(file, tmplStr, data, funcs); err != nil {
				return err
			}
		}

		// Only generate static response-based templates for methods with outputs
		if len(method.Outputs) > 0 {
			for _, tmplStr := range []string{
				setResponseTemplate,
				enqueueStaticTemplate,
				enqueueStaticWithDelayTemplate,
			} {
				if err := writeTemplate(file, tmplStr, data, funcs); err != nil {
					return err
				}
			}
		}

	}

	return nil
}
