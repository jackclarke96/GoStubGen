package generator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateStructs(structs []StructSpec, concreteType string) error {
	const structTemplate = `package generated

// {{ .Name }} is a user-defined struct
type {{ .Name }} struct {
{{- range .Fields }}
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

		filePath := fmt.Sprintf("generated/%s.go", structDef.Name)
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		tmpl, err := template.New("struct").Parse(structTemplate)
		if err != nil {
			return fmt.Errorf("failed to parse template: %w", err)
		}

		if err := tmpl.Execute(file, structDef); err != nil {
			return fmt.Errorf("failed to write struct file: %w", err)
		}
	}

	return nil
}
