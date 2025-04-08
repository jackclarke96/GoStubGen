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

// driverless driving
func (s *RoboCar) DriveSelf(endLocation string) error {
	return nil
}
