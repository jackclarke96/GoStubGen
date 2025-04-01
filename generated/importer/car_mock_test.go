package driver

	import "github.com/jackclarke/GoStubGen/generated/vehicle"

// mockCarConfig stores mock flags and responses
type mockCarConfig struct {

	// OpenTrunk flag and mock response
	mockOpenTrunk     bool
	OpenTrunkResponse func() (string)

	// Stop flag and mock response
	mockStop     bool
	StopResponse func() (string)

	// Drive flag and mock response
	mockDrive     bool
	DriveResponse func() (string)

}

// mockCar embeds a concrete Car and its mocks
type mockCar struct {
	real   vehicle.Car
	mocked mockCarConfig
}

// carMock returns a new mock
func carMock(v vehicle.Car) *mockCar {
	return &mockCar{
		real:   v,
		mocked: mockCarConfig{},
	}
}
/* -------------------------- OpenTrunk Mock Helpers --------------------------- */

// OpenTrunk overrides the method to return the mock response
func (m *mockCar) OpenTrunk() (string) {
	if m.mocked.mockOpenTrunk {
		return m.mocked.OpenTrunkResponse()
	}
	return m.real.OpenTrunk()
}

// setOpenTrunkFunc sets the function for OpenTrunk
func (m *mockCar) setOpenTrunkFunc(f func() (string)) {
	m.mocked.OpenTrunkResponse = f
}


// setOpenTrunkResponse sets the response for OpenTrunk
func (m *mockCar) setOpenTrunkResponse (output0 string) {
	m.setOpenTrunkFunc(func() (string) {
		return output0
	})
}


// enableOpenTrunkMock turns the mock on
func (m *mockCar) enableOpenTrunkMock() {
	m.mocked.mockOpenTrunk = true
}

// disableOpenTrunkMock turns the mock off
func (m *mockCar) disableOpenTrunkMock() {
	m.mocked.mockOpenTrunk = false
}
/* -------------------------- Stop Mock Helpers --------------------------- */

// Stop overrides the method to return the mock response
func (m *mockCar) Stop() (string) {
	if m.mocked.mockStop {
		return m.mocked.StopResponse()
	}
	return m.real.Stop()
}

// setStopFunc sets the function for Stop
func (m *mockCar) setStopFunc(f func() (string)) {
	m.mocked.StopResponse = f
}


// setStopResponse sets the response for Stop
func (m *mockCar) setStopResponse (output0 string) {
	m.setStopFunc(func() (string) {
		return output0
	})
}


// enableStopMock turns the mock on
func (m *mockCar) enableStopMock() {
	m.mocked.mockStop = true
}

// disableStopMock turns the mock off
func (m *mockCar) disableStopMock() {
	m.mocked.mockStop = false
}
/* -------------------------- Drive Mock Helpers --------------------------- */

// Drive overrides the method to return the mock response
func (m *mockCar) Drive() (string) {
	if m.mocked.mockDrive {
		return m.mocked.DriveResponse()
	}
	return m.real.Drive()
}

// setDriveFunc sets the function for Drive
func (m *mockCar) setDriveFunc(f func() (string)) {
	m.mocked.DriveResponse = f
}


// setDriveResponse sets the response for Drive
func (m *mockCar) setDriveResponse (output0 string) {
	m.setDriveFunc(func() (string) {
		return output0
	})
}


// enableDriveMock turns the mock on
func (m *mockCar) enableDriveMock() {
	m.mocked.mockDrive = true
}

// disableDriveMock turns the mock off
func (m *mockCar) disableDriveMock() {
	m.mocked.mockDrive = false
}
