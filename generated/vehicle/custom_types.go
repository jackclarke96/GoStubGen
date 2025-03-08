package vehicle

// A mapping of string keys to integer values
type MyMapType map[string]int

// Alias for a standard string
type StringAlias string

// An example struct
type VehicleStatus struct {
	// The speed at which the vehicle is currently travelling
    Speed int
	// The direction in which the vehicle is currently travelling
    Direction string
	// Boolean indicating whether an engine is turned on
    EngineOn bool
}
