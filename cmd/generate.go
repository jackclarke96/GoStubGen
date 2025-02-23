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
	Package        string                 `yaml:"package"`
	Importer       string                 `yaml:"importer"`
	PackageImports []string               `yaml:"package_imports"`
	CustomStructs  []generator.StructSpec `yaml:"custom_structs"`
	Name           string                 `yaml:"name"`
	Methods        []generator.Method     `yaml:"methods"`
	Concrete       string                 `yaml:"concrete"`
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

		// Generate the interface
		interfaceSpec := generator.InterfaceSpec{
			Name:     config.Name,
			Methods:  config.Methods,
			Concrete: config.Concrete,
		}

		commonSpec := generator.CommonSpec{
			Package:  config.Package,
			Importer: config.Importer,
		}

		// Generate custom structs (excluding the one implementing the interface)
		if err := generator.GenerateStructs(config.CustomStructs, commonSpec, config.Concrete); err != nil {
			log.Fatalf("Error generating custom structs: %v", err)
		}

		// generate the interface
		if err := generator.GenerateInterface(interfaceSpec, commonSpec); err != nil {
			log.Fatalf("Error generating interface: %v", err)
		}

		// Find the struct that matches the concrete implementation (Car)
		var structSpec generator.StructSpec
		found := false
		for _, s := range config.CustomStructs {
			if s.Name == config.Concrete {
				structSpec = s
				found = true
				break
			}
		}

		if !found {
			log.Fatalf("Error: No struct definition found for concrete type '%s'", config.Concrete)
		}

		// Generate the struct and methods together in one file
		if err := generator.GenerateConcreteType(interfaceSpec, structSpec, commonSpec); err != nil {
			log.Fatalf("Error generating struct: %v", err)
		}

		// Generate the mock implementation
		customTypes := make(map[string]bool)
		for _, cs := range config.CustomStructs {
			customTypes[cs.Name] = true
		}

		// Prefix types for each custom struct field if needed.
		for mIdx, m := range interfaceSpec.Methods {
			for iIdx, inp := range m.Inputs {
				if customTypes[inp.Type] {
					interfaceSpec.Methods[mIdx].Inputs[iIdx].Type = fmt.Sprintf("%s.%s", commonSpec.Package, inp.Type)
				}
			}
			for oIdx, out := range m.Outputs {
				if customTypes[out.Type] {
					interfaceSpec.Methods[mIdx].Outputs[oIdx].Type = fmt.Sprintf("%s.%s", commonSpec.Package, out.Type)
				}
			}
		}

		if err := generator.GenerateMock(interfaceSpec, structSpec, commonSpec); err != nil {
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
