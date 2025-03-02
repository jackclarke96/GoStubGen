package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jackclarke/GoStubGen/internal/generator"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Config structure to match the YAML file format
type Config struct {
	Package        string                      `yaml:"package"`
	Importer       string                      `yaml:"importer"`
	PackageImports []string                    `yaml:"package_imports"`
	CustomStructs  []generator.StructSpec      `yaml:"custom_structs"`
	CustomTypes    []generator.CustomTypesSpec `yaml:"custom_types"`
	Implementers   []generator.StructSpec      `yaml:"implementers"`
	Name           string                      `yaml:"name"`
	Methods        []generator.Method          `yaml:"methods"`
}

var configPath string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go code from a YAML config",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		var config Config
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("Invalid YAML format: %v", err)
		}

		interfaceSpec := generator.InterfaceSpec{
			Name:    config.Name,
			Methods: config.Methods,
		}

		commonSpec := generator.CommonSpec{
			Package:  config.Package,
			Importer: config.Importer,
		}

		// Generate custom structs (excluding the ones implementing the interface)
		if err := generator.GenerateTypesAndStructs(config.CustomStructs, config.CustomTypes, commonSpec); err != nil {
			log.Fatalf("Error generating custom structs: %v", err)
		}

		// Generate the interface
		if err := generator.GenerateInterface(interfaceSpec, commonSpec); err != nil {
			log.Fatalf("Error generating interface: %v", err)
		}

		// Generate concrete types structs
		if err := generator.GenerateConcreteTypes(interfaceSpec, config.Implementers, commonSpec); err != nil {
			log.Fatalf("Error generating custom structs: %v", err)
		}

		// Generate the mocks
		mockInterfaceSpec := prefixTypesWithPackageName(config, interfaceSpec, commonSpec.Package)
		if err := generator.GenerateMock(mockInterfaceSpec, generator.StructSpec{}, commonSpec); err != nil {
			log.Fatalf("Error generating mock: %v", err)
		}

		fmt.Println("Code generation complete!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to YAML config file")
	generateCmd.MarkFlagRequired("config")
}

func prefixTypesWithPackageName(config Config, spec generator.InterfaceSpec, packageName string) generator.InterfaceSpec {
	// Generate the mock implementation
	customTypes := make(map[string]bool)
	for _, cs := range config.CustomStructs {
		customTypes[cs.Name] = true
	}

	// Prefix types with package name so that they are suitable for import into external package when added to mocks.
	for mIdx, m := range spec.Methods {
		for iIdx, inp := range m.Inputs {
			if customTypes[inp.Type] {
				spec.Methods[mIdx].Inputs[iIdx].Type = fmt.Sprintf("%s.%s", packageName, inp.Type)
			}
		}
		for oIdx, out := range m.Outputs {
			if customTypes[out.Type] {
				spec.Methods[mIdx].Outputs[oIdx].Type = fmt.Sprintf("%s.%s", packageName, out.Type)
			}
		}
	}
	return spec

}
