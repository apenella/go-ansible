package exec

import (
	"io"

	"github.com/stretchr/testify/mock"
)

// MockCmd struct is a mock of exec.Cmd
type MockCmd struct {
	mock.Mock
}

// NewMockCmd return a Mock for exec.Cmd
func NewMockCmd() *MockCmd {
	return &MockCmd{}
}

// CombinedOutput is a mock of exec.Cmd CombinedOutput method
func (c *MockCmd) CombinedOutput() ([]byte, error) {
	args := c.Mock.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// Environ is a mock of exec.Cmd Environ method
func (c *MockCmd) Environ() []string {
	args := c.Mock.Called()
	return args.Get(0).([]string)
}

// Output is a mock of exec.Cmd Output method
func (c *MockCmd) Output() ([]byte, error) {
	args := c.Mock.Called()
	return args.Get(0).([]byte), args.Error(1)
}

// Run is a mock of exec.Cmd Run method
func (c *MockCmd) Run() error {
	args := c.Mock.Called()
	return args.Error(0)
}

// Start is a mock of exec.Cmd Start method
func (c *MockCmd) Start() error {
	args := c.Mock.Called()
	return args.Error(0)
}

// StderrPipe is a mock of exec.Cmd StderrPipe method
func (c *MockCmd) StderrPipe() (io.ReadCloser, error) {
	args := c.Mock.Called()
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

// StdinPipe is a mock of exec.Cmd StdinPipe method
func (c *MockCmd) StdinPipe() (io.WriteCloser, error) {
	args := c.Mock.Called()
	return args.Get(0).(io.WriteCloser), args.Error(1)
}

// StdoutPipe is a mock of exec.Cmd StdoutPipe method
func (c *MockCmd) StdoutPipe() (io.ReadCloser, error) {
	args := c.Mock.Called()
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

// String is a mock of exec.Cmd String method
func (c *MockCmd) String() string {
	args := c.Mock.Called()
	return args.String(0)
}

// Wait is a mock of exec.Cmd Wait method
func (c *MockCmd) Wait() error {
	args := c.Mock.Called()
	return args.Error(0)
}

// func (c *MockCmd) Dir(dir string) {
// 	c.Mock.Called(dir)
// }
// func (c *MockCmd) Env(env []string) {
// 	c.Mock.Called(env)
// }
// func (c *MockCmd) Stdin(file *os.File) {
// 	c.Mock.Called(file)
// }
