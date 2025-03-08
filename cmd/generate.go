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
	Interfaces     []generator.InterfaceSpec   `yaml:"interfaces"`
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

		// fmt.Println("Custom structs")
		// fmt.Printf("%+v", config.CustomStructs)

		// fmt.Println("Implementers")
		// fmt.Printf("%+v", config.Implementers)

		nameToSpec := NameToSpec(config.Interfaces)

		itm := interfaceToMethods{}
		for k := range nameToSpec {
			itm.makeMap(nameToSpec, k)
		}

		fmt.Println("\nITM\n")
		for k, v := range itm {
			fmt.Printf("Interface: %s\n", k)
			for _, method := range v {
				fmt.Printf("  - %s\n", method.Name)
			}
			fmt.Println("-----------------------------")
		}

		// Add methods to implementers
		for i := range config.Implementers { // ✅ Use index-based iteration
			for _, implemented := range config.Implementers[i].Implements {
				fmt.Println(itm[implemented]) // Debugging output
				config.Implementers[i].Methods = append(itm[implemented], config.Implementers[i].Methods...)
			}
		}

		fmt.Println(config.Implementers)

		// Print all interfaces and their methods
		fmt.Println("\nInterfaces and their Methods:\n")
		for _, is := range config.Interfaces {
			fmt.Printf("Interface: %s\n", is.Name)
			for _, method := range is.Methods {
				fmt.Printf("  - %s\n", method.Name)
			}
			fmt.Println("-----------------------------")
		}

		// Print all implementers and their methods
		fmt.Println("\nImplementers and their Methods:\n")
		for _, implementer := range config.Implementers {
			fmt.Printf("Implementer: %s (Implements: %v)\n", implementer.Name, implementer.Implements)
			for _, method := range implementer.Methods {
				fmt.Printf("  - %s\n", method.Name)
			}
			fmt.Println("-----------------------------")
		}

		// todo the

		// Generate custom structs (excluding the ones implementing the interface)
		if err := generator.GenerateTypesAndStructs(config.CustomStructs, config.CustomTypes, commonSpec); err != nil {
			log.Fatalf("Error generating custom structs: %v", err)
		}

		// Generate the interface
		if err := generator.GenerateInterface(config.Interfaces, commonSpec); err != nil {
			log.Fatalf("Error generating interface: %v", err)
		}

		// Generate concrete types structs
		if err := generator.GenerateConcreteTypes(interfaceSpec, config.Implementers, commonSpec); err != nil {
			log.Fatalf("Error generating concrete types: %v", err)
		}

		// Generate the mocks

		for i := range config.Interfaces {
			fmt.Println("name = ", config.Interfaces[i].Name)
			fmt.Println("methods before = ", config.Interfaces[i].Methods)
			config.Interfaces[i].Methods = itm[config.Interfaces[i].Name]
			fmt.Println("methods after = ", config.Interfaces[i].Methods)
		}

		for _, i := range config.Interfaces {
			mockInterfaceSpec := prefixTypesWithPackageName(config, i, commonSpec.Package)
			if err := generator.GenerateMock(mockInterfaceSpec, generator.StructSpec{}, commonSpec); err != nil {
				log.Fatalf("Error generating mock: %v", err)
			}
		}
		// mockInterfaceSpec := prefixTypesWithPackageName(config, interfaceSpec, commonSpec.Package)
		// fmt.Printf("interface spec %+v\n", interfaceSpec)
		// fmt.Printf("mock interface spec %+v\n", mockInterfaceSpec)

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

type interfaceToMethods map[string][]generator.Method
type nameToSpec map[string]generator.InterfaceSpec

func NameToSpec(s []generator.InterfaceSpec) nameToSpec {
	nts := nameToSpec{}
	for _, i := range s {
		nts[i.Name] = i
	}
	return nts
}

func (itm interfaceToMethods) makeMap(nts nameToSpec, name string) {
	// If the interface has already been fully explored, return early
	if _, exists := itm[name]; exists && len(itm[name]) > 0 {
		return
	}

	// Initialize the slice for the current interface if not set
	itm[name] = []generator.Method{}

	// Get the interface spec
	spec, found := nts[name]
	if !found {
		fmt.Printf("Warning: Interface %s not found\n", name)
		return
	}

	// Add its own methods first
	itm[name] = append(itm[name], spec.Methods...)

	// Recursively process embedded interfaces
	for _, embeddedName := range spec.Embedded {
		// Check if embedded interface has already been processed
		if _, exists := itm[embeddedName]; !exists {
			// Recursive call if not explored
			itm.makeMap(nts, embeddedName)
		}
		// Merge methods only if they were not already added
		itm[name] = append(itm[name], itm[embeddedName]...)
	}
}
