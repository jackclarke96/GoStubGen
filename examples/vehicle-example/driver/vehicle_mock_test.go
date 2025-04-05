package driver

import "github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"

// mockVehicleConfig stores mock flags and responses
type mockVehicleConfig struct {
	GetTopSpeed      methodConfig[func() int]
	IsMoving         methodConfig[func() bool]
	GetEngineSpecs   methodConfig[func() (int, string)]
	ApplyBrakes      methodConfig[func(float64) bool]
	Telemetry        methodConfig[func() map[string]float64]
	GetPassengers    methodConfig[func() []string]
	Accelerate       methodConfig[func(int, string) (int, error)]
	Honk             methodConfig[func(int)]
	LoadCargo        methodConfig[func([]string) (int, error)]
	GetVehicleStatus methodConfig[func() vehicle.VehicleStatus]
	Turn             methodConfig[func(string) string]
	Reverse          methodConfig[func() (string, error)]
	ChangeGears      methodConfig[func(int) (int, int)]
	UpdateStatus     methodConfig[func(vehicle.VehicleStatus) error]
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
	if m.mocked.GetTopSpeed.enabled {
		return m.mocked.GetTopSpeed.response()
	}
	return m.real.GetTopSpeed()
}

// setGetTopSpeedFunc sets the function for GetTopSpeed
func (m *mockVehicle) setGetTopSpeedFunc(f func() int) {
	m.mocked.GetTopSpeed.response = f
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockVehicle) setGetTopSpeedResponse(output0 int) {
	m.setGetTopSpeedFunc(func() int {
		return output0
	})
}

// enableGetTopSpeedMock turns the mock on
func (m *mockVehicle) enableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.enabled = true
}

// disableGetTopSpeedMock turns the mock off
func (m *mockVehicle) disableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.enabled = false
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// IsMoving overrides the method to return the mock response
func (m *mockVehicle) IsMoving() bool {
	if m.mocked.IsMoving.enabled {
		return m.mocked.IsMoving.response()
	}
	return m.real.IsMoving()
}

// setIsMovingFunc sets the function for IsMoving
func (m *mockVehicle) setIsMovingFunc(f func() bool) {
	m.mocked.IsMoving.response = f
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockVehicle) setIsMovingResponse(output0 bool) {
	m.setIsMovingFunc(func() bool {
		return output0
	})
}

// enableIsMovingMock turns the mock on
func (m *mockVehicle) enableIsMovingMock() {
	m.mocked.IsMoving.enabled = true
}

// disableIsMovingMock turns the mock off
func (m *mockVehicle) disableIsMovingMock() {
	m.mocked.IsMoving.enabled = false
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// GetEngineSpecs overrides the method to return the mock response
func (m *mockVehicle) GetEngineSpecs() (int, string) {
	if m.mocked.GetEngineSpecs.enabled {
		return m.mocked.GetEngineSpecs.response()
	}
	return m.real.GetEngineSpecs()
}

// setGetEngineSpecsFunc sets the function for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.response = f
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockVehicle) setGetEngineSpecsResponse(output0 int, output1 string) {
	m.setGetEngineSpecsFunc(func() (int, string) {
		return output0, output1
	})
}

// enableGetEngineSpecsMock turns the mock on
func (m *mockVehicle) enableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.enabled = true
}

// disableGetEngineSpecsMock turns the mock off
func (m *mockVehicle) disableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.enabled = false
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// ApplyBrakes overrides the method to return the mock response
func (m *mockVehicle) ApplyBrakes(force float64) bool {
	if m.mocked.ApplyBrakes.enabled {
		return m.mocked.ApplyBrakes.response(force)
	}
	return m.real.ApplyBrakes(force)
}

// setApplyBrakesFunc sets the function for ApplyBrakes
func (m *mockVehicle) setApplyBrakesFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.response = f
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockVehicle) setApplyBrakesResponse(output0 bool) {
	m.setApplyBrakesFunc(func(float64) bool {
		return output0
	})
}

// enableApplyBrakesMock turns the mock on
func (m *mockVehicle) enableApplyBrakesMock() {
	m.mocked.ApplyBrakes.enabled = true
}

// disableApplyBrakesMock turns the mock off
func (m *mockVehicle) disableApplyBrakesMock() {
	m.mocked.ApplyBrakes.enabled = false
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// Telemetry overrides the method to return the mock response
func (m *mockVehicle) Telemetry() map[string]float64 {
	if m.mocked.Telemetry.enabled {
		return m.mocked.Telemetry.response()
	}
	return m.real.Telemetry()
}

// setTelemetryFunc sets the function for Telemetry
func (m *mockVehicle) setTelemetryFunc(f func() map[string]float64) {
	m.mocked.Telemetry.response = f
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockVehicle) setTelemetryResponse(output0 map[string]float64) {
	m.setTelemetryFunc(func() map[string]float64 {
		return output0
	})
}

// enableTelemetryMock turns the mock on
func (m *mockVehicle) enableTelemetryMock() {
	m.mocked.Telemetry.enabled = true
}

// disableTelemetryMock turns the mock off
func (m *mockVehicle) disableTelemetryMock() {
	m.mocked.Telemetry.enabled = false
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// GetPassengers overrides the method to return the mock response
func (m *mockVehicle) GetPassengers() []string {
	if m.mocked.GetPassengers.enabled {
		return m.mocked.GetPassengers.response()
	}
	return m.real.GetPassengers()
}

// setGetPassengersFunc sets the function for GetPassengers
func (m *mockVehicle) setGetPassengersFunc(f func() []string) {
	m.mocked.GetPassengers.response = f
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockVehicle) setGetPassengersResponse(output0 []string) {
	m.setGetPassengersFunc(func() []string {
		return output0
	})
}

// enableGetPassengersMock turns the mock on
func (m *mockVehicle) enableGetPassengersMock() {
	m.mocked.GetPassengers.enabled = true
}

// disableGetPassengersMock turns the mock off
func (m *mockVehicle) disableGetPassengersMock() {
	m.mocked.GetPassengers.enabled = false
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// Accelerate overrides the method to return the mock response
func (m *mockVehicle) Accelerate(speed int, unit string) (int, error) {
	if m.mocked.Accelerate.enabled {
		return m.mocked.Accelerate.response(speed, unit)
	}
	return m.real.Accelerate(speed, unit)
}

// setAccelerateFunc sets the function for Accelerate
func (m *mockVehicle) setAccelerateFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.response = f
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockVehicle) setAccelerateResponse(output0 int, output1 error) {
	m.setAccelerateFunc(func(int, string) (int, error) {
		return output0, output1
	})
}

// enableAccelerateMock turns the mock on
func (m *mockVehicle) enableAccelerateMock() {
	m.mocked.Accelerate.enabled = true
}

// disableAccelerateMock turns the mock off
func (m *mockVehicle) disableAccelerateMock() {
	m.mocked.Accelerate.enabled = false
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// Honk overrides the method to return the mock response
func (m *mockVehicle) Honk(times int) {
	if m.mocked.Honk.enabled {
		m.mocked.Honk.response(times)
	}
	m.real.Honk(times)
}

// setHonkFunc sets the function for Honk
func (m *mockVehicle) setHonkFunc(f func(int)) {
	m.mocked.Honk.response = f
}

// enableHonkMock turns the mock on
func (m *mockVehicle) enableHonkMock() {
	m.mocked.Honk.enabled = true
}

// disableHonkMock turns the mock off
func (m *mockVehicle) disableHonkMock() {
	m.mocked.Honk.enabled = false
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// LoadCargo overrides the method to return the mock response
func (m *mockVehicle) LoadCargo(items []string) (int, error) {
	if m.mocked.LoadCargo.enabled {
		return m.mocked.LoadCargo.response(items)
	}
	return m.real.LoadCargo(items)
}

// setLoadCargoFunc sets the function for LoadCargo
func (m *mockVehicle) setLoadCargoFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.response = f
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockVehicle) setLoadCargoResponse(output0 int, output1 error) {
	m.setLoadCargoFunc(func([]string) (int, error) {
		return output0, output1
	})
}

// enableLoadCargoMock turns the mock on
func (m *mockVehicle) enableLoadCargoMock() {
	m.mocked.LoadCargo.enabled = true
}

// disableLoadCargoMock turns the mock off
func (m *mockVehicle) disableLoadCargoMock() {
	m.mocked.LoadCargo.enabled = false
}

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// GetVehicleStatus overrides the method to return the mock response
func (m *mockVehicle) GetVehicleStatus() vehicle.VehicleStatus {
	if m.mocked.GetVehicleStatus.enabled {
		return m.mocked.GetVehicleStatus.response()
	}
	return m.real.GetVehicleStatus()
}

// setGetVehicleStatusFunc sets the function for GetVehicleStatus
func (m *mockVehicle) setGetVehicleStatusFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.response = f
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockVehicle) setGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.setGetVehicleStatusFunc(func() vehicle.VehicleStatus {
		return output0
	})
}

// enableGetVehicleStatusMock turns the mock on
func (m *mockVehicle) enableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.enabled = true
}

// disableGetVehicleStatusMock turns the mock off
func (m *mockVehicle) disableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.enabled = false
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// Turn overrides the method to return the mock response
func (m *mockVehicle) Turn(dir string) string {
	if m.mocked.Turn.enabled {
		return m.mocked.Turn.response(dir)
	}
	return m.real.Turn(dir)
}

// setTurnFunc sets the function for Turn
func (m *mockVehicle) setTurnFunc(f func(string) string) {
	m.mocked.Turn.response = f
}

// setTurnResponse sets the response for Turn
func (m *mockVehicle) setTurnResponse(output0 string) {
	m.setTurnFunc(func(string) string {
		return output0
	})
}

// enableTurnMock turns the mock on
func (m *mockVehicle) enableTurnMock() {
	m.mocked.Turn.enabled = true
}

// disableTurnMock turns the mock off
func (m *mockVehicle) disableTurnMock() {
	m.mocked.Turn.enabled = false
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// Reverse overrides the method to return the mock response
func (m *mockVehicle) Reverse() (string, error) {
	if m.mocked.Reverse.enabled {
		return m.mocked.Reverse.response()
	}
	return m.real.Reverse()
}

// setReverseFunc sets the function for Reverse
func (m *mockVehicle) setReverseFunc(f func() (string, error)) {
	m.mocked.Reverse.response = f
}

// setReverseResponse sets the response for Reverse
func (m *mockVehicle) setReverseResponse(output0 string, output1 error) {
	m.setReverseFunc(func() (string, error) {
		return output0, output1
	})
}

// enableReverseMock turns the mock on
func (m *mockVehicle) enableReverseMock() {
	m.mocked.Reverse.enabled = true
}

// disableReverseMock turns the mock off
func (m *mockVehicle) disableReverseMock() {
	m.mocked.Reverse.enabled = false
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// ChangeGears overrides the method to return the mock response
func (m *mockVehicle) ChangeGears(gear int) (int, int) {
	if m.mocked.ChangeGears.enabled {
		return m.mocked.ChangeGears.response(gear)
	}
	return m.real.ChangeGears(gear)
}

// setChangeGearsFunc sets the function for ChangeGears
func (m *mockVehicle) setChangeGearsFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.response = f
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockVehicle) setChangeGearsResponse(output0 int, output1 int) {
	m.setChangeGearsFunc(func(int) (int, int) {
		return output0, output1
	})
}

// enableChangeGearsMock turns the mock on
func (m *mockVehicle) enableChangeGearsMock() {
	m.mocked.ChangeGears.enabled = true
}

// disableChangeGearsMock turns the mock off
func (m *mockVehicle) disableChangeGearsMock() {
	m.mocked.ChangeGears.enabled = false
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// UpdateStatus overrides the method to return the mock response
func (m *mockVehicle) UpdateStatus(status vehicle.VehicleStatus) error {
	if m.mocked.UpdateStatus.enabled {
		return m.mocked.UpdateStatus.response(status)
	}
	return m.real.UpdateStatus(status)
}

// setUpdateStatusFunc sets the function for UpdateStatus
func (m *mockVehicle) setUpdateStatusFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.response = f
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockVehicle) setUpdateStatusResponse(output0 error) {
	m.setUpdateStatusFunc(func(vehicle.VehicleStatus) error {
		return output0
	})
}

// enableUpdateStatusMock turns the mock on
func (m *mockVehicle) enableUpdateStatusMock() {
	m.mocked.UpdateStatus.enabled = true
}

// disableUpdateStatusMock turns the mock off
func (m *mockVehicle) disableUpdateStatusMock() {
	m.mocked.UpdateStatus.enabled = false
}
