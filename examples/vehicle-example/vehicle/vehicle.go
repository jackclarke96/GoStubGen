package vehicle

// Vehicle defines the interface
type Vehicle interface {
	// Returns top speed of vehicle
	GetTopSpeed() int
	// Turns the vehicle
	Turn(dir string) string
	// Reverse the vehicle
	Reverse() (string, error)
	// Increase speed in curent direction
	Accelerate(speed int, unit string) (int, error)
	// Returns boolean indicating whether vehicle is currently moving
	IsMoving() bool
	// Honks the provided number of times
	Honk(times int)
	// Returns data about the engine specs
	GetEngineSpecs() (int, string)
	// slows down the vehicle
	ApplyBrakes(force float64) bool
	// change gears
	ChangeGears(gear int) (int, int)
	// get telemetry data in a map
	Telemetry() map[string]float64
	// returns names of passengers currently in/on vehicle
	GetPassengers() []string
	// Loads cargo onto vehicle. Returns remaining capacity of vehicle or error otherwise.
	LoadCargo(items []string) (int, error)
	// Returns current status of vehicle.
	GetVehicleStatus() VehicleStatus
	// Updates current status of vehicle.
	UpdateStatus(status VehicleStatus) error
}
