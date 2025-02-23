package vehicle

// Vehicle defines the interface
type Vehicle interface {
	GetTopSpeed() (int)
	Turn(dir string) (string)
	Reverse() (string, error)
	Accelerate(speed int, unit string) (int, error)
	IsMoving() (bool)
	Honk(times int) ()
	GetEngineSpecs() (int, string)
	ApplyBrakes(force float64) (bool)
	ChangeGears(gear int) (int, int)
	Telemetry() (map[string]float64)
	GetPassengers() ([]string)
	LoadCargo(items []string) (int, error)
	GetVehicleStatus() (VehicleStatus)
	UpdateStatus(status VehicleStatus) (error)
}
