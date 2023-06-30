package encrypt

import "github.com/stretchr/testify/mock"

type MockEncryptString struct {
	mock.Mock
}

func NewMockEncryptString() *MockEncryptString {
	return &MockEncryptString{}
}

func (e *MockEncryptString) Encrypt(plainText string) (string, error) {
	args := e.Called(plainText)

	return args.String(0), args.Error(1)
}
