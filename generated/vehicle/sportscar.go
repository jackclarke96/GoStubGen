package vehicle

// A high-performance sports car. 
type SportsCar struct {
    topSpeed int
}

// New SportsCar creates a new instance of SportsCar with default values
func NewSportsCar() *SportsCar {
	return &SportsCar{
		topSpeed: 0,
	}
}

// Drives the vehicle forward.
func (s *SportsCar) Drive() (string) {
	return ""
}
// Stops the vehicle.
func (s *SportsCar) Stop() (string) {
	return ""
}
// Opens the car trunk.
func (s *SportsCar) OpenTrunk() (string) {
	return ""
}

