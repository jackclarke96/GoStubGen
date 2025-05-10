package driver

import (
	"testing"
	"time"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
	"github.com/jackclarke/GoStubGen/stubs"
)

// mockVehicleConfig stores mock flags and responses
type mockVehicleConfig struct {
	GetTopSpeed      stubs.MethodConfig[func() int]
	Turn             stubs.MethodConfig[func(string) string]
	Reverse          stubs.MethodConfig[func() (string, error)]
	IsMoving         stubs.MethodConfig[func() bool]
	GetEngineSpecs   stubs.MethodConfig[func() (int, string)]
	ApplyBrakes      stubs.MethodConfig[func(float64) bool]
	ChangeGears      stubs.MethodConfig[func(int) (int, int)]
	Telemetry        stubs.MethodConfig[func() map[string]float64]
	Accelerate       stubs.MethodConfig[func(int, string) (int, error)]
	Honk             stubs.MethodConfig[func(int)]
	GetPassengers    stubs.MethodConfig[func() []string]
	LoadCargo        stubs.MethodConfig[func([]string) (int, error)]
	GetVehicleStatus stubs.MethodConfig[func() vehicle.VehicleStatus]
	UpdateStatus     stubs.MethodConfig[func(vehicle.VehicleStatus) error]
}

// mockVehicle embeds a concrete Vehicle and its mocks
type mockVehicle struct {
	real          vehicle.Vehicle
	mocked        mockVehicleConfig
	responseChans map[string]any
}

// newVehicleMock returns a new mock
func newVehicleMock(v vehicle.Vehicle) *mockVehicle {
	return &mockVehicle{
		real:          v,
		mocked:        mockVehicleConfig{},
		responseChans: make(map[string]any),
	}
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// enableGetTopSpeedMock turns the spy on
func (m *mockVehicle) enableGetTopSpeedSpy() {
	m.mocked.GetTopSpeed.SpyEnabled = true
}

// getGetTopSpeedCalls returns recorded calls to GetTopSpeed
func (m *mockVehicle) getGetTopSpeedCalls() []stubs.MethodCall {
	return m.mocked.GetTopSpeed.Calls()
}

// enableGetTopSpeedMock turns the spy off
func (m *mockVehicle) disableGetTopSpeedSpy() {
	m.mocked.GetTopSpeed.SpyEnabled = false
}

// GetTopSpeed overrides the method to return the mock response
func (m *mockVehicle) GetTopSpeed() int {
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
func (m *mockVehicle) setGetTopSpeedFunc(f func() int) {
	m.mocked.GetTopSpeed.Fallback = f
}

// enableGetTopSpeedMock turns the mock on
func (m *mockVehicle) enableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.Enabled = true
}

// disableGetTopSpeedMock turns the mock off
func (m *mockVehicle) disableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.Enabled = false
}

// enqueueGetTopSpeedResponseFunc enqueues a function response for GetTopSpeed
func (m *mockVehicle) enqueueGetTopSpeedResponseFunc(f func() int) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(f, 0)
}

// enqueueGetTopSpeedResponseFuncWithDelay enqueues a function response with delay for GetTopSpeed
func (m *mockVehicle) enqueueGetTopSpeedResponseFuncWithDelay(f func() int, d time.Duration) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(f, d)
}

// captureGetTopSpeedResult sets up a channel to capture GetTopSpeed results.
func (m *mockVehicle) captureGetTopSpeedResult() <-chan int {
	ch := make(chan int, 1)
	m.responseChans["GetTopSpeed"] = ch
	return ch
}

// captureGetTopSpeedCallSpy starts watching for GetTopSpeed spy calls and sends them into a channel.
func (m *mockVehicle) captureGetTopSpeedCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetTopSpeedCalls, timeout)
		ch <- m.getGetTopSpeedCalls()
	}()
	return ch
}

type mockVehicleGetTopSpeedResult struct {
	Output0 int
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockVehicle) setGetTopSpeedResponse(output0 int) {
	m.setGetTopSpeedFunc(func() int {
		return output0
	})
}

// enqueueGetTopSpeedResponse enqueues a static response for GetTopSpeed
func (m *mockVehicle) enqueueGetTopSpeedResponse(output0 int) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(func() int {
		return output0
	}, 0)
}

// enqueueGetTopSpeedResponseWithDelay enqueues a static response with delay for GetTopSpeed
func (m *mockVehicle) enqueueGetTopSpeedResponseWithDelay(output0 int, d time.Duration) {
	m.mocked.GetTopSpeed.EnqueueWithDelay(func() int {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// enableTurnMock turns the spy on
func (m *mockVehicle) enableTurnSpy() {
	m.mocked.Turn.SpyEnabled = true
}

// getTurnCalls returns recorded calls to Turn
func (m *mockVehicle) getTurnCalls() []stubs.MethodCall {
	return m.mocked.Turn.Calls()
}

// enableTurnMock turns the spy off
func (m *mockVehicle) disableTurnSpy() {
	m.mocked.Turn.SpyEnabled = false
}

// Turn overrides the method to return the mock response
func (m *mockVehicle) Turn(dir string) string {
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
func (m *mockVehicle) setTurnFunc(f func(string) string) {
	m.mocked.Turn.Fallback = f
}

// enableTurnMock turns the mock on
func (m *mockVehicle) enableTurnMock() {
	m.mocked.Turn.Enabled = true
}

// disableTurnMock turns the mock off
func (m *mockVehicle) disableTurnMock() {
	m.mocked.Turn.Enabled = false
}

// enqueueTurnResponseFunc enqueues a function response for Turn
func (m *mockVehicle) enqueueTurnResponseFunc(f func(string) string) {
	m.mocked.Turn.EnqueueWithDelay(f, 0)
}

// enqueueTurnResponseFuncWithDelay enqueues a function response with delay for Turn
func (m *mockVehicle) enqueueTurnResponseFuncWithDelay(f func(string) string, d time.Duration) {
	m.mocked.Turn.EnqueueWithDelay(f, d)
}

// captureTurnResult sets up a channel to capture Turn results.
func (m *mockVehicle) captureTurnResult() <-chan string {
	ch := make(chan string, 1)
	m.responseChans["Turn"] = ch
	return ch
}

// captureTurnCallSpy starts watching for Turn spy calls and sends them into a channel.
func (m *mockVehicle) captureTurnCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getTurnCalls, timeout)
		ch <- m.getTurnCalls()
	}()
	return ch
}

type mockVehicleTurnResult struct {
	Output0 string
}

// setTurnResponse sets the response for Turn
func (m *mockVehicle) setTurnResponse(output0 string) {
	m.setTurnFunc(func(string) string {
		return output0
	})
}

// enqueueTurnResponse enqueues a static response for Turn
func (m *mockVehicle) enqueueTurnResponse(output0 string) {
	m.mocked.Turn.EnqueueWithDelay(func(string) string {
		return output0
	}, 0)
}

// enqueueTurnResponseWithDelay enqueues a static response with delay for Turn
func (m *mockVehicle) enqueueTurnResponseWithDelay(output0 string, d time.Duration) {
	m.mocked.Turn.EnqueueWithDelay(func(string) string {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// enableReverseMock turns the spy on
func (m *mockVehicle) enableReverseSpy() {
	m.mocked.Reverse.SpyEnabled = true
}

// getReverseCalls returns recorded calls to Reverse
func (m *mockVehicle) getReverseCalls() []stubs.MethodCall {
	return m.mocked.Reverse.Calls()
}

// enableReverseMock turns the spy off
func (m *mockVehicle) disableReverseSpy() {
	m.mocked.Reverse.SpyEnabled = false
}

// Reverse overrides the method to return the mock response
func (m *mockVehicle) Reverse() (string, error) {
	m.mocked.Reverse.RecordCall()
	var (
		out0 string
		out1 error

		result mockVehicleReverseResult
	)

	if m.mocked.Reverse.Enabled {
		out0, out1 = m.mocked.Reverse.NextResponse(func() (string, error) {
			return m.real.Reverse()
		})()

	} else {
		out0, out1 = m.real.Reverse()

	}

	result = mockVehicleReverseResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["Reverse"]; ok {
		chTyped := ch.(chan mockVehicleReverseResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setReverseFunc sets the function for Reverse
func (m *mockVehicle) setReverseFunc(f func() (string, error)) {
	m.mocked.Reverse.Fallback = f
}

// enableReverseMock turns the mock on
func (m *mockVehicle) enableReverseMock() {
	m.mocked.Reverse.Enabled = true
}

// disableReverseMock turns the mock off
func (m *mockVehicle) disableReverseMock() {
	m.mocked.Reverse.Enabled = false
}

// enqueueReverseResponseFunc enqueues a function response for Reverse
func (m *mockVehicle) enqueueReverseResponseFunc(f func() (string, error)) {
	m.mocked.Reverse.EnqueueWithDelay(f, 0)
}

// enqueueReverseResponseFuncWithDelay enqueues a function response with delay for Reverse
func (m *mockVehicle) enqueueReverseResponseFuncWithDelay(f func() (string, error), d time.Duration) {
	m.mocked.Reverse.EnqueueWithDelay(f, d)
}

// captureReverseResult sets up a channel to capture Reverse results.
func (m *mockVehicle) captureReverseResult() <-chan mockVehicleReverseResult {
	ch := make(chan mockVehicleReverseResult, 1)
	m.responseChans["Reverse"] = ch
	return ch
}

// captureReverseCallSpy starts watching for Reverse spy calls and sends them into a channel.
func (m *mockVehicle) captureReverseCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getReverseCalls, timeout)
		ch <- m.getReverseCalls()
	}()
	return ch
}

type mockVehicleReverseResult struct {
	Output0 string
	Output1 error
}

// setReverseResponse sets the response for Reverse
func (m *mockVehicle) setReverseResponse(output0 string, output1 error) {
	m.setReverseFunc(func() (string, error) {
		return output0, output1
	})
}

// enqueueReverseResponse enqueues a static response for Reverse
func (m *mockVehicle) enqueueReverseResponse(output0 string, output1 error) {
	m.mocked.Reverse.EnqueueWithDelay(func() (string, error) {
		return output0, output1
	}, 0)
}

// enqueueReverseResponseWithDelay enqueues a static response with delay for Reverse
func (m *mockVehicle) enqueueReverseResponseWithDelay(output0 string, output1 error, d time.Duration) {
	m.mocked.Reverse.EnqueueWithDelay(func() (string, error) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// enableIsMovingMock turns the spy on
func (m *mockVehicle) enableIsMovingSpy() {
	m.mocked.IsMoving.SpyEnabled = true
}

// getIsMovingCalls returns recorded calls to IsMoving
func (m *mockVehicle) getIsMovingCalls() []stubs.MethodCall {
	return m.mocked.IsMoving.Calls()
}

// enableIsMovingMock turns the spy off
func (m *mockVehicle) disableIsMovingSpy() {
	m.mocked.IsMoving.SpyEnabled = false
}

// IsMoving overrides the method to return the mock response
func (m *mockVehicle) IsMoving() bool {
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
func (m *mockVehicle) setIsMovingFunc(f func() bool) {
	m.mocked.IsMoving.Fallback = f
}

// enableIsMovingMock turns the mock on
func (m *mockVehicle) enableIsMovingMock() {
	m.mocked.IsMoving.Enabled = true
}

// disableIsMovingMock turns the mock off
func (m *mockVehicle) disableIsMovingMock() {
	m.mocked.IsMoving.Enabled = false
}

// enqueueIsMovingResponseFunc enqueues a function response for IsMoving
func (m *mockVehicle) enqueueIsMovingResponseFunc(f func() bool) {
	m.mocked.IsMoving.EnqueueWithDelay(f, 0)
}

// enqueueIsMovingResponseFuncWithDelay enqueues a function response with delay for IsMoving
func (m *mockVehicle) enqueueIsMovingResponseFuncWithDelay(f func() bool, d time.Duration) {
	m.mocked.IsMoving.EnqueueWithDelay(f, d)
}

// captureIsMovingResult sets up a channel to capture IsMoving results.
func (m *mockVehicle) captureIsMovingResult() <-chan bool {
	ch := make(chan bool, 1)
	m.responseChans["IsMoving"] = ch
	return ch
}

// captureIsMovingCallSpy starts watching for IsMoving spy calls and sends them into a channel.
func (m *mockVehicle) captureIsMovingCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getIsMovingCalls, timeout)
		ch <- m.getIsMovingCalls()
	}()
	return ch
}

type mockVehicleIsMovingResult struct {
	Output0 bool
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockVehicle) setIsMovingResponse(output0 bool) {
	m.setIsMovingFunc(func() bool {
		return output0
	})
}

// enqueueIsMovingResponse enqueues a static response for IsMoving
func (m *mockVehicle) enqueueIsMovingResponse(output0 bool) {
	m.mocked.IsMoving.EnqueueWithDelay(func() bool {
		return output0
	}, 0)
}

// enqueueIsMovingResponseWithDelay enqueues a static response with delay for IsMoving
func (m *mockVehicle) enqueueIsMovingResponseWithDelay(output0 bool, d time.Duration) {
	m.mocked.IsMoving.EnqueueWithDelay(func() bool {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// enableGetEngineSpecsMock turns the spy on
func (m *mockVehicle) enableGetEngineSpecsSpy() {
	m.mocked.GetEngineSpecs.SpyEnabled = true
}

// getGetEngineSpecsCalls returns recorded calls to GetEngineSpecs
func (m *mockVehicle) getGetEngineSpecsCalls() []stubs.MethodCall {
	return m.mocked.GetEngineSpecs.Calls()
}

// enableGetEngineSpecsMock turns the spy off
func (m *mockVehicle) disableGetEngineSpecsSpy() {
	m.mocked.GetEngineSpecs.SpyEnabled = false
}

// GetEngineSpecs overrides the method to return the mock response
func (m *mockVehicle) GetEngineSpecs() (int, string) {
	m.mocked.GetEngineSpecs.RecordCall()
	var (
		out0 int
		out1 string

		result mockVehicleGetEngineSpecsResult
	)

	if m.mocked.GetEngineSpecs.Enabled {
		out0, out1 = m.mocked.GetEngineSpecs.NextResponse(func() (int, string) {
			return m.real.GetEngineSpecs()
		})()

	} else {
		out0, out1 = m.real.GetEngineSpecs()

	}

	result = mockVehicleGetEngineSpecsResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["GetEngineSpecs"]; ok {
		chTyped := ch.(chan mockVehicleGetEngineSpecsResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setGetEngineSpecsFunc sets the function for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.Fallback = f
}

// enableGetEngineSpecsMock turns the mock on
func (m *mockVehicle) enableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.Enabled = true
}

// disableGetEngineSpecsMock turns the mock off
func (m *mockVehicle) disableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.Enabled = false
}

// enqueueGetEngineSpecsResponseFunc enqueues a function response for GetEngineSpecs
func (m *mockVehicle) enqueueGetEngineSpecsResponseFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(f, 0)
}

// enqueueGetEngineSpecsResponseFuncWithDelay enqueues a function response with delay for GetEngineSpecs
func (m *mockVehicle) enqueueGetEngineSpecsResponseFuncWithDelay(f func() (int, string), d time.Duration) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(f, d)
}

// captureGetEngineSpecsResult sets up a channel to capture GetEngineSpecs results.
func (m *mockVehicle) captureGetEngineSpecsResult() <-chan mockVehicleGetEngineSpecsResult {
	ch := make(chan mockVehicleGetEngineSpecsResult, 1)
	m.responseChans["GetEngineSpecs"] = ch
	return ch
}

// captureGetEngineSpecsCallSpy starts watching for GetEngineSpecs spy calls and sends them into a channel.
func (m *mockVehicle) captureGetEngineSpecsCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetEngineSpecsCalls, timeout)
		ch <- m.getGetEngineSpecsCalls()
	}()
	return ch
}

type mockVehicleGetEngineSpecsResult struct {
	Output0 int
	Output1 string
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsResponse(output0 int, output1 string) {
	m.setGetEngineSpecsFunc(func() (int, string) {
		return output0, output1
	})
}

// enqueueGetEngineSpecsResponse enqueues a static response for GetEngineSpecs
func (m *mockVehicle) enqueueGetEngineSpecsResponse(output0 int, output1 string) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(func() (int, string) {
		return output0, output1
	}, 0)
}

// enqueueGetEngineSpecsResponseWithDelay enqueues a static response with delay for GetEngineSpecs
func (m *mockVehicle) enqueueGetEngineSpecsResponseWithDelay(output0 int, output1 string, d time.Duration) {
	m.mocked.GetEngineSpecs.EnqueueWithDelay(func() (int, string) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// enableApplyBrakesMock turns the spy on
func (m *mockVehicle) enableApplyBrakesSpy() {
	m.mocked.ApplyBrakes.SpyEnabled = true
}

// getApplyBrakesCalls returns recorded calls to ApplyBrakes
func (m *mockVehicle) getApplyBrakesCalls() []stubs.MethodCall {
	return m.mocked.ApplyBrakes.Calls()
}

// enableApplyBrakesMock turns the spy off
func (m *mockVehicle) disableApplyBrakesSpy() {
	m.mocked.ApplyBrakes.SpyEnabled = false
}

// ApplyBrakes overrides the method to return the mock response
func (m *mockVehicle) ApplyBrakes(force float64) bool {
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
func (m *mockVehicle) setApplyBrakesFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.Fallback = f
}

// enableApplyBrakesMock turns the mock on
func (m *mockVehicle) enableApplyBrakesMock() {
	m.mocked.ApplyBrakes.Enabled = true
}

// disableApplyBrakesMock turns the mock off
func (m *mockVehicle) disableApplyBrakesMock() {
	m.mocked.ApplyBrakes.Enabled = false
}

// enqueueApplyBrakesResponseFunc enqueues a function response for ApplyBrakes
func (m *mockVehicle) enqueueApplyBrakesResponseFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(f, 0)
}

// enqueueApplyBrakesResponseFuncWithDelay enqueues a function response with delay for ApplyBrakes
func (m *mockVehicle) enqueueApplyBrakesResponseFuncWithDelay(f func(float64) bool, d time.Duration) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(f, d)
}

// captureApplyBrakesResult sets up a channel to capture ApplyBrakes results.
func (m *mockVehicle) captureApplyBrakesResult() <-chan bool {
	ch := make(chan bool, 1)
	m.responseChans["ApplyBrakes"] = ch
	return ch
}

// captureApplyBrakesCallSpy starts watching for ApplyBrakes spy calls and sends them into a channel.
func (m *mockVehicle) captureApplyBrakesCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getApplyBrakesCalls, timeout)
		ch <- m.getApplyBrakesCalls()
	}()
	return ch
}

type mockVehicleApplyBrakesResult struct {
	Output0 bool
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockVehicle) setApplyBrakesResponse(output0 bool) {
	m.setApplyBrakesFunc(func(float64) bool {
		return output0
	})
}

// enqueueApplyBrakesResponse enqueues a static response for ApplyBrakes
func (m *mockVehicle) enqueueApplyBrakesResponse(output0 bool) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(func(float64) bool {
		return output0
	}, 0)
}

// enqueueApplyBrakesResponseWithDelay enqueues a static response with delay for ApplyBrakes
func (m *mockVehicle) enqueueApplyBrakesResponseWithDelay(output0 bool, d time.Duration) {
	m.mocked.ApplyBrakes.EnqueueWithDelay(func(float64) bool {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// enableChangeGearsMock turns the spy on
func (m *mockVehicle) enableChangeGearsSpy() {
	m.mocked.ChangeGears.SpyEnabled = true
}

// getChangeGearsCalls returns recorded calls to ChangeGears
func (m *mockVehicle) getChangeGearsCalls() []stubs.MethodCall {
	return m.mocked.ChangeGears.Calls()
}

// enableChangeGearsMock turns the spy off
func (m *mockVehicle) disableChangeGearsSpy() {
	m.mocked.ChangeGears.SpyEnabled = false
}

// ChangeGears overrides the method to return the mock response
func (m *mockVehicle) ChangeGears(gear int) (int, int) {
	m.mocked.ChangeGears.RecordCall(gear)
	var (
		out0 int
		out1 int

		result mockVehicleChangeGearsResult
	)

	if m.mocked.ChangeGears.Enabled {
		out0, out1 = m.mocked.ChangeGears.NextResponse(func(gear int) (int, int) {
			return m.real.ChangeGears(gear)
		})(gear)

	} else {
		out0, out1 = m.real.ChangeGears(gear)

	}

	result = mockVehicleChangeGearsResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["ChangeGears"]; ok {
		chTyped := ch.(chan mockVehicleChangeGearsResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setChangeGearsFunc sets the function for ChangeGears
func (m *mockVehicle) setChangeGearsFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.Fallback = f
}

// enableChangeGearsMock turns the mock on
func (m *mockVehicle) enableChangeGearsMock() {
	m.mocked.ChangeGears.Enabled = true
}

// disableChangeGearsMock turns the mock off
func (m *mockVehicle) disableChangeGearsMock() {
	m.mocked.ChangeGears.Enabled = false
}

// enqueueChangeGearsResponseFunc enqueues a function response for ChangeGears
func (m *mockVehicle) enqueueChangeGearsResponseFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.EnqueueWithDelay(f, 0)
}

// enqueueChangeGearsResponseFuncWithDelay enqueues a function response with delay for ChangeGears
func (m *mockVehicle) enqueueChangeGearsResponseFuncWithDelay(f func(int) (int, int), d time.Duration) {
	m.mocked.ChangeGears.EnqueueWithDelay(f, d)
}

// captureChangeGearsResult sets up a channel to capture ChangeGears results.
func (m *mockVehicle) captureChangeGearsResult() <-chan mockVehicleChangeGearsResult {
	ch := make(chan mockVehicleChangeGearsResult, 1)
	m.responseChans["ChangeGears"] = ch
	return ch
}

// captureChangeGearsCallSpy starts watching for ChangeGears spy calls and sends them into a channel.
func (m *mockVehicle) captureChangeGearsCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getChangeGearsCalls, timeout)
		ch <- m.getChangeGearsCalls()
	}()
	return ch
}

type mockVehicleChangeGearsResult struct {
	Output0 int
	Output1 int
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockVehicle) setChangeGearsResponse(output0 int, output1 int) {
	m.setChangeGearsFunc(func(int) (int, int) {
		return output0, output1
	})
}

// enqueueChangeGearsResponse enqueues a static response for ChangeGears
func (m *mockVehicle) enqueueChangeGearsResponse(output0 int, output1 int) {
	m.mocked.ChangeGears.EnqueueWithDelay(func(int) (int, int) {
		return output0, output1
	}, 0)
}

// enqueueChangeGearsResponseWithDelay enqueues a static response with delay for ChangeGears
func (m *mockVehicle) enqueueChangeGearsResponseWithDelay(output0 int, output1 int, d time.Duration) {
	m.mocked.ChangeGears.EnqueueWithDelay(func(int) (int, int) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// enableTelemetryMock turns the spy on
func (m *mockVehicle) enableTelemetrySpy() {
	m.mocked.Telemetry.SpyEnabled = true
}

// getTelemetryCalls returns recorded calls to Telemetry
func (m *mockVehicle) getTelemetryCalls() []stubs.MethodCall {
	return m.mocked.Telemetry.Calls()
}

// enableTelemetryMock turns the spy off
func (m *mockVehicle) disableTelemetrySpy() {
	m.mocked.Telemetry.SpyEnabled = false
}

// Telemetry overrides the method to return the mock response
func (m *mockVehicle) Telemetry() map[string]float64 {
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
func (m *mockVehicle) setTelemetryFunc(f func() map[string]float64) {
	m.mocked.Telemetry.Fallback = f
}

// enableTelemetryMock turns the mock on
func (m *mockVehicle) enableTelemetryMock() {
	m.mocked.Telemetry.Enabled = true
}

// disableTelemetryMock turns the mock off
func (m *mockVehicle) disableTelemetryMock() {
	m.mocked.Telemetry.Enabled = false
}

// enqueueTelemetryResponseFunc enqueues a function response for Telemetry
func (m *mockVehicle) enqueueTelemetryResponseFunc(f func() map[string]float64) {
	m.mocked.Telemetry.EnqueueWithDelay(f, 0)
}

// enqueueTelemetryResponseFuncWithDelay enqueues a function response with delay for Telemetry
func (m *mockVehicle) enqueueTelemetryResponseFuncWithDelay(f func() map[string]float64, d time.Duration) {
	m.mocked.Telemetry.EnqueueWithDelay(f, d)
}

// captureTelemetryResult sets up a channel to capture Telemetry results.
func (m *mockVehicle) captureTelemetryResult() <-chan map[string]float64 {
	ch := make(chan map[string]float64, 1)
	m.responseChans["Telemetry"] = ch
	return ch
}

// captureTelemetryCallSpy starts watching for Telemetry spy calls and sends them into a channel.
func (m *mockVehicle) captureTelemetryCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getTelemetryCalls, timeout)
		ch <- m.getTelemetryCalls()
	}()
	return ch
}

type mockVehicleTelemetryResult struct {
	Output0 map[string]float64
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockVehicle) setTelemetryResponse(output0 map[string]float64) {
	m.setTelemetryFunc(func() map[string]float64 {
		return output0
	})
}

// enqueueTelemetryResponse enqueues a static response for Telemetry
func (m *mockVehicle) enqueueTelemetryResponse(output0 map[string]float64) {
	m.mocked.Telemetry.EnqueueWithDelay(func() map[string]float64 {
		return output0
	}, 0)
}

// enqueueTelemetryResponseWithDelay enqueues a static response with delay for Telemetry
func (m *mockVehicle) enqueueTelemetryResponseWithDelay(output0 map[string]float64, d time.Duration) {
	m.mocked.Telemetry.EnqueueWithDelay(func() map[string]float64 {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// enableAccelerateMock turns the spy on
func (m *mockVehicle) enableAccelerateSpy() {
	m.mocked.Accelerate.SpyEnabled = true
}

// getAccelerateCalls returns recorded calls to Accelerate
func (m *mockVehicle) getAccelerateCalls() []stubs.MethodCall {
	return m.mocked.Accelerate.Calls()
}

// enableAccelerateMock turns the spy off
func (m *mockVehicle) disableAccelerateSpy() {
	m.mocked.Accelerate.SpyEnabled = false
}

// Accelerate overrides the method to return the mock response
func (m *mockVehicle) Accelerate(speed int, unit string) (int, error) {
	m.mocked.Accelerate.RecordCall(speed, unit)
	var (
		out0 int
		out1 error

		result mockVehicleAccelerateResult
	)

	if m.mocked.Accelerate.Enabled {
		out0, out1 = m.mocked.Accelerate.NextResponse(func(speed int, unit string) (int, error) {
			return m.real.Accelerate(speed, unit)
		})(speed, unit)

	} else {
		out0, out1 = m.real.Accelerate(speed, unit)

	}

	result = mockVehicleAccelerateResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["Accelerate"]; ok {
		chTyped := ch.(chan mockVehicleAccelerateResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setAccelerateFunc sets the function for Accelerate
func (m *mockVehicle) setAccelerateFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.Fallback = f
}

// enableAccelerateMock turns the mock on
func (m *mockVehicle) enableAccelerateMock() {
	m.mocked.Accelerate.Enabled = true
}

// disableAccelerateMock turns the mock off
func (m *mockVehicle) disableAccelerateMock() {
	m.mocked.Accelerate.Enabled = false
}

// enqueueAccelerateResponseFunc enqueues a function response for Accelerate
func (m *mockVehicle) enqueueAccelerateResponseFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.EnqueueWithDelay(f, 0)
}

// enqueueAccelerateResponseFuncWithDelay enqueues a function response with delay for Accelerate
func (m *mockVehicle) enqueueAccelerateResponseFuncWithDelay(f func(int, string) (int, error), d time.Duration) {
	m.mocked.Accelerate.EnqueueWithDelay(f, d)
}

// captureAccelerateResult sets up a channel to capture Accelerate results.
func (m *mockVehicle) captureAccelerateResult() <-chan mockVehicleAccelerateResult {
	ch := make(chan mockVehicleAccelerateResult, 1)
	m.responseChans["Accelerate"] = ch
	return ch
}

// captureAccelerateCallSpy starts watching for Accelerate spy calls and sends them into a channel.
func (m *mockVehicle) captureAccelerateCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getAccelerateCalls, timeout)
		ch <- m.getAccelerateCalls()
	}()
	return ch
}

type mockVehicleAccelerateResult struct {
	Output0 int
	Output1 error
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockVehicle) setAccelerateResponse(output0 int, output1 error) {
	m.setAccelerateFunc(func(int, string) (int, error) {
		return output0, output1
	})
}

// enqueueAccelerateResponse enqueues a static response for Accelerate
func (m *mockVehicle) enqueueAccelerateResponse(output0 int, output1 error) {
	m.mocked.Accelerate.EnqueueWithDelay(func(int, string) (int, error) {
		return output0, output1
	}, 0)
}

// enqueueAccelerateResponseWithDelay enqueues a static response with delay for Accelerate
func (m *mockVehicle) enqueueAccelerateResponseWithDelay(output0 int, output1 error, d time.Duration) {
	m.mocked.Accelerate.EnqueueWithDelay(func(int, string) (int, error) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// enableHonkMock turns the spy on
func (m *mockVehicle) enableHonkSpy() {
	m.mocked.Honk.SpyEnabled = true
}

// getHonkCalls returns recorded calls to Honk
func (m *mockVehicle) getHonkCalls() []stubs.MethodCall {
	return m.mocked.Honk.Calls()
}

// enableHonkMock turns the spy off
func (m *mockVehicle) disableHonkSpy() {
	m.mocked.Honk.SpyEnabled = false
}

// Honk overrides the method to return the mock response
func (m *mockVehicle) Honk(times int) {
	m.mocked.Honk.RecordCall(times)
	var ()

	if m.mocked.Honk.Enabled {

	} else {

	}

	return

}

// setHonkFunc sets the function for Honk
func (m *mockVehicle) setHonkFunc(f func(int)) {
	m.mocked.Honk.Fallback = f
}

// enableHonkMock turns the mock on
func (m *mockVehicle) enableHonkMock() {
	m.mocked.Honk.Enabled = true
}

// disableHonkMock turns the mock off
func (m *mockVehicle) disableHonkMock() {
	m.mocked.Honk.Enabled = false
}

// enqueueHonkResponseFunc enqueues a function response for Honk
func (m *mockVehicle) enqueueHonkResponseFunc(f func(int)) {
	m.mocked.Honk.EnqueueWithDelay(f, 0)
}

// enqueueHonkResponseFuncWithDelay enqueues a function response with delay for Honk
func (m *mockVehicle) enqueueHonkResponseFuncWithDelay(f func(int), d time.Duration) {
	m.mocked.Honk.EnqueueWithDelay(f, d)
}

// captureHonkResult sets up a channel to capture Honk results.
func (m *mockVehicle) captureHonkResult() <-chan struct{} {
	ch := make(chan struct{}, 1)
	m.responseChans["Honk"] = ch
	return ch
}

// captureHonkCallSpy starts watching for Honk spy calls and sends them into a channel.
func (m *mockVehicle) captureHonkCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getHonkCalls, timeout)
		ch <- m.getHonkCalls()
	}()
	return ch
}

type mockVehicleHonkResult struct {
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// enableGetPassengersMock turns the spy on
func (m *mockVehicle) enableGetPassengersSpy() {
	m.mocked.GetPassengers.SpyEnabled = true
}

// getGetPassengersCalls returns recorded calls to GetPassengers
func (m *mockVehicle) getGetPassengersCalls() []stubs.MethodCall {
	return m.mocked.GetPassengers.Calls()
}

// enableGetPassengersMock turns the spy off
func (m *mockVehicle) disableGetPassengersSpy() {
	m.mocked.GetPassengers.SpyEnabled = false
}

// GetPassengers overrides the method to return the mock response
func (m *mockVehicle) GetPassengers() []string {
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
func (m *mockVehicle) setGetPassengersFunc(f func() []string) {
	m.mocked.GetPassengers.Fallback = f
}

// enableGetPassengersMock turns the mock on
func (m *mockVehicle) enableGetPassengersMock() {
	m.mocked.GetPassengers.Enabled = true
}

// disableGetPassengersMock turns the mock off
func (m *mockVehicle) disableGetPassengersMock() {
	m.mocked.GetPassengers.Enabled = false
}

// enqueueGetPassengersResponseFunc enqueues a function response for GetPassengers
func (m *mockVehicle) enqueueGetPassengersResponseFunc(f func() []string) {
	m.mocked.GetPassengers.EnqueueWithDelay(f, 0)
}

// enqueueGetPassengersResponseFuncWithDelay enqueues a function response with delay for GetPassengers
func (m *mockVehicle) enqueueGetPassengersResponseFuncWithDelay(f func() []string, d time.Duration) {
	m.mocked.GetPassengers.EnqueueWithDelay(f, d)
}

// captureGetPassengersResult sets up a channel to capture GetPassengers results.
func (m *mockVehicle) captureGetPassengersResult() <-chan []string {
	ch := make(chan []string, 1)
	m.responseChans["GetPassengers"] = ch
	return ch
}

// captureGetPassengersCallSpy starts watching for GetPassengers spy calls and sends them into a channel.
func (m *mockVehicle) captureGetPassengersCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetPassengersCalls, timeout)
		ch <- m.getGetPassengersCalls()
	}()
	return ch
}

type mockVehicleGetPassengersResult struct {
	Output0 []string
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockVehicle) setGetPassengersResponse(output0 []string) {
	m.setGetPassengersFunc(func() []string {
		return output0
	})
}

// enqueueGetPassengersResponse enqueues a static response for GetPassengers
func (m *mockVehicle) enqueueGetPassengersResponse(output0 []string) {
	m.mocked.GetPassengers.EnqueueWithDelay(func() []string {
		return output0
	}, 0)
}

// enqueueGetPassengersResponseWithDelay enqueues a static response with delay for GetPassengers
func (m *mockVehicle) enqueueGetPassengersResponseWithDelay(output0 []string, d time.Duration) {
	m.mocked.GetPassengers.EnqueueWithDelay(func() []string {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// enableLoadCargoMock turns the spy on
func (m *mockVehicle) enableLoadCargoSpy() {
	m.mocked.LoadCargo.SpyEnabled = true
}

// getLoadCargoCalls returns recorded calls to LoadCargo
func (m *mockVehicle) getLoadCargoCalls() []stubs.MethodCall {
	return m.mocked.LoadCargo.Calls()
}

// enableLoadCargoMock turns the spy off
func (m *mockVehicle) disableLoadCargoSpy() {
	m.mocked.LoadCargo.SpyEnabled = false
}

// LoadCargo overrides the method to return the mock response
func (m *mockVehicle) LoadCargo(items []string) (int, error) {
	m.mocked.LoadCargo.RecordCall(items)
	var (
		out0 int
		out1 error

		result mockVehicleLoadCargoResult
	)

	if m.mocked.LoadCargo.Enabled {
		out0, out1 = m.mocked.LoadCargo.NextResponse(func(items []string) (int, error) {
			return m.real.LoadCargo(items)
		})(items)

	} else {
		out0, out1 = m.real.LoadCargo(items)

	}

	result = mockVehicleLoadCargoResult{
		Output0: out0, Output1: out1,
	}
	if ch, ok := m.responseChans["LoadCargo"]; ok {
		chTyped := ch.(chan mockVehicleLoadCargoResult)
		chTyped <- result
	}
	return result.Output0, result.Output1

}

// setLoadCargoFunc sets the function for LoadCargo
func (m *mockVehicle) setLoadCargoFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.Fallback = f
}

// enableLoadCargoMock turns the mock on
func (m *mockVehicle) enableLoadCargoMock() {
	m.mocked.LoadCargo.Enabled = true
}

// disableLoadCargoMock turns the mock off
func (m *mockVehicle) disableLoadCargoMock() {
	m.mocked.LoadCargo.Enabled = false
}

// enqueueLoadCargoResponseFunc enqueues a function response for LoadCargo
func (m *mockVehicle) enqueueLoadCargoResponseFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.EnqueueWithDelay(f, 0)
}

// enqueueLoadCargoResponseFuncWithDelay enqueues a function response with delay for LoadCargo
func (m *mockVehicle) enqueueLoadCargoResponseFuncWithDelay(f func([]string) (int, error), d time.Duration) {
	m.mocked.LoadCargo.EnqueueWithDelay(f, d)
}

// captureLoadCargoResult sets up a channel to capture LoadCargo results.
func (m *mockVehicle) captureLoadCargoResult() <-chan mockVehicleLoadCargoResult {
	ch := make(chan mockVehicleLoadCargoResult, 1)
	m.responseChans["LoadCargo"] = ch
	return ch
}

// captureLoadCargoCallSpy starts watching for LoadCargo spy calls and sends them into a channel.
func (m *mockVehicle) captureLoadCargoCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getLoadCargoCalls, timeout)
		ch <- m.getLoadCargoCalls()
	}()
	return ch
}

type mockVehicleLoadCargoResult struct {
	Output0 int
	Output1 error
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockVehicle) setLoadCargoResponse(output0 int, output1 error) {
	m.setLoadCargoFunc(func([]string) (int, error) {
		return output0, output1
	})
}

// enqueueLoadCargoResponse enqueues a static response for LoadCargo
func (m *mockVehicle) enqueueLoadCargoResponse(output0 int, output1 error) {
	m.mocked.LoadCargo.EnqueueWithDelay(func([]string) (int, error) {
		return output0, output1
	}, 0)
}

// enqueueLoadCargoResponseWithDelay enqueues a static response with delay for LoadCargo
func (m *mockVehicle) enqueueLoadCargoResponseWithDelay(output0 int, output1 error, d time.Duration) {
	m.mocked.LoadCargo.EnqueueWithDelay(func([]string) (int, error) {
		time.Sleep(d)
		return output0, output1
	}, d)
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// enableGetVehicleStatusMock turns the spy on
func (m *mockVehicle) enableGetVehicleStatusSpy() {
	m.mocked.GetVehicleStatus.SpyEnabled = true
}

// getGetVehicleStatusCalls returns recorded calls to GetVehicleStatus
func (m *mockVehicle) getGetVehicleStatusCalls() []stubs.MethodCall {
	return m.mocked.GetVehicleStatus.Calls()
}

// enableGetVehicleStatusMock turns the spy off
func (m *mockVehicle) disableGetVehicleStatusSpy() {
	m.mocked.GetVehicleStatus.SpyEnabled = false
}

// GetVehicleStatus overrides the method to return the mock response
func (m *mockVehicle) GetVehicleStatus() vehicle.VehicleStatus {
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
func (m *mockVehicle) setGetVehicleStatusFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.Fallback = f
}

// enableGetVehicleStatusMock turns the mock on
func (m *mockVehicle) enableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.Enabled = true
}

// disableGetVehicleStatusMock turns the mock off
func (m *mockVehicle) disableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.Enabled = false
}

// enqueueGetVehicleStatusResponseFunc enqueues a function response for GetVehicleStatus
func (m *mockVehicle) enqueueGetVehicleStatusResponseFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(f, 0)
}

// enqueueGetVehicleStatusResponseFuncWithDelay enqueues a function response with delay for GetVehicleStatus
func (m *mockVehicle) enqueueGetVehicleStatusResponseFuncWithDelay(f func() vehicle.VehicleStatus, d time.Duration) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(f, d)
}

// captureGetVehicleStatusResult sets up a channel to capture GetVehicleStatus results.
func (m *mockVehicle) captureGetVehicleStatusResult() <-chan vehicle.VehicleStatus {
	ch := make(chan vehicle.VehicleStatus, 1)
	m.responseChans["GetVehicleStatus"] = ch
	return ch
}

// captureGetVehicleStatusCallSpy starts watching for GetVehicleStatus spy calls and sends them into a channel.
func (m *mockVehicle) captureGetVehicleStatusCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getGetVehicleStatusCalls, timeout)
		ch <- m.getGetVehicleStatusCalls()
	}()
	return ch
}

type mockVehicleGetVehicleStatusResult struct {
	Output0 vehicle.VehicleStatus
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockVehicle) setGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.setGetVehicleStatusFunc(func() vehicle.VehicleStatus {
		return output0
	})
}

// enqueueGetVehicleStatusResponse enqueues a static response for GetVehicleStatus
func (m *mockVehicle) enqueueGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(func() vehicle.VehicleStatus {
		return output0
	}, 0)
}

// enqueueGetVehicleStatusResponseWithDelay enqueues a static response with delay for GetVehicleStatus
func (m *mockVehicle) enqueueGetVehicleStatusResponseWithDelay(output0 vehicle.VehicleStatus, d time.Duration) {
	m.mocked.GetVehicleStatus.EnqueueWithDelay(func() vehicle.VehicleStatus {
		time.Sleep(d)
		return output0
	}, d)
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// enableUpdateStatusMock turns the spy on
func (m *mockVehicle) enableUpdateStatusSpy() {
	m.mocked.UpdateStatus.SpyEnabled = true
}

// getUpdateStatusCalls returns recorded calls to UpdateStatus
func (m *mockVehicle) getUpdateStatusCalls() []stubs.MethodCall {
	return m.mocked.UpdateStatus.Calls()
}

// enableUpdateStatusMock turns the spy off
func (m *mockVehicle) disableUpdateStatusSpy() {
	m.mocked.UpdateStatus.SpyEnabled = false
}

// UpdateStatus overrides the method to return the mock response
func (m *mockVehicle) UpdateStatus(status vehicle.VehicleStatus) error {
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
func (m *mockVehicle) setUpdateStatusFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.Fallback = f
}

// enableUpdateStatusMock turns the mock on
func (m *mockVehicle) enableUpdateStatusMock() {
	m.mocked.UpdateStatus.Enabled = true
}

// disableUpdateStatusMock turns the mock off
func (m *mockVehicle) disableUpdateStatusMock() {
	m.mocked.UpdateStatus.Enabled = false
}

// enqueueUpdateStatusResponseFunc enqueues a function response for UpdateStatus
func (m *mockVehicle) enqueueUpdateStatusResponseFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.EnqueueWithDelay(f, 0)
}

// enqueueUpdateStatusResponseFuncWithDelay enqueues a function response with delay for UpdateStatus
func (m *mockVehicle) enqueueUpdateStatusResponseFuncWithDelay(f func(vehicle.VehicleStatus) error, d time.Duration) {
	m.mocked.UpdateStatus.EnqueueWithDelay(f, d)
}

// captureUpdateStatusResult sets up a channel to capture UpdateStatus results.
func (m *mockVehicle) captureUpdateStatusResult() <-chan error {
	ch := make(chan error, 1)
	m.responseChans["UpdateStatus"] = ch
	return ch
}

// captureUpdateStatusCallSpy starts watching for UpdateStatus spy calls and sends them into a channel.
func (m *mockVehicle) captureUpdateStatusCallSpy(t *testing.T, timeout time.Duration) <-chan []stubs.MethodCall {
	ch := make(chan []stubs.MethodCall, 1)
	go func() {
		stubs.WaitForSpyCall(t, m.getUpdateStatusCalls, timeout)
		ch <- m.getUpdateStatusCalls()
	}()
	return ch
}

type mockVehicleUpdateStatusResult struct {
	Output0 error
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockVehicle) setUpdateStatusResponse(output0 error) {
	m.setUpdateStatusFunc(func(vehicle.VehicleStatus) error {
		return output0
	})
}

// enqueueUpdateStatusResponse enqueues a static response for UpdateStatus
func (m *mockVehicle) enqueueUpdateStatusResponse(output0 error) {
	m.mocked.UpdateStatus.EnqueueWithDelay(func(vehicle.VehicleStatus) error {
		return output0
	}, 0)
}

// enqueueUpdateStatusResponseWithDelay enqueues a static response with delay for UpdateStatus
func (m *mockVehicle) enqueueUpdateStatusResponseWithDelay(output0 error, d time.Duration) {
	m.mocked.UpdateStatus.EnqueueWithDelay(func(vehicle.VehicleStatus) error {
		time.Sleep(d)
		return output0
	}, d)
}
