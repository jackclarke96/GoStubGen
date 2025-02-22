package generator

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func GenerateStructs(structs []StructSpec, common CommonSpec, concreteType string) error {
	const structTemplate = `package {{ .Common.Package }}

// {{ .Struct.Name }} is a user-defined struct
type {{ .Struct.Name }} struct {
{{- range .Struct.Fields }}
    {{ .Name }} {{ .Type }}
{{- end }}
}`

	// Ensure the "generated/" directory exists
	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	for _, structDef := range structs {
		// Skip the struct that implements the interface (e.g., "Car" in this case)
		if structDef.Name == concreteType {
			continue
		}

		filePath := fmt.Sprintf("generated/%s.go", strings.ToLower(structDef.Name))
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		tmpl, err := template.New("struct").Parse(structTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		combinedTemplate := struct {
			Struct StructSpec
			Common CommonSpec
		}{
			Struct: structDef,
			Common: common,
		}

		if err := tmpl.Execute(file, combinedTemplate); err != nil {
			return fmt.Errorf("failed to write struct file: %w", err)
		}
	}

	return nil
}
