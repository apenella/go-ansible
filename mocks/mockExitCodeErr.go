package mocks

type MockExitCodeErr struct {
	error   //nolint:golint,unused
	Code    int
	Message string
}

func (e *MockExitCodeErr) Error() string {
	return e.Message
}

func (e *MockExitCodeErr) ExitCode() int {
	return e.Code
}
