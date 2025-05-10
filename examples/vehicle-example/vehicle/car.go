package vehicle

// A struct that will act as a concrete type for the interface
type Car struct {
	doors   int
	insured bool
}

// New Car creates a new instance of Car with default values
func NewCar() *Car {
	return &Car{
		doors:   0,
		insured: false,
	}
}

// Returns boolean indicating whether vehicle is currently moving
func (s *Car) IsMoving() bool {
	return false
}

// Reverse the vehicle
func (s *Car) Reverse() (string, error) {
	return "", nil
}

// returns names of passengers currently in/on vehicle
func (s *Car) GetPassengers() []string {
	return nil
}

// Turns the vehicle
func (s *Car) Turn(dir string) string {
	return ""
}

// slows down the vehicle
func (s *Car) ApplyBrakes(force float64) bool {
	return false
}

// change gears
func (s *Car) ChangeGears(gear int) (int, int) {
	return 0, 0
}

// Loads cargo onto vehicle. Returns remaining capacity of vehicle or error otherwise.
func (s *Car) LoadCargo(items []string) (int, error) {
	return 0, nil
}

// Honks the provided number of times
func (s *Car) Honk(times int) {
}

// Increase speed in curent direction
func (s *Car) Accelerate(speed int, unit string) (int, error) {
	return 0, nil
}

// Updates current status of vehicle.
func (s *Car) UpdateStatus(status VehicleStatus) error {
	return nil
}

// Returns top speed of vehicle
func (s *Car) GetTopSpeed() int {
	return 0
}

// Returns current status of vehicle.
func (s *Car) GetVehicleStatus() VehicleStatus {
	return VehicleStatus{}
}

// get telemetry data in a map
func (s *Car) Telemetry() map[string]float64 {
	return nil
}

// Returns data about the engine specs
func (s *Car) GetEngineSpecs() (int, string) {
	return 0, ""
}
