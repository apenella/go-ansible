package resolve

import (
	"errors"
)

// OptionsFunc is a function used to configure ReadPasswordResolve
type OptionsFunc func(*ReadPasswordResolve)

// ReadPasswordResolve contains multiple methods that can resolve a password
type ReadPasswordResolve struct {
	reader []PasswordReader
}

// NewReadPasswordResolve return a ReadPasswordResolve
func NewReadPasswordResolve(options ...OptionsFunc) *ReadPasswordResolve {
	secret := &ReadPasswordResolve{}
	secret.Options(options...)

	return secret
}

// WithReader allow you to set a list of readers that you can use to resolve a password
func WithReader(reader ...PasswordReader) OptionsFunc {
	return func(s *ReadPasswordResolve) {
		if s.reader == nil {
			s.reader = []PasswordReader{}
		}

		s.reader = append(s.reader, reader...)
	}
}

// Options configure the ReadPasswordResolve
func (s *ReadPasswordResolve) Options(opts ...OptionsFunc) {
	for _, opt := range opts {
		opt(s)
	}
}

// Read looks for the first reader defined into the reader attribute which returns a password
func (s *ReadPasswordResolve) Read() (string, error) {
	if s == nil {
		return "", errors.New("The component to resolve read password mechanism has not been initialized.")
	}

	for _, reader := range s.reader {
		secret, err := reader.Read()
		if err == nil {
			return secret, nil
		}
	}

	return "", errors.New("The component to resolve read password does not found a password.")
}
