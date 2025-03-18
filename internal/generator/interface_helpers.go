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

// Loops through an array of interface YAML defs and maps the interface against its own name
func INameToSpec(s []InterfaceSpec) iNameToSpec {
	nts := iNameToSpec{}
	for _, i := range s {
		nts[i.Name] = i
	}
	return nts
}

// Loops through an array of struct YAML defs and maps the struct against its own name
func SNameToSpec(s []StructSpec) sNameToSpec {
	nts := sNameToSpec{}
	for _, v := range s {
		nts[v.Name] = v
	}
	return nts
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
func (itm InterfaceToMethods) combineMethods(interfaces []string) []Method {
	combined := []Method{}
	methodAlreadyPresent := map[string]bool{}
	for i := range interfaces {
		iFace := interfaces[i]
		if interfaceMethodSet, found := itm[iFace]; found {
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

// takes the map of struct names to their method set and returns a unique set of combined methods of provided subset of structs
func (stm StructToMethods) combineMethods(structs []string) []Method {
	combined := []Method{}
	methodAlreadyPresent := map[string]bool{}
	for _, structName := range structs {
		if structMethodSet, found := stm[structName]; found {
			for _, method := range structMethodSet {
				if !methodAlreadyPresent[method.Name] {
					combined = append(combined, method)
					methodAlreadyPresent[method.Name] = true
				}
			}
		} else {
			fmt.Printf("Warning: Struct %s could not be found in struct set. Is YAML correct?\n", structName)
		}
	}
	return combined
}

// gets the set of methods required for an interface MINUS the ones provided by the embedded interfaces
func (i InterfaceSpec) GetUniqueMethods(itm InterfaceToMethods) []Method {
	fullSet := itm[i.Name]
	embeddedMethods := itm.combineMethods(i.Embedded)
	unique := []Method{}

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
	return unique
}

// gets the set of methods required for an interface MINUS the ones provided by the embedded interfaces
func GetUniqueMethods(ss []StructSpec, is []InterfaceSpec) (map[string][]Method, map[string][]Method) {
	interfaceToUniqueMethods := map[string][]Method{}
	structToUniqueMethods := map[string][]Method{}

	// map each struct to its method set
	stm := StructToMethods{}
	itm := InterfaceToMethods{}

	// map each struct spec against its own name
	sNameToSpec := SNameToSpec(ss)
	iNameToSpec := INameToSpec(is)

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

		embeddedMethods := itm.combineMethods(i.Embedded)

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
		fmt.Printf("Struct: %s\n", s.Name)
		// init unique set of methods for struct
		unique := []Method{}

		// extract full set of methods for that struct
		fullSet := stm[s.Name]

		// get set of methods provided by embedded structs
		embeddedMethods := stm.combineMethods(s.Embedded)

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
		structToUniqueMethods[s.Name] = unique

	}

	return interfaceToUniqueMethods, structToUniqueMethods
}

// so at this point for any given combination of interfaces.
// next, we should repeat for structs. For each struct, we can extract its "implements" to get its full method set. Just use combineMethods again to get the combo
// then for each embedded struct, get its "implements" and remove it's method set.

// need a function such that if a struct embeds another struct, then can tell which extra methods need to be added in order to satisfy an interface...

// so, collect the set of all methods of all methods the struct must have as receiver, then work through the structs it embeds and remove methods.
