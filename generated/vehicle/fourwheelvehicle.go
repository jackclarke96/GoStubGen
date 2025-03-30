package vehicle

// FourWheelVehicle defines the interface
type FourWheelVehicle interface {
	// Drives the vehicle forward.
	Drive() (string)
	// Stops the vehicle.
	Stop() (string)
}
