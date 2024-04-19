package mocks

type MockExitCodeErr struct {
	error
	Code    int
	Message string
}

func (e *MockExitCodeErr) Error() string {
	return e.Message
}

func (e *MockExitCodeErr) ExitCode() int {
	return e.Code
}
