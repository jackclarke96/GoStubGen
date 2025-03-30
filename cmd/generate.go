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
	Package       string                      `yaml:"package"`
	Importer      string                      `yaml:"importer"`
	CustomStructs []generator.StructSpec      `yaml:"custom_structs"`
	CustomTypes   []generator.CustomTypesSpec `yaml:"custom_types"`
	Implementers  []generator.StructSpec      `yaml:"implementers"`
	Interfaces    []generator.InterfaceSpec   `yaml:"interfaces"`
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

		commonSpec := generator.CommonSpec{
			Package:  config.Package,
			Importer: config.Importer,
		}

		// get unique methods of each interface and struct (with methods provided by embedded interfaces/structs removed)
		uniqueInterfaceMethods, uniqueStructMethods := generator.GetUniqueMethods(config.Implementers, config.Interfaces)

		// Update interfaces and implementers to have unique methods calculated in previous step
		for i := range config.Interfaces {
			config.Interfaces[i].Methods = uniqueInterfaceMethods[config.Interfaces[i].Name]
		}

		// Update implementing structs to only have minimal set of methods required to satisfy the method set
		for i := range config.Implementers {
			config.Implementers[i].Methods = uniqueStructMethods[config.Implementers[i].Name]
		}

		// Generate custom structs (excluding the ones implementing the interface).
		if err := generator.GenerateTypesAndStructs(config.CustomStructs, config.CustomTypes, commonSpec); err != nil {
			log.Fatalf("Error generating custom structs: %v", err)
		}

		// Generate interfaces.
		if err := generator.GenerateInterfaces(config.Interfaces, commonSpec); err != nil {
			log.Fatalf("Error generating interface: %v", err)
		}

		// Generate implementer structs.
		if err := generator.GenerateConcreteTypes(config.Implementers, commonSpec); err != nil {
			log.Fatalf("Error generating concrete types: %v", err)
		}

		// Generate mocks.
		for _, i := range config.Interfaces {
			mockInterfaceSpec := prefixTypesWithPackageName(config, i, commonSpec.Package)
			if err := generator.GenerateMock(mockInterfaceSpec, generator.StructSpec{}, commonSpec); err != nil {
				log.Fatalf("Error generating mock: %v", err)
			}
		}

		fmt.Println("Code generation complete!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to YAML config file")
	generateCmd.MarkFlagRequired("config")
}
