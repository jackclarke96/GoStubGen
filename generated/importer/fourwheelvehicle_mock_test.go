package driver

	import "github.com/jackclarke/GoStubGen/generated/vehicle"

// mockFourWheelVehicleConfig stores mock flags and responses
type mockFourWheelVehicleConfig struct {

	// Drive flag and mock response
	mockDrive     bool
	DriveResponse func() (string)

	// Stop flag and mock response
	mockStop     bool
	StopResponse func() (string)

}

// mockFourWheelVehicle embeds a concrete FourWheelVehicle and its mocks
type mockFourWheelVehicle struct {
	real   vehicle.FourWheelVehicle
	mocked mockFourWheelVehicleConfig
}

// fourWheelVehicleMock returns a new mock
func fourWheelVehicleMock(v vehicle.FourWheelVehicle) *mockFourWheelVehicle {
	return &mockFourWheelVehicle{
		real:   v,
		mocked: mockFourWheelVehicleConfig{},
	}
}
/* -------------------------- Drive Mock Helpers --------------------------- */

// Drive overrides the method to return the mock response
func (m *mockFourWheelVehicle) Drive() (string) {
	if m.mocked.mockDrive {
		return m.mocked.DriveResponse()
	}
	return m.real.Drive()
}

// setDriveFunc sets the function for Drive
func (m *mockFourWheelVehicle) setDriveFunc(f func() (string)) {
	m.mocked.DriveResponse = f
}


// setDriveResponse sets the response for Drive
func (m *mockFourWheelVehicle) setDriveResponse (output0 string) {
	m.setDriveFunc(func() (string) {
		return output0
	})
}


// enableDriveMock turns the mock on
func (m *mockFourWheelVehicle) enableDriveMock() {
	m.mocked.mockDrive = true
}

// disableDriveMock turns the mock off
func (m *mockFourWheelVehicle) disableDriveMock() {
	m.mocked.mockDrive = false
}
/* -------------------------- Stop Mock Helpers --------------------------- */

// Stop overrides the method to return the mock response
func (m *mockFourWheelVehicle) Stop() (string) {
	if m.mocked.mockStop {
		return m.mocked.StopResponse()
	}
	return m.real.Stop()
}

// setStopFunc sets the function for Stop
func (m *mockFourWheelVehicle) setStopFunc(f func() (string)) {
	m.mocked.StopResponse = f
}


// setStopResponse sets the response for Stop
func (m *mockFourWheelVehicle) setStopResponse (output0 string) {
	m.setStopFunc(func() (string) {
		return output0
	})
}


// enableStopMock turns the mock on
func (m *mockFourWheelVehicle) enableStopMock() {
	m.mocked.mockStop = true
}

// disableStopMock turns the mock off
func (m *mockFourWheelVehicle) disableStopMock() {
	m.mocked.mockStop = false
}
