package driver

	import "github.com/jackclarke/GoStubGen/generated/vehicle"

// mockCarConfig stores mock flags and responses
type mockCarConfig struct {

	// OpenTrunk flag and mock response
	mockOpenTrunk     bool
	OpenTrunkResponse func() (string)

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
