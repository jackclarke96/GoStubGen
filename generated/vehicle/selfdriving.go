package vehicle

// SelfDriving defines the interface
type SelfDriving interface {
	// embedded Car interface
	Car
	// Activates self-driving mode.
	ActivateAutopilot() (string)
}
