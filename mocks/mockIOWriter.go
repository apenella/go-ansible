package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockIOWriter struct {
	mock.Mock
}

func NewMockIOWriter() *MockIOWriter {
	return &MockIOWriter{}
}

func (m *MockIOWriter) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Get(0).(int), args.Get(1).(error)
}
