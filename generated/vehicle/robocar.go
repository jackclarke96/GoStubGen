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

