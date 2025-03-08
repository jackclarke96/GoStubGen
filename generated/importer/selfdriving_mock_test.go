package driver

import "github.com/jackclarke/GoStubGen/generated/vehicle"

// mockSelfDrivingConfig stores mock flags and responses
type mockSelfDrivingConfig struct {

	// DriveSelf flag and mock response
	mockDriveSelf     bool
	DriveSelfResponse func(string) error

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

// mockSelfDriving embeds a concrete SelfDriving and its mocks
type mockSelfDriving struct {
	real   vehicle.SelfDriving
	mocked mockSelfDrivingConfig
}

// selfdrivingMock returns a new mock
func selfdrivingMock(v vehicle.SelfDriving) *mockSelfDriving {
	return &mockSelfDriving{
		real:   v,
		mocked: mockSelfDrivingConfig{},
	}
}

/* -------------------------- DriveSelf Mock Helpers --------------------------- */

// DriveSelf overrides the method to return the mock response
func (m *mockSelfDriving) DriveSelf(endLocation string) error {
	if m.mocked.mockDriveSelf {
		return m.mocked.DriveSelfResponse(endLocation)
	}
	return m.real.DriveSelf(endLocation)
}

// setDriveSelfFunc sets the function for DriveSelf
func (m *mockSelfDriving) setDriveSelfFunc(f func(string) error) {
	m.mocked.DriveSelfResponse = f
}

// setDriveSelfResponse sets the response for DriveSelf
func (m *mockSelfDriving) setDriveSelfResponse(output0 error) {
	m.setDriveSelfFunc(func(string) error {
		return output0
	})
}

// enableDriveSelfMock turns the mock on
func (m *mockSelfDriving) enableDriveSelfMock() {
	m.mocked.mockDriveSelf = true
}

// disableDriveSelfMock turns the mock off
func (m *mockSelfDriving) disableDriveSelfMock() {
	m.mocked.mockDriveSelf = false
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// GetTopSpeed overrides the method to return the mock response
func (m *mockSelfDriving) GetTopSpeed() int {
	if m.mocked.mockGetTopSpeed {
		return m.mocked.GetTopSpeedResponse()
	}
	return m.real.GetTopSpeed()
}

// setGetTopSpeedFunc sets the function for GetTopSpeed
func (m *mockSelfDriving) setGetTopSpeedFunc(f func() int) {
	m.mocked.GetTopSpeedResponse = f
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockSelfDriving) setGetTopSpeedResponse(output0 int) {
	m.setGetTopSpeedFunc(func() int {
		return output0
	})
}

// enableGetTopSpeedMock turns the mock on
func (m *mockSelfDriving) enableGetTopSpeedMock() {
	m.mocked.mockGetTopSpeed = true
}

// disableGetTopSpeedMock turns the mock off
func (m *mockSelfDriving) disableGetTopSpeedMock() {
	m.mocked.mockGetTopSpeed = false
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// Turn overrides the method to return the mock response
func (m *mockSelfDriving) Turn(dir string) string {
	if m.mocked.mockTurn {
		return m.mocked.TurnResponse(dir)
	}
	return m.real.Turn(dir)
}

// setTurnFunc sets the function for Turn
func (m *mockSelfDriving) setTurnFunc(f func(string) string) {
	m.mocked.TurnResponse = f
}

// setTurnResponse sets the response for Turn
func (m *mockSelfDriving) setTurnResponse(output0 string) {
	m.setTurnFunc(func(string) string {
		return output0
	})
}

// enableTurnMock turns the mock on
func (m *mockSelfDriving) enableTurnMock() {
	m.mocked.mockTurn = true
}

// disableTurnMock turns the mock off
func (m *mockSelfDriving) disableTurnMock() {
	m.mocked.mockTurn = false
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// Reverse overrides the method to return the mock response
func (m *mockSelfDriving) Reverse() (string, error) {
	if m.mocked.mockReverse {
		return m.mocked.ReverseResponse()
	}
	return m.real.Reverse()
}

// setReverseFunc sets the function for Reverse
func (m *mockSelfDriving) setReverseFunc(f func() (string, error)) {
	m.mocked.ReverseResponse = f
}

// setReverseResponse sets the response for Reverse
func (m *mockSelfDriving) setReverseResponse(output0 string, output1 error) {
	m.setReverseFunc(func() (string, error) {
		return output0, output1
	})
}

// enableReverseMock turns the mock on
func (m *mockSelfDriving) enableReverseMock() {
	m.mocked.mockReverse = true
}

// disableReverseMock turns the mock off
func (m *mockSelfDriving) disableReverseMock() {
	m.mocked.mockReverse = false
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// Accelerate overrides the method to return the mock response
func (m *mockSelfDriving) Accelerate(speed int, unit string) (int, error) {
	if m.mocked.mockAccelerate {
		return m.mocked.AccelerateResponse(speed, unit)
	}
	return m.real.Accelerate(speed, unit)
}

// setAccelerateFunc sets the function for Accelerate
func (m *mockSelfDriving) setAccelerateFunc(f func(int, string) (int, error)) {
	m.mocked.AccelerateResponse = f
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockSelfDriving) setAccelerateResponse(output0 int, output1 error) {
	m.setAccelerateFunc(func(int, string) (int, error) {
		return output0, output1
	})
}

// enableAccelerateMock turns the mock on
func (m *mockSelfDriving) enableAccelerateMock() {
	m.mocked.mockAccelerate = true
}

// disableAccelerateMock turns the mock off
func (m *mockSelfDriving) disableAccelerateMock() {
	m.mocked.mockAccelerate = false
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// IsMoving overrides the method to return the mock response
func (m *mockSelfDriving) IsMoving() bool {
	if m.mocked.mockIsMoving {
		return m.mocked.IsMovingResponse()
	}
	return m.real.IsMoving()
}

// setIsMovingFunc sets the function for IsMoving
func (m *mockSelfDriving) setIsMovingFunc(f func() bool) {
	m.mocked.IsMovingResponse = f
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockSelfDriving) setIsMovingResponse(output0 bool) {
	m.setIsMovingFunc(func() bool {
		return output0
	})
}

// enableIsMovingMock turns the mock on
func (m *mockSelfDriving) enableIsMovingMock() {
	m.mocked.mockIsMoving = true
}

// disableIsMovingMock turns the mock off
func (m *mockSelfDriving) disableIsMovingMock() {
	m.mocked.mockIsMoving = false
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// Honk overrides the method to return the mock response
func (m *mockSelfDriving) Honk(times int) {
	if m.mocked.mockHonk {
		m.mocked.HonkResponse(times)
	}
	m.real.Honk(times)
}

// setHonkFunc sets the function for Honk
func (m *mockSelfDriving) setHonkFunc(f func(int)) {
	m.mocked.HonkResponse = f
}

// enableHonkMock turns the mock on
func (m *mockSelfDriving) enableHonkMock() {
	m.mocked.mockHonk = true
}

// disableHonkMock turns the mock off
func (m *mockSelfDriving) disableHonkMock() {
	m.mocked.mockHonk = false
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// GetEngineSpecs overrides the method to return the mock response
func (m *mockSelfDriving) GetEngineSpecs() (int, string) {
	if m.mocked.mockGetEngineSpecs {
		return m.mocked.GetEngineSpecsResponse()
	}
	return m.real.GetEngineSpecs()
}

// setGetEngineSpecsFunc sets the function for GetEngineSpecs
func (m *mockSelfDriving) setGetEngineSpecsFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecsResponse = f
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockSelfDriving) setGetEngineSpecsResponse(output0 int, output1 string) {
	m.setGetEngineSpecsFunc(func() (int, string) {
		return output0, output1
	})
}

// enableGetEngineSpecsMock turns the mock on
func (m *mockSelfDriving) enableGetEngineSpecsMock() {
	m.mocked.mockGetEngineSpecs = true
}

// disableGetEngineSpecsMock turns the mock off
func (m *mockSelfDriving) disableGetEngineSpecsMock() {
	m.mocked.mockGetEngineSpecs = false
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// ApplyBrakes overrides the method to return the mock response
func (m *mockSelfDriving) ApplyBrakes(force float64) bool {
	if m.mocked.mockApplyBrakes {
		return m.mocked.ApplyBrakesResponse(force)
	}
	return m.real.ApplyBrakes(force)
}

// setApplyBrakesFunc sets the function for ApplyBrakes
func (m *mockSelfDriving) setApplyBrakesFunc(f func(float64) bool) {
	m.mocked.ApplyBrakesResponse = f
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockSelfDriving) setApplyBrakesResponse(output0 bool) {
	m.setApplyBrakesFunc(func(float64) bool {
		return output0
	})
}

// enableApplyBrakesMock turns the mock on
func (m *mockSelfDriving) enableApplyBrakesMock() {
	m.mocked.mockApplyBrakes = true
}

// disableApplyBrakesMock turns the mock off
func (m *mockSelfDriving) disableApplyBrakesMock() {
	m.mocked.mockApplyBrakes = false
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// ChangeGears overrides the method to return the mock response
func (m *mockSelfDriving) ChangeGears(gear int) (int, int) {
	if m.mocked.mockChangeGears {
		return m.mocked.ChangeGearsResponse(gear)
	}
	return m.real.ChangeGears(gear)
}

// setChangeGearsFunc sets the function for ChangeGears
func (m *mockSelfDriving) setChangeGearsFunc(f func(int) (int, int)) {
	m.mocked.ChangeGearsResponse = f
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockSelfDriving) setChangeGearsResponse(output0 int, output1 int) {
	m.setChangeGearsFunc(func(int) (int, int) {
		return output0, output1
	})
}

// enableChangeGearsMock turns the mock on
func (m *mockSelfDriving) enableChangeGearsMock() {
	m.mocked.mockChangeGears = true
}

// disableChangeGearsMock turns the mock off
func (m *mockSelfDriving) disableChangeGearsMock() {
	m.mocked.mockChangeGears = false
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// Telemetry overrides the method to return the mock response
func (m *mockSelfDriving) Telemetry() map[string]float64 {
	if m.mocked.mockTelemetry {
		return m.mocked.TelemetryResponse()
	}
	return m.real.Telemetry()
}

// setTelemetryFunc sets the function for Telemetry
func (m *mockSelfDriving) setTelemetryFunc(f func() map[string]float64) {
	m.mocked.TelemetryResponse = f
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockSelfDriving) setTelemetryResponse(output0 map[string]float64) {
	m.setTelemetryFunc(func() map[string]float64 {
		return output0
	})
}

// enableTelemetryMock turns the mock on
func (m *mockSelfDriving) enableTelemetryMock() {
	m.mocked.mockTelemetry = true
}

// disableTelemetryMock turns the mock off
func (m *mockSelfDriving) disableTelemetryMock() {
	m.mocked.mockTelemetry = false
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// GetPassengers overrides the method to return the mock response
func (m *mockSelfDriving) GetPassengers() []string {
	if m.mocked.mockGetPassengers {
		return m.mocked.GetPassengersResponse()
	}
	return m.real.GetPassengers()
}

// setGetPassengersFunc sets the function for GetPassengers
func (m *mockSelfDriving) setGetPassengersFunc(f func() []string) {
	m.mocked.GetPassengersResponse = f
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockSelfDriving) setGetPassengersResponse(output0 []string) {
	m.setGetPassengersFunc(func() []string {
		return output0
	})
}

// enableGetPassengersMock turns the mock on
func (m *mockSelfDriving) enableGetPassengersMock() {
	m.mocked.mockGetPassengers = true
}

// disableGetPassengersMock turns the mock off
func (m *mockSelfDriving) disableGetPassengersMock() {
	m.mocked.mockGetPassengers = false
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// LoadCargo overrides the method to return the mock response
func (m *mockSelfDriving) LoadCargo(items []string) (int, error) {
	if m.mocked.mockLoadCargo {
		return m.mocked.LoadCargoResponse(items)
	}
	return m.real.LoadCargo(items)
}

// setLoadCargoFunc sets the function for LoadCargo
func (m *mockSelfDriving) setLoadCargoFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargoResponse = f
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockSelfDriving) setLoadCargoResponse(output0 int, output1 error) {
	m.setLoadCargoFunc(func([]string) (int, error) {
		return output0, output1
	})
}

// enableLoadCargoMock turns the mock on
func (m *mockSelfDriving) enableLoadCargoMock() {
	m.mocked.mockLoadCargo = true
}

// disableLoadCargoMock turns the mock off
func (m *mockSelfDriving) disableLoadCargoMock() {
	m.mocked.mockLoadCargo = false
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// GetVehicleStatus overrides the method to return the mock response
func (m *mockSelfDriving) GetVehicleStatus() vehicle.VehicleStatus {
	if m.mocked.mockGetVehicleStatus {
		return m.mocked.GetVehicleStatusResponse()
	}
	return m.real.GetVehicleStatus()
}

// setGetVehicleStatusFunc sets the function for GetVehicleStatus
func (m *mockSelfDriving) setGetVehicleStatusFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatusResponse = f
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockSelfDriving) setGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.setGetVehicleStatusFunc(func() vehicle.VehicleStatus {
		return output0
	})
}

// enableGetVehicleStatusMock turns the mock on
func (m *mockSelfDriving) enableGetVehicleStatusMock() {
	m.mocked.mockGetVehicleStatus = true
}

// disableGetVehicleStatusMock turns the mock off
func (m *mockSelfDriving) disableGetVehicleStatusMock() {
	m.mocked.mockGetVehicleStatus = false
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// UpdateStatus overrides the method to return the mock response
func (m *mockSelfDriving) UpdateStatus(status vehicle.VehicleStatus) error {
	if m.mocked.mockUpdateStatus {
		return m.mocked.UpdateStatusResponse(status)
	}
	return m.real.UpdateStatus(status)
}

// setUpdateStatusFunc sets the function for UpdateStatus
func (m *mockSelfDriving) setUpdateStatusFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatusResponse = f
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockSelfDriving) setUpdateStatusResponse(output0 error) {
	m.setUpdateStatusFunc(func(vehicle.VehicleStatus) error {
		return output0
	})
}

// enableUpdateStatusMock turns the mock on
func (m *mockSelfDriving) enableUpdateStatusMock() {
	m.mocked.mockUpdateStatus = true
}

// disableUpdateStatusMock turns the mock off
func (m *mockSelfDriving) disableUpdateStatusMock() {
	m.mocked.mockUpdateStatus = false
}
