package vehicle

// Another struct that will act as a concrete type for the SelfDriving interface
type RoboCar struct {
	Car
	doors   int
	insured bool
}

// New RoboCar creates a new instance of RoboCar with default values
func NewRoboCar() *RoboCar {
	return &RoboCar{
		doors:   0,
		insured: false,
	}
}

// Should be used after driving
func (s *RoboCar) TurnOffAC() error {
	return nil
}

// Should be used after driving
func (s *RoboCar) LockDoors() error {
	return nil
}

// driverless driving
func (s *RoboCar) DriveSelf(endLocation string) error {
	return nil
}

// Should be used after driving
func (s *RoboCar) CloseWindows() error {
	return nil
}

// Should be used after driving
func (s *RoboCar) ParkSelf() error {
	return nil
}

// Should be used after driving
func (s *RoboCar) TurnOffMusic() error {
	return nil
}
