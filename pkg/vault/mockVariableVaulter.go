package vault

import "github.com/stretchr/testify/mock"

// MockVariableVaulter is a mock for VariableVaulter
type MockVariableVaulter struct {
	mock.Mock
}

// NewMockVariableVaulter returns a new MockVariableVaulter
func NewMockVariableVaulter() *MockVariableVaulter {
	return &MockVariableVaulter{}
}

func (v *MockVariableVaulter) Vault(value string) (*VaultVariableValue, error) {
	args := v.Called(value)
	return args.Get(0).(*VaultVariableValue), args.Error(1)
}
