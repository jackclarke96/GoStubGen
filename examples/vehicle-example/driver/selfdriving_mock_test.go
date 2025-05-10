package driver

import (
	"testing"
	"time"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
	"github.com/jackclarke/GoStubGen/stubs"
)

// mockSelfDrivingConfig stores mock flags and responses
type mockSelfDrivingConfig struct {
	UpdateStatus     stubs.MethodConfig[func(vehicle.VehicleStatus) error]
	LockDoors        stubs.MethodConfig[func() error]
	GetEngineSpecs   stubs.MethodConfig[func() (int, string)]
	ApplyBrakes      stubs.MethodConfig[func(float64) bool]
	GetTopSpeed      stubs.MethodConfig[func() int]
	ParkSelf         stubs.MethodConfig[func() error]
	Honk             stubs.MethodConfig[func(int)]
	LoadCargo        stubs.MethodConfig[func([]string) (int, error)]
	GetVehicleStatus stubs.MethodConfig[func() vehicle.VehicleStatus]
	TurnOffAC        stubs.MethodConfig[func() error]
	TurnOffMusic     stubs.MethodConfig[func() error]
	CloseWindows     stubs.MethodConfig[func() error]
	Reverse          stubs.MethodConfig[func() (string, error)]
	IsMoving         stubs.MethodConfig[func() bool]
	ChangeGears      stubs.MethodConfig[func(int) (int, int)]
	Telemetry        stubs.MethodConfig[func() map[string]float64]
	Accelerate       stubs.MethodConfig[func(int, string) (int, error)]
	DriveSelf        stubs.MethodConfig[func(string) error]
	Turn             stubs.MethodConfig[func(string) string]
	GetPassengers    stubs.MethodConfig[func() []string]
}

// mockSelfDriving embeds a concrete SelfDriving and its mocks
type mockSelfDriving struct {
	real          vehicle.SelfDriving
	mocked        mockSelfDrivingConfig
	responseChans map[string]any
}

// newSelfDrivingMock returns a new mock
func newSelfDrivingMock(v vehicle.SelfDriving) *mockSelfDriving {
	return &mockSelfDriving{
		real:          v,
		mocked:        mockSelfDrivingConfig{},
		responseChans: make(map[string]any),
	}
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// enableUpdateStatusMock turns the spy on
func (m *mockSelfDriving) enableUpdateStatusSpy() {
	m.mocked.UpdateStatus.SpyEnabled = true
}

// getUpdateStatusCalls returns recorded calls to UpdateStatus
func (m *mockSelfDriving) getUpdateStatusCalls() []stubs.MethodCall {
	return m.mocked.UpdateStatus.Calls()
}

// enableUpdateStatusMock turns the spy off
func (m *mockSelfDriving) disableUpdateStatusSpy() {
	m.mocked.UpdateStatus.SpyEnabled = false
}

// UpdateStatus overrides the method to return the mock response
func (m *mockSelfDriving) UpdateStatus(status vehicle.VehicleStatus) error {
	m.mocked.UpdateStatus.RecordCall(status)
	var (
		out0 error
	)

	if m.mocked.UpdateStatus.Enabled {
		out0 = m.mocked.UpdateStatus.NextResponse(func(status vehicle.VehicleStatus) error {
			return m.real.UpdateStatus(status)
		})(status)

	} else {
		out0 = m.real.UpdateStatus(status)

	}

	if ch, ok := m.responseChans["UpdateStatus"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setUpdateStatusFunc sets the function for UpdateStatus
func (m *mockSelfDriving) setUpdateStatusFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.Fallback = f
}

// enableUpdateStatusMock turns the mock on
func (m *mockSelfDriving) enableUpdateStatusMock() {
	m.mocked.UpdateStatus.Enabled = true
}

// disableUpdateStatusMock turns the mock off
func (m *mockSelfDriving) disableUpdateStatusMock() {
	m.mocked.UpdateStatus.Enabled = false
}

// enqueueUpdateStatusResponseFunc enqueues a function response for UpdateStatus
func (m *mockSelfDriving) enqueueUpdateStatusResponseFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.EnqueueWithDelay(f, 0)
}

// enqueueUpdateStatusResponseFuncWithDelay enqueues a function response with delay for UpdateStatus
func (m *mockSelfDriving) enqueueUpdateStatusResponseFuncWithDelay(f func(vehicle.VehicleStatus) error, d time.Duration) {
	m.mocked.UpdateStatus.EnqueueWithDelay(f, d)
}

// captureUpdateStatusResult sets up a channel to capture UpdateStatus results.
func (m *mockSelfDriving) captureUpdateStatusResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["UpdateStatus"] = ch
	return ch
}

// captureUpdateStatusCallSpy starts watching for UpdateStatus spy calls and sends them into a channel.
func (m *mockSelfDriving) captureUpdateStatusCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getUpdateStatusCalls, timeout)
		ch <- m.getUpdateStatusCalls()
	}()
	return ch
}

type mockSelfDrivingUpdateStatusResult struct {
	Output0 error
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockSelfDriving) setUpdateStatusResponse(output0 error) {
	m.setUpdateStatusFunc(func(vehicle.VehicleStatus) error {
		return output0
	})
}

// enqueueUpdateStatusResponse enqueues a static response for UpdateStatus
func (m *mockSelfDriving) enqueueUpdateStatusResponse(output0 error) {
	m.mocked.UpdateStatus.EnqueueWithDelay(func(vehicle.VehicleStatus) error {
		return output0
	}, 0)
}

// enqueueUpdateStatusResponseWithDelay enqueues a static response with delay for UpdateStatus
func (m *mockSelfDriving) enqueueUpdateStatusResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.UpdateStatus.EnqueueWithDelay(func(vehicle.VehicleStatus) error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- LockDoors Mock Helpers --------------------------- */

// enableLockDoorsMock turns the spy on
func (m *mockSelfDriving) enableLockDoorsSpy() {
	m.mocked.LockDoors.SpyEnabled = true
}

// getLockDoorsCalls returns recorded calls to LockDoors
func (m *mockSelfDriving) getLockDoorsCalls() []stubs.MethodCall {
	return m.mocked.LockDoors.Calls()
}

// enableLockDoorsMock turns the spy off
func (m *mockSelfDriving) disableLockDoorsSpy() {
	m.mocked.LockDoors.SpyEnabled = false
}

// LockDoors overrides the method to return the mock response
func (m *mockSelfDriving) LockDoors() error {
	m.mocked.LockDoors.RecordCall()
	var (
		out0 error
	)

	if m.mocked.LockDoors.Enabled {
		out0 = m.mocked.LockDoors.NextResponse(func() error {
			return m.real.LockDoors()
		})()

	} else {
		out0 = m.real.LockDoors()

	}

	if ch, ok := m.responseChans["LockDoors"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setLockDoorsFunc sets the function for LockDoors
func (m *mockSelfDriving) setLockDoorsFunc(f func() error) {
	m.mocked.LockDoors.Fallback = f
}

// enableLockDoorsMock turns the mock on
func (m *mockSelfDriving) enableLockDoorsMock() {
	m.mocked.LockDoors.Enabled = true
}

// disableLockDoorsMock turns the mock off
func (m *mockSelfDriving) disableLockDoorsMock() {
	m.mocked.LockDoors.Enabled = false
}

// enqueueLockDoorsResponseFunc enqueues a function response for LockDoors
func (m *mockSelfDriving) enqueueLockDoorsResponseFunc(f func() error) {
	m.mocked.LockDoors.EnqueueWithDelay(f, 0)
}

// enqueueLockDoorsResponseFuncWithDelay enqueues a function response with delay for LockDoors
func (m *mockSelfDriving) enqueueLockDoorsResponseFuncWithDelay(f func() error, d time.Duration) {
	m.mocked.LockDoors.EnqueueWithDelay(f, d)
}

// captureLockDoorsResult sets up a channel to capture LockDoors results.
func (m *mockSelfDriving) captureLockDoorsResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["LockDoors"] = ch
	return ch
}

// captureLockDoorsCallSpy starts watching for LockDoors spy calls and sends them into a channel.
func (m *mockSelfDriving) captureLockDoorsCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getLockDoorsCalls, timeout)
		ch <- m.getLockDoorsCalls()
	}()
	return ch
}

type mockSelfDrivingLockDoorsResult struct {
	Output0 error
}

// setLockDoorsResponse sets the response for LockDoors
func (m *mockSelfDriving) setLockDoorsResponse(output0 error) {
	m.setLockDoorsFunc(func() error {
		return output0
	})
}

// enqueueLockDoorsResponse enqueues a static response for LockDoors
func (m *mockSelfDriving) enqueueLockDoorsResponse(output0 error) {
	m.mocked.LockDoors.EnqueueWithDelay(func() error {
		return output0
	}, 0)
}

// enqueueLockDoorsResponseWithDelay enqueues a static response with delay for LockDoors
func (m *mockSelfDriving) enqueueLockDoorsResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.LockDoors.EnqueueWithDelay(func() error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// enableGetEngineSpecsMock turns the spy on
func (m *mockSelfDriving) enableGetEngineSpecsSpy() {
	m.mocked.GetEngineSpecs.SpyEnabled = true
}

// getGetEngineSpecsCalls returns recorded calls to GetEngineSpecs
func (m *mockSelfDriving) getGetEngineSpecsCalls() []stubs.MethodCall {
	return m.mocked.GetEngineSpecs.Calls()
}

// enableGetEngineSpecsMock turns the spy off
func (m *mockSelfDriving) disableGetEngineSpecsSpy() {
	m.mocked.GetEngineSpecs.SpyEnabled = false
}

// GetEngineSpecs overrides the method to return the mock response
func (m *mockSelfDriving) GetEngineSpecs() (int, string) {
	m.mocked.GetEngineSpecs.RecordCall()
	var (
		out0 int
		out1 string

		result mockSelfDrivingGetEngineSpecsResult
	)

	if m.mocked.GetEngineSpecs.Enabled {
		out0, out1 = m.mocked.GetEngineSpecs.NextResponse(func() (int, string) {
			return m.real.GetEngineSpecs()
		})()

	} else {
		out0, out1 = m.real.GetEngineSpecs()

	}

	result = mockSelfDrivingGetEngineSpecsResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["GetEngineSpecs"]; ok {
		chTyped := ch.(chan mockSelfDrivingGetEngineSpecsResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setGetEngineSpecsFunc sets the function for GetEngineSpecs
func (m *mockSelfDriving) setGetEngineSpecsFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.Fallback = f
}

// enableGetEngineSpecsMock turns the mock on
func (m *mockSelfDriving) enableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.Enabled = true
}

// disableGetEngineSpecsMock turns the mock off
func (m *mockSelfDriving) disableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.Enabled = false
}

// enqueueGetEngineSpecsResponseFunc enqueues a function response for GetEngineSpecs
func (m *mockSelfDriving) enqueueGetEngineSpecsResponseFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(f, 0)
}

// enqueueGetEngineSpecsResponseFuncWithDelay enqueues a function response with delay for GetEngineSpecs
func (m *mockSelfDriving) enqueueGetEngineSpecsResponseFuncWithDelay(f func() (int, string), d time.Duration) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(f, d)
}

// captureGetEngineSpecsResult sets up a channel to capture GetEngineSpecs results.
func (m *mockSelfDriving) captureGetEngineSpecsResult() <-chan mockSelfDrivingGetEngineSpecsResult {
	ch := make(chan mockSelfDrivingGetEngineSpecsResult, 1)
	m.responseChans["GetEngineSpecs"] = ch
	return ch
}

// captureGetEngineSpecsCallSpy starts watching for GetEngineSpecs spy calls and sends them into a channel.
func (m *mockSelfDriving) captureGetEngineSpecsCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetEngineSpecsCalls, timeout)
		ch <- m.getGetEngineSpecsCalls()
	}()
	return ch
}

type mockSelfDrivingGetEngineSpecsResult struct {
	Output0 int
	Output1 string
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockSelfDriving) setGetEngineSpecsResponse(output0 int, output1 string) {
	m.setGetEngineSpecsFunc(func() (int, string) {
		return output0, output1
	})
}

// enqueueGetEngineSpecsResponse enqueues a static response for GetEngineSpecs
func (m *mockSelfDriving) enqueueGetEngineSpecsResponse(output0 int, output1 string) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(func() (int, string) {
		return output0, output1
	}, 0)
}

// enqueueGetEngineSpecsResponseWithDelay enqueues a static response with delay for GetEngineSpecs
func (m *mockSelfDriving) enqueueGetEngineSpecsResponseWithDelay(output0 int, output1 string, d time.Duration) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(func() (int, string) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// enableApplyBrakesMock turns the spy on
func (m *mockSelfDriving) enableApplyBrakesSpy() {
	m.mocked.ApplyBrakes.SpyEnabled = true
}

// getApplyBrakesCalls returns recorded calls to ApplyBrakes
func (m *mockSelfDriving) getApplyBrakesCalls() []stubs.MethodCall {
	return m.mocked.ApplyBrakes.Calls()
}

// enableApplyBrakesMock turns the spy off
func (m *mockSelfDriving) disableApplyBrakesSpy() {
	m.mocked.ApplyBrakes.SpyEnabled = false
}

// ApplyBrakes overrides the method to return the mock response
func (m *mockSelfDriving) ApplyBrakes(force float64) bool {
	m.mocked.ApplyBrakes.RecordCall(force)
	var (
		out0 bool
	)

	if m.mocked.ApplyBrakes.Enabled {
		out0 = m.mocked.ApplyBrakes.NextResponse(func(force float64) bool {
			return m.real.ApplyBrakes(force)
		})(force)

	} else {
		out0 = m.real.ApplyBrakes(force)

	}

	if ch, ok := m.responseChans["ApplyBrakes"]; ok {
		chTyped := ch.(chan bool)
		chTyped <- out0
	}
	return out0

}

// setApplyBrakesFunc sets the function for ApplyBrakes
func (m *mockSelfDriving) setApplyBrakesFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.Fallback = f
}

// enableApplyBrakesMock turns the mock on
func (m *mockSelfDriving) enableApplyBrakesMock() {
	m.mocked.ApplyBrakes.Enabled = true
}

// disableApplyBrakesMock turns the mock off
func (m *mockSelfDriving) disableApplyBrakesMock() {
	m.mocked.ApplyBrakes.Enabled = false
}

// enqueueApplyBrakesResponseFunc enqueues a function response for ApplyBrakes
func (m *mockSelfDriving) enqueueApplyBrakesResponseFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(f, 0)
}

// enqueueApplyBrakesResponseFuncWithDelay enqueues a function response with delay for ApplyBrakes
func (m *mockSelfDriving) enqueueApplyBrakesResponseFuncWithDelay(f func(float64) bool, d time.Duration) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(f, d)
}

// captureApplyBrakesResult sets up a channel to capture ApplyBrakes results.
func (m *mockSelfDriving) captureApplyBrakesResult() <-chan bool {
	ch := make(chan bool, 1)
	m.responseChans["ApplyBrakes"] = ch
	return ch
}

// captureApplyBrakesCallSpy starts watching for ApplyBrakes spy calls and sends them into a channel.
func (m *mockSelfDriving) captureApplyBrakesCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getApplyBrakesCalls, timeout)
		ch <- m.getApplyBrakesCalls()
	}()
	return ch
}

type mockSelfDrivingApplyBrakesResult struct {
	Output0 bool
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockSelfDriving) setApplyBrakesResponse(output0 bool) {
	m.setApplyBrakesFunc(func(float64) bool {
		return output0
	})
}

// enqueueApplyBrakesResponse enqueues a static response for ApplyBrakes
func (m *mockSelfDriving) enqueueApplyBrakesResponse(output0 bool) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(func(float64) bool {
		return output0
	}, 0)
}

// enqueueApplyBrakesResponseWithDelay enqueues a static response with delay for ApplyBrakes
func (m *mockSelfDriving) enqueueApplyBrakesResponseWithDelay(output0 bool, d time.Duration) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(func(float64) bool {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// enableGetTopSpeedMock turns the spy on
func (m *mockSelfDriving) enableGetTopSpeedSpy() {
	m.mocked.GetTopSpeed.SpyEnabled = true
}

// getGetTopSpeedCalls returns recorded calls to GetTopSpeed
func (m *mockSelfDriving) getGetTopSpeedCalls() []stubs.MethodCall {
	return m.mocked.GetTopSpeed.Calls()
}

// enableGetTopSpeedMock turns the spy off
func (m *mockSelfDriving) disableGetTopSpeedSpy() {
	m.mocked.GetTopSpeed.SpyEnabled = false
}

// GetTopSpeed overrides the method to return the mock response
func (m *mockSelfDriving) GetTopSpeed() int {
	m.mocked.GetTopSpeed.RecordCall()
	var (
		out0 int
	)

	if m.mocked.GetTopSpeed.Enabled {
		out0 = m.mocked.GetTopSpeed.NextResponse(func() int {
			return m.real.GetTopSpeed()
		})()

	} else {
		out0 = m.real.GetTopSpeed()

	}

	if ch, ok := m.responseChans["GetTopSpeed"]; ok {
		chTyped := ch.(chan int)
		chTyped <- out0
	}
	return out0

}

// setGetTopSpeedFunc sets the function for GetTopSpeed
func (m *mockSelfDriving) setGetTopSpeedFunc(f func() int) {
	m.mocked.GetTopSpeed.Fallback = f
}

// enableGetTopSpeedMock turns the mock on
func (m *mockSelfDriving) enableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.Enabled = true
}

// disableGetTopSpeedMock turns the mock off
func (m *mockSelfDriving) disableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.Enabled = false
}

// enqueueGetTopSpeedResponseFunc enqueues a function response for GetTopSpeed
func (m *mockSelfDriving) enqueueGetTopSpeedResponseFunc(f func() int) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(f, 0)
}

// enqueueGetTopSpeedResponseFuncWithDelay enqueues a function response with delay for GetTopSpeed
func (m *mockSelfDriving) enqueueGetTopSpeedResponseFuncWithDelay(f func() int, d time.Duration) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(f, d)
}

// captureGetTopSpeedResult sets up a channel to capture GetTopSpeed results.
func (m *mockSelfDriving) captureGetTopSpeedResult() <-chan int {
	ch := make(chan int, 1)
	m.responseChans["GetTopSpeed"] = ch
	return ch
}

// captureGetTopSpeedCallSpy starts watching for GetTopSpeed spy calls and sends them into a channel.
func (m *mockSelfDriving) captureGetTopSpeedCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetTopSpeedCalls, timeout)
		ch <- m.getGetTopSpeedCalls()
	}()
	return ch
}

type mockSelfDrivingGetTopSpeedResult struct {
	Output0 int
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockSelfDriving) setGetTopSpeedResponse(output0 int) {
	m.setGetTopSpeedFunc(func() int {
		return output0
	})
}

// enqueueGetTopSpeedResponse enqueues a static response for GetTopSpeed
func (m *mockSelfDriving) enqueueGetTopSpeedResponse(output0 int) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(func() int {
		return output0
	}, 0)
}

// enqueueGetTopSpeedResponseWithDelay enqueues a static response with delay for GetTopSpeed
func (m *mockSelfDriving) enqueueGetTopSpeedResponseWithDelay(output0 int, d time.Duration) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(func() int {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- ParkSelf Mock Helpers --------------------------- */

// enableParkSelfMock turns the spy on
func (m *mockSelfDriving) enableParkSelfSpy() {
	m.mocked.ParkSelf.SpyEnabled = true
}

// getParkSelfCalls returns recorded calls to ParkSelf
func (m *mockSelfDriving) getParkSelfCalls() []stubs.MethodCall {
	return m.mocked.ParkSelf.Calls()
}

// enableParkSelfMock turns the spy off
func (m *mockSelfDriving) disableParkSelfSpy() {
	m.mocked.ParkSelf.SpyEnabled = false
}

// ParkSelf overrides the method to return the mock response
func (m *mockSelfDriving) ParkSelf() error {
	m.mocked.ParkSelf.RecordCall()
	var (
		out0 error
	)

	if m.mocked.ParkSelf.Enabled {
		out0 = m.mocked.ParkSelf.NextResponse(func() error {
			return m.real.ParkSelf()
		})()

	} else {
		out0 = m.real.ParkSelf()

	}

	if ch, ok := m.responseChans["ParkSelf"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setParkSelfFunc sets the function for ParkSelf
func (m *mockSelfDriving) setParkSelfFunc(f func() error) {
	m.mocked.ParkSelf.Fallback = f
}

// enableParkSelfMock turns the mock on
func (m *mockSelfDriving) enableParkSelfMock() {
	m.mocked.ParkSelf.Enabled = true
}

// disableParkSelfMock turns the mock off
func (m *mockSelfDriving) disableParkSelfMock() {
	m.mocked.ParkSelf.Enabled = false
}

// enqueueParkSelfResponseFunc enqueues a function response for ParkSelf
func (m *mockSelfDriving) enqueueParkSelfResponseFunc(f func() error) {
	m.mocked.ParkSelf.EnqueueWithDelay(f, 0)
}

// enqueueParkSelfResponseFuncWithDelay enqueues a function response with delay for ParkSelf
func (m *mockSelfDriving) enqueueParkSelfResponseFuncWithDelay(f func() error, d time.Duration) {
	m.mocked.ParkSelf.EnqueueWithDelay(f, d)
}

// captureParkSelfResult sets up a channel to capture ParkSelf results.
func (m *mockSelfDriving) captureParkSelfResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["ParkSelf"] = ch
	return ch
}

// captureParkSelfCallSpy starts watching for ParkSelf spy calls and sends them into a channel.
func (m *mockSelfDriving) captureParkSelfCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getParkSelfCalls, timeout)
		ch <- m.getParkSelfCalls()
	}()
	return ch
}

type mockSelfDrivingParkSelfResult struct {
	Output0 error
}

// setParkSelfResponse sets the response for ParkSelf
func (m *mockSelfDriving) setParkSelfResponse(output0 error) {
	m.setParkSelfFunc(func() error {
		return output0
	})
}

// enqueueParkSelfResponse enqueues a static response for ParkSelf
func (m *mockSelfDriving) enqueueParkSelfResponse(output0 error) {
	m.mocked.ParkSelf.EnqueueWithDelay(func() error {
		return output0
	}, 0)
}

// enqueueParkSelfResponseWithDelay enqueues a static response with delay for ParkSelf
func (m *mockSelfDriving) enqueueParkSelfResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.ParkSelf.EnqueueWithDelay(func() error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// enableHonkMock turns the spy on
func (m *mockSelfDriving) enableHonkSpy() {
	m.mocked.Honk.SpyEnabled = true
}

// getHonkCalls returns recorded calls to Honk
func (m *mockSelfDriving) getHonkCalls() []stubs.MethodCall {
	return m.mocked.Honk.Calls()
}

// enableHonkMock turns the spy off
func (m *mockSelfDriving) disableHonkSpy() {
	m.mocked.Honk.SpyEnabled = false
}

// Honk overrides the method to return the mock response
func (m *mockSelfDriving) Honk(times int) {
	m.mocked.Honk.RecordCall(times)
	var ()

	if m.mocked.Honk.Enabled {

	} else {

	}

	return

}

// setHonkFunc sets the function for Honk
func (m *mockSelfDriving) setHonkFunc(f func(int)) {
	m.mocked.Honk.Fallback = f
}

// enableHonkMock turns the mock on
func (m *mockSelfDriving) enableHonkMock() {
	m.mocked.Honk.Enabled = true
}

// disableHonkMock turns the mock off
func (m *mockSelfDriving) disableHonkMock() {
	m.mocked.Honk.Enabled = false
}

// enqueueHonkResponseFunc enqueues a function response for Honk
func (m *mockSelfDriving) enqueueHonkResponseFunc(f func(int)) {
	m.mocked.Honk.EnqueueWithDelay(f, 0)
}

// enqueueHonkResponseFuncWithDelay enqueues a function response with delay for Honk
func (m *mockSelfDriving) enqueueHonkResponseFuncWithDelay(f func(int), d time.Duration) {
	m.mocked.Honk.EnqueueWithDelay(f, d)
}

// captureHonkResult sets up a channel to capture Honk results.
func (m *mockSelfDriving) captureHonkResult() <-chan struct{} {
	ch := make(chan struct{}, 1)
	m.responseChans["Honk"] = ch
	return ch
}

// captureHonkCallSpy starts watching for Honk spy calls and sends them into a channel.
func (m *mockSelfDriving) captureHonkCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getHonkCalls, timeout)
		ch <- m.getHonkCalls()
	}()
	return ch
}

type mockSelfDrivingHonkResult struct {
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// enableLoadCargoMock turns the spy on
func (m *mockSelfDriving) enableLoadCargoSpy() {
	m.mocked.LoadCargo.SpyEnabled = true
}

// getLoadCargoCalls returns recorded calls to LoadCargo
func (m *mockSelfDriving) getLoadCargoCalls() []stubs.MethodCall {
	return m.mocked.LoadCargo.Calls()
}

// enableLoadCargoMock turns the spy off
func (m *mockSelfDriving) disableLoadCargoSpy() {
	m.mocked.LoadCargo.SpyEnabled = false
}

// LoadCargo overrides the method to return the mock response
func (m *mockSelfDriving) LoadCargo(items []string) (int, error) {
	m.mocked.LoadCargo.RecordCall(items)
	var (
		out0 int
		out1 error

		result mockSelfDrivingLoadCargoResult
	)

	if m.mocked.LoadCargo.Enabled {
		out0, out1 = m.mocked.LoadCargo.NextResponse(func(items []string) (int, error) {
			return m.real.LoadCargo(items)
		})(items)

	} else {
		out0, out1 = m.real.LoadCargo(items)

	}

	result = mockSelfDrivingLoadCargoResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["LoadCargo"]; ok {
		chTyped := ch.(chan mockSelfDrivingLoadCargoResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setLoadCargoFunc sets the function for LoadCargo
func (m *mockSelfDriving) setLoadCargoFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.Fallback = f
}

// enableLoadCargoMock turns the mock on
func (m *mockSelfDriving) enableLoadCargoMock() {
	m.mocked.LoadCargo.Enabled = true
}

// disableLoadCargoMock turns the mock off
func (m *mockSelfDriving) disableLoadCargoMock() {
	m.mocked.LoadCargo.Enabled = false
}

// enqueueLoadCargoResponseFunc enqueues a function response for LoadCargo
func (m *mockSelfDriving) enqueueLoadCargoResponseFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.EnqueueWithDelay(f, 0)
}

// enqueueLoadCargoResponseFuncWithDelay enqueues a function response with delay for LoadCargo
func (m *mockSelfDriving) enqueueLoadCargoResponseFuncWithDelay(f func([]string) (int, error), d time.Duration) {
	m.mocked.LoadCargo.EnqueueWithDelay(f, d)
}

// captureLoadCargoResult sets up a channel to capture LoadCargo results.
func (m *mockSelfDriving) captureLoadCargoResult() <-chan mockSelfDrivingLoadCargoResult {
	ch := make(chan mockSelfDrivingLoadCargoResult, 1)
	m.responseChans["LoadCargo"] = ch
	return ch
}

// captureLoadCargoCallSpy starts watching for LoadCargo spy calls and sends them into a channel.
func (m *mockSelfDriving) captureLoadCargoCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getLoadCargoCalls, timeout)
		ch <- m.getLoadCargoCalls()
	}()
	return ch
}

type mockSelfDrivingLoadCargoResult struct {
	Output0 int
	Output1 error
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockSelfDriving) setLoadCargoResponse(output0 int, output1 error) {
	m.setLoadCargoFunc(func([]string) (int, error) {
		return output0, output1
	})
}

// enqueueLoadCargoResponse enqueues a static response for LoadCargo
func (m *mockSelfDriving) enqueueLoadCargoResponse(output0 int, output1 error) {
	m.mocked.LoadCargo.EnqueueWithDelay(func([]string) (int, error) {
		return output0, output1
	}, 0)
}

// enqueueLoadCargoResponseWithDelay enqueues a static response with delay for LoadCargo
func (m *mockSelfDriving) enqueueLoadCargoResponseWithDelay(output0 int, output1 error, d time.Duration) {
	m.mocked.LoadCargo.EnqueueWithDelay(func([]string) (int, error) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// enableGetVehicleStatusMock turns the spy on
func (m *mockSelfDriving) enableGetVehicleStatusSpy() {
	m.mocked.GetVehicleStatus.SpyEnabled = true
}

// getGetVehicleStatusCalls returns recorded calls to GetVehicleStatus
func (m *mockSelfDriving) getGetVehicleStatusCalls() []stubs.MethodCall {
	return m.mocked.GetVehicleStatus.Calls()
}

// enableGetVehicleStatusMock turns the spy off
func (m *mockSelfDriving) disableGetVehicleStatusSpy() {
	m.mocked.GetVehicleStatus.SpyEnabled = false
}

// GetVehicleStatus overrides the method to return the mock response
func (m *mockSelfDriving) GetVehicleStatus() vehicle.VehicleStatus {
	m.mocked.GetVehicleStatus.RecordCall()
	var (
		out0 vehicle.VehicleStatus
	)

	if m.mocked.GetVehicleStatus.Enabled {
		out0 = m.mocked.GetVehicleStatus.NextResponse(func() vehicle.VehicleStatus {
			return m.real.GetVehicleStatus()
		})()

	} else {
		out0 = m.real.GetVehicleStatus()

	}

	if ch, ok := m.responseChans["GetVehicleStatus"]; ok {
		chTyped := ch.(chan vehicle.VehicleStatus)
		chTyped <- out0
	}
	return out0

}

// setGetVehicleStatusFunc sets the function for GetVehicleStatus
func (m *mockSelfDriving) setGetVehicleStatusFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.Fallback = f
}

// enableGetVehicleStatusMock turns the mock on
func (m *mockSelfDriving) enableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.Enabled = true
}

// disableGetVehicleStatusMock turns the mock off
func (m *mockSelfDriving) disableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.Enabled = false
}

// enqueueGetVehicleStatusResponseFunc enqueues a function response for GetVehicleStatus
func (m *mockSelfDriving) enqueueGetVehicleStatusResponseFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(f, 0)
}

// enqueueGetVehicleStatusResponseFuncWithDelay enqueues a function response with delay for GetVehicleStatus
func (m *mockSelfDriving) enqueueGetVehicleStatusResponseFuncWithDelay(f func() vehicle.VehicleStatus, d time.Duration) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(f, d)
}

// captureGetVehicleStatusResult sets up a channel to capture GetVehicleStatus results.
func (m *mockSelfDriving) captureGetVehicleStatusResult() <-chan vehicle.VehicleStatus {
	ch := make(chan vehicle.VehicleStatus, 1)
	m.responseChans["GetVehicleStatus"] = ch
	return ch
}

// captureGetVehicleStatusCallSpy starts watching for GetVehicleStatus spy calls and sends them into a channel.
func (m *mockSelfDriving) captureGetVehicleStatusCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetVehicleStatusCalls, timeout)
		ch <- m.getGetVehicleStatusCalls()
	}()
	return ch
}

type mockSelfDrivingGetVehicleStatusResult struct {
	Output0 vehicle.VehicleStatus
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockSelfDriving) setGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.setGetVehicleStatusFunc(func() vehicle.VehicleStatus {
		return output0
	})
}

// enqueueGetVehicleStatusResponse enqueues a static response for GetVehicleStatus
func (m *mockSelfDriving) enqueueGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(func() vehicle.VehicleStatus {
		return output0
	}, 0)
}

// enqueueGetVehicleStatusResponseWithDelay enqueues a static response with delay for GetVehicleStatus
func (m *mockSelfDriving) enqueueGetVehicleStatusResponseWithDelay(output0 vehicle.VehicleStatus, d time.Duration) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(func() vehicle.VehicleStatus {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- TurnOffAC Mock Helpers --------------------------- */

// enableTurnOffACMock turns the spy on
func (m *mockSelfDriving) enableTurnOffACSpy() {
	m.mocked.TurnOffAC.SpyEnabled = true
}

// getTurnOffACCalls returns recorded calls to TurnOffAC
func (m *mockSelfDriving) getTurnOffACCalls() []stubs.MethodCall {
	return m.mocked.TurnOffAC.Calls()
}

// enableTurnOffACMock turns the spy off
func (m *mockSelfDriving) disableTurnOffACSpy() {
	m.mocked.TurnOffAC.SpyEnabled = false
}

// TurnOffAC overrides the method to return the mock response
func (m *mockSelfDriving) TurnOffAC() error {
	m.mocked.TurnOffAC.RecordCall()
	var (
		out0 error
	)

	if m.mocked.TurnOffAC.Enabled {
		out0 = m.mocked.TurnOffAC.NextResponse(func() error {
			return m.real.TurnOffAC()
		})()

	} else {
		out0 = m.real.TurnOffAC()

	}

	if ch, ok := m.responseChans["TurnOffAC"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setTurnOffACFunc sets the function for TurnOffAC
func (m *mockSelfDriving) setTurnOffACFunc(f func() error) {
	m.mocked.TurnOffAC.Fallback = f
}

// enableTurnOffACMock turns the mock on
func (m *mockSelfDriving) enableTurnOffACMock() {
	m.mocked.TurnOffAC.Enabled = true
}

// disableTurnOffACMock turns the mock off
func (m *mockSelfDriving) disableTurnOffACMock() {
	m.mocked.TurnOffAC.Enabled = false
}

// enqueueTurnOffACResponseFunc enqueues a function response for TurnOffAC
func (m *mockSelfDriving) enqueueTurnOffACResponseFunc(f func() error) {
	m.mocked.TurnOffAC.EnqueueWithDelay(f, 0)
}

// enqueueTurnOffACResponseFuncWithDelay enqueues a function response with delay for TurnOffAC
func (m *mockSelfDriving) enqueueTurnOffACResponseFuncWithDelay(f func() error, d time.Duration) {
	m.mocked.TurnOffAC.EnqueueWithDelay(f, d)
}

// captureTurnOffACResult sets up a channel to capture TurnOffAC results.
func (m *mockSelfDriving) captureTurnOffACResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["TurnOffAC"] = ch
	return ch
}

// captureTurnOffACCallSpy starts watching for TurnOffAC spy calls and sends them into a channel.
func (m *mockSelfDriving) captureTurnOffACCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getTurnOffACCalls, timeout)
		ch <- m.getTurnOffACCalls()
	}()
	return ch
}

type mockSelfDrivingTurnOffACResult struct {
	Output0 error
}

// setTurnOffACResponse sets the response for TurnOffAC
func (m *mockSelfDriving) setTurnOffACResponse(output0 error) {
	m.setTurnOffACFunc(func() error {
		return output0
	})
}

// enqueueTurnOffACResponse enqueues a static response for TurnOffAC
func (m *mockSelfDriving) enqueueTurnOffACResponse(output0 error) {
	m.mocked.TurnOffAC.EnqueueWithDelay(func() error {
		return output0
	}, 0)
}

// enqueueTurnOffACResponseWithDelay enqueues a static response with delay for TurnOffAC
func (m *mockSelfDriving) enqueueTurnOffACResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.TurnOffAC.EnqueueWithDelay(func() error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- TurnOffMusic Mock Helpers --------------------------- */

// enableTurnOffMusicMock turns the spy on
func (m *mockSelfDriving) enableTurnOffMusicSpy() {
	m.mocked.TurnOffMusic.SpyEnabled = true
}

// getTurnOffMusicCalls returns recorded calls to TurnOffMusic
func (m *mockSelfDriving) getTurnOffMusicCalls() []stubs.MethodCall {
	return m.mocked.TurnOffMusic.Calls()
}

// enableTurnOffMusicMock turns the spy off
func (m *mockSelfDriving) disableTurnOffMusicSpy() {
	m.mocked.TurnOffMusic.SpyEnabled = false
}

// TurnOffMusic overrides the method to return the mock response
func (m *mockSelfDriving) TurnOffMusic() error {
	m.mocked.TurnOffMusic.RecordCall()
	var (
		out0 error
	)

	if m.mocked.TurnOffMusic.Enabled {
		out0 = m.mocked.TurnOffMusic.NextResponse(func() error {
			return m.real.TurnOffMusic()
		})()

	} else {
		out0 = m.real.TurnOffMusic()

	}

	if ch, ok := m.responseChans["TurnOffMusic"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setTurnOffMusicFunc sets the function for TurnOffMusic
func (m *mockSelfDriving) setTurnOffMusicFunc(f func() error) {
	m.mocked.TurnOffMusic.Fallback = f
}

// enableTurnOffMusicMock turns the mock on
func (m *mockSelfDriving) enableTurnOffMusicMock() {
	m.mocked.TurnOffMusic.Enabled = true
}

// disableTurnOffMusicMock turns the mock off
func (m *mockSelfDriving) disableTurnOffMusicMock() {
	m.mocked.TurnOffMusic.Enabled = false
}

// enqueueTurnOffMusicResponseFunc enqueues a function response for TurnOffMusic
func (m *mockSelfDriving) enqueueTurnOffMusicResponseFunc(f func() error) {
	m.mocked.TurnOffMusic.EnqueueWithDelay(f, 0)
}

// enqueueTurnOffMusicResponseFuncWithDelay enqueues a function response with delay for TurnOffMusic
func (m *mockSelfDriving) enqueueTurnOffMusicResponseFuncWithDelay(f func() error, d time.Duration) {
	m.mocked.TurnOffMusic.EnqueueWithDelay(f, d)
}

// captureTurnOffMusicResult sets up a channel to capture TurnOffMusic results.
func (m *mockSelfDriving) captureTurnOffMusicResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["TurnOffMusic"] = ch
	return ch
}

// captureTurnOffMusicCallSpy starts watching for TurnOffMusic spy calls and sends them into a channel.
func (m *mockSelfDriving) captureTurnOffMusicCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getTurnOffMusicCalls, timeout)
		ch <- m.getTurnOffMusicCalls()
	}()
	return ch
}

type mockSelfDrivingTurnOffMusicResult struct {
	Output0 error
}

// setTurnOffMusicResponse sets the response for TurnOffMusic
func (m *mockSelfDriving) setTurnOffMusicResponse(output0 error) {
	m.setTurnOffMusicFunc(func() error {
		return output0
	})
}

// enqueueTurnOffMusicResponse enqueues a static response for TurnOffMusic
func (m *mockSelfDriving) enqueueTurnOffMusicResponse(output0 error) {
	m.mocked.TurnOffMusic.EnqueueWithDelay(func() error {
		return output0
	}, 0)
}

// enqueueTurnOffMusicResponseWithDelay enqueues a static response with delay for TurnOffMusic
func (m *mockSelfDriving) enqueueTurnOffMusicResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.TurnOffMusic.EnqueueWithDelay(func() error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- CloseWindows Mock Helpers --------------------------- */

// enableCloseWindowsMock turns the spy on
func (m *mockSelfDriving) enableCloseWindowsSpy() {
	m.mocked.CloseWindows.SpyEnabled = true
}

// getCloseWindowsCalls returns recorded calls to CloseWindows
func (m *mockSelfDriving) getCloseWindowsCalls() []stubs.MethodCall {
	return m.mocked.CloseWindows.Calls()
}

// enableCloseWindowsMock turns the spy off
func (m *mockSelfDriving) disableCloseWindowsSpy() {
	m.mocked.CloseWindows.SpyEnabled = false
}

// CloseWindows overrides the method to return the mock response
func (m *mockSelfDriving) CloseWindows() error {
	m.mocked.CloseWindows.RecordCall()
	var (
		out0 error
	)

	if m.mocked.CloseWindows.Enabled {
		out0 = m.mocked.CloseWindows.NextResponse(func() error {
			return m.real.CloseWindows()
		})()

	} else {
		out0 = m.real.CloseWindows()

	}

	if ch, ok := m.responseChans["CloseWindows"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setCloseWindowsFunc sets the function for CloseWindows
func (m *mockSelfDriving) setCloseWindowsFunc(f func() error) {
	m.mocked.CloseWindows.Fallback = f
}

// enableCloseWindowsMock turns the mock on
func (m *mockSelfDriving) enableCloseWindowsMock() {
	m.mocked.CloseWindows.Enabled = true
}

// disableCloseWindowsMock turns the mock off
func (m *mockSelfDriving) disableCloseWindowsMock() {
	m.mocked.CloseWindows.Enabled = false
}

// enqueueCloseWindowsResponseFunc enqueues a function response for CloseWindows
func (m *mockSelfDriving) enqueueCloseWindowsResponseFunc(f func() error) {
	m.mocked.CloseWindows.EnqueueWithDelay(f, 0)
}

// enqueueCloseWindowsResponseFuncWithDelay enqueues a function response with delay for CloseWindows
func (m *mockSelfDriving) enqueueCloseWindowsResponseFuncWithDelay(f func() error, d time.Duration) {
	m.mocked.CloseWindows.EnqueueWithDelay(f, d)
}

// captureCloseWindowsResult sets up a channel to capture CloseWindows results.
func (m *mockSelfDriving) captureCloseWindowsResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["CloseWindows"] = ch
	return ch
}

// captureCloseWindowsCallSpy starts watching for CloseWindows spy calls and sends them into a channel.
func (m *mockSelfDriving) captureCloseWindowsCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getCloseWindowsCalls, timeout)
		ch <- m.getCloseWindowsCalls()
	}()
	return ch
}

type mockSelfDrivingCloseWindowsResult struct {
	Output0 error
}

// setCloseWindowsResponse sets the response for CloseWindows
func (m *mockSelfDriving) setCloseWindowsResponse(output0 error) {
	m.setCloseWindowsFunc(func() error {
		return output0
	})
}

// enqueueCloseWindowsResponse enqueues a static response for CloseWindows
func (m *mockSelfDriving) enqueueCloseWindowsResponse(output0 error) {
	m.mocked.CloseWindows.EnqueueWithDelay(func() error {
		return output0
	}, 0)
}

// enqueueCloseWindowsResponseWithDelay enqueues a static response with delay for CloseWindows
func (m *mockSelfDriving) enqueueCloseWindowsResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.CloseWindows.EnqueueWithDelay(func() error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// enableReverseMock turns the spy on
func (m *mockSelfDriving) enableReverseSpy() {
	m.mocked.Reverse.SpyEnabled = true
}

// getReverseCalls returns recorded calls to Reverse
func (m *mockSelfDriving) getReverseCalls() []stubs.MethodCall {
	return m.mocked.Reverse.Calls()
}

// enableReverseMock turns the spy off
func (m *mockSelfDriving) disableReverseSpy() {
	m.mocked.Reverse.SpyEnabled = false
}

// Reverse overrides the method to return the mock response
func (m *mockSelfDriving) Reverse() (string, error) {
	m.mocked.Reverse.RecordCall()
	var (
		out0 string
		out1 error

		result mockSelfDrivingReverseResult
	)

	if m.mocked.Reverse.Enabled {
		out0, out1 = m.mocked.Reverse.NextResponse(func() (string, error) {
			return m.real.Reverse()
		})()

	} else {
		out0, out1 = m.real.Reverse()

	}

	result = mockSelfDrivingReverseResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["Reverse"]; ok {
		chTyped := ch.(chan mockSelfDrivingReverseResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setReverseFunc sets the function for Reverse
func (m *mockSelfDriving) setReverseFunc(f func() (string, error)) {
	m.mocked.Reverse.Fallback = f
}

// enableReverseMock turns the mock on
func (m *mockSelfDriving) enableReverseMock() {
	m.mocked.Reverse.Enabled = true
}

// disableReverseMock turns the mock off
func (m *mockSelfDriving) disableReverseMock() {
	m.mocked.Reverse.Enabled = false
}

// enqueueReverseResponseFunc enqueues a function response for Reverse
func (m *mockSelfDriving) enqueueReverseResponseFunc(f func() (string, error)) {
	m.mocked.Reverse.EnqueueWithDelay(f, 0)
}

// enqueueReverseResponseFuncWithDelay enqueues a function response with delay for Reverse
func (m *mockSelfDriving) enqueueReverseResponseFuncWithDelay(f func() (string, error), d time.Duration) {
	m.mocked.Reverse.EnqueueWithDelay(f, d)
}

// captureReverseResult sets up a channel to capture Reverse results.
func (m *mockSelfDriving) captureReverseResult() <-chan mockSelfDrivingReverseResult {
	ch := make(chan mockSelfDrivingReverseResult, 1)
	m.responseChans["Reverse"] = ch
	return ch
}

// captureReverseCallSpy starts watching for Reverse spy calls and sends them into a channel.
func (m *mockSelfDriving) captureReverseCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getReverseCalls, timeout)
		ch <- m.getReverseCalls()
	}()
	return ch
}

type mockSelfDrivingReverseResult struct {
	Output0 string
	Output1 error
}

// setReverseResponse sets the response for Reverse
func (m *mockSelfDriving) setReverseResponse(output0 string, output1 error) {
	m.setReverseFunc(func() (string, error) {
		return output0, output1
	})
}

// enqueueReverseResponse enqueues a static response for Reverse
func (m *mockSelfDriving) enqueueReverseResponse(output0 string, output1 error) {
	m.mocked.Reverse.EnqueueWithDelay(func() (string, error) {
		return output0, output1
	}, 0)
}

// enqueueReverseResponseWithDelay enqueues a static response with delay for Reverse
func (m *mockSelfDriving) enqueueReverseResponseWithDelay(output0 string, output1 error, d time.Duration) {
	m.mocked.Reverse.EnqueueWithDelay(func() (string, error) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// enableIsMovingMock turns the spy on
func (m *mockSelfDriving) enableIsMovingSpy() {
	m.mocked.IsMoving.SpyEnabled = true
}

// getIsMovingCalls returns recorded calls to IsMoving
func (m *mockSelfDriving) getIsMovingCalls() []stubs.MethodCall {
	return m.mocked.IsMoving.Calls()
}

// enableIsMovingMock turns the spy off
func (m *mockSelfDriving) disableIsMovingSpy() {
	m.mocked.IsMoving.SpyEnabled = false
}

// IsMoving overrides the method to return the mock response
func (m *mockSelfDriving) IsMoving() bool {
	m.mocked.IsMoving.RecordCall()
	var (
		out0 bool
	)

	if m.mocked.IsMoving.Enabled {
		out0 = m.mocked.IsMoving.NextResponse(func() bool {
			return m.real.IsMoving()
		})()

	} else {
		out0 = m.real.IsMoving()

	}

	if ch, ok := m.responseChans["IsMoving"]; ok {
		chTyped := ch.(chan bool)
		chTyped <- out0
	}
	return out0

}

// setIsMovingFunc sets the function for IsMoving
func (m *mockSelfDriving) setIsMovingFunc(f func() bool) {
	m.mocked.IsMoving.Fallback = f
}

// enableIsMovingMock turns the mock on
func (m *mockSelfDriving) enableIsMovingMock() {
	m.mocked.IsMoving.Enabled = true
}

// disableIsMovingMock turns the mock off
func (m *mockSelfDriving) disableIsMovingMock() {
	m.mocked.IsMoving.Enabled = false
}

// enqueueIsMovingResponseFunc enqueues a function response for IsMoving
func (m *mockSelfDriving) enqueueIsMovingResponseFunc(f func() bool) {
	m.mocked.IsMoving.EnqueueWithDelay(f, 0)
}

// enqueueIsMovingResponseFuncWithDelay enqueues a function response with delay for IsMoving
func (m *mockSelfDriving) enqueueIsMovingResponseFuncWithDelay(f func() bool, d time.Duration) {
	m.mocked.IsMoving.EnqueueWithDelay(f, d)
}

// captureIsMovingResult sets up a channel to capture IsMoving results.
func (m *mockSelfDriving) captureIsMovingResult() <-chan bool {
	ch := make(chan bool, 1)
	m.responseChans["IsMoving"] = ch
	return ch
}

// captureIsMovingCallSpy starts watching for IsMoving spy calls and sends them into a channel.
func (m *mockSelfDriving) captureIsMovingCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getIsMovingCalls, timeout)
		ch <- m.getIsMovingCalls()
	}()
	return ch
}

type mockSelfDrivingIsMovingResult struct {
	Output0 bool
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockSelfDriving) setIsMovingResponse(output0 bool) {
	m.setIsMovingFunc(func() bool {
		return output0
	})
}

// enqueueIsMovingResponse enqueues a static response for IsMoving
func (m *mockSelfDriving) enqueueIsMovingResponse(output0 bool) {
	m.mocked.IsMoving.EnqueueWithDelay(func() bool {
		return output0
	}, 0)
}

// enqueueIsMovingResponseWithDelay enqueues a static response with delay for IsMoving
func (m *mockSelfDriving) enqueueIsMovingResponseWithDelay(output0 bool, d time.Duration) {
	m.mocked.IsMoving.EnqueueWithDelay(func() bool {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// enableChangeGearsMock turns the spy on
func (m *mockSelfDriving) enableChangeGearsSpy() {
	m.mocked.ChangeGears.SpyEnabled = true
}

// getChangeGearsCalls returns recorded calls to ChangeGears
func (m *mockSelfDriving) getChangeGearsCalls() []stubs.MethodCall {
	return m.mocked.ChangeGears.Calls()
}

// enableChangeGearsMock turns the spy off
func (m *mockSelfDriving) disableChangeGearsSpy() {
	m.mocked.ChangeGears.SpyEnabled = false
}

// ChangeGears overrides the method to return the mock response
func (m *mockSelfDriving) ChangeGears(gear int) (int, int) {
	m.mocked.ChangeGears.RecordCall(gear)
	var (
		out0 int
		out1 int

		result mockSelfDrivingChangeGearsResult
	)

	if m.mocked.ChangeGears.Enabled {
		out0, out1 = m.mocked.ChangeGears.NextResponse(func(gear int) (int, int) {
			return m.real.ChangeGears(gear)
		})(gear)

	} else {
		out0, out1 = m.real.ChangeGears(gear)

	}

	result = mockSelfDrivingChangeGearsResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["ChangeGears"]; ok {
		chTyped := ch.(chan mockSelfDrivingChangeGearsResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setChangeGearsFunc sets the function for ChangeGears
func (m *mockSelfDriving) setChangeGearsFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.Fallback = f
}

// enableChangeGearsMock turns the mock on
func (m *mockSelfDriving) enableChangeGearsMock() {
	m.mocked.ChangeGears.Enabled = true
}

// disableChangeGearsMock turns the mock off
func (m *mockSelfDriving) disableChangeGearsMock() {
	m.mocked.ChangeGears.Enabled = false
}

// enqueueChangeGearsResponseFunc enqueues a function response for ChangeGears
func (m *mockSelfDriving) enqueueChangeGearsResponseFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.EnqueueWithDelay(f, 0)
}

// enqueueChangeGearsResponseFuncWithDelay enqueues a function response with delay for ChangeGears
func (m *mockSelfDriving) enqueueChangeGearsResponseFuncWithDelay(f func(int) (int, int), d time.Duration) {
	m.mocked.ChangeGears.EnqueueWithDelay(f, d)
}

// captureChangeGearsResult sets up a channel to capture ChangeGears results.
func (m *mockSelfDriving) captureChangeGearsResult() <-chan mockSelfDrivingChangeGearsResult {
	ch := make(chan mockSelfDrivingChangeGearsResult, 1)
	m.responseChans["ChangeGears"] = ch
	return ch
}

// captureChangeGearsCallSpy starts watching for ChangeGears spy calls and sends them into a channel.
func (m *mockSelfDriving) captureChangeGearsCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getChangeGearsCalls, timeout)
		ch <- m.getChangeGearsCalls()
	}()
	return ch
}

type mockSelfDrivingChangeGearsResult struct {
	Output0 int
	Output1 int
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockSelfDriving) setChangeGearsResponse(output0 int, output1 int) {
	m.setChangeGearsFunc(func(int) (int, int) {
		return output0, output1
	})
}

// enqueueChangeGearsResponse enqueues a static response for ChangeGears
func (m *mockSelfDriving) enqueueChangeGearsResponse(output0 int, output1 int) {
	m.mocked.ChangeGears.EnqueueWithDelay(func(int) (int, int) {
		return output0, output1
	}, 0)
}

// enqueueChangeGearsResponseWithDelay enqueues a static response with delay for ChangeGears
func (m *mockSelfDriving) enqueueChangeGearsResponseWithDelay(output0 int, output1 int, d time.Duration) {
	m.mocked.ChangeGears.EnqueueWithDelay(func(int) (int, int) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// enableTelemetryMock turns the spy on
func (m *mockSelfDriving) enableTelemetrySpy() {
	m.mocked.Telemetry.SpyEnabled = true
}

// getTelemetryCalls returns recorded calls to Telemetry
func (m *mockSelfDriving) getTelemetryCalls() []stubs.MethodCall {
	return m.mocked.Telemetry.Calls()
}

// enableTelemetryMock turns the spy off
func (m *mockSelfDriving) disableTelemetrySpy() {
	m.mocked.Telemetry.SpyEnabled = false
}

// Telemetry overrides the method to return the mock response
func (m *mockSelfDriving) Telemetry() map[string]float64 {
	m.mocked.Telemetry.RecordCall()
	var (
		out0 map[string]float64
	)

	if m.mocked.Telemetry.Enabled {
		out0 = m.mocked.Telemetry.NextResponse(func() map[string]float64 {
			return m.real.Telemetry()
		})()

	} else {
		out0 = m.real.Telemetry()

	}

	if ch, ok := m.responseChans["Telemetry"]; ok {
		chTyped := ch.(chan map[string]float64)
		chTyped <- out0
	}
	return out0

}

// setTelemetryFunc sets the function for Telemetry
func (m *mockSelfDriving) setTelemetryFunc(f func() map[string]float64) {
	m.mocked.Telemetry.Fallback = f
}

// enableTelemetryMock turns the mock on
func (m *mockSelfDriving) enableTelemetryMock() {
	m.mocked.Telemetry.Enabled = true
}

// disableTelemetryMock turns the mock off
func (m *mockSelfDriving) disableTelemetryMock() {
	m.mocked.Telemetry.Enabled = false
}

// enqueueTelemetryResponseFunc enqueues a function response for Telemetry
func (m *mockSelfDriving) enqueueTelemetryResponseFunc(f func() map[string]float64) {
	m.mocked.Telemetry.EnqueueWithDelay(f, 0)
}

// enqueueTelemetryResponseFuncWithDelay enqueues a function response with delay for Telemetry
func (m *mockSelfDriving) enqueueTelemetryResponseFuncWithDelay(f func() map[string]float64, d time.Duration) {
	m.mocked.Telemetry.EnqueueWithDelay(f, d)
}

// captureTelemetryResult sets up a channel to capture Telemetry results.
func (m *mockSelfDriving) captureTelemetryResult() <-chan map[string]float64 {
	ch := make(chan map[string]float64, 1)
	m.responseChans["Telemetry"] = ch
	return ch
}

// captureTelemetryCallSpy starts watching for Telemetry spy calls and sends them into a channel.
func (m *mockSelfDriving) captureTelemetryCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getTelemetryCalls, timeout)
		ch <- m.getTelemetryCalls()
	}()
	return ch
}

type mockSelfDrivingTelemetryResult struct {
	Output0 map[string]float64
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockSelfDriving) setTelemetryResponse(output0 map[string]float64) {
	m.setTelemetryFunc(func() map[string]float64 {
		return output0
	})
}

// enqueueTelemetryResponse enqueues a static response for Telemetry
func (m *mockSelfDriving) enqueueTelemetryResponse(output0 map[string]float64) {
	m.mocked.Telemetry.EnqueueWithDelay(func() map[string]float64 {
		return output0
	}, 0)
}

// enqueueTelemetryResponseWithDelay enqueues a static response with delay for Telemetry
func (m *mockSelfDriving) enqueueTelemetryResponseWithDelay(output0 map[string]float64, d time.Duration) {
	m.mocked.Telemetry.EnqueueWithDelay(func() map[string]float64 {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// enableAccelerateMock turns the spy on
func (m *mockSelfDriving) enableAccelerateSpy() {
	m.mocked.Accelerate.SpyEnabled = true
}

// getAccelerateCalls returns recorded calls to Accelerate
func (m *mockSelfDriving) getAccelerateCalls() []stubs.MethodCall {
	return m.mocked.Accelerate.Calls()
}

// enableAccelerateMock turns the spy off
func (m *mockSelfDriving) disableAccelerateSpy() {
	m.mocked.Accelerate.SpyEnabled = false
}

// Accelerate overrides the method to return the mock response
func (m *mockSelfDriving) Accelerate(speed int, unit string) (int, error) {
	m.mocked.Accelerate.RecordCall(speed, unit)
	var (
		out0 int
		out1 error

		result mockSelfDrivingAccelerateResult
	)

	if m.mocked.Accelerate.Enabled {
		out0, out1 = m.mocked.Accelerate.NextResponse(func(speed int, unit string) (int, error) {
			return m.real.Accelerate(speed, unit)
		})(speed, unit)

	} else {
		out0, out1 = m.real.Accelerate(speed, unit)

	}

	result = mockSelfDrivingAccelerateResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["Accelerate"]; ok {
		chTyped := ch.(chan mockSelfDrivingAccelerateResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setAccelerateFunc sets the function for Accelerate
func (m *mockSelfDriving) setAccelerateFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.Fallback = f
}

// enableAccelerateMock turns the mock on
func (m *mockSelfDriving) enableAccelerateMock() {
	m.mocked.Accelerate.Enabled = true
}

// disableAccelerateMock turns the mock off
func (m *mockSelfDriving) disableAccelerateMock() {
	m.mocked.Accelerate.Enabled = false
}

// enqueueAccelerateResponseFunc enqueues a function response for Accelerate
func (m *mockSelfDriving) enqueueAccelerateResponseFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.EnqueueWithDelay(f, 0)
}

// enqueueAccelerateResponseFuncWithDelay enqueues a function response with delay for Accelerate
func (m *mockSelfDriving) enqueueAccelerateResponseFuncWithDelay(f func(int, string) (int, error), d time.Duration) {
	m.mocked.Accelerate.EnqueueWithDelay(f, d)
}

// captureAccelerateResult sets up a channel to capture Accelerate results.
func (m *mockSelfDriving) captureAccelerateResult() <-chan mockSelfDrivingAccelerateResult {
	ch := make(chan mockSelfDrivingAccelerateResult, 1)
	m.responseChans["Accelerate"] = ch
	return ch
}

// captureAccelerateCallSpy starts watching for Accelerate spy calls and sends them into a channel.
func (m *mockSelfDriving) captureAccelerateCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getAccelerateCalls, timeout)
		ch <- m.getAccelerateCalls()
	}()
	return ch
}

type mockSelfDrivingAccelerateResult struct {
	Output0 int
	Output1 error
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockSelfDriving) setAccelerateResponse(output0 int, output1 error) {
	m.setAccelerateFunc(func(int, string) (int, error) {
		return output0, output1
	})
}

// enqueueAccelerateResponse enqueues a static response for Accelerate
func (m *mockSelfDriving) enqueueAccelerateResponse(output0 int, output1 error) {
	m.mocked.Accelerate.EnqueueWithDelay(func(int, string) (int, error) {
		return output0, output1
	}, 0)
}

// enqueueAccelerateResponseWithDelay enqueues a static response with delay for Accelerate
func (m *mockSelfDriving) enqueueAccelerateResponseWithDelay(output0 int, output1 error, d time.Duration) {
	m.mocked.Accelerate.EnqueueWithDelay(func(int, string) (int, error) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- DriveSelf Mock Helpers --------------------------- */

// enableDriveSelfMock turns the spy on
func (m *mockSelfDriving) enableDriveSelfSpy() {
	m.mocked.DriveSelf.SpyEnabled = true
}

// getDriveSelfCalls returns recorded calls to DriveSelf
func (m *mockSelfDriving) getDriveSelfCalls() []stubs.MethodCall {
	return m.mocked.DriveSelf.Calls()
}

// enableDriveSelfMock turns the spy off
func (m *mockSelfDriving) disableDriveSelfSpy() {
	m.mocked.DriveSelf.SpyEnabled = false
}

// DriveSelf overrides the method to return the mock response
func (m *mockSelfDriving) DriveSelf(endLocation string) error {
	m.mocked.DriveSelf.RecordCall(endLocation)
	var out0 error

	if m.mocked.DriveSelf.Enabled {
		out0 = m.mocked.DriveSelf.NextResponse(func(endLocation string) error {
			return m.real.DriveSelf(endLocation)
		})(endLocation)

	} else {
		out0 = m.real.DriveSelf(endLocation)

	}

	if ch, ok := m.responseChans["DriveSelf"]; ok {
		chTyped := ch.(chan error)
		chTyped <- out0
	}
	return out0

}

// setDriveSelfFunc sets the function for DriveSelf
func (m *mockSelfDriving) setDriveSelfFunc(f func(string) error) {
	m.mocked.DriveSelf.Fallback = f
}

// enableDriveSelfMock turns the mock on
func (m *mockSelfDriving) enableDriveSelfMock() {
	m.mocked.DriveSelf.Enabled = true
}

// disableDriveSelfMock turns the mock off
func (m *mockSelfDriving) disableDriveSelfMock() {
	m.mocked.DriveSelf.Enabled = false
}

// enqueueDriveSelfResponseFunc enqueues a function response for DriveSelf
func (m *mockSelfDriving) enqueueDriveSelfResponseFunc(f func(string) error) {
	m.mocked.DriveSelf.EnqueueWithDelay(f, 0)
}

// enqueueDriveSelfResponseFuncWithDelay enqueues a function response with delay for DriveSelf
func (m *mockSelfDriving) enqueueDriveSelfResponseFuncWithDelay(f func(string) error, d time.Duration) {
	m.mocked.DriveSelf.EnqueueWithDelay(f, d)
}

// captureDriveSelfResult sets up a channel to capture DriveSelf results.
func (m *mockSelfDriving) captureDriveSelfResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["DriveSelf"] = ch
	return ch
}

// captureDriveSelfCallSpy starts watching for DriveSelf spy calls and sends them into a channel.
func (m *mockSelfDriving) captureDriveSelfCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getDriveSelfCalls, timeout)
		ch <- m.getDriveSelfCalls()
	}()
	return ch
}

type mockSelfDrivingDriveSelfResult struct {
	Output0 error
}

// setDriveSelfResponse sets the response for DriveSelf
func (m *mockSelfDriving) setDriveSelfResponse(output0 error) {
	m.setDriveSelfFunc(func(string) error {
		return output0
	})
}

// enqueueDriveSelfResponse enqueues a static response for DriveSelf
func (m *mockSelfDriving) enqueueDriveSelfResponse(output0 error) {
	m.mocked.DriveSelf.EnqueueWithDelay(func(string) error {
		return output0
	}, 0)
}

// enqueueDriveSelfResponseWithDelay enqueues a static response with delay for DriveSelf
func (m *mockSelfDriving) enqueueDriveSelfResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.DriveSelf.EnqueueWithDelay(func(string) error {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// enableTurnMock turns the spy on
func (m *mockSelfDriving) enableTurnSpy() {
	m.mocked.Turn.SpyEnabled = true
}

// getTurnCalls returns recorded calls to Turn
func (m *mockSelfDriving) getTurnCalls() []stubs.MethodCall {
	return m.mocked.Turn.Calls()
}

// enableTurnMock turns the spy off
func (m *mockSelfDriving) disableTurnSpy() {
	m.mocked.Turn.SpyEnabled = false
}

// Turn overrides the method to return the mock response
func (m *mockSelfDriving) Turn(dir string) string {
	m.mocked.Turn.RecordCall(dir)
	var (
		out0 string
	)

	if m.mocked.Turn.Enabled {
		out0 = m.mocked.Turn.NextResponse(func(dir string) string {
			return m.real.Turn(dir)
		})(dir)

	} else {
		out0 = m.real.Turn(dir)

	}

	if ch, ok := m.responseChans["Turn"]; ok {
		chTyped := ch.(chan string)
		chTyped <- out0
	}
	return out0

}

// setTurnFunc sets the function for Turn
func (m *mockSelfDriving) setTurnFunc(f func(string) string) {
	m.mocked.Turn.Fallback = f
}

// enableTurnMock turns the mock on
func (m *mockSelfDriving) enableTurnMock() {
	m.mocked.Turn.Enabled = true
}

// disableTurnMock turns the mock off
func (m *mockSelfDriving) disableTurnMock() {
	m.mocked.Turn.Enabled = false
}

// enqueueTurnResponseFunc enqueues a function response for Turn
func (m *mockSelfDriving) enqueueTurnResponseFunc(f func(string) string) {
	m.mocked.Turn.EnqueueWithDelay(f, 0)
}

// enqueueTurnResponseFuncWithDelay enqueues a function response with delay for Turn
func (m *mockSelfDriving) enqueueTurnResponseFuncWithDelay(f func(string) string, d time.Duration) {
	m.mocked.Turn.EnqueueWithDelay(f, d)
}

// captureTurnResult sets up a channel to capture Turn results.
func (m *mockSelfDriving) captureTurnResult() <-chan string {
	ch := make(chan string, 1)
	m.responseChans["Turn"] = ch
	return ch
}

// captureTurnCallSpy starts watching for Turn spy calls and sends them into a channel.
func (m *mockSelfDriving) captureTurnCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getTurnCalls, timeout)
		ch <- m.getTurnCalls()
	}()
	return ch
}

type mockSelfDrivingTurnResult struct {
	Output0 string
}

// setTurnResponse sets the response for Turn
func (m *mockSelfDriving) setTurnResponse(output0 string) {
	m.setTurnFunc(func(string) string {
		return output0
	})
}

// enqueueTurnResponse enqueues a static response for Turn
func (m *mockSelfDriving) enqueueTurnResponse(output0 string) {
	m.mocked.Turn.EnqueueWithDelay(func(string) string {
		return output0
	}, 0)
}

// enqueueTurnResponseWithDelay enqueues a static response with delay for Turn
func (m *mockSelfDriving) enqueueTurnResponseWithDelay(output0 string, d time.Duration) {
	m.mocked.Turn.EnqueueWithDelay(func(string) string {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// enableGetPassengersMock turns the spy on
func (m *mockSelfDriving) enableGetPassengersSpy() {
	m.mocked.GetPassengers.SpyEnabled = true
}

// getGetPassengersCalls returns recorded calls to GetPassengers
func (m *mockSelfDriving) getGetPassengersCalls() []stubs.MethodCall {
	return m.mocked.GetPassengers.Calls()
}

// enableGetPassengersMock turns the spy off
func (m *mockSelfDriving) disableGetPassengersSpy() {
	m.mocked.GetPassengers.SpyEnabled = false
}

// GetPassengers overrides the method to return the mock response
func (m *mockSelfDriving) GetPassengers() []string {
	m.mocked.GetPassengers.RecordCall()
	var (
		out0 []string
	)

	if m.mocked.GetPassengers.Enabled {
		out0 = m.mocked.GetPassengers.NextResponse(func() []string {
			return m.real.GetPassengers()
		})()

	} else {
		out0 = m.real.GetPassengers()

	}

	if ch, ok := m.responseChans["GetPassengers"]; ok {
		chTyped := ch.(chan []string)
		chTyped <- out0
	}
	return out0

}

// setGetPassengersFunc sets the function for GetPassengers
func (m *mockSelfDriving) setGetPassengersFunc(f func() []string) {
	m.mocked.GetPassengers.Fallback = f
}

// enableGetPassengersMock turns the mock on
func (m *mockSelfDriving) enableGetPassengersMock() {
	m.mocked.GetPassengers.Enabled = true
}

// disableGetPassengersMock turns the mock off
func (m *mockSelfDriving) disableGetPassengersMock() {
	m.mocked.GetPassengers.Enabled = false
}

// enqueueGetPassengersResponseFunc enqueues a function response for GetPassengers
func (m *mockSelfDriving) enqueueGetPassengersResponseFunc(f func() []string) {
	m.mocked.GetPassengers.EnqueueWithDelay(f, 0)
}

// enqueueGetPassengersResponseFuncWithDelay enqueues a function response with delay for GetPassengers
func (m *mockSelfDriving) enqueueGetPassengersResponseFuncWithDelay(f func() []string, d time.Duration) {
	m.mocked.GetPassengers.EnqueueWithDelay(f, d)
}

// captureGetPassengersResult sets up a channel to capture GetPassengers results.
func (m *mockSelfDriving) captureGetPassengersResult() <-chan []string {
	ch := make(chan []string, 1)
	m.responseChans["GetPassengers"] = ch
	return ch
}

// captureGetPassengersCallSpy starts watching for GetPassengers spy calls and sends them into a channel.
func (m *mockSelfDriving) captureGetPassengersCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetPassengersCalls, timeout)
		ch <- m.getGetPassengersCalls()
	}()
	return ch
}

type mockSelfDrivingGetPassengersResult struct {
	Output0 []string
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockSelfDriving) setGetPassengersResponse(output0 []string) {
	m.setGetPassengersFunc(func() []string {
		return output0
	})
}

// enqueueGetPassengersResponse enqueues a static response for GetPassengers
func (m *mockSelfDriving) enqueueGetPassengersResponse(output0 []string) {
	m.mocked.GetPassengers.EnqueueWithDelay(func() []string {
		return output0
	}, 0)
}

// enqueueGetPassengersResponseWithDelay enqueues a static response with delay for GetPassengers
func (m *mockSelfDriving) enqueueGetPassengersResponseWithDelay(output0 []string, d time.Duration) {
	m.mocked.GetPassengers.EnqueueWithDelay(func() []string {
		time.Sleep(d)
		return output0
	}, d)
}
