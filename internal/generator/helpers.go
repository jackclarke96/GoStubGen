package generator

import "fmt"

// TODO improve this
func getDefaultReturnValue(paramType string) string {
	switch paramType {
	case "int", "int32", "int64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "false"
	case "string":
		return `""`
	case "error":
		return "nil"
	case "map[string]float64":
		return "map[string]float64{}"
	case "[]string":
		return "[]string{}"
	default:
		if paramType != "" {
			return fmt.Sprintf("%s{}", paramType) // Assume custom struct initialization
		}
		return "nil"
	}
}
