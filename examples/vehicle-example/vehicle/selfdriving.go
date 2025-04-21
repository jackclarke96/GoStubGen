package vehicle

// SelfDriving defines the interface
type SelfDriving interface {
	// embedded Vehicle interface
	Vehicle
	// Should be used after driving
	TurnOffAC() error
	// Should be used after driving
	ParkSelf() error
	// Should be used after driving
	TurnOffMusic() error
	// driverless driving
	DriveSelf(endLocation string) error
	// Should be used after driving
	LockDoors() error
	// Should be used after driving
	CloseWindows() error
}
