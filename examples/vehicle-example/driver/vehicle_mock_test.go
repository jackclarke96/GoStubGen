package driver

import "github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"

// mockVehicleConfig stores mock flags and responses
type mockVehicleConfig struct {
	IsMoving         methodConfig[func() bool]
	Honk             methodConfig[func(int)]
	GetEngineSpecs   methodConfig[func() (int, string)]
	GetPassengers    methodConfig[func() []string]
	GetVehicleStatus methodConfig[func() vehicle.VehicleStatus]
	Reverse          methodConfig[func() (string, error)]
	Accelerate       methodConfig[func(int, string) (int, error)]
	ApplyBrakes      methodConfig[func(float64) bool]
	ChangeGears      methodConfig[func(int) (int, int)]
	Telemetry        methodConfig[func() map[string]float64]
	LoadCargo        methodConfig[func([]string) (int, error)]
	UpdateStatus     methodConfig[func(vehicle.VehicleStatus) error]
	GetTopSpeed      methodConfig[func() int]
	Turn             methodConfig[func(string) string]
}

// mockVehicle embeds a concrete Vehicle and its mocks
type mockVehicle struct {
	real   vehicle.Vehicle
	mocked mockVehicleConfig
}

// newVehicleMock returns a new mock
func newVehicleMock(v vehicle.Vehicle) *mockVehicle {
	return &mockVehicle{
		real:   v,
		mocked: mockVehicleConfig{},
	}
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// IsMoving overrides the method to return the mock response
func (m *mockVehicle) IsMoving() bool {
	if m.mocked.IsMoving.enabled {
		return m.mocked.IsMoving.fallback.(bool)
	}
	return m.real.IsMoving()
}

// setIsMovingFunc sets the function for IsMoving
func (m *mockVehicle) setIsMovingFunc(f func() bool) {
	m.mocked.IsMoving.fallback = f
}

// enableIsMovingMock turns the mock on
func (m *mockVehicle) enableIsMovingMock() {
	m.mocked.IsMoving.enabled = true
}

// disableIsMovingMock turns the mock off
func (m *mockVehicle) disableIsMovingMock() {
	m.mocked.IsMoving.enabled = false
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockVehicle) setIsMovingResponse(output0 bool) {
	m.setIsMovingFunc(func() bool {
		return output0
	})
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// Honk overrides the method to return the mock response
func (m *mockVehicle) Honk(times int) {
	if m.mocked.Honk.enabled {
		m.mocked.Honk.fallback(times)
	}
	m.real.Honk(times)
}

// setHonkFunc sets the function for Honk
func (m *mockVehicle) setHonkFunc(f func(int)) {
	m.mocked.Honk.fallback = f
}

// enableHonkMock turns the mock on
func (m *mockVehicle) enableHonkMock() {
	m.mocked.Honk.enabled = true
}

// disableHonkMock turns the mock off
func (m *mockVehicle) disableHonkMock() {
	m.mocked.Honk.enabled = false
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// GetEngineSpecs overrides the method to return the mock response
func (m *mockVehicle) GetEngineSpecs() (int, string) {
	if m.mocked.GetEngineSpecs.enabled {
		return m.mocked.GetEngineSpecs.fallback()
	}
	return m.real.GetEngineSpecs()
}

// setGetEngineSpecsFunc sets the function for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.fallback = f
}

// enableGetEngineSpecsMock turns the mock on
func (m *mockVehicle) enableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.enabled = true
}

// disableGetEngineSpecsMock turns the mock off
func (m *mockVehicle) disableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.enabled = false
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsResponse(output0 int, output1 string) {
	m.setGetEngineSpecsFunc(func() (int, string) {
		return output0, output1
	})
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// GetPassengers overrides the method to return the mock response
func (m *mockVehicle) GetPassengers() []string {
	if m.mocked.GetPassengers.enabled {
		return m.mocked.GetPassengers.fallback()
	}
	return m.real.GetPassengers()
}

// setGetPassengersFunc sets the function for GetPassengers
func (m *mockVehicle) setGetPassengersFunc(f func() []string) {
	m.mocked.GetPassengers.fallback = f
}

// enableGetPassengersMock turns the mock on
func (m *mockVehicle) enableGetPassengersMock() {
	m.mocked.GetPassengers.enabled = true
}

// disableGetPassengersMock turns the mock off
func (m *mockVehicle) disableGetPassengersMock() {
	m.mocked.GetPassengers.enabled = false
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockVehicle) setGetPassengersResponse(output0 []string) {
	m.setGetPassengersFunc(func() []string {
		return output0
	})
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// GetVehicleStatus overrides the method to return the mock response
func (m *mockVehicle) GetVehicleStatus() vehicle.VehicleStatus {
	if m.mocked.GetVehicleStatus.enabled {
		return m.mocked.GetVehicleStatus.fallback()
	}
	return m.real.GetVehicleStatus()
}

// setGetVehicleStatusFunc sets the function for GetVehicleStatus
func (m *mockVehicle) setGetVehicleStatusFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.fallback = f
}

// enableGetVehicleStatusMock turns the mock on
func (m *mockVehicle) enableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.enabled = true
}

// disableGetVehicleStatusMock turns the mock off
func (m *mockVehicle) disableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.enabled = false
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockVehicle) setGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.setGetVehicleStatusFunc(func() vehicle.VehicleStatus {
		return output0
	})
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// Reverse overrides the method to return the mock response
func (m *mockVehicle) Reverse() (string, error) {
	if m.mocked.Reverse.enabled {
		return m.mocked.Reverse.fallback()
	}
	return m.real.Reverse()
}

// setReverseFunc sets the function for Reverse
func (m *mockVehicle) setReverseFunc(f func() (string, error)) {
	m.mocked.Reverse.fallback = f
}

// enableReverseMock turns the mock on
func (m *mockVehicle) enableReverseMock() {
	m.mocked.Reverse.enabled = true
}

// disableReverseMock turns the mock off
func (m *mockVehicle) disableReverseMock() {
	m.mocked.Reverse.enabled = false
}

// setReverseResponse sets the response for Reverse
func (m *mockVehicle) setReverseResponse(output0 string, output1 error) {
	m.setReverseFunc(func() (string, error) {
		return output0, output1
	})
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// Accelerate overrides the method to return the mock response
func (m *mockVehicle) Accelerate(speed int, unit string) (int, error) {
	if m.mocked.Accelerate.enabled {
		return m.mocked.Accelerate.fallback(speed, unit)
	}
	return m.real.Accelerate(speed, unit)
}

// setAccelerateFunc sets the function for Accelerate
func (m *mockVehicle) setAccelerateFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.fallback = f
}

// enableAccelerateMock turns the mock on
func (m *mockVehicle) enableAccelerateMock() {
	m.mocked.Accelerate.enabled = true
}

// disableAccelerateMock turns the mock off
func (m *mockVehicle) disableAccelerateMock() {
	m.mocked.Accelerate.enabled = false
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockVehicle) setAccelerateResponse(output0 int, output1 error) {
	m.setAccelerateFunc(func(int, string) (int, error) {
		return output0, output1
	})
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// ApplyBrakes overrides the method to return the mock response
func (m *mockVehicle) ApplyBrakes(force float64) bool {
	if m.mocked.ApplyBrakes.enabled {
		return m.mocked.ApplyBrakes.fallback(force)
	}
	return m.real.ApplyBrakes(force)
}

// setApplyBrakesFunc sets the function for ApplyBrakes
func (m *mockVehicle) setApplyBrakesFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.fallback = f
}

// enableApplyBrakesMock turns the mock on
func (m *mockVehicle) enableApplyBrakesMock() {
	m.mocked.ApplyBrakes.enabled = true
}

// disableApplyBrakesMock turns the mock off
func (m *mockVehicle) disableApplyBrakesMock() {
	m.mocked.ApplyBrakes.enabled = false
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockVehicle) setApplyBrakesResponse(output0 bool) {
	m.setApplyBrakesFunc(func(float64) bool {
		return output0
	})
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// ChangeGears overrides the method to return the mock response
func (m *mockVehicle) ChangeGears(gear int) (int, int) {
	if m.mocked.ChangeGears.enabled {
		return m.mocked.ChangeGears.fallback(gear)
	}
	return m.real.ChangeGears(gear)
}

// setChangeGearsFunc sets the function for ChangeGears
func (m *mockVehicle) setChangeGearsFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.fallback = f
}

// enableChangeGearsMock turns the mock on
func (m *mockVehicle) enableChangeGearsMock() {
	m.mocked.ChangeGears.enabled = true
}

// disableChangeGearsMock turns the mock off
func (m *mockVehicle) disableChangeGearsMock() {
	m.mocked.ChangeGears.enabled = false
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockVehicle) setChangeGearsResponse(output0 int, output1 int) {
	m.setChangeGearsFunc(func(int) (int, int) {
		return output0, output1
	})
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// Telemetry overrides the method to return the mock response
func (m *mockVehicle) Telemetry() map[string]float64 {
	if m.mocked.Telemetry.enabled {
		return m.mocked.Telemetry.fallback()
	}
	return m.real.Telemetry()
}

// setTelemetryFunc sets the function for Telemetry
func (m *mockVehicle) setTelemetryFunc(f func() map[string]float64) {
	m.mocked.Telemetry.fallback = f
}

// enableTelemetryMock turns the mock on
func (m *mockVehicle) enableTelemetryMock() {
	m.mocked.Telemetry.enabled = true
}

// disableTelemetryMock turns the mock off
func (m *mockVehicle) disableTelemetryMock() {
	m.mocked.Telemetry.enabled = false
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockVehicle) setTelemetryResponse(output0 map[string]float64) {
	m.setTelemetryFunc(func() map[string]float64 {
		return output0
	})
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// LoadCargo overrides the method to return the mock response
func (m *mockVehicle) LoadCargo(items []string) (int, error) {
	if m.mocked.LoadCargo.enabled {
		return m.mocked.LoadCargo.fallback(items)
	}
	return m.real.LoadCargo(items)
}

// setLoadCargoFunc sets the function for LoadCargo
func (m *mockVehicle) setLoadCargoFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.fallback = f
}

// enableLoadCargoMock turns the mock on
func (m *mockVehicle) enableLoadCargoMock() {
	m.mocked.LoadCargo.enabled = true
}

// disableLoadCargoMock turns the mock off
func (m *mockVehicle) disableLoadCargoMock() {
	m.mocked.LoadCargo.enabled = false
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockVehicle) setLoadCargoResponse(output0 int, output1 error) {
	m.setLoadCargoFunc(func([]string) (int, error) {
		return output0, output1
	})
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// UpdateStatus overrides the method to return the mock response
func (m *mockVehicle) UpdateStatus(status vehicle.VehicleStatus) error {
	if m.mocked.UpdateStatus.enabled {
		return m.mocked.UpdateStatus.fallback(status)
	}
	return m.real.UpdateStatus(status)
}

// setUpdateStatusFunc sets the function for UpdateStatus
func (m *mockVehicle) setUpdateStatusFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.fallback = f
}

// enableUpdateStatusMock turns the mock on
func (m *mockVehicle) enableUpdateStatusMock() {
	m.mocked.UpdateStatus.enabled = true
}

// disableUpdateStatusMock turns the mock off
func (m *mockVehicle) disableUpdateStatusMock() {
	m.mocked.UpdateStatus.enabled = false
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockVehicle) setUpdateStatusResponse(output0 error) {
	m.setUpdateStatusFunc(func(vehicle.VehicleStatus) error {
		return output0
	})
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// GetTopSpeed overrides the method to return the mock response
func (m *mockVehicle) GetTopSpeed() int {
	if m.mocked.GetTopSpeed.enabled {
		return m.mocked.GetTopSpeed.fallback()
	}
	return m.real.GetTopSpeed()
}

// setGetTopSpeedFunc sets the function for GetTopSpeed
func (m *mockVehicle) setGetTopSpeedFunc(f func() int) {
	m.mocked.GetTopSpeed.fallback = f
}

// enableGetTopSpeedMock turns the mock on
func (m *mockVehicle) enableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.enabled = true
}

// disableGetTopSpeedMock turns the mock off
func (m *mockVehicle) disableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.enabled = false
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockVehicle) setGetTopSpeedResponse(output0 int) {
	m.setGetTopSpeedFunc(func() int {
		return output0
	})
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// Turn overrides the method to return the mock response
func (m *mockVehicle) Turn(dir string) string {
	if m.mocked.Turn.enabled {
		return m.mocked.Turn.fallback(dir)
	}
	return m.real.Turn(dir)
}

// setTurnFunc sets the function for Turn
func (m *mockVehicle) setTurnFunc(f func(string) string) {
	m.mocked.Turn.fallback = f
}

// enableTurnMock turns the mock on
func (m *mockVehicle) enableTurnMock() {
	m.mocked.Turn.enabled = true
}

// disableTurnMock turns the mock off
func (m *mockVehicle) disableTurnMock() {
	m.mocked.Turn.enabled = false
}

// setTurnResponse sets the response for Turn
func (m *mockVehicle) setTurnResponse(output0 string) {
	m.setTurnFunc(func(string) string {
		return output0
	})
}
