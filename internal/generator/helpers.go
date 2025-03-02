package generator

import (
	"fmt"
	"strings"
)

func getZeroVal(paramType string) string {
	paramType = strings.TrimSpace(paramType)

	// Predefined defaults for common basic types.
	defaults := map[string]string{
		"int":        "0",
		"int8":       "0",
		"int16":      "0",
		"int32":      "0",
		"int64":      "0",
		"uint":       "0",
		"uint8":      "0",
		"uint16":     "0",
		"uint32":     "0",
		"uint64":     "0",
		"uintptr":    "0",
		"float32":    "0.0",
		"float64":    "0.0",
		"complex64":  "0",
		"complex128": "0",
		"bool":       "false",
		"string":     `""`,
		"error":      "nil",
	}

	if val, ok := defaults[paramType]; ok {
		return val
	}

	// Handle pointer types (e.g. *MyType)
	if strings.HasPrefix(paramType, "*") {
		return "nil"
	}

	// Handle slices (e.g. []int)
	if strings.HasPrefix(paramType, "[]") {
		return "nil"
	}

	// Handle arrays (e.g. [10]int). Arrays are composite types, so we return a literal.
	if len(paramType) > 0 && paramType[0] == '[' && (len(paramType) < 2 || paramType[1] != ']') {
		return fmt.Sprintf("%s{}", paramType)
	}

	// Handle maps (e.g. map[string]int)
	if strings.HasPrefix(paramType, "map[") {
		return "nil"
	}

	// Handle channels (e.g. chan int, <-chan int, chan<- int)
	if strings.HasPrefix(paramType, "chan ") ||
		strings.HasPrefix(paramType, "<-chan") ||
		strings.HasPrefix(paramType, "chan<-") {
		return "nil"
	}

	// Handle function types (e.g. func(int) string)
	if strings.HasPrefix(paramType, "func(") {
		return "nil"
	}

	// Handle interface types (e.g. interface{} or interface{...})
	if strings.HasPrefix(paramType, "interface{") || paramType == "interface{}" {
		return "nil"
	}

	// Fallback: assume custom types, structs, etc. can be initialised via a composite literal.
	if paramType != "" {
		return fmt.Sprintf("%s{}", paramType)
	}

	return "nil"
}

func prefixTypeIfCustom(typeName, pkg string, customTypes map[string]bool) string {
	if customTypes[typeName] {
		return pkg + "." + typeName
	}
	return typeName
}
