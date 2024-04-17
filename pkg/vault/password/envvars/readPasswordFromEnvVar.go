package envvars

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

// OptionsFunc is a function used to configure ReadPasswordFromEnvVar
type OptionsFunc func(*ReadPasswordFromEnvVar)

// ReadPasswordFromEnvVar returns a password set in an environment variable
type ReadPasswordFromEnvVar struct {
	envvar string
}

// NewReadPasswordFromEnvVar returns a ReadPasswordFromEnvVar
func NewReadPasswordFromEnvVar(options ...OptionsFunc) *ReadPasswordFromEnvVar {
	secret := &ReadPasswordFromEnvVar{}
	secret.Options(options...)

	return secret
}

// WithEnvVar set the environment variable that contains the password defined on it
func WithEnvVar(envvar string) OptionsFunc {
	return func(s *ReadPasswordFromEnvVar) {
		s.envvar = envvar
	}
}

// Options configure the ReadPasswordFromEnvVar
func (s *ReadPasswordFromEnvVar) Options(opts ...OptionsFunc) {
	for _, opt := range opts {
		opt(s)
	}
}

// Read returns the password that is set in the environment variable defined on the envvar attribute
func (s *ReadPasswordFromEnvVar) Read() (string, error) {
	if s == nil {
		return "", errors.New("Read password from an environment variable component has not been initialized.")
	}

	secret := os.Getenv(s.envvar)
	if len(secret) <= 0 {
		log.Printf("The environment variable '%s' is not set.", s.envvar)
	}

	return secret, nil
}
