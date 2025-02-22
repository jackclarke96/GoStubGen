package generator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateStructs(structs []StructSpec, common CommonSpec, concreteType string) error {
	// Create the file once and write the package declaration first.
	filePath := "generated/custom_types.go"
	// Ensure the "generated/" directory exists
	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the package declaration only once.
	if _, err := file.WriteString(fmt.Sprintf("package %s\n\n", common.Package)); err != nil {
		return fmt.Errorf("failed to write package declaration: %w", err)
	}

	// Define the template without the package line.
	const structTemplate = `// {{ .Struct.Name }} is a user-defined struct
type {{ .Struct.Name }} struct {
{{- range .Struct.Fields }}
    {{ .Name }} {{ .Type }}
{{- end }}
}

`

	tmpl, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Loop through the structs and write each one to the file.
	for _, structDef := range structs {
		// Skip the struct that implements the interface (e.g., "Car")
		if structDef.Name == concreteType {
			continue
		}

		combinedTemplate := struct {
			Struct StructSpec
			Common CommonSpec
		}{
			Struct: structDef,
			Common: common,
		}

		if err := tmpl.Execute(file, combinedTemplate); err != nil {
			return fmt.Errorf("failed to write struct to file: %w", err)
		}
	}

	return nil
}
