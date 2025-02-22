package config

var PackageImports = []string{
	"fmt",
	"errors",
	"math",
}

type Car struct {
	doors   int
	insured bool
}

// interface
type Vehicle interface {
	getTopSpeed() int                                            // No params, single return
	turn(dir string) (newDir string)                             // One param, named return
	reverse() (string, error)                                    // No params, multiple returns
	accelerate(speed int, unit string) (newSpeed int, err error) // Multiple params, named returns
	isMoving() bool                                              // Boolean return
	honk(times int)                                              // Takes param, no return
	getEngineSpecs() (power int, fuelType string)                // Named multi-value return
	applyBrakes(force float64) bool                              // Float input, boolean return
	changeGears(gear int) (prevGear, newGear int)                // Takes param, two return values
	telemetry() map[string]float64                               // Returns a map
	getPassengers() []string                                     // Returns a slice
	loadCargo(items []string) (remainingCapacity int, err error) // Takes slice, returns values
	getVehicleStatus() VehicleStatus                             // Returns a struct
	updateStatus(status VehicleStatus) error                     // Takes a struct as input
}

// Custom Nested Type
type VehicleStatus struct {
	Speed     int
	Direction string
	EngineOn  bool
}

type car struct {
	surnames []string
	age      int
}
