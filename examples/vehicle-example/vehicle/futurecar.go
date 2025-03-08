package vehicle

// Another struct that will act as a concrete type for the SelfDriving interface
type FutureCar struct {
	doors   int
	insured bool
}

// New FutureCar creates a new instance of FutureCar with default values
func NewFutureCar() *FutureCar {
	return &FutureCar{
		doors:   0,
		insured: false,
	}
}

// driverless driving
func (s *FutureCar) DriveSelf(endLocation string) error {
	return nil
}

// Returns top speed of vehicle
func (s *FutureCar) GetTopSpeed() int {
	return 0
}

// Turns the vehicle
func (s *FutureCar) Turn(dir string) string {
	return ""
}

// Reverse the vehicle
func (s *FutureCar) Reverse() (string, error) {
	return "", nil
}

// Increase speed in curent direction
func (s *FutureCar) Accelerate(speed int, unit string) (int, error) {
	return 0, nil
}

// Returns boolean indicating whether vehicle is currently moving
func (s *FutureCar) IsMoving() bool {
	return false
}

// Honks the provided number of times
func (s *FutureCar) Honk(times int) {
}

// Returns data about the engine specs
func (s *FutureCar) GetEngineSpecs() (int, string) {
	return 0, ""
}

// slows down the vehicle
func (s *FutureCar) ApplyBrakes(force float64) bool {
	return false
}

// change gears
func (s *FutureCar) ChangeGears(gear int) (int, int) {
	return 0, 0
}

// get telemetry data in a map
func (s *FutureCar) Telemetry() map[string]float64 {
	return nil
}

// returns names of passengers currently in/on vehicle
func (s *FutureCar) GetPassengers() []string {
	return nil
}

// Loads cargo onto vehicle. Returns remaining capacity of vehicle or error otherwise.
func (s *FutureCar) LoadCargo(items []string) (int, error) {
	return 0, nil
}

// Returns current status of vehicle.
func (s *FutureCar) GetVehicleStatus() VehicleStatus {
	return VehicleStatus{}
}

// Updates current status of vehicle.
func (s *FutureCar) UpdateStatus(status VehicleStatus) error {
	return nil
}
