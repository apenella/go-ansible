package exec

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockExec struct is wrapper of os.Exec that enable you to mock os.Cmd
type MockExec struct {
	mock.Mock
}

// NewMockExec returns a MockExec to mock to Cmd
func NewMockExec() *MockExec {
	return &MockExec{}
}

// Command is a wrapper of exec.Command
func (e *MockExec) Command(name string, arg ...string) Cmder {
	ret := e.Mock.Called(name, append([]string{}, arg...))
	return ret.Get(0).(*MockCmd)
}

// CommandContext is a wrapper of exec.CommandContext
func (e *MockExec) CommandContext(ctx context.Context, name string, arg ...string) Cmder {
	ret := e.Mock.Called(ctx, name, append([]string{}, arg...))
	return ret.Get(0).(*MockCmd)
}
