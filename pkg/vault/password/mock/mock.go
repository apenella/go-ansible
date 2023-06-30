package mock

import (
	"github.com/stretchr/testify/mock"
)

// MockReadPassword mocks the PasswordReader
type MockReadPassword struct {
	mock.Mock
}

// NewMockReadPassword return a MockReadPassword
func NewMockReadPassword() *MockReadPassword {
	return &MockReadPassword{}
}

// Read returns a mocked password
func (s *MockReadPassword) Read() (string, error) {
	args := s.Called()

	return args.String(0), args.Error(1)
}
