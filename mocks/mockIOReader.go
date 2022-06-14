package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockIOReader struct {
	mock.Mock
}

func NewMockIOReader() *MockIOReader {
	return &MockIOReader{}
}

func (m *MockIOReader) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Get(0).(int), args.Get(1).(error)
}
