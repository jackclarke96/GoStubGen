package driver

import (
	"fmt"

	"github.com/jackclarke/GoStubGen/templates/vehicle"
)

// mockVehicleConfig stores mock flags and responses
type mockVehicleConfig struct {

	// GetTopSpeed flag and mock response
	mockGetTopSpeed     bool
	GetTopSpeedResponse func() int

	// Turn flag and mock response
	mockTurn     bool
	TurnResponse func(string) string

	// Reverse flag and mock response
	mockReverse     bool
	ReverseResponse func() (string, error)

	// Accelerate flag and mock response
	mockAccelerate     bool
	AccelerateResponse func(int, string) (int, error)

	// IsMoving flag and mock response
	mockIsMoving     bool
	IsMovingResponse func() bool

	// Honk flag and mock response
	mockHonk     bool
	HonkResponse func(int)

	// GetEngineSpecs flag and mock response
	mockGetEngineSpecs     bool
	GetEngineSpecsResponse func() (int, string)

	// ApplyBrakes flag and mock response
	mockApplyBrakes     bool
	ApplyBrakesResponse func(float64) bool

	// ChangeGears flag and mock response
	mockChangeGears     bool
	ChangeGearsResponse func(int) (int, int)

	// Telemetry flag and mock response
	mockTelemetry     bool
	TelemetryResponse func() map[string]float64

	// GetPassengers flag and mock response
	mockGetPassengers     bool
	GetPassengersResponse func() []string

	// LoadCargo flag and mock response
	mockLoadCargo     bool
	LoadCargoResponse func([]string) (int, error)

	// GetVehicleStatus flag and mock response
	mockGetVehicleStatus     bool
	GetVehicleStatusResponse func() vehicle.VehicleStatus

	// UpdateStatus flag and mock response
	mockUpdateStatus     bool
	UpdateStatusResponse func(vehicle.VehicleStatus) error
}

// mockVehicle embeds Car and its mocks
type mockVehicle struct {
	car   vehicle.Car
	mocks mockVehicleConfig
}

// vehicleMock returns a new mock
func vehicleMock() *mockVehicle {
	car := vehicle.NewCar()
	return &mockVehicle{
		car:   *car,
		mocks: mockVehicleConfig{},
	}
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// GetTopSpeed overrides the method to return the mock response
func (m mockVehicle) GetTopSpeed() int {
	if m.mocks.mockGetTopSpeed {
		return m.mocks.GetTopSpeedResponse()
	}
	return m.car.GetTopSpeed()
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m mockVehicle) setGetTopSpeedResponse(resp func() int) {
	m.mocks.GetTopSpeedResponse = resp
}

// enableGetTopSpeedResponse turns the mock on
func (m mockVehicle) enableGetTopSpeedResponse() {
	m.mocks.mockGetTopSpeed = true
}

// disableGetTopSpeedResponse turns the mock off
func (m mockVehicle) disableGetTopSpeedResponse() {
	m.mocks.mockGetTopSpeed = false
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// Turn overrides the method to return the mock response
func (m mockVehicle) Turn(dir string) string {
	if m.mocks.mockTurn {
		return m.mocks.TurnResponse(dir)
	}
	return m.car.Turn(dir)
}

// setTurnResponse sets the response for Turn
func (m mockVehicle) setTurnResponse(resp func(string) string) {
	m.mocks.TurnResponse = resp
}

// enableTurnResponse turns the mock on
func (m mockVehicle) enableTurnResponse() {
	m.mocks.mockTurn = true
}

// disableTurnResponse turns the mock off
func (m mockVehicle) disableTurnResponse() {
	m.mocks.mockTurn = false
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// Reverse overrides the method to return the mock response
func (m mockVehicle) Reverse() (string, error) {
	if m.mocks.mockReverse {
		return m.mocks.ReverseResponse()
	}
	return m.car.Reverse()
}

// setReverseResponse sets the response for Reverse
func (m mockVehicle) setReverseResponse(resp func() (string, error)) {
	m.mocks.ReverseResponse = resp
}

// enableReverseResponse turns the mock on
func (m mockVehicle) enableReverseResponse() {
	m.mocks.mockReverse = true
}

// disableReverseResponse turns the mock off
func (m mockVehicle) disableReverseResponse() {
	m.mocks.mockReverse = false
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// Accelerate overrides the method to return the mock response
func (m mockVehicle) Accelerate(speed int, unit string) (int, error) {
	if m.mocks.mockAccelerate {
		return m.mocks.AccelerateResponse(speed, unit)
	}
	return m.car.Accelerate(speed, unit)
}

// setAccelerateResponse sets the response for Accelerate
func (m mockVehicle) setAccelerateResponse(resp func(int, string) (int, error)) {
	m.mocks.AccelerateResponse = resp
}

// enableAccelerateResponse turns the mock on
func (m mockVehicle) enableAccelerateResponse() {
	m.mocks.mockAccelerate = true
}

// disableAccelerateResponse turns the mock off
func (m mockVehicle) disableAccelerateResponse() {
	m.mocks.mockAccelerate = false
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// IsMoving overrides the method to return the mock response
func (m mockVehicle) IsMoving() bool {
	if m.mocks.mockIsMoving {
		return m.mocks.IsMovingResponse()
	}
	return m.car.IsMoving()
}

// setIsMovingResponse sets the response for IsMoving
func (m mockVehicle) setIsMovingResponse(resp func() bool) {
	m.mocks.IsMovingResponse = resp
}

// enableIsMovingResponse turns the mock on
func (m mockVehicle) enableIsMovingResponse() {
	m.mocks.mockIsMoving = true
}

// disableIsMovingResponse turns the mock off
func (m mockVehicle) disableIsMovingResponse() {
	m.mocks.mockIsMoving = false
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// Honk overrides the method to return the mock response
func (m mockVehicle) Honk(times int) {
	if m.mocks.mockHonk {
		m.mocks.HonkResponse(times)
	}
	m.car.Honk(times)
}

// setHonkResponse sets the response for Honk
func (m mockVehicle) setHonkResponse(resp func(int)) {
	m.mocks.HonkResponse = resp
}

// enableHonkResponse turns the mock on
func (m mockVehicle) enableHonkResponse() {
	m.mocks.mockHonk = true
}

// disableHonkResponse turns the mock off
func (m mockVehicle) disableHonkResponse() {
	m.mocks.mockHonk = false
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// GetEngineSpecs overrides the method to return the mock response
func (m mockVehicle) GetEngineSpecs() (int, string) {
	if m.mocks.mockGetEngineSpecs {
		return m.mocks.GetEngineSpecsResponse()
	}
	return m.car.GetEngineSpecs()
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m mockVehicle) setGetEngineSpecsResponse(resp func() (int, string)) {
	m.mocks.GetEngineSpecsResponse = resp
}

// enableGetEngineSpecsResponse turns the mock on
func (m mockVehicle) enableGetEngineSpecsResponse() {
	m.mocks.mockGetEngineSpecs = true
}

// disableGetEngineSpecsResponse turns the mock off
func (m mockVehicle) disableGetEngineSpecsResponse() {
	m.mocks.mockGetEngineSpecs = false
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// ApplyBrakes overrides the method to return the mock response
func (m mockVehicle) ApplyBrakes(force float64) bool {
	if m.mocks.mockApplyBrakes {
		return m.mocks.ApplyBrakesResponse(force)
	}
	return m.car.ApplyBrakes(force)
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m mockVehicle) setApplyBrakesResponse(resp func(float64) bool) {
	m.mocks.ApplyBrakesResponse = resp
}

// enableApplyBrakesResponse turns the mock on
func (m mockVehicle) enableApplyBrakesResponse() {
	m.mocks.mockApplyBrakes = true
}

// disableApplyBrakesResponse turns the mock off
func (m mockVehicle) disableApplyBrakesResponse() {
	m.mocks.mockApplyBrakes = false
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// ChangeGears overrides the method to return the mock response
func (m mockVehicle) ChangeGears(gear int) (int, int) {
	if m.mocks.mockChangeGears {
		return m.mocks.ChangeGearsResponse(gear)
	}
	return m.car.ChangeGears(gear)
}

// setChangeGearsResponse sets the response for ChangeGears
func (m mockVehicle) setChangeGearsResponse(resp func(int) (int, int)) {
	m.mocks.ChangeGearsResponse = resp
}

// enableChangeGearsResponse turns the mock on
func (m mockVehicle) enableChangeGearsResponse() {
	m.mocks.mockChangeGears = true
}

// disableChangeGearsResponse turns the mock off
func (m mockVehicle) disableChangeGearsResponse() {
	m.mocks.mockChangeGears = false
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// Telemetry overrides the method to return the mock response
func (m mockVehicle) Telemetry() map[string]float64 {
	if m.mocks.mockTelemetry {
		return m.mocks.TelemetryResponse()
	}
	return m.car.Telemetry()
}

// setTelemetryResponse sets the response for Telemetry
func (m mockVehicle) setTelemetryResponse(resp func() map[string]float64) {
	m.mocks.TelemetryResponse = resp
}

// enableTelemetryResponse turns the mock on
func (m mockVehicle) enableTelemetryResponse() {
	m.mocks.mockTelemetry = true
}

// disableTelemetryResponse turns the mock off
func (m mockVehicle) disableTelemetryResponse() {
	m.mocks.mockTelemetry = false
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// GetPassengers overrides the method to return the mock response
func (m mockVehicle) GetPassengers() []string {
	if m.mocks.mockGetPassengers {
		return m.mocks.GetPassengersResponse()
	}
	return m.car.GetPassengers()
}

// setGetPassengersResponse sets the response for GetPassengers
func (m mockVehicle) setGetPassengersResponse(resp func() []string) {
	m.mocks.GetPassengersResponse = resp
}

// enableGetPassengersResponse turns the mock on
func (m mockVehicle) enableGetPassengersResponse() {
	m.mocks.mockGetPassengers = true
}

// disableGetPassengersResponse turns the mock off
func (m mockVehicle) disableGetPassengersResponse() {
	m.mocks.mockGetPassengers = false
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// LoadCargo overrides the method to return the mock response
func (m *mockVehicle) LoadCargo(items []string) (int, error) {
	if m.mocks.mockLoadCargo {
		fmt.Println("returning mock")
		return m.mocks.LoadCargoResponse(items)
	}
	return m.car.LoadCargo(items)
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockVehicle) setLoadCargoResponse(resp func([]string) (int, error)) {
	m.mocks.LoadCargoResponse = resp
}

// enableLoadCargoResponse turns the mock on
func (m *mockVehicle) enableLoadCargoResponse() {
	fmt.Println("enabling mock")
	m.mocks.mockLoadCargo = true
}

// disableLoadCargoResponse turns the mock off
func (m *mockVehicle) disableLoadCargoResponse() {
	m.mocks.mockLoadCargo = false
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// GetVehicleStatus overrides the method to return the mock response
func (m mockVehicle) GetVehicleStatus() vehicle.VehicleStatus {
	if m.mocks.mockGetVehicleStatus {
		return m.mocks.GetVehicleStatusResponse()
	}
	return m.car.GetVehicleStatus()
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m mockVehicle) setGetVehicleStatusResponse(resp func() vehicle.VehicleStatus) {
	m.mocks.GetVehicleStatusResponse = resp
}

// enableGetVehicleStatusResponse turns the mock on
func (m mockVehicle) enableGetVehicleStatusResponse() {
	m.mocks.mockGetVehicleStatus = true
}

// disableGetVehicleStatusResponse turns the mock off
func (m mockVehicle) disableGetVehicleStatusResponse() {
	m.mocks.mockGetVehicleStatus = false
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// UpdateStatus overrides the method to return the mock response
func (m mockVehicle) UpdateStatus(status vehicle.VehicleStatus) error {
	if m.mocks.mockUpdateStatus {
		return m.mocks.UpdateStatusResponse(status)
	}
	return m.car.UpdateStatus(status)
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m mockVehicle) setUpdateStatusResponse(resp func(vehicle.VehicleStatus) error) {
	m.mocks.UpdateStatusResponse = resp
}

// enableUpdateStatusResponse turns the mock on
func (m mockVehicle) enableUpdateStatusResponse() {
	m.mocks.mockUpdateStatus = true
}

// disableUpdateStatusResponse turns the mock off
func (m mockVehicle) disableUpdateStatusResponse() {
	m.mocks.mockUpdateStatus = false
}

//todo:
// - mocks should import the package they are mocking.
