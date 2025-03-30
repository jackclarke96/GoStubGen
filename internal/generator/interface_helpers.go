package generator

import (
	"fmt"
)

// InterfaceNameToSpec maps each interface name to its corresponding InterfaceSpec
type InterfaceNameToSpec map[string]InterfaceSpec

// StructNameToSpec maps each struct name to its corresponding StructSpec
type StructNameToSpec map[string]StructSpec

// MethodNameToMethodMap maps each method name to its Method definition
type MethodNameToMethodMap map[string]Method

// InterfaceNameToMethodMap maps interface names to the full set of methods they require (including embedded interfaces)
type InterfaceNameToMethodMap map[string]MethodNameToMethodMap

// StructNameToMethodMap maps struct names to the full set of methods required by the interfaces they implement (including embedded structs)
type StructNameToMethodMap map[string]MethodNameToMethodMap

// mapSpecsByName takes a slice of specs (struct or interface) and returns a map from name to spec for quick lookup
func mapSpecsByName[T interface{ getName() string }](items []T) map[string]T {
	m := make(map[string]T)
	for _, item := range items {
		m[item.getName()] = item
	}
	return m
}

// BuildMethodSetMap recursively builds the method set for an interface, including methods from embedded interfaces
func (interfaceMethods InterfaceNameToMethodMap) BuildMethodSetMap(interfaceSpecs InterfaceNameToSpec, name string) {
	if _, exists := interfaceMethods[name]; exists {
		return
	}

	interfaceMethods[name] = MethodNameToMethodMap{}

	spec, found := interfaceSpecs[name]
	if !found {
		fmt.Printf("Warning: Interface %s not found\n", name)
		return
	}

	// Add methods declared directly in the interface
	for _, method := range spec.Methods {
		interfaceMethods[name][method.Name] = method
	}

	// Recursively add methods from embedded interfaces
	for _, embeddedName := range spec.Embedded {
		if _, exists := interfaceMethods[embeddedName]; !exists {
			interfaceMethods.BuildMethodSetMap(interfaceSpecs, embeddedName)
		}
		for methodName, method := range interfaceMethods[embeddedName] {
			interfaceMethods[name][methodName] = method
		}
	}
}

// BuildMethodSetMap recursively builds the method set for a struct, including methods from embedded structs and implemented interfaces
func (structMethods StructNameToMethodMap) BuildMethodSetMap(structSpecs StructNameToSpec, interfaceMethods InterfaceNameToMethodMap, name string) {
	if _, exists := structMethods[name]; exists {
		return
	}

	structMethods[name] = MethodNameToMethodMap{}

	spec, found := structSpecs[name]
	if !found {
		fmt.Printf("Warning: Struct %s not found\n", name)
		return
	}

	// Add methods required by implemented interfaces
	for _, interfaceName := range spec.Implements {
		if methods, found := interfaceMethods[interfaceName]; found {
			for methodName, method := range methods {
				structMethods[name][methodName] = method
			}
		} else {
			fmt.Printf("Warning: Interface %s not found for struct %s\n", interfaceName, name)
		}
	}

	// Recursively add methods from embedded structs
	for _, embeddedName := range spec.Embedded {
		if _, exists := structMethods[embeddedName]; !exists {
			structMethods.BuildMethodSetMap(structSpecs, interfaceMethods, embeddedName)
		}
		for methodName, method := range structMethods[embeddedName] {
			structMethods[name][methodName] = method
		}
	}
}

// combineMethods aggregates and deduplicates methods from a list of interface or struct names
func combineMethods[T ~map[string]Method](methodSets map[string]T, names []string) []Method {
	combined := []Method{}
	methodAlreadyPresent := map[string]bool{}
	for _, name := range names {
		if methodSet, found := methodSets[name]; found {
			for _, method := range methodSet {
				if _, alreadyPresent := methodAlreadyPresent[method.Name]; !alreadyPresent {
					combined = append(combined, method)
					methodAlreadyPresent[method.Name] = true
				}
			}
		} else {
			fmt.Printf("%s interface could not be found in interface set. Is YAML correct?", name)
		}
	}
	return combined
}

/*
GetUniqueMethods computes the set of methods that must be implemented by each interface and struct,
excluding those inherited from embedded interfaces or structs.

Algorithm Overview:

1. Convert slices of InterfaceSpec and StructSpec to name-based maps for quick lookup.
2. Build a complete method set for each interface (including embedded interfaces).
3. Build a complete method set for each struct (including methods from embedded structs and implemented interfaces).
4. For each interface, subtract methods inherited from embedded interfaces to find its unique methods.
5. For each struct, subtract methods inherited from embedded structs to find its unique required methods.
*/
func GetUniqueMethods(structSpecs []StructSpec, interfaceSpecs []InterfaceSpec) (map[string][]Method, map[string][]Method) {
	interfaceToUniqueMethods := map[string][]Method{}
	structToUniqueMethods := map[string][]Method{}

	interfaceNameToSpec := mapSpecsByName(interfaceSpecs)
	structNameToSpec := mapSpecsByName(structSpecs)

	interfaceNameToMethods := InterfaceNameToMethodMap{}
	structNameToMethods := StructNameToMethodMap{}

	for name := range interfaceNameToSpec {
		interfaceNameToMethods.BuildMethodSetMap(interfaceNameToSpec, name)
	}

	for name := range structNameToSpec {
		structNameToMethods.BuildMethodSetMap(structNameToSpec, interfaceNameToMethods, name)
	}

	// Compute unique methods for each interface
	for _, iface := range interfaceSpecs {
		fullSet := interfaceNameToMethods[iface.Name]
		embedded := combineMethods(interfaceNameToMethods, iface.Embedded)

		embeddedMethodNames := map[string]bool{}
		for _, method := range embedded {
			embeddedMethodNames[method.Name] = true
		}

		unique := []Method{}
		for _, method := range fullSet {
			if !embeddedMethodNames[method.Name] {
				unique = append(unique, method)
			}
		}
		interfaceToUniqueMethods[iface.Name] = unique
	}

	// Compute unique methods required for each struct (not inherited via embedding)
	for _, s := range structSpecs {
		fmt.Printf("\n------------------------ Struct:%s ------------------------\n", s.Name)
		fmt.Println("\n  full method set required by implemented interfaces:\n")

		fullSet := structNameToMethods[s.Name]
		for _, meth := range fullSet {
			fmt.Println("  -", meth.Name)
		}

		embedded := combineMethods(structNameToMethods, s.Embedded)
		fmt.Println("\n  method set provided by embedded structs:\n")
		for _, meth := range embedded {
			fmt.Println("  -", meth.Name)
		}

		embeddedMethodNames := map[string]bool{}
		for _, method := range embedded {
			embeddedMethodNames[method.Name] = true
		}

		unique := []Method{}
		for _, method := range fullSet {
			if !embeddedMethodNames[method.Name] {
				unique = append(unique, method)
			}
		}

		fmt.Println("\n  unique set required in addition to embedded methods:\n")
		for _, meth := range unique {
			fmt.Println("  -", meth.Name)
		}
		structToUniqueMethods[s.Name] = unique
	}

	return interfaceToUniqueMethods, structToUniqueMethods
}
