package vehicle

// Bike is the concrete implementation of Vehicle
type Bike struct {
    kind string
}

// NewBike creates a new instance of Bike with default values
func NewBike() *Bike {
	return &Bike{
		kind: "",
	}
}


func (s *Bike) GetTopSpeed() (int) {
	// TODO: Implement
	return 0
}

func (s *Bike) Turn(dir string) (string) {
	// TODO: Implement
	return ""
}

func (s *Bike) Reverse() (string, error) {
	// TODO: Implement
	return "", nil
}

func (s *Bike) Accelerate(speed int, unit string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *Bike) IsMoving() (bool) {
	// TODO: Implement
	return false
}

func (s *Bike) Honk(times int) () {
	// TODO: Implement
}

func (s *Bike) GetEngineSpecs() (int, string) {
	// TODO: Implement
	return 0, ""
}

func (s *Bike) ApplyBrakes(force float64) (bool) {
	// TODO: Implement
	return false
}

func (s *Bike) ChangeGears(gear int) (int, int) {
	// TODO: Implement
	return 0, 0
}

func (s *Bike) Telemetry() (map[string]float64) {
	// TODO: Implement
	return nil
}

func (s *Bike) GetPassengers() ([]string) {
	// TODO: Implement
	return nil
}

func (s *Bike) LoadCargo(items []string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *Bike) GetVehicleStatus() (VehicleStatus) {
	// TODO: Implement
	return VehicleStatus{}
}

func (s *Bike) UpdateStatus(status VehicleStatus) (error) {
	// TODO: Implement
	return nil
}

