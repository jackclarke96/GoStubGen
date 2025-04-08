package generator

import (
	"fmt"
	"os"
	"text/template"
)

func GenerateMockHelpers(importer string) error {
	const helpersTemplate = `package {{ .Importer }}

// MethodConfig is a generic container for mocking method behavior.
type methodConfig[T any] struct {
	mu            sync.Mutex
	response      T
	enabled       bool
	useChannel    bool
	staticValue   T
	responseFunc  func() T
	responseQueue chan func() T

	callCount int
	callLog   []struct{} // or extend with args, timestamps, etc
}
`

	// Create the output directory (e.g., generated/internal/)
	outputDir := "generated/importer"
	fmt.Println("output dir", outputDir)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", outputDir, err)
	}

	// Output file path
	filePath := fmt.Sprintf("%s/mock_helpers_test.go", outputDir)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Parse and render the template
	tmpl, err := template.New("helpers").Parse(helpersTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Only passing the importer since that’s all we need
	return tmpl.Execute(file, struct {
		Importer string
	}{
		Importer: importer,
	})
}
