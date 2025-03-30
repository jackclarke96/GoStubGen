package driver

	import "github.com/jackclarke/GoStubGen/generated/vehicle"

// mockSelfDrivingConfig stores mock flags and responses
type mockSelfDrivingConfig struct {

	// ActivateAutopilot flag and mock response
	mockActivateAutopilot     bool
	ActivateAutopilotResponse func() (string)

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
