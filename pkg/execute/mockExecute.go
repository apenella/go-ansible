package execute

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute/result"
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

// Quiet is a mock
func (e *MockExecute) Quiet() {
	e.Called()
}

// Execute is a mock
func (e *MockExecute) Execute(ctx context.Context) error {
	args := e.Called(ctx)
	return args.Error(0)
}

// AddEnvVar is a mock
func (e *MockExecute) AddEnvVar(key, value string) {
	e.Called(key, value)
}

// WithOutput is a mock
func (e *MockExecute) WithOutput(output result.ResultsOutputer) {
	e.Called(output)
}
