package driver

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
	"github.com/jackclarke/GoStubGen/stubs"
)

// TODO: enable mock for repsponses 1,2 and 3 but use real for 4
// TODO: setResponseTimes (same response three times in a row for example).
// TODO: channels for background stuff
// TODO: can things be more generic?

func TestDriverDriveWithMock(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := newVehicleMock(vehicle.NewCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	// For example, always return the number of items as 42.
	mockVeh.enableLoadCargoMock()
	mockVeh.setLoadCargoResponse(0, errors.New("mock error!"))

	// Call the driver's drive method which uses LoadCargo.
	_, err := d.drive()
	if err == nil {
		t.Fatal("expected an error")
	}
}

// Future Car embeds car so should inherit methods and therefore work with newVehiclemock
func TestDriverDriveWithMockFunc(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := newVehicleMock(vehicle.NewRoboCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	mockVeh.enableLoadCargoMock()
	mockVeh.setLoadCargoFunc(func([]string) (int, error) {
		fmt.Println("mock func being used!")
		return 15, nil
	})
	// Call the driver's drive method which uses LoadCargo.
	_, err := d.drive()
	if err != nil {
		t.Fatalf("Did not expect an error. Got %s", err)
	}
}

func TestDriverDriveWithMultipleResponses(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := newVehicleMock(vehicle.NewRoboCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	// todo: lock queue and clear queue
	mockVeh.enableLoadCargoMock()
	mockVeh.enqueueLoadCargoResponse(10, nil)
	mockVeh.enqueueLoadCargoResponseFunc(func(_ []string) (int, error) {
		return 12, nil
	})
	mockVeh.enqueueLoadCargoResponse(14, nil)

	// Call the driver's drive method which uses LoadCargo.
	resp, err := d.drive()
	if err != nil {
		t.Fatalf("Did not expect an error. Got %s", err)
	}
	if resp != 12 {
		t.Fatalf("Expected 12. Got %v", resp)
	}
}

func TestDriverDriveWithMultipleResponsesAndDelay(t *testing.T) {

	// Create a new mock vehicle.
	mockVeh := newVehicleMock(vehicle.NewRoboCar())

	// Inject the mock vehicle into the Driver.
	d := NewDriver(WithVehicle(mockVeh))

	// Set up the expected behavior for LoadCargo:
	// todo: lock queue and clear queue
	mockVeh.enableLoadCargoMock()
	mockVeh.enqueueLoadCargoResponseWithDelay(16, nil, 5*time.Second)
	mockVeh.enqueueLoadCargoResponseFuncWithDelay(func(_ []string) (int, error) {
		return 18, nil
	}, 5*time.Second)
	mockVeh.enqueueLoadCargoResponse(20, nil)

	// Call the driver's drive method which uses LoadCargo.
	resp, err := d.drive()
	if err != nil {
		t.Fatalf("Did not expect an error. Got %s", err)
	}
	if resp != 18 {
		t.Fatalf("Expected 18. Got %v", resp)
	}
}

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
	expected := []any{[]string{"clothes", "toiletries", "electronics"}, []string{"clothes", "toiletries", "electronics", "more stuff"}}
	for i, call := range calls {
		fmt.Printf("Call #%d: Args = %+v\n", i+1, call.Args)
		if !call.ArgsEqual(expected[i]) {
			t.Errorf("call %d arguments did not match expected.\nGot:  %+v\nWant: %+v", i+1, call.Args, expected)
		}
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

func TestInstructSelfDriver_TriggersDeferredPark(t *testing.T) {
	mock := newSelfDrivingMock(vehicle.NewRoboCar()) // assume realVehicle is a dummy or another mock
	driver := &Driver{vehicle: mock}
	mock.enableParkSelfMock()

	// This will receive the result of ParkSelf (from inside defer)
	parkResult := make(chan error, 1)
	mock.setParkSelfFunc(func() error {
		time.Sleep(10 * time.Millisecond)
		parkResult <- nil
		return nil
	})

	expectedErr := errors.New("Uh oh")

	mock.setParkSelfFunc(func() error {
		return expectedErr
	})

	done := mock.captureParkSelfCallFunc(func() error {
		return expectedErr
	})

	// you can mock DriveSelf too if needed
	mock.enableDriveSelfMock()
	mock.setDriveSelfResponse(nil)

	// call the method
	_, err := driver.instructSelfDriver("garage", "mall")
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	// Wait for deferred ParkSelf to complete
	// got := stubs.WaitForResult(t, parkResult, time.Second)
	got := stubs.WaitForResult(t, done, time.Second)
	if got == nil {
		t.Fatalf("expected error from ParkSelf, got nil")
	}
	if got != expectedErr {
		t.Fatalf("expected %s error from ParkSelf, got %s", expectedErr, got)
	}
}

func TestInstructSelfDriver_TriggersDeferredParkPanic(t *testing.T) {
	mock := newSelfDrivingMock(vehicle.NewRoboCar())
	driver := &Driver{vehicle: mock}
	mock.enableParkSelfMock()

	// todo abstract setup of channel or not
	parkResult := make(chan error, 1)
	mock.setParkSelfFunc(func() error {
		time.Sleep(10 * time.Millisecond)
		parkResult <- nil
		return nil
	})

	// panic!
	mock.enableDriveSelfMock()
	mock.setDriveSelfFunc(func(string) error {
		panic("ahhhh!")
	})

	// Catch panic so test doesn't fail prematurely
	stubs.MustPanic(t, func() {
		_, _ = driver.instructSelfDriver("garage", "mall")
	})

	// Wait for deferred ParkSelf to complete
	got := stubs.WaitForResult(t, parkResult, time.Second)
	if got != nil {
		t.Fatalf("expected no error from ParkSelf, got: %v", got)
	}
}

func (m *mockSelfDriving) captureParkSelfCall(delay time.Duration, resp error) <-chan error {
	ch := make(chan error, 1)
	m.setParkSelfFunc(func() error {
		time.Sleep(delay)
		ch <- resp
		return resp
	})
	return ch
}

func (m *mockSelfDriving) captureParkSelfCallFunc(parkSelfFunc func() error) <-chan error {
	ch := make(chan error, 1)
	m.setParkSelfFunc(func() error {
		err := parkSelfFunc()
		ch <- err
		return err
	})
	return ch
}

/*
plan for void return types/spies: create channel to pass spy results into for deferred/background calls

func (m *{{ .MockName }}) capture{{ title .Name }}CallSpy(timeout time.Duration) []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(m, m.get{{ title .Name }}Calls, timeout)
		ch <- m.get{{ title .Name }}Calls()
	}()
	return ch
}

plan for returned types: Add generic done := mock.captureParkSelfResult() which

type mockSelfDriving struct {
	real   vehicle.SelfDriving
	mocked mockSelfDrivingConfig

	// new: hook for capturing ParkSelf() results
	parkSelfChan chan error
}

func (m *mockSelfDriving) ParkSelf() error {
	m.mocked.ParkSelf.RecordCall()
	var result error
	if m.mocked.ParkSelf.Enabled {
		result = m.mocked.ParkSelf.NextResponse(func() error {
			return m.real.ParkSelf()
		})()
	} else {
		result = m.real.ParkSelf()
	}

	// new: Send result to channel if it's set
	if m.parkSelfChan != nil {
		m.parkSelfChan <- result
	}
	return result
}

func (m *mockSelfDriving) captureParkSelfResult() <-chan error {
	ch := make(chan error, 1)
	m.parkSelfChan = ch
	return ch
}

done:= captureParkSelfResult(expectedResponseTimes, expectedResults) could also work

then just wait for the result:

	got := stubs.WaitForResult(t, done, time.Second)
	or
	got := stubs.WaitForResults(t, done, time.Second)
*/
