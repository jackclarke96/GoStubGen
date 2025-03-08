package vehicle

// SelfDriving defines the interface
type SelfDriving interface {
	// embedded Vehicle interface
	Vehicle
	// driverless driving
	DriveSelf(endLocation string) (error)
}
