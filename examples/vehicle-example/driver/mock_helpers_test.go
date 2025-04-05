package driver

// MethodConfig is a generic container for mocking method behavior.
type methodConfig[F any] struct {
	enabled  bool
	response F
}
