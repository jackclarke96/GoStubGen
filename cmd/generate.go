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

		// commonSpec := generator.CommonSpec{
		// 	Package:  config.Package,
		// 	Importer: config.Importer,
		// }

		uniqueInterfaceMethods, uniqueStructMethods := generator.GetUniqueMethods(config.Implementers, config.Interfaces)
		fmt.Printf("UniqueInterfaceMethods %+v\n", uniqueInterfaceMethods)
		fmt.Printf("UniqueStructMethods %+v\n", uniqueStructMethods)

		// iNameToSpec := generator.INameToSpec(config.Interfaces)
		// sNameToSpec := generator.SNameToSpec(config.Implementers)

		// itm := generator.InterfaceToMethods{}
		// stm := generator.StructToMethods{}

		// // 1a: For each interface, get itm gives a map of interfaceName[[]all-methods]
		// for k := range iNameToSpec {
		// 	itm.MakeMap(iNameToSpec, k)
		// }

		// for k, v := range itm {
		// 	fmt.Printf("Interface: %s\n", k)
		// 	for _, method := range v {
		// 		fmt.Printf("  - %s\n", method.Name)
		// 	}
		// 	fmt.Println("-----------------------------")
		// }

		// // 2a: For each struct, get itm gives a map of structName[[]all-methods]
		// for k := range sNameToSpec {
		// 	stm.MakeMap(sNameToSpec, itm, k)
		// }

		// for k, v := range stm {
		// 	fmt.Printf("Struct: %s\n", k)
		// 	for _, method := range v {
		// 		fmt.Printf("  - %s\n", method.Name)
		// 	}
		// 	fmt.Println("-----------------------------")
		// }

		// // 1b: get unique struct methods
		// for _, structSpec := range config.Implementers {
		// 	fmt.Printf("Struct: %s\n", structSpec.Name)
		// 	uniqueMethods := structSpec.GetUniqueMethods(stm)
		// 	for _, method := range uniqueMethods {
		// 		fmt.Printf("  - %s\n", method.Name)
		// 	}
		// }

		// // 2b: get unique interface methods
		// for _, intSpec := range config.Interfaces {
		// 	fmt.Printf("intSpec: %s\n", intSpec.Name)
		// 	uniqueMethods := intSpec.GetUniqueMethods(itm)
		// 	for _, method := range uniqueMethods {
		// 		fmt.Printf("  - %s\n", method.Name)
		// 	}
		// }
		// Add methods to implementers
		// for i := range config.Implementers {
		// 	for _, implemented := range config.Implementers[i].Implements {
		// 		fmt.Println(itm[implemented])
		// 		config.Implementers[i].Methods = append(itm[implemented], config.Implementers[i].Methods...)
		// 	}
		// }

		// // // Print all interfaces and their methods
		// // fmt.Println("\nInterfaces and their Methods:\n")
		// // for _, is := range config.Interfaces {
		// // 	fmt.Printf("Interface: %s\n", is.Name)
		// // 	for _, method := range is.Methods {
		// // 		fmt.Printf("  - %s\n", method.Name)
		// // 	}
		// // 	fmt.Println("-----------------------------")
		// // }

		// // // Print all implementers and their methods
		// // fmt.Println("\nImplementers and their Methods:\n")
		// // for _, implementer := range config.Implementers {
		// // 	fmt.Printf("Implementer: %s (Implements: %v)\n", implementer.Name, implementer.Implements)
		// // 	for _, method := range implementer.Methods {
		// // 		fmt.Printf("  - %s\n", method.Name)
		// // 	}
		// // 	fmt.Println("-----------------------------")
		// // }

		// // Generate custom structs (excluding the ones implementing the interface)
		// if err := generator.GenerateTypesAndStructs(config.CustomStructs, config.CustomTypes, commonSpec); err != nil {
		// 	log.Fatalf("Error generating custom structs: %v", err)
		// }

		// // Generate the interface
		// if err := generator.GenerateInterface(config.Interfaces, commonSpec); err != nil {
		// 	log.Fatalf("Error generating interface: %v", err)
		// }

		// // Generate concrete types structs
		// if err := generator.GenerateConcreteTypes(config.Implementers, commonSpec); err != nil {
		// 	log.Fatalf("Error generating concrete types: %v", err)
		// }

		// // Generate the mocks
		// for i := range config.Interfaces {
		// 	config.Interfaces[i].Methods = itm[config.Interfaces[i].Name]
		// }

		// for _, i := range config.Interfaces {
		// 	mockInterfaceSpec := prefixTypesWithPackageName(config, i, commonSpec.Package)
		// 	if err := generator.GenerateMock(mockInterfaceSpec, generator.StructSpec{}, commonSpec); err != nil {
		// 		log.Fatalf("Error generating mock: %v", err)
		// 	}
		// }

		// fmt.Println("Code generation complete!")
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
