package vehicle

// Car is the concrete implementation of Vehicle
type Car struct {
    doors int
    insured bool
}

// NewCar creates a new instance of Car with default values
func NewCar() *Car {
	return &Car{
		doors: 0,
		insured: false,
	}
}


func (s *Car) GetTopSpeed() (int) {
	// TODO: Implement
	return 0
}

func (s *Car) Turn(dir string) (string) {
	// TODO: Implement
	return ""
}

func (s *Car) Reverse() (string, error) {
	// TODO: Implement
	return "", nil
}

func (s *Car) Accelerate(speed int, unit string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *Car) IsMoving() (bool) {
	// TODO: Implement
	return false
}

func (s *Car) Honk(times int) () {
	// TODO: Implement
}

func (s *Car) GetEngineSpecs() (int, string) {
	// TODO: Implement
	return 0, ""
}

func (s *Car) ApplyBrakes(force float64) (bool) {
	// TODO: Implement
	return false
}

func (s *Car) ChangeGears(gear int) (int, int) {
	// TODO: Implement
	return 0, 0
}

func (s *Car) Telemetry() (map[string]float64) {
	// TODO: Implement
	return nil
}

func (s *Car) GetPassengers() ([]string) {
	// TODO: Implement
	return nil
}

func (s *Car) LoadCargo(items []string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *Car) GetVehicleStatus() (VehicleStatus) {
	// TODO: Implement
	return VehicleStatus{}
}

func (s *Car) UpdateStatus(status VehicleStatus) (error) {
	// TODO: Implement
	return nil
}

