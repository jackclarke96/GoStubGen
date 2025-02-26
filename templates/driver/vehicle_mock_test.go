package driver

import "github.com/jackclarke/GoStubGen/templates/vehicle"

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

// mockVehicle embeds a concrete Vehicle and its mocks
type mockVehicle struct {
	real   vehicle.Vehicle
	mocked mockVehicleConfig
}

// vehicleMock returns a new mock
func vehicleMock(v vehicle.Vehicle) *mockVehicle {
	return &mockVehicle{
		real:   v,
		mocked: mockVehicleConfig{},
	}
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// GetTopSpeed overrides the method to return the mock response
func (m *mockVehicle) GetTopSpeed() int {
	if m.mocked.mockGetTopSpeed {
		return m.mocked.GetTopSpeedResponse()
	}
	return m.real.GetTopSpeed()
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockVehicle) setGetTopSpeedResponse(resp func() int) {
	m.mocked.GetTopSpeedResponse = resp
}

// enableGetTopSpeedResponse turns the mock on
func (m *mockVehicle) enableGetTopSpeedResponse() {
	m.mocked.mockGetTopSpeed = true
}

// disableGetTopSpeedResponse turns the mock off
func (m *mockVehicle) disableGetTopSpeedResponse() {
	m.mocked.mockGetTopSpeed = false
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// Turn overrides the method to return the mock response
func (m *mockVehicle) Turn(dir string) string {
	if m.mocked.mockTurn {
		return m.mocked.TurnResponse(dir)
	}
	return m.real.Turn(dir)
}

// setTurnResponse sets the response for Turn
func (m *mockVehicle) setTurnResponse(resp func(string) string) {
	m.mocked.TurnResponse = resp
}

// enableTurnResponse turns the mock on
func (m *mockVehicle) enableTurnResponse() {
	m.mocked.mockTurn = true
}

// disableTurnResponse turns the mock off
func (m *mockVehicle) disableTurnResponse() {
	m.mocked.mockTurn = false
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// Reverse overrides the method to return the mock response
func (m *mockVehicle) Reverse() (string, error) {
	if m.mocked.mockReverse {
		return m.mocked.ReverseResponse()
	}
	return m.real.Reverse()
}

// setReverseResponse sets the response for Reverse
func (m *mockVehicle) setReverseResponse(resp func() (string, error)) {
	m.mocked.ReverseResponse = resp
}

// enableReverseResponse turns the mock on
func (m *mockVehicle) enableReverseResponse() {
	m.mocked.mockReverse = true
}

// disableReverseResponse turns the mock off
func (m *mockVehicle) disableReverseResponse() {
	m.mocked.mockReverse = false
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// Accelerate overrides the method to return the mock response
func (m *mockVehicle) Accelerate(speed int, unit string) (int, error) {
	if m.mocked.mockAccelerate {
		return m.mocked.AccelerateResponse(speed, unit)
	}
	return m.real.Accelerate(speed, unit)
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockVehicle) setAccelerateResponse(resp func(int, string) (int, error)) {
	m.mocked.AccelerateResponse = resp
}

// enableAccelerateResponse turns the mock on
func (m *mockVehicle) enableAccelerateResponse() {
	m.mocked.mockAccelerate = true
}

// disableAccelerateResponse turns the mock off
func (m *mockVehicle) disableAccelerateResponse() {
	m.mocked.mockAccelerate = false
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// IsMoving overrides the method to return the mock response
func (m *mockVehicle) IsMoving() bool {
	if m.mocked.mockIsMoving {
		return m.mocked.IsMovingResponse()
	}
	return m.real.IsMoving()
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockVehicle) setIsMovingResponse(resp func() bool) {
	m.mocked.IsMovingResponse = resp
}

// enableIsMovingResponse turns the mock on
func (m *mockVehicle) enableIsMovingResponse() {
	m.mocked.mockIsMoving = true
}

// disableIsMovingResponse turns the mock off
func (m *mockVehicle) disableIsMovingResponse() {
	m.mocked.mockIsMoving = false
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// Honk overrides the method to return the mock response
func (m *mockVehicle) Honk(times int) {
	if m.mocked.mockHonk {
		m.mocked.HonkResponse(times)
	}
	m.real.Honk(times)
}

// setHonkResponse sets the response for Honk
func (m *mockVehicle) setHonkResponse(resp func(int)) {
	m.mocked.HonkResponse = resp
}

// enableHonkResponse turns the mock on
func (m *mockVehicle) enableHonkResponse() {
	m.mocked.mockHonk = true
}

// disableHonkResponse turns the mock off
func (m *mockVehicle) disableHonkResponse() {
	m.mocked.mockHonk = false
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// GetEngineSpecs overrides the method to return the mock response
func (m *mockVehicle) GetEngineSpecs() (int, string) {
	if m.mocked.mockGetEngineSpecs {
		return m.mocked.GetEngineSpecsResponse()
	}
	return m.real.GetEngineSpecs()
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsResponse(resp func() (int, string)) {
	m.mocked.GetEngineSpecsResponse = resp
}

// enableGetEngineSpecsResponse turns the mock on
func (m *mockVehicle) enableGetEngineSpecsResponse() {
	m.mocked.mockGetEngineSpecs = true
}

// disableGetEngineSpecsResponse turns the mock off
func (m *mockVehicle) disableGetEngineSpecsResponse() {
	m.mocked.mockGetEngineSpecs = false
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// ApplyBrakes overrides the method to return the mock response
func (m *mockVehicle) ApplyBrakes(force float64) bool {
	if m.mocked.mockApplyBrakes {
		return m.mocked.ApplyBrakesResponse(force)
	}
	return m.real.ApplyBrakes(force)
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockVehicle) setApplyBrakesResponse(resp func(float64) bool) {
	m.mocked.ApplyBrakesResponse = resp
}

// enableApplyBrakesResponse turns the mock on
func (m *mockVehicle) enableApplyBrakesResponse() {
	m.mocked.mockApplyBrakes = true
}

// disableApplyBrakesResponse turns the mock off
func (m *mockVehicle) disableApplyBrakesResponse() {
	m.mocked.mockApplyBrakes = false
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// ChangeGears overrides the method to return the mock response
func (m *mockVehicle) ChangeGears(gear int) (int, int) {
	if m.mocked.mockChangeGears {
		return m.mocked.ChangeGearsResponse(gear)
	}
	return m.real.ChangeGears(gear)
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockVehicle) setChangeGearsResponse(resp func(int) (int, int)) {
	m.mocked.ChangeGearsResponse = resp
}

// enableChangeGearsResponse turns the mock on
func (m *mockVehicle) enableChangeGearsResponse() {
	m.mocked.mockChangeGears = true
}

// disableChangeGearsResponse turns the mock off
func (m *mockVehicle) disableChangeGearsResponse() {
	m.mocked.mockChangeGears = false
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// Telemetry overrides the method to return the mock response
func (m *mockVehicle) Telemetry() map[string]float64 {
	if m.mocked.mockTelemetry {
		return m.mocked.TelemetryResponse()
	}
	return m.real.Telemetry()
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockVehicle) setTelemetryResponse(resp func() map[string]float64) {
	m.mocked.TelemetryResponse = resp
}

// enableTelemetryResponse turns the mock on
func (m *mockVehicle) enableTelemetryResponse() {
	m.mocked.mockTelemetry = true
}

// disableTelemetryResponse turns the mock off
func (m *mockVehicle) disableTelemetryResponse() {
	m.mocked.mockTelemetry = false
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// GetPassengers overrides the method to return the mock response
func (m *mockVehicle) GetPassengers() []string {
	if m.mocked.mockGetPassengers {
		return m.mocked.GetPassengersResponse()
	}
	return m.real.GetPassengers()
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockVehicle) setGetPassengersResponse(resp func() []string) {
	m.mocked.GetPassengersResponse = resp
}

// enableGetPassengersResponse turns the mock on
func (m *mockVehicle) enableGetPassengersResponse() {
	m.mocked.mockGetPassengers = true
}

// disableGetPassengersResponse turns the mock off
func (m *mockVehicle) disableGetPassengersResponse() {
	m.mocked.mockGetPassengers = false
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// LoadCargo overrides the method to return the mock response
func (m *mockVehicle) LoadCargo(items []string) (int, error) {
	if m.mocked.mockLoadCargo {
		return m.mocked.LoadCargoResponse(items)
	}
	return m.real.LoadCargo(items)
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockVehicle) setLoadCargoResponse(resp func([]string) (int, error)) {
	m.mocked.LoadCargoResponse = resp
}

// enableLoadCargoResponse turns the mock on
func (m *mockVehicle) enableLoadCargoResponse() {
	m.mocked.mockLoadCargo = true
}

// disableLoadCargoResponse turns the mock off
func (m *mockVehicle) disableLoadCargoResponse() {
	m.mocked.mockLoadCargo = false
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// GetVehicleStatus overrides the method to return the mock response
func (m *mockVehicle) GetVehicleStatus() vehicle.VehicleStatus {
	if m.mocked.mockGetVehicleStatus {
		return m.mocked.GetVehicleStatusResponse()
	}
	return m.real.GetVehicleStatus()
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockVehicle) setGetVehicleStatusResponse(resp func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatusResponse = resp
}

// enableGetVehicleStatusResponse turns the mock on
func (m *mockVehicle) enableGetVehicleStatusResponse() {
	m.mocked.mockGetVehicleStatus = true
}

// disableGetVehicleStatusResponse turns the mock off
func (m *mockVehicle) disableGetVehicleStatusResponse() {
	m.mocked.mockGetVehicleStatus = false
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// UpdateStatus overrides the method to return the mock response
func (m *mockVehicle) UpdateStatus(status vehicle.VehicleStatus) error {
	if m.mocked.mockUpdateStatus {
		return m.mocked.UpdateStatusResponse(status)
	}
	return m.real.UpdateStatus(status)
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockVehicle) setUpdateStatusResponse(resp func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatusResponse = resp
}

// enableUpdateStatusResponse turns the mock on
func (m *mockVehicle) enableUpdateStatusResponse() {
	m.mocked.mockUpdateStatus = true
}

// disableUpdateStatusResponse turns the mock off
func (m *mockVehicle) disableUpdateStatusResponse() {
	m.mocked.mockUpdateStatus = false
}
