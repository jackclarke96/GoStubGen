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


func (s *Car) getTopSpeed() (int) {
	// TODO: Implement
	return 0
}

func (s *Car) turn(dir string) (string) {
	// TODO: Implement
	return ""
}

func (s *Car) reverse() (string, error) {
	// TODO: Implement
	return "", nil
}

func (s *Car) accelerate(speed int, unit string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *Car) isMoving() (bool) {
	// TODO: Implement
	return false
}

func (s *Car) honk(times int) () {
	// TODO: Implement
}

func (s *Car) getEngineSpecs() (int, string) {
	// TODO: Implement
	return 0, ""
}

func (s *Car) applyBrakes(force float64) (bool) {
	// TODO: Implement
	return false
}

func (s *Car) changeGears(gear int) (int, int) {
	// TODO: Implement
	return 0, 0
}

func (s *Car) telemetry() (map[string]float64) {
	// TODO: Implement
	return nil
}

func (s *Car) getPassengers() ([]string) {
	// TODO: Implement
	return nil
}

func (s *Car) loadCargo(items []string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *Car) getVehicleStatus() (VehicleStatus) {
	// TODO: Implement
	return VehicleStatus{}
}

func (s *Car) updateStatus(status VehicleStatus) (error) {
	// TODO: Implement
	return nil
}

