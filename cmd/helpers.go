package cmd

import (
	"fmt"

	"github.com/jackclarke/GoStubGen/internal/generator"
)

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
