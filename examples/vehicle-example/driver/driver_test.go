package driver

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
)

// TODO: enable mock for repsponses 1,2 and 3 but use real for 4
// TODO: spy
// TODO: setResponseTimes (same response three times in a row for example).
// TODO: channels for background stuff
// TODO: can things be more generic?

// func TestDriverDriveWithMock(t *testing.T) {

// 	// Create a new mock vehicle.
// 	mockVeh := newVehicleMock(vehicle.NewCar())

// 	// Inject the mock vehicle into the Driver.
// 	d := NewDriver(WithVehicle(mockVeh))

// 	// Set up the expected behavior for LoadCargo:
// 	// For example, always return the number of items as 42.
// 	mockVeh.enableLoadCargoMock()
// 	mockVeh.setLoadCargoResponse(0, errors.New("mock error!"))

// 	// Call the driver's drive method which uses LoadCargo.
// 	_, err := d.drive()
// 	if err == nil {
// 		t.Fatal("expected an error")
// 	}
// }

// // Future Car embeds car so should inherit methods and therefore work with newVehiclemock
// func TestDriverDriveWithMockFunc(t *testing.T) {

// 	// Create a new mock vehicle.
// 	mockVeh := newVehicleMock(vehicle.NewRoboCar())

// 	// Inject the mock vehicle into the Driver.
// 	d := NewDriver(WithVehicle(mockVeh))

// 	// Set up the expected behavior for LoadCargo:
// 	mockVeh.enableLoadCargoMock()
// 	mockVeh.setLoadCargoFunc(func([]string) (int, error) {
// 		fmt.Println("mock func being used!")
// 		return 15, nil
// 	})
// 	// Call the driver's drive method which uses LoadCargo.
// 	_, err := d.drive()
// 	if err != nil {
// 		t.Fatalf("Did not expect an error. Got %s", err)
// 	}
// }

// func TestDriverDriveWithMultipleResponses(t *testing.T) {

// 	// Create a new mock vehicle.
// 	mockVeh := newVehicleMock(vehicle.NewRoboCar())

// 	// Inject the mock vehicle into the Driver.
// 	d := NewDriver(WithVehicle(mockVeh))

// 	// Set up the expected behavior for LoadCargo:
// 	// todo: lock queue and clear queue
// 	mockVeh.enableLoadCargoMock()
// 	mockVeh.enqueueLoadCargoResponse(10, nil)
// 	mockVeh.enqueueLoadCargoResponseFunc(func(_ []string) (int, error) {
// 		return 12, nil
// 	})
// 	mockVeh.enqueueLoadCargoResponse(14, nil)

// 	// Call the driver's drive method which uses LoadCargo.
// 	resp, err := d.drive()
// 	if err != nil {
// 		t.Fatalf("Did not expect an error. Got %s", err)
// 	}
// 	if resp != 12 {
// 		t.Fatalf("Expected 12. Got %v", resp)
// 	}
// }

// func TestDriverDriveWithMultipleResponsesAndDelay(t *testing.T) {

// 	// Create a new mock vehicle.
// 	mockVeh := newVehicleMock(vehicle.NewRoboCar())

// 	// Inject the mock vehicle into the Driver.
// 	d := NewDriver(WithVehicle(mockVeh))

// 	// Set up the expected behavior for LoadCargo:
// 	// todo: lock queue and clear queue
// 	mockVeh.enableLoadCargoMock()
// 	mockVeh.enqueueLoadCargoResponseWithDelay(16, nil, 5*time.Second)
// 	mockVeh.enqueueLoadCargoResponseFuncWithDelay(func(_ []string) (int, error) {
// 		fmt.Println("delay!!!!")
// 		return 18, nil
// 	}, 5*time.Second)
// 	mockVeh.enqueueLoadCargoResponse(20, nil)

// 	// Call the driver's drive method which uses LoadCargo.
// 	resp, err := d.drive()
// 	if err != nil {
// 		t.Fatalf("Did not expect an error. Got %s", err)
// 	}
// 	if resp != 12 {
// 		t.Fatalf("Expected 12. Got %v", resp)
// 	}
// }

func TestDriverDrive_spyLoadCargo(t *testing.T) {
	realCar := vehicle.NewCar()
	mock := newVehicleMock(realCar)

	// Enable spying but not mocking
	mock.enableLoadCargoSpy()

	driver := NewDriver(WithVehicle(mock))

	_, err := driver.drive()
	if err != nil {
		t.Fatalf("unexpected error from driver.drive(): %v", err)
	}

	// Validate number of calls
	calls := mock.getLoadCargoCalls()
	if len(calls) != 2 {
		t.Fatalf("expected 2 calls to LoadCargo, got %d", len(calls))
	}

	// Validate arguments
	// expected := []any{[]string{"clothes", "toiletries", "electronics"}}
	for i, call := range calls {
		fmt.Printf("Call #%d: Args = %+v\n", i+1, call.Args)
		// if !call.ArgsEqual(expected...) {
		// 	t.Errorf("call %d arguments did not match expected.\nGot:  %+v\nWant: %+v", i+1, call.Args, expected)
		// }
	}
}

// Future Car embeds car so should inherit methods and therefore work with newVehiclemock
func TestSelfDriverMethod(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := newSelfDrivingMock(vehicle.NewRoboCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for DriveSelf:
	mockVeh.enableDriveSelfMock()
	mockVeh.setDriveSelfResponse(errors.New("mock error oh no!"))

	// Call the driver's drive method which uses DriveSelf.
	err := d.setVehicleDriveSelf()
	if err == nil {
		t.Fatal("expected an error!")
	}

	// disable and test again
	mockVeh.disableDriveSelfMock()
	err = d.setVehicleDriveSelf()
	if err != nil {
		t.Fatalf("did not expect an error! got %s", err)
	}

}
