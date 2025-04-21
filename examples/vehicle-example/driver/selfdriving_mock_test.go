package driver

import (
	"time"

	"github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"
	"github.com/jackclarke/GoStubGen/stubs"
)

// mockSelfDrivingConfig stores mock flags and responses
type mockSelfDrivingConfig struct {
	UpdateStatus     stubs.MethodConfig[func(vehicle.VehicleStatus) error]
	Reverse          stubs.MethodConfig[func() (string, error)]
	TurnOffAC        stubs.MethodConfig[func() error]
	ApplyBrakes      stubs.MethodConfig[func(float64) bool]
	Turn             stubs.MethodConfig[func(string) string]
	ParkSelf         stubs.MethodConfig[func() error]
	TurnOffMusic     stubs.MethodConfig[func() error]
	GetPassengers    stubs.MethodConfig[func() []string]
	GetVehicleStatus stubs.MethodConfig[func() vehicle.VehicleStatus]
	GetTopSpeed      stubs.MethodConfig[func() int]
	DriveSelf        stubs.MethodConfig[func(string) error]
	LockDoors        stubs.MethodConfig[func() error]
	CloseWindows     stubs.MethodConfig[func() error]
	Accelerate       stubs.MethodConfig[func(int, string) (int, error)]
	IsMoving         stubs.MethodConfig[func() bool]
	Honk             stubs.MethodConfig[func(int)]
	ChangeGears      stubs.MethodConfig[func(int) (int, int)]
	Telemetry        stubs.MethodConfig[func() map[string]float64]
	GetEngineSpecs   stubs.MethodConfig[func() (int, string)]
	LoadCargo        stubs.MethodConfig[func([]string) (int, error)]
}

// mockSelfDriving embeds a concrete SelfDriving and its mocks
type mockSelfDriving struct {
	real   vehicle.SelfDriving
	mocked mockSelfDrivingConfig
}

// newSelfDrivingMock returns a new mock
func newSelfDrivingMock(v vehicle.SelfDriving) *mockSelfDriving {
	return &mockSelfDriving{
		real:   v,
		mocked: mockSelfDrivingConfig{},
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
	if m.mocked.UpdateStatus.Enabled {
		return m.mocked.UpdateStatus.NextResponse(func(status vehicle.VehicleStatus) error {
			return m.real.UpdateStatus(status)
		})(status)
	}
	return m.real.UpdateStatus(status)
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
	if m.mocked.Reverse.Enabled {
		return m.mocked.Reverse.NextResponse(func() (string, error) {
			return m.real.Reverse()
		})()
	}
	return m.real.Reverse()
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
	if m.mocked.TurnOffAC.Enabled {
		return m.mocked.TurnOffAC.NextResponse(func() error {
			return m.real.TurnOffAC()
		})()
	}
	return m.real.TurnOffAC()
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
	if m.mocked.ApplyBrakes.Enabled {
		return m.mocked.ApplyBrakes.NextResponse(func(force float64) bool {
			return m.real.ApplyBrakes(force)
		})(force)
	}
	return m.real.ApplyBrakes(force)
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
	if m.mocked.Turn.Enabled {
		return m.mocked.Turn.NextResponse(func(dir string) string {
			return m.real.Turn(dir)
		})(dir)
	}
	return m.real.Turn(dir)
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
	if m.mocked.ParkSelf.Enabled {
		return m.mocked.ParkSelf.NextResponse(func() error {
			return m.real.ParkSelf()
		})()
	}
	return m.real.ParkSelf()
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
	if m.mocked.TurnOffMusic.Enabled {
		return m.mocked.TurnOffMusic.NextResponse(func() error {
			return m.real.TurnOffMusic()
		})()
	}
	return m.real.TurnOffMusic()
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
	if m.mocked.GetPassengers.Enabled {
		return m.mocked.GetPassengers.NextResponse(func() []string {
			return m.real.GetPassengers()
		})()
	}
	return m.real.GetPassengers()
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
	if m.mocked.GetVehicleStatus.Enabled {
		return m.mocked.GetVehicleStatus.NextResponse(func() vehicle.VehicleStatus {
			return m.real.GetVehicleStatus()
		})()
	}
	return m.real.GetVehicleStatus()
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
	if m.mocked.GetTopSpeed.Enabled {
		return m.mocked.GetTopSpeed.NextResponse(func() int {
			return m.real.GetTopSpeed()
		})()
	}
	return m.real.GetTopSpeed()
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
	if m.mocked.DriveSelf.Enabled {
		return m.mocked.DriveSelf.NextResponse(func(endLocation string) error {
			return m.real.DriveSelf(endLocation)
		})(endLocation)
	}
	return m.real.DriveSelf(endLocation)
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
	if m.mocked.LockDoors.Enabled {
		return m.mocked.LockDoors.NextResponse(func() error {
			return m.real.LockDoors()
		})()
	}
	return m.real.LockDoors()
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
	if m.mocked.CloseWindows.Enabled {
		return m.mocked.CloseWindows.NextResponse(func() error {
			return m.real.CloseWindows()
		})()
	}
	return m.real.CloseWindows()
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
	if m.mocked.Accelerate.Enabled {
		return m.mocked.Accelerate.NextResponse(func(speed int, unit string) (int, error) {
			return m.real.Accelerate(speed, unit)
		})(speed, unit)
	}
	return m.real.Accelerate(speed, unit)
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
	if m.mocked.IsMoving.Enabled {
		return m.mocked.IsMoving.NextResponse(func() bool {
			return m.real.IsMoving()
		})()
	}
	return m.real.IsMoving()
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
	if m.mocked.Honk.Enabled {
		m.mocked.Honk.NextResponse(func(times int) {
			m.real.Honk(times)
		})(times)
		return
	}
	m.real.Honk(times)
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
	if m.mocked.ChangeGears.Enabled {
		return m.mocked.ChangeGears.NextResponse(func(gear int) (int, int) {
			return m.real.ChangeGears(gear)
		})(gear)
	}
	return m.real.ChangeGears(gear)
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
	if m.mocked.Telemetry.Enabled {
		return m.mocked.Telemetry.NextResponse(func() map[string]float64 {
			return m.real.Telemetry()
		})()
	}
	return m.real.Telemetry()
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
	if m.mocked.GetEngineSpecs.Enabled {
		return m.mocked.GetEngineSpecs.NextResponse(func() (int, string) {
			return m.real.GetEngineSpecs()
		})()
	}
	return m.real.GetEngineSpecs()
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
	if m.mocked.LoadCargo.Enabled {
		return m.mocked.LoadCargo.NextResponse(func(items []string) (int, error) {
			return m.real.LoadCargo(items)
		})(items)
	}
	return m.real.LoadCargo(items)
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
