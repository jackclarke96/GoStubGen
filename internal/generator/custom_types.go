package generator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateTypesAndStructs(structs []StructSpec, types []CustomTypesSpec, common CommonSpec) error {
	// Create the file once and write the package declaration first.
	filePath := fmt.Sprintf("generated/%s/custom_types.go", common.Package)
	if err := os.MkdirAll(fmt.Sprintf("generated/%s", common.Package), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create 'generated' directory: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("package %s\n\n", common.Package)); err != nil {
		return fmt.Errorf("failed to write package declaration: %w", err)
	}

	// Define struct and custom type templates
	const structTemplate = `// {{ .Struct.Description }}
type {{ .Struct.Name }} struct {
{{- range .Struct.Fields }}
	{{ if .Description }}// {{ .Description }} {{- end }}
    {{ .Name }} {{ .Type }}
{{- end }}
}
`
	const typeTemplate = `// {{ .Type.Description }}
type {{ .Type.Name }} {{ .Type.Definition }}

`

	tmplStruct, err := template.New("struct").Parse(structTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse struct template: %w", err)
	}

	tmplType, err := template.New("type").Parse(typeTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse type template: %w", err)
	}

	// Write custom types first
	for _, typeDef := range types {
		combinedTypeTemplate := struct {
			Type CustomTypesSpec
		}{
			Type: typeDef,
		}
		if err := tmplType.Execute(file, combinedTypeTemplate); err != nil {
			return fmt.Errorf("failed to write type to file: %w", err)
		}
	}

	// Write struct definitions
	for _, structDef := range structs {
		combinedStructTemplate := struct {
			Struct StructSpec
			Common CommonSpec
		}{
			Struct: structDef,
			Common: common,
		}

		if err := tmplStruct.Execute(file, combinedStructTemplate); err != nil {
			return fmt.Errorf("failed to write struct to file: %w", err)
		}
	}
	return nil
}
