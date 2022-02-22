package execute

import (
	"context"

	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	"github.com/stretchr/testify/mock"
)

// MockExecute is a mock of Execute interface
type MockExecute struct {
	mock.Mock
}

// NewMockExecute returns a new instance of MockExecute
func NewMockExecute() *MockExecute {
	return &MockExecute{}
}

// Execute is a mock
func (e *MockExecute) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error {
	args := e.Called(ctx, command, resultsFunc, options)
	return args.Error(0)
}
