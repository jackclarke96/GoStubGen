package vehicle

// Another struct that will act as a concrete type for the interface
type Bike struct {
	kind string
}

// New Bike creates a new instance of Bike with default values
func NewBike() *Bike {
	return &Bike{
		kind: "",
	}
}

// change gears
func (s *Bike) ChangeGears(gear int) (int, int) {
	return 0, 0
}

// get telemetry data in a map
func (s *Bike) Telemetry() map[string]float64 {
	return nil
}

// Increase speed in curent direction
func (s *Bike) Accelerate(speed int, unit string) (int, error) {
	return 0, nil
}

// Updates current status of vehicle.
func (s *Bike) UpdateStatus(status VehicleStatus) error {
	return nil
}

// Turns the vehicle
func (s *Bike) Turn(dir string) string {
	return ""
}

// returns names of passengers currently in/on vehicle
func (s *Bike) GetPassengers() []string {
	return nil
}

// Returns current status of vehicle.
func (s *Bike) GetVehicleStatus() VehicleStatus {
	return VehicleStatus{}
}

// Reverse the vehicle
func (s *Bike) Reverse() (string, error) {
	return "", nil
}

// slows down the vehicle
func (s *Bike) ApplyBrakes(force float64) bool {
	return false
}

// Returns boolean indicating whether vehicle is currently moving
func (s *Bike) IsMoving() bool {
	return false
}

// Loads cargo onto vehicle. Returns remaining capacity of vehicle or error otherwise.
func (s *Bike) LoadCargo(items []string) (int, error) {
	return 0, nil
}

// Returns top speed of vehicle
func (s *Bike) GetTopSpeed() int {
	return 0
}

// Honks the provided number of times
func (s *Bike) Honk(times int) {
}

// Returns data about the engine specs
func (s *Bike) GetEngineSpecs() (int, string) {
	return 0, ""
}
