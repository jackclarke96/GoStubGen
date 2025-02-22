package vehicle

// Vehicle defines the interface
type Vehicle interface {
	getTopSpeed() int
	turn(dir string) string
	reverse() (string, error)
	accelerate(speed int, unit string) (int, error)
	isMoving() bool
	honk(times int)
	getEngineSpecs() (int, string)
	applyBrakes(force float64) bool
	changeGears(gear int) (int, int)
	telemetry() map[string]float64
	getPassengers() []string
	loadCargo(items []string) (int, error)
	getVehicleStatus() VehicleStatus
	updateStatus(status VehicleStatus) error
}
