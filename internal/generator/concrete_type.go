package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// GenerateStruct generates a Go struct file with stub methods
func GenerateConcreteTypes(implementers []StructSpec, common CommonSpec) error {
	const structTemplate = `package {{ .Common.Package }}

{{ if .Struct.Description }}// {{ .Struct.Description }} {{ end }}
type {{ .Struct.Name }} struct {
{{- range .Struct.Fields }}
    {{ .Name }} {{ .Type }}
{{- end }}
}

// New {{ .Struct.Name }} creates a new instance of {{ .Struct.Name }} with default values
func New{{ .Struct.Name }}() *{{ .Struct.Name }} {
	return &{{ .Struct.Name }}{
		{{- range .Struct.Fields }}
		{{ .Name }}: {{ getDefaultReturnValue .Type }},
		{{- end }}
	}
}

{{ range .Struct.Methods }}
{{- if .Description }}// {{ .Description }} {{- end }}
func (s *{{ $.Struct.Name }}) {{ .Name }}({{ range $index, $param := .Inputs }}{{ if $index }}, {{ end }}{{ $param.Name }} {{ $param.Type }}{{ end }}) ({{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ $param.Type }}{{ end }}) {
	{{- if gt (len .Outputs) 0 }}
	return {{ range $index, $param := .Outputs }}{{ if $index }}, {{ end }}{{ getDefaultReturnValue .Type }}{{ end }}
	{{- end }}
}
{{ end }}
`

	if err := os.MkdirAll(fmt.Sprintf("generated/%s", common.Package), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	for _, structDef := range implementers {
		filePath := fmt.Sprintf("generated/%s/%s.go", common.Package, strings.ToLower(structDef.Name))
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		tmpl, err := template.New("struct").Funcs(template.FuncMap{
			"getDefaultReturnValue": getZeroVal,
		}).Parse(structTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		err = tmpl.Execute(file, struct {
			Struct StructSpec
			Common CommonSpec
		}{
			Struct: structDef,
			Common: common,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
