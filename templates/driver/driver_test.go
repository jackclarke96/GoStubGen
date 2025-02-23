package driver

import (
	"errors"
	"testing"
)

func TestDriverDriveWithMock(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := vehicleMock()

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	// For example, always return the number of items as 42.
	mockVeh.enableLoadCargoResponse()
	mockVeh.setLoadCargoResponse(func(items []string) (int, error) {
		return 0, errors.New("mock error!")
	})

	// Call the driver's drive method which uses LoadCargo.
	err := d.drive() // note: you might want to export the method as Drive (capitalized)
	if err == nil {
		t.Fatal("expected an error")
	}

	// Here you could also capture stdout or other side effects to verify behavior.
	// For example, if your drive method printed cargo, you might redirect os.Stdout during the test.
}
