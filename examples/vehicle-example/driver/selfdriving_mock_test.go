package driver

import "github.com/jackclarke/GoStubGen/examples/vehicle-example/vehicle"

// mockSelfDrivingConfig stores mock flags and responses
type mockSelfDrivingConfig struct {
	GetVehicleStatus methodConfig[func() vehicle.VehicleStatus]
	Reverse          methodConfig[func() (string, error)]
	ApplyBrakes      methodConfig[func(float64) bool]
	LoadCargo        methodConfig[func([]string) (int, error)]
	UpdateStatus     methodConfig[func(vehicle.VehicleStatus) error]
	Honk             methodConfig[func(int)]
	GetPassengers    methodConfig[func() []string]
	Accelerate       methodConfig[func(int, string) (int, error)]
	Telemetry        methodConfig[func() map[string]float64]
	DriveSelf        methodConfig[func(string) error]
	IsMoving         methodConfig[func() bool]
	GetTopSpeed      methodConfig[func() int]
	ChangeGears      methodConfig[func(int) (int, int)]
	Turn             methodConfig[func(string) string]
	GetEngineSpecs   methodConfig[func() (int, string)]
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

/* -------------------------- GetVehicleStatus Mock Helpers --------------------------- */

// GetVehicleStatus overrides the method to return the mock response
func (m *mockSelfDriving) GetVehicleStatus() vehicle.VehicleStatus {
	if m.mocked.GetVehicleStatus.enabled {
		return m.mocked.GetVehicleStatus.response()
	}
	return m.real.GetVehicleStatus()
}

// setGetVehicleStatusFunc sets the function for GetVehicleStatus
func (m *mockSelfDriving) setGetVehicleStatusFunc(f func() vehicle.VehicleStatus) {
	m.mocked.GetVehicleStatus.response = f
}

// enableGetVehicleStatusMock turns the mock on
func (m *mockSelfDriving) enableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.enabled = true
}

// disableGetVehicleStatusMock turns the mock off
func (m *mockSelfDriving) disableGetVehicleStatusMock() {
	m.mocked.GetVehicleStatus.enabled = false
}

// setGetVehicleStatusResponse sets the response for GetVehicleStatus
func (m *mockSelfDriving) setGetVehicleStatusResponse(output0 vehicle.VehicleStatus) {
	m.setGetVehicleStatusFunc(func() vehicle.VehicleStatus {
		return output0
	})
}

/* -------------------------- Reverse Mock Helpers --------------------------- */

// Reverse overrides the method to return the mock response
func (m *mockSelfDriving) Reverse() (string, error) {
	if m.mocked.Reverse.enabled {
		return m.mocked.Reverse.response()
	}
	return m.real.Reverse()
}

// setReverseFunc sets the function for Reverse
func (m *mockSelfDriving) setReverseFunc(f func() (string, error)) {
	m.mocked.Reverse.response = f
}

// enableReverseMock turns the mock on
func (m *mockSelfDriving) enableReverseMock() {
	m.mocked.Reverse.enabled = true
}

// disableReverseMock turns the mock off
func (m *mockSelfDriving) disableReverseMock() {
	m.mocked.Reverse.enabled = false
}

// setReverseResponse sets the response for Reverse
func (m *mockSelfDriving) setReverseResponse(output0 string, output1 error) {
	m.setReverseFunc(func() (string, error) {
		return output0, output1
	})
}

/* -------------------------- ApplyBrakes Mock Helpers --------------------------- */

// ApplyBrakes overrides the method to return the mock response
func (m *mockSelfDriving) ApplyBrakes(force float64) bool {
	if m.mocked.ApplyBrakes.enabled {
		return m.mocked.ApplyBrakes.response(force)
	}
	return m.real.ApplyBrakes(force)
}

// setApplyBrakesFunc sets the function for ApplyBrakes
func (m *mockSelfDriving) setApplyBrakesFunc(f func(float64) bool) {
	m.mocked.ApplyBrakes.response = f
}

// enableApplyBrakesMock turns the mock on
func (m *mockSelfDriving) enableApplyBrakesMock() {
	m.mocked.ApplyBrakes.enabled = true
}

// disableApplyBrakesMock turns the mock off
func (m *mockSelfDriving) disableApplyBrakesMock() {
	m.mocked.ApplyBrakes.enabled = false
}

// setApplyBrakesResponse sets the response for ApplyBrakes
func (m *mockSelfDriving) setApplyBrakesResponse(output0 bool) {
	m.setApplyBrakesFunc(func(float64) bool {
		return output0
	})
}

/* -------------------------- LoadCargo Mock Helpers --------------------------- */

// LoadCargo overrides the method to return the mock response
func (m *mockSelfDriving) LoadCargo(items []string) (int, error) {
	if m.mocked.LoadCargo.enabled {
		return m.mocked.LoadCargo.response(items)
	}
	return m.real.LoadCargo(items)
}

// setLoadCargoFunc sets the function for LoadCargo
func (m *mockSelfDriving) setLoadCargoFunc(f func([]string) (int, error)) {
	m.mocked.LoadCargo.response = f
}

// enableLoadCargoMock turns the mock on
func (m *mockSelfDriving) enableLoadCargoMock() {
	m.mocked.LoadCargo.enabled = true
}

// disableLoadCargoMock turns the mock off
func (m *mockSelfDriving) disableLoadCargoMock() {
	m.mocked.LoadCargo.enabled = false
}

// setLoadCargoResponse sets the response for LoadCargo
func (m *mockSelfDriving) setLoadCargoResponse(output0 int, output1 error) {
	m.setLoadCargoFunc(func([]string) (int, error) {
		return output0, output1
	})
}

/* -------------------------- UpdateStatus Mock Helpers --------------------------- */

// UpdateStatus overrides the method to return the mock response
func (m *mockSelfDriving) UpdateStatus(status vehicle.VehicleStatus) error {
	if m.mocked.UpdateStatus.enabled {
		return m.mocked.UpdateStatus.response(status)
	}
	return m.real.UpdateStatus(status)
}

// setUpdateStatusFunc sets the function for UpdateStatus
func (m *mockSelfDriving) setUpdateStatusFunc(f func(vehicle.VehicleStatus) error) {
	m.mocked.UpdateStatus.response = f
}

// enableUpdateStatusMock turns the mock on
func (m *mockSelfDriving) enableUpdateStatusMock() {
	m.mocked.UpdateStatus.enabled = true
}

// disableUpdateStatusMock turns the mock off
func (m *mockSelfDriving) disableUpdateStatusMock() {
	m.mocked.UpdateStatus.enabled = false
}

// setUpdateStatusResponse sets the response for UpdateStatus
func (m *mockSelfDriving) setUpdateStatusResponse(output0 error) {
	m.setUpdateStatusFunc(func(vehicle.VehicleStatus) error {
		return output0
	})
}

/* -------------------------- Honk Mock Helpers --------------------------- */

// Honk overrides the method to return the mock response
func (m *mockSelfDriving) Honk(times int) {
	if m.mocked.Honk.enabled {
		m.mocked.Honk.response(times)
	}
	m.real.Honk(times)
}

// setHonkFunc sets the function for Honk
func (m *mockSelfDriving) setHonkFunc(f func(int)) {
	m.mocked.Honk.response = f
}

// enableHonkMock turns the mock on
func (m *mockSelfDriving) enableHonkMock() {
	m.mocked.Honk.enabled = true
}

// disableHonkMock turns the mock off
func (m *mockSelfDriving) disableHonkMock() {
	m.mocked.Honk.enabled = false
}

/* -------------------------- GetPassengers Mock Helpers --------------------------- */

// GetPassengers overrides the method to return the mock response
func (m *mockSelfDriving) GetPassengers() []string {
	if m.mocked.GetPassengers.enabled {
		return m.mocked.GetPassengers.response()
	}
	return m.real.GetPassengers()
}

// setGetPassengersFunc sets the function for GetPassengers
func (m *mockSelfDriving) setGetPassengersFunc(f func() []string) {
	m.mocked.GetPassengers.response = f
}

// enableGetPassengersMock turns the mock on
func (m *mockSelfDriving) enableGetPassengersMock() {
	m.mocked.GetPassengers.enabled = true
}

// disableGetPassengersMock turns the mock off
func (m *mockSelfDriving) disableGetPassengersMock() {
	m.mocked.GetPassengers.enabled = false
}

// setGetPassengersResponse sets the response for GetPassengers
func (m *mockSelfDriving) setGetPassengersResponse(output0 []string) {
	m.setGetPassengersFunc(func() []string {
		return output0
	})
}

/* -------------------------- Accelerate Mock Helpers --------------------------- */

// Accelerate overrides the method to return the mock response
func (m *mockSelfDriving) Accelerate(speed int, unit string) (int, error) {
	if m.mocked.Accelerate.enabled {
		return m.mocked.Accelerate.response(speed, unit)
	}
	return m.real.Accelerate(speed, unit)
}

// setAccelerateFunc sets the function for Accelerate
func (m *mockSelfDriving) setAccelerateFunc(f func(int, string) (int, error)) {
	m.mocked.Accelerate.response = f
}

// enableAccelerateMock turns the mock on
func (m *mockSelfDriving) enableAccelerateMock() {
	m.mocked.Accelerate.enabled = true
}

// disableAccelerateMock turns the mock off
func (m *mockSelfDriving) disableAccelerateMock() {
	m.mocked.Accelerate.enabled = false
}

// setAccelerateResponse sets the response for Accelerate
func (m *mockSelfDriving) setAccelerateResponse(output0 int, output1 error) {
	m.setAccelerateFunc(func(int, string) (int, error) {
		return output0, output1
	})
}

/* -------------------------- Telemetry Mock Helpers --------------------------- */

// Telemetry overrides the method to return the mock response
func (m *mockSelfDriving) Telemetry() map[string]float64 {
	if m.mocked.Telemetry.enabled {
		return m.mocked.Telemetry.response()
	}
	return m.real.Telemetry()
}

// setTelemetryFunc sets the function for Telemetry
func (m *mockSelfDriving) setTelemetryFunc(f func() map[string]float64) {
	m.mocked.Telemetry.response = f
}

// enableTelemetryMock turns the mock on
func (m *mockSelfDriving) enableTelemetryMock() {
	m.mocked.Telemetry.enabled = true
}

// disableTelemetryMock turns the mock off
func (m *mockSelfDriving) disableTelemetryMock() {
	m.mocked.Telemetry.enabled = false
}

// setTelemetryResponse sets the response for Telemetry
func (m *mockSelfDriving) setTelemetryResponse(output0 map[string]float64) {
	m.setTelemetryFunc(func() map[string]float64 {
		return output0
	})
}

/* -------------------------- DriveSelf Mock Helpers --------------------------- */

// DriveSelf overrides the method to return the mock response
func (m *mockSelfDriving) DriveSelf(endLocation string) error {
	if m.mocked.DriveSelf.enabled {
		return m.mocked.DriveSelf.response(endLocation)
	}
	return m.real.DriveSelf(endLocation)
}

// setDriveSelfFunc sets the function for DriveSelf
func (m *mockSelfDriving) setDriveSelfFunc(f func(string) error) {
	m.mocked.DriveSelf.response = f
}

// enableDriveSelfMock turns the mock on
func (m *mockSelfDriving) enableDriveSelfMock() {
	m.mocked.DriveSelf.enabled = true
}

// disableDriveSelfMock turns the mock off
func (m *mockSelfDriving) disableDriveSelfMock() {
	m.mocked.DriveSelf.enabled = false
}

// setDriveSelfResponse sets the response for DriveSelf
func (m *mockSelfDriving) setDriveSelfResponse(output0 error) {
	m.setDriveSelfFunc(func(string) error {
		return output0
	})
}

/* -------------------------- IsMoving Mock Helpers --------------------------- */

// IsMoving overrides the method to return the mock response
func (m *mockSelfDriving) IsMoving() bool {
	if m.mocked.IsMoving.enabled {
		return m.mocked.IsMoving.response()
	}
	return m.real.IsMoving()
}

// setIsMovingFunc sets the function for IsMoving
func (m *mockSelfDriving) setIsMovingFunc(f func() bool) {
	m.mocked.IsMoving.response = f
}

// enableIsMovingMock turns the mock on
func (m *mockSelfDriving) enableIsMovingMock() {
	m.mocked.IsMoving.enabled = true
}

// disableIsMovingMock turns the mock off
func (m *mockSelfDriving) disableIsMovingMock() {
	m.mocked.IsMoving.enabled = false
}

// setIsMovingResponse sets the response for IsMoving
func (m *mockSelfDriving) setIsMovingResponse(output0 bool) {
	m.setIsMovingFunc(func() bool {
		return output0
	})
}

/* -------------------------- GetTopSpeed Mock Helpers --------------------------- */

// GetTopSpeed overrides the method to return the mock response
func (m *mockSelfDriving) GetTopSpeed() int {
	if m.mocked.GetTopSpeed.enabled {
		return m.mocked.GetTopSpeed.response()
	}
	return m.real.GetTopSpeed()
}

// setGetTopSpeedFunc sets the function for GetTopSpeed
func (m *mockSelfDriving) setGetTopSpeedFunc(f func() int) {
	m.mocked.GetTopSpeed.response = f
}

// enableGetTopSpeedMock turns the mock on
func (m *mockSelfDriving) enableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.enabled = true
}

// disableGetTopSpeedMock turns the mock off
func (m *mockSelfDriving) disableGetTopSpeedMock() {
	m.mocked.GetTopSpeed.enabled = false
}

// setGetTopSpeedResponse sets the response for GetTopSpeed
func (m *mockSelfDriving) setGetTopSpeedResponse(output0 int) {
	m.setGetTopSpeedFunc(func() int {
		return output0
	})
}

/* -------------------------- ChangeGears Mock Helpers --------------------------- */

// ChangeGears overrides the method to return the mock response
func (m *mockSelfDriving) ChangeGears(gear int) (int, int) {
	if m.mocked.ChangeGears.enabled {
		return m.mocked.ChangeGears.response(gear)
	}
	return m.real.ChangeGears(gear)
}

// setChangeGearsFunc sets the function for ChangeGears
func (m *mockSelfDriving) setChangeGearsFunc(f func(int) (int, int)) {
	m.mocked.ChangeGears.response = f
}

// enableChangeGearsMock turns the mock on
func (m *mockSelfDriving) enableChangeGearsMock() {
	m.mocked.ChangeGears.enabled = true
}

// disableChangeGearsMock turns the mock off
func (m *mockSelfDriving) disableChangeGearsMock() {
	m.mocked.ChangeGears.enabled = false
}

// setChangeGearsResponse sets the response for ChangeGears
func (m *mockSelfDriving) setChangeGearsResponse(output0 int, output1 int) {
	m.setChangeGearsFunc(func(int) (int, int) {
		return output0, output1
	})
}

/* -------------------------- Turn Mock Helpers --------------------------- */

// Turn overrides the method to return the mock response
func (m *mockSelfDriving) Turn(dir string) string {
	if m.mocked.Turn.enabled {
		return m.mocked.Turn.response(dir)
	}
	return m.real.Turn(dir)
}

// setTurnFunc sets the function for Turn
func (m *mockSelfDriving) setTurnFunc(f func(string) string) {
	m.mocked.Turn.response = f
}

// enableTurnMock turns the mock on
func (m *mockSelfDriving) enableTurnMock() {
	m.mocked.Turn.enabled = true
}

// disableTurnMock turns the mock off
func (m *mockSelfDriving) disableTurnMock() {
	m.mocked.Turn.enabled = false
}

// setTurnResponse sets the response for Turn
func (m *mockSelfDriving) setTurnResponse(output0 string) {
	m.setTurnFunc(func(string) string {
		return output0
	})
}

/* -------------------------- GetEngineSpecs Mock Helpers --------------------------- */

// GetEngineSpecs overrides the method to return the mock response
func (m *mockSelfDriving) GetEngineSpecs() (int, string) {
	if m.mocked.GetEngineSpecs.enabled {
		return m.mocked.GetEngineSpecs.response()
	}
	return m.real.GetEngineSpecs()
}

// setGetEngineSpecsFunc sets the function for GetEngineSpecs
func (m *mockSelfDriving) setGetEngineSpecsFunc(f func() (int, string)) {
	m.mocked.GetEngineSpecs.response = f
}

// enableGetEngineSpecsMock turns the mock on
func (m *mockSelfDriving) enableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.enabled = true
}

// disableGetEngineSpecsMock turns the mock off
func (m *mockSelfDriving) disableGetEngineSpecsMock() {
	m.mocked.GetEngineSpecs.enabled = false
}

// setGetEngineSpecsResponse sets the response for GetEngineSpecs
func (m *mockSelfDriving) setGetEngineSpecsResponse(output0 int, output1 string) {
	m.setGetEngineSpecsFunc(func() (int, string) {
		return output0, output1
	})
}
