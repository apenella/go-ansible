package mock

import (
	"github.com/stretchr/testify/mock"
)

type MockReadPassword struct {
	mock.Mock
}

func NewMockReadPassword() *MockReadPassword {
	return &MockReadPassword{}
}

// Secret returns a secrets
func (s *MockReadPassword) Read() (string, error) {
	args := s.Called()

	return args.String(0), args.Error(1)
}
