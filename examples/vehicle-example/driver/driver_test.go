package driver

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
)

func TestDriverDriveWithMock(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := vehicleMock(vehicle.NewCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	// For example, always return the number of items as 42.
	mockVeh.enableLoadCargoMock()
	mockVeh.setLoadCargoResponse(0, errors.New("mock error!"))

	// Call the driver's drive method which uses LoadCargo.
	err := d.drive()
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestDriverDriveWithMockFunc(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := vehicleMock(vehicle.NewCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	mockVeh.enableLoadCargoMock()
	mockVeh.setLoadCargoFunc(func([]string) (int, error) {
		fmt.Println("mock func being used!")
		return 15, nil
	})
	// Call the driver's drive method which uses LoadCargo.
	err := d.drive()
	if err != nil {
		t.Fatalf("Did not expect an error. Got %s", err)
	}
}
