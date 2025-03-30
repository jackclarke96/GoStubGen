package vehicle

// A standard four-door sedan car. 
type Sedan struct {
    color string
    seats int
}

// New Sedan creates a new instance of Sedan with default values
func NewSedan() *Sedan {
	return &Sedan{
		color: "",
		seats: 0,
	}
}

// Opens the car trunk.
func (s *Sedan) OpenTrunk() (string) {
	return ""
}
// Drives the vehicle forward.
func (s *Sedan) Drive() (string) {
	return ""
}
// Stops the vehicle.
func (s *Sedan) Stop() (string) {
	return ""
}

