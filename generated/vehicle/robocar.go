package vehicle

// A self-driving robotic car. 
type RoboCar struct {
	Sedan
    aiVersion string
}

// New RoboCar creates a new instance of RoboCar with default values
func NewRoboCar() *RoboCar {
	return &RoboCar{
		aiVersion: "",
	}
}

// Activates self-driving mode.
func (s *RoboCar) ActivateAutopilot() (string) {
	return ""
}
// Drives the vehicle forward.
func (s *RoboCar) Drive() (string) {
	return ""
}
// Opens the car trunk.
func (s *RoboCar) OpenTrunk() (string) {
	return ""
}
// Stops the vehicle.
func (s *RoboCar) Stop() (string) {
	return ""
}

