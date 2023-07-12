package text

import (
	"errors"
)

// OptionsFunc is a function used to configure ReadPasswordFromText
type OptionsFunc func(*ReadPasswordFromText)

// ReadPasswordFromText allow to read a password as a plain text
type ReadPasswordFromText struct {
	text string
}

// NewReadPasswordFromText generates a ReadPasswordFromText struct
func NewReadPasswordFromText(options ...OptionsFunc) *ReadPasswordFromText {
	secret := &ReadPasswordFromText{}
	secret.Options(options...)

	return secret
}

// WithText set the text (password) into the ReadPasswordFromText
func WithText(text string) OptionsFunc {
	return func(s *ReadPasswordFromText) {
		s.text = text
	}
}

// Options configure the ReadPasswordFromText
func (s *ReadPasswordFromText) Options(opts ...OptionsFunc) {
	for _, opt := range opts {
		opt(s)
	}
}

// Read returns a password save the text attribute from the ReadPasswordFromText struct
func (s *ReadPasswordFromText) Read() (string, error) {
	if s == nil {
		return "", errors.New("Password input from text has not been initialized.")
	}

	if len(s.text) <= 0 {
		return "", errors.New("Text must be specified to use the password input from text.")
	}

	return s.text, nil
}
