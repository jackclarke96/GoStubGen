package generator

import (
	"fmt"
)

// maps each interface name to its spec
type iNameToSpec map[string]InterfaceSpec

// maps each struct name to its spec
type sNameToSpec map[string]StructSpec

// maps each method name to the method spec itself
type MethodNameToMethod map[string]Method

// maps each interface name to its method set
type InterfaceToMethods map[string]MethodNameToMethod

// maps each struct to its method set
type StructToMethods map[string]MethodNameToMethod

// Loops through an array of interface or struct YAML defs and maps the spec against its own name
func mapByName[T interface{ getName() string }](items []T) map[string]T {
	m := make(map[string]T)
	for _, item := range items {
		m[item.getName()] = item
	}
	return m
}

// Maps an interface name to its methods, including embedded methods
func (itm InterfaceToMethods) MakeMap(nts iNameToSpec, name string) {
	if _, exists := itm[name]; exists {
		return
	}

	itm[name] = MethodNameToMethod{}

	spec, found := nts[name]
	if !found {
		fmt.Printf("Warning: Interface %s not found\n", name)
		return
	}

	// Add its own methods
	for _, method := range spec.Methods {
		itm[name][method.Name] = method
	}

	// Recursively process embedded interfaces
	for _, embeddedName := range spec.Embedded {
		if _, exists := itm[embeddedName]; !exists {
			itm.MakeMap(nts, embeddedName)
		}
		for methodName, method := range itm[embeddedName] {
			itm[name][methodName] = method
		}
	}
}

// Maps a struct name to its methods, including methods from embedded structs
func (stm StructToMethods) MakeMap(nts sNameToSpec, itm InterfaceToMethods, name string) {
	if _, exists := stm[name]; exists {
		return
	}

	stm[name] = MethodNameToMethod{}

	spec, found := nts[name]
	if !found {
		fmt.Printf("Warning: Struct %s not found\n", name)
		return
	}

	// Add methods from the interfaces it implements
	for _, iFaceName := range spec.Implements {
		if methods, found := itm[iFaceName]; found {
			for methodName, method := range methods {
				stm[name][methodName] = method
			}
		} else {
			fmt.Printf("Warning: Interface %s not found for struct %s\n", iFaceName, name)
		}
	}

	// Recursively process embedded structs
	for _, embeddedName := range spec.Embedded {
		if _, exists := stm[embeddedName]; !exists {
			stm.MakeMap(nts, itm, embeddedName)
		}
		for methodName, method := range stm[embeddedName] {
			stm[name][methodName] = method
		}
	}
}

// takes the map of interface names to their method set and returns unique set of combined methods of provided subset of interfaces
func combineMethods[T ~map[string]Method](methodSets map[string]T, names []string) []Method {
	combined := []Method{}
	methodAlreadyPresent := map[string]bool{}
	for i := range names {
		iFace := names[i]
		if interfaceMethodSet, found := methodSets[iFace]; found {
			for _, method := range interfaceMethodSet {
				if _, alreadyPresent := methodAlreadyPresent[method.Name]; !alreadyPresent {
					combined = append(combined, method)
				}
			}
		} else {
			fmt.Printf("%s interface could not be found in interface set. Is YAML correct?", iFace)
		}
	}
	return combined
}

// gets the set of methods required for an interface MINUS the ones provided by the embedded interfaces
func GetUniqueMethods(ss []StructSpec, is []InterfaceSpec) (map[string][]Method, map[string][]Method) {
	interfaceToUniqueMethods := map[string][]Method{}
	structToUniqueMethods := map[string][]Method{}

	// map each struct to its method set
	stm := StructToMethods{}
	itm := InterfaceToMethods{}

	// map each struct spec against its own name
	sNameToSpec := mapByName(ss)
	iNameToSpec := mapByName(is)

	// hydrate itm with full method set of each interface
	for k := range iNameToSpec {
		itm.MakeMap(iNameToSpec, k)
	}

	// hydrate stm with full method set of each implementing struct
	for k := range sNameToSpec {
		stm.MakeMap(sNameToSpec, itm, k)
	}

	for _, i := range is {
		unique := []Method{}
		// extract full set of methods for that struct

		fullSet := itm[i.Name]

		embeddedMethods := combineMethods(itm, i.Embedded)

		embeddedMethodNames := make(map[string]bool)
		for _, method := range embeddedMethods {
			embeddedMethodNames[method.Name] = true
		}

		for _, method := range fullSet {
			if !embeddedMethodNames[method.Name] {
				unique = append(unique, method)
			}
		}

		interfaceToUniqueMethods[i.Name] = unique
	}

	// for each struct spec
	for _, s := range ss {
		fmt.Printf("\n------------------------ Struct:%s ------------------------\n", s.Name)
		fmt.Println("\n  full method set required by implemented interfaces:\n")
		// init unique set of methods for struct
		unique := []Method{}

		// extract full set of methods for that struct
		fullSet := stm[s.Name]
		for _, meth := range fullSet {
			fmt.Println("  -", meth.Name)
		}

		// get set of methods provided by embedded structs
		embeddedMethods := combineMethods(stm, s.Embedded)

		fmt.Println("\n  method set provided by embedded structs:\n")
		for _, meth := range embeddedMethods {
			fmt.Println("  -", meth.Name)
		}

		// Create a set of embedded method names for quick lookup
		embeddedMethodNames := make(map[string]bool)
		for _, method := range embeddedMethods {
			embeddedMethodNames[method.Name] = true
		}

		// Filter out methods that are in the embedded set
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
