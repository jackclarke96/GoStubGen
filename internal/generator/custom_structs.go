package generator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateStructs(structs []StructSpec, common CommonSpec) error {
	// Create the file once and write the package declaration first.
	filePath := fmt.Sprintf("generated/%s/custom_types.go", common.Package)
	// Ensure the "generated/" directory exists
	if err := os.MkdirAll(fmt.Sprintf("generated/%s", common.Package), os.ModePerm); err != nil {
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
	const structTemplate = `// {{ .Struct.Description }}
type {{ .Struct.Name }} struct {
{{- range .Struct.Fields }}
	{{ if .Description }}// {{ .Description }} {{- end }}
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
