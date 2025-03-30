package vehicle

// Car defines the interface
type Car interface {
	// embedded FourWheelVehicle interface
	FourWheelVehicle
	// Opens the car trunk.
	OpenTrunk() (string)
}
