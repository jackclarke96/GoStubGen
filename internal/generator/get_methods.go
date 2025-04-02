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

// each entry maps an interface or struct name to its method set
type MethodSets struct {
	// full set of methods a struct must have to implement interfaces specified in yaml
	FullSets map[string][]Method
	// full set of methods a struct must have to implement interfaces specified in yaml minus methods provided by embedded structs
	UniqueSets map[string][]Method
	// methods provided by embedded structs
	EmbeddedSets map[string][]Method
}

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

	// Recursively add methods from embedded interfaces. Todo MAKE OPTIONAL
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

	// Recursively add methods from embedded structs. Could probably remove this, though might be useful later. Todo MAKE OPTIONAL. Useful for getting whole set
	for _, embeddedName := range spec.Embedded {
		if _, exists := structMethods[embeddedName]; !exists {
			structMethods.BuildMethodSetMap(structSpecs, interfaceMethods, embeddedName)
		}
		for methodName, method := range structMethods[embeddedName] {
			structMethods[name][methodName] = method
		}
	}
}

// mergeMethodMaps merges multiple method maps into a deduplicated slice
func mergeMethodMaps[T ~map[string]Method](sets ...T) []Method {
	result := []Method{}
	seen := map[string]bool{}

	for _, set := range sets {
		for _, method := range set {
			if !seen[method.Name] {
				result = append(result, method)
				seen[method.Name] = true
			}
		}
	}
	return result
}

// mergeMethodMapsByProviderName finds the unique set of methods of a combination of structs or interfaces.
// Aggregates and deduplicates methods from a list of interface or struct names.
func mergeMethodMapsByProviderName[T ~map[string]Method](methodSets map[string]T, names []string) []Method {
	collected := []T{}
	for _, name := range names {
		if set, ok := methodSets[name]; ok {
			collected = append(collected, set)
		} else {
			fmt.Printf("Warning: %s not found in method set. Is YAML correct?\n", name)
		}
	}
	return mergeMethodMaps(collected...)
}

// subtractMethods takes a map of methodNames to Method objects and removes methods provided in method slice
func subtractMethods(full map[string]Method, toRemove []Method) []Method {
	removeSet := map[string]bool{}
	for _, method := range toRemove {
		removeSet[method.Name] = true
	}

	unique := []Method{}
	for name, method := range full {
		if !removeSet[name] {
			unique = append(unique, method)
		}
	}
	return unique
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

func GetMethods(structSpecs []StructSpec, interfaceSpecs []InterfaceSpec) (MethodSets, MethodSets) {

	// init map mapping interface names to their method sets
	interfaceNameToMethods := InterfaceNameToMethodMap{}
	// init map mapping struct names to their method sets
	structNameToMethods := StructNameToMethodMap{}

	// create map mapping interface names to their full spec
	interfaceNameToSpec := mapSpecsByName(interfaceSpecs)
	// Hydrate interfaceNameToMethods so that each interface name is mapped against its full method set
	for name := range interfaceNameToSpec {
		interfaceNameToMethods.BuildMethodSetMap(interfaceNameToSpec, name)
	}

	// create map mapping struct names to their full spec
	structNameToSpec := mapSpecsByName(structSpecs)
	// Hydrate structNameToMethods so that each struct name is mapped against its full method set
	for name := range structNameToSpec {
		structNameToMethods.BuildMethodSetMap(structNameToSpec, interfaceNameToMethods, name)
	}

	// Compute unique methods for each interface by taking union of methods across interface and any embedded interfaces, then removing methods already present due to embedded interfaces
	interfaceMethodSets := MethodSets{
		FullSets:     map[string][]Method{},
		UniqueSets:   map[string][]Method{},
		EmbeddedSets: map[string][]Method{},
	}
	for _, iface := range interfaceSpecs {
		fullSet := interfaceNameToMethods[iface.Name]
		embedded := mergeMethodMapsByProviderName(interfaceNameToMethods, iface.Embedded)
		unique := subtractMethods(fullSet, embedded)

		interfaceMethodSets.FullSets[iface.Name] = mapsToSlice(fullSet)
		interfaceMethodSets.UniqueSets[iface.Name] = unique
		interfaceMethodSets.EmbeddedSets[iface.Name] = embedded
	}

	// Compute unique methods for each struct by taking union of methods across struct and any embedded structs, then removing methods already present due to embedded structs
	structMethodSets := MethodSets{
		FullSets:     map[string][]Method{},
		UniqueSets:   map[string][]Method{},
		EmbeddedSets: map[string][]Method{},
	}

	for _, s := range structSpecs {
		fullSet := structNameToMethods[s.Name]
		embedded := mergeMethodMapsByProviderName(structNameToMethods, s.Embedded)
		unique := subtractMethods(fullSet, embedded)

		structMethodSets.FullSets[s.Name] = mapsToSlice(fullSet)
		structMethodSets.UniqueSets[s.Name] = unique
		structMethodSets.EmbeddedSets[s.Name] = embedded
	}

	return interfaceMethodSets, structMethodSets
}

func mapsToSlice(m map[string]Method) []Method {
	methods := make([]Method, 0, len(m))
	for _, method := range m {
		methods = append(methods, method)
	}
	return methods
}
