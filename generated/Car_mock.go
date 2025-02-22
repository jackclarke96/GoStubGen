package vehicle

// mockVehicleConfig stores mock flags and responses
type mockVehicleConfig struct {
	// getTopSpeed flag and mock response
	mockGetTopSpeed     bool
	getTopSpeedResponse func() int
	// turn flag and mock response
	mockTurn     bool
	turnResponse func(string) string
	// reverse flag and mock response
	mockReverse     bool
	reverseResponse func() (string, error)
	// accelerate flag and mock response
	mockAccelerate     bool
	accelerateResponse func(int, string) (int, error)
	// isMoving flag and mock response
	mockIsMoving     bool
	isMovingResponse func() bool
	// honk flag and mock response
	mockHonk     bool
	honkResponse func(int)
	// getEngineSpecs flag and mock response
	mockGetEngineSpecs     bool
	getEngineSpecsResponse func() (int, string)
	// applyBrakes flag and mock response
	mockApplyBrakes     bool
	applyBrakesResponse func(float64) bool
	// changeGears flag and mock response
	mockChangeGears     bool
	changeGearsResponse func(int) (int, int)
	// telemetry flag and mock response
	mockTelemetry     bool
	telemetryResponse func() map[string]float64
	// getPassengers flag and mock response
	mockGetPassengers     bool
	getPassengersResponse func() []string
	// loadCargo flag and mock response
	mockLoadCargo     bool
	loadCargoResponse func([]string) (int, error)
	// getVehicleStatus flag and mock response
	mockGetVehicleStatus     bool
	getVehicleStatusResponse func() VehicleStatus
	// updateStatus flag and mock response
	mockUpdateStatus     bool
	updateStatusResponse func(VehicleStatus) error
}

// mockVehicle embeds Car and its mocks
type mockVehicle struct {
	car   Car
	mocks mockVehicleConfig
}

// vehicleMock returns a new mock
func vehicleMock() Vehicle {
	car := NewCar()
	return mockVehicle{
		car:   *car,
		mocks: mockVehicleConfig{},
	}
}

/* -------------------------- getTopSpeed Mock Helpers --------------------------- */

// getTopSpeed overrides the method to return the mock response
func (m mockVehicle) getTopSpeed() int {
	if m.mocks.mockGetTopSpeed {

		return m.mocks.getTopSpeedResponse()

	}

	return m.car.getTopSpeed()

}

// setGetTopSpeedResponse sets the response for getTopSpeed
func (m mockVehicle) setGetTopSpeedResponse(resp func() int) {
	m.mocks.getTopSpeedResponse = resp
}

// enableGetTopSpeedResponse turns the mock on
func (m mockVehicle) enableGetTopSpeedResponse() {
	m.mocks.mockGetTopSpeed = true
}

// disableGetTopSpeedResponse turns the mock off
func (m mockVehicle) disableGetTopSpeedResponse() {
	m.mocks.mockGetTopSpeed = false
}

/* -------------------------- turn Mock Helpers --------------------------- */

// turn overrides the method to return the mock response
func (m mockVehicle) turn(dir string) string {
	if m.mocks.mockTurn {

		return m.mocks.turnResponse(dir)

	}

	return m.car.turn(dir)

}

// setTurnResponse sets the response for turn
func (m mockVehicle) setTurnResponse(resp func(string) string) {
	m.mocks.turnResponse = resp
}

// enableTurnResponse turns the mock on
func (m mockVehicle) enableTurnResponse() {
	m.mocks.mockTurn = true
}

// disableTurnResponse turns the mock off
func (m mockVehicle) disableTurnResponse() {
	m.mocks.mockTurn = false
}

/* -------------------------- reverse Mock Helpers --------------------------- */

// reverse overrides the method to return the mock response
func (m mockVehicle) reverse() (string, error) {
	if m.mocks.mockReverse {

		return m.mocks.reverseResponse()

	}

	return m.car.reverse()

}

// setReverseResponse sets the response for reverse
func (m mockVehicle) setReverseResponse(resp func() (string, error)) {
	m.mocks.reverseResponse = resp
}

// enableReverseResponse turns the mock on
func (m mockVehicle) enableReverseResponse() {
	m.mocks.mockReverse = true
}

// disableReverseResponse turns the mock off
func (m mockVehicle) disableReverseResponse() {
	m.mocks.mockReverse = false
}

/* -------------------------- accelerate Mock Helpers --------------------------- */

// accelerate overrides the method to return the mock response
func (m mockVehicle) accelerate(speed int, unit string) (int, error) {
	if m.mocks.mockAccelerate {

		return m.mocks.accelerateResponse(speed, unit)

	}

	return m.car.accelerate(speed, unit)

}

// setAccelerateResponse sets the response for accelerate
func (m mockVehicle) setAccelerateResponse(resp func(int, string) (int, error)) {
	m.mocks.accelerateResponse = resp
}

// enableAccelerateResponse turns the mock on
func (m mockVehicle) enableAccelerateResponse() {
	m.mocks.mockAccelerate = true
}

// disableAccelerateResponse turns the mock off
func (m mockVehicle) disableAccelerateResponse() {
	m.mocks.mockAccelerate = false
}

/* -------------------------- isMoving Mock Helpers --------------------------- */

// isMoving overrides the method to return the mock response
func (m mockVehicle) isMoving() bool {
	if m.mocks.mockIsMoving {

		return m.mocks.isMovingResponse()

	}

	return m.car.isMoving()

}

// setIsMovingResponse sets the response for isMoving
func (m mockVehicle) setIsMovingResponse(resp func() bool) {
	m.mocks.isMovingResponse = resp
}

// enableIsMovingResponse turns the mock on
func (m mockVehicle) enableIsMovingResponse() {
	m.mocks.mockIsMoving = true
}

// disableIsMovingResponse turns the mock off
func (m mockVehicle) disableIsMovingResponse() {
	m.mocks.mockIsMoving = false
}

/* -------------------------- honk Mock Helpers --------------------------- */

// honk overrides the method to return the mock response
func (m mockVehicle) honk(times int) {
	if m.mocks.mockHonk {

		m.mocks.honkResponse(times)

	}

	m.car.honk(times)

}

// setHonkResponse sets the response for honk
func (m mockVehicle) setHonkResponse(resp func(int)) {
	m.mocks.honkResponse = resp
}

// enableHonkResponse turns the mock on
func (m mockVehicle) enableHonkResponse() {
	m.mocks.mockHonk = true
}

// disableHonkResponse turns the mock off
func (m mockVehicle) disableHonkResponse() {
	m.mocks.mockHonk = false
}

/* -------------------------- getEngineSpecs Mock Helpers --------------------------- */

// getEngineSpecs overrides the method to return the mock response
func (m mockVehicle) getEngineSpecs() (int, string) {
	if m.mocks.mockGetEngineSpecs {

		return m.mocks.getEngineSpecsResponse()

	}

	return m.car.getEngineSpecs()

}

// setGetEngineSpecsResponse sets the response for getEngineSpecs
func (m mockVehicle) setGetEngineSpecsResponse(resp func() (int, string)) {
	m.mocks.getEngineSpecsResponse = resp
}

// enableGetEngineSpecsResponse turns the mock on
func (m mockVehicle) enableGetEngineSpecsResponse() {
	m.mocks.mockGetEngineSpecs = true
}

// disableGetEngineSpecsResponse turns the mock off
func (m mockVehicle) disableGetEngineSpecsResponse() {
	m.mocks.mockGetEngineSpecs = false
}

/* -------------------------- applyBrakes Mock Helpers --------------------------- */

// applyBrakes overrides the method to return the mock response
func (m mockVehicle) applyBrakes(force float64) bool {
	if m.mocks.mockApplyBrakes {

		return m.mocks.applyBrakesResponse(force)

	}

	return m.car.applyBrakes(force)

}

// setApplyBrakesResponse sets the response for applyBrakes
func (m mockVehicle) setApplyBrakesResponse(resp func(float64) bool) {
	m.mocks.applyBrakesResponse = resp
}

// enableApplyBrakesResponse turns the mock on
func (m mockVehicle) enableApplyBrakesResponse() {
	m.mocks.mockApplyBrakes = true
}

// disableApplyBrakesResponse turns the mock off
func (m mockVehicle) disableApplyBrakesResponse() {
	m.mocks.mockApplyBrakes = false
}

/* -------------------------- changeGears Mock Helpers --------------------------- */

// changeGears overrides the method to return the mock response
func (m mockVehicle) changeGears(gear int) (int, int) {
	if m.mocks.mockChangeGears {

		return m.mocks.changeGearsResponse(gear)

	}

	return m.car.changeGears(gear)

}

// setChangeGearsResponse sets the response for changeGears
func (m mockVehicle) setChangeGearsResponse(resp func(int) (int, int)) {
	m.mocks.changeGearsResponse = resp
}

// enableChangeGearsResponse turns the mock on
func (m mockVehicle) enableChangeGearsResponse() {
	m.mocks.mockChangeGears = true
}

// disableChangeGearsResponse turns the mock off
func (m mockVehicle) disableChangeGearsResponse() {
	m.mocks.mockChangeGears = false
}

/* -------------------------- telemetry Mock Helpers --------------------------- */

// telemetry overrides the method to return the mock response
func (m mockVehicle) telemetry() map[string]float64 {
	if m.mocks.mockTelemetry {

		return m.mocks.telemetryResponse()

	}

	return m.car.telemetry()

}

// setTelemetryResponse sets the response for telemetry
func (m mockVehicle) setTelemetryResponse(resp func() map[string]float64) {
	m.mocks.telemetryResponse = resp
}

// enableTelemetryResponse turns the mock on
func (m mockVehicle) enableTelemetryResponse() {
	m.mocks.mockTelemetry = true
}

// disableTelemetryResponse turns the mock off
func (m mockVehicle) disableTelemetryResponse() {
	m.mocks.mockTelemetry = false
}

/* -------------------------- getPassengers Mock Helpers --------------------------- */

// getPassengers overrides the method to return the mock response
func (m mockVehicle) getPassengers() []string {
	if m.mocks.mockGetPassengers {

		return m.mocks.getPassengersResponse()

	}

	return m.car.getPassengers()

}

// setGetPassengersResponse sets the response for getPassengers
func (m mockVehicle) setGetPassengersResponse(resp func() []string) {
	m.mocks.getPassengersResponse = resp
}

// enableGetPassengersResponse turns the mock on
func (m mockVehicle) enableGetPassengersResponse() {
	m.mocks.mockGetPassengers = true
}

// disableGetPassengersResponse turns the mock off
func (m mockVehicle) disableGetPassengersResponse() {
	m.mocks.mockGetPassengers = false
}

/* -------------------------- loadCargo Mock Helpers --------------------------- */

// loadCargo overrides the method to return the mock response
func (m mockVehicle) loadCargo(items []string) (int, error) {
	if m.mocks.mockLoadCargo {

		return m.mocks.loadCargoResponse(items)

	}

	return m.car.loadCargo(items)

}

// setLoadCargoResponse sets the response for loadCargo
func (m mockVehicle) setLoadCargoResponse(resp func([]string) (int, error)) {
	m.mocks.loadCargoResponse = resp
}

// enableLoadCargoResponse turns the mock on
func (m mockVehicle) enableLoadCargoResponse() {
	m.mocks.mockLoadCargo = true
}

// disableLoadCargoResponse turns the mock off
func (m mockVehicle) disableLoadCargoResponse() {
	m.mocks.mockLoadCargo = false
}

/* -------------------------- getVehicleStatus Mock Helpers --------------------------- */

// getVehicleStatus overrides the method to return the mock response
func (m mockVehicle) getVehicleStatus() VehicleStatus {
	if m.mocks.mockGetVehicleStatus {

		return m.mocks.getVehicleStatusResponse()

	}

	return m.car.getVehicleStatus()

}

// setGetVehicleStatusResponse sets the response for getVehicleStatus
func (m mockVehicle) setGetVehicleStatusResponse(resp func() VehicleStatus) {
	m.mocks.getVehicleStatusResponse = resp
}

// enableGetVehicleStatusResponse turns the mock on
func (m mockVehicle) enableGetVehicleStatusResponse() {
	m.mocks.mockGetVehicleStatus = true
}

// disableGetVehicleStatusResponse turns the mock off
func (m mockVehicle) disableGetVehicleStatusResponse() {
	m.mocks.mockGetVehicleStatus = false
}

/* -------------------------- updateStatus Mock Helpers --------------------------- */

// updateStatus overrides the method to return the mock response
func (m mockVehicle) updateStatus(status VehicleStatus) error {
	if m.mocks.mockUpdateStatus {

		return m.mocks.updateStatusResponse(status)

	}

	return m.car.updateStatus(status)

}

// setUpdateStatusResponse sets the response for updateStatus
func (m mockVehicle) setUpdateStatusResponse(resp func(VehicleStatus) error) {
	m.mocks.updateStatusResponse = resp
}

// enableUpdateStatusResponse turns the mock on
func (m mockVehicle) enableUpdateStatusResponse() {
	m.mocks.mockUpdateStatus = true
}

// disableUpdateStatusResponse turns the mock off
func (m mockVehicle) disableUpdateStatusResponse() {
	m.mocks.mockUpdateStatus = false
}
