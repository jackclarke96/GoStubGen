package driver

	import "github.com/jackclarke/GoStubGen/generated/vehicle"

// mockSelfDrivingConfig stores mock flags and responses
type mockSelfDrivingConfig struct {

	// ActivateAutopilot flag and mock response
	mockActivateAutopilot     bool
	ActivateAutopilotResponse func() (string)

	// Drive flag and mock response
	mockDrive     bool
	DriveResponse func() (string)

	// OpenTrunk flag and mock response
	mockOpenTrunk     bool
	OpenTrunkResponse func() (string)

	// Stop flag and mock response
	mockStop     bool
	StopResponse func() (string)

}

// mockSelfDriving embeds a concrete SelfDriving and its mocks
type mockSelfDriving struct {
	real   vehicle.SelfDriving
	mocked mockSelfDrivingConfig
}

// selfDrivingMock returns a new mock
func selfDrivingMock(v vehicle.SelfDriving) *mockSelfDriving {
	return &mockSelfDriving{
		real:   v,
		mocked: mockSelfDrivingConfig{},
	}
}
/* -------------------------- ActivateAutopilot Mock Helpers --------------------------- */

// ActivateAutopilot overrides the method to return the mock response
func (m *mockSelfDriving) ActivateAutopilot() (string) {
	if m.mocked.mockActivateAutopilot {
		return m.mocked.ActivateAutopilotResponse()
	}
	return m.real.ActivateAutopilot()
}

// setActivateAutopilotFunc sets the function for ActivateAutopilot
func (m *mockSelfDriving) setActivateAutopilotFunc(f func() (string)) {
	m.mocked.ActivateAutopilotResponse = f
}


// setActivateAutopilotResponse sets the response for ActivateAutopilot
func (m *mockSelfDriving) setActivateAutopilotResponse (output0 string) {
	m.setActivateAutopilotFunc(func() (string) {
		return output0
	})
}


// enableActivateAutopilotMock turns the mock on
func (m *mockSelfDriving) enableActivateAutopilotMock() {
	m.mocked.mockActivateAutopilot = true
}

// disableActivateAutopilotMock turns the mock off
func (m *mockSelfDriving) disableActivateAutopilotMock() {
	m.mocked.mockActivateAutopilot = false
}
/* -------------------------- Drive Mock Helpers --------------------------- */

// Drive overrides the method to return the mock response
func (m *mockSelfDriving) Drive() (string) {
	if m.mocked.mockDrive {
		return m.mocked.DriveResponse()
	}
	return m.real.Drive()
}

// setDriveFunc sets the function for Drive
func (m *mockSelfDriving) setDriveFunc(f func() (string)) {
	m.mocked.DriveResponse = f
}


// setDriveResponse sets the response for Drive
func (m *mockSelfDriving) setDriveResponse (output0 string) {
	m.setDriveFunc(func() (string) {
		return output0
	})
}


// enableDriveMock turns the mock on
func (m *mockSelfDriving) enableDriveMock() {
	m.mocked.mockDrive = true
}

// disableDriveMock turns the mock off
func (m *mockSelfDriving) disableDriveMock() {
	m.mocked.mockDrive = false
}
/* -------------------------- OpenTrunk Mock Helpers --------------------------- */

// OpenTrunk overrides the method to return the mock response
func (m *mockSelfDriving) OpenTrunk() (string) {
	if m.mocked.mockOpenTrunk {
		return m.mocked.OpenTrunkResponse()
	}
	return m.real.OpenTrunk()
}

// setOpenTrunkFunc sets the function for OpenTrunk
func (m *mockSelfDriving) setOpenTrunkFunc(f func() (string)) {
	m.mocked.OpenTrunkResponse = f
}


// setOpenTrunkResponse sets the response for OpenTrunk
func (m *mockSelfDriving) setOpenTrunkResponse (output0 string) {
	m.setOpenTrunkFunc(func() (string) {
		return output0
	})
}


// enableOpenTrunkMock turns the mock on
func (m *mockSelfDriving) enableOpenTrunkMock() {
	m.mocked.mockOpenTrunk = true
}

// disableOpenTrunkMock turns the mock off
func (m *mockSelfDriving) disableOpenTrunkMock() {
	m.mocked.mockOpenTrunk = false
}
/* -------------------------- Stop Mock Helpers --------------------------- */

// Stop overrides the method to return the mock response
func (m *mockSelfDriving) Stop() (string) {
	if m.mocked.mockStop {
		return m.mocked.StopResponse()
	}
	return m.real.Stop()
}

// setStopFunc sets the function for Stop
func (m *mockSelfDriving) setStopFunc(f func() (string)) {
	m.mocked.StopResponse = f
}


// setStopResponse sets the response for Stop
func (m *mockSelfDriving) setStopResponse (output0 string) {
	m.setStopFunc(func() (string) {
		return output0
	})
}


// enableStopMock turns the mock on
func (m *mockSelfDriving) enableStopMock() {
	m.mocked.mockStop = true
}

// disableStopMock turns the mock off
func (m *mockSelfDriving) disableStopMock() {
	m.mocked.mockStop = false
}
