package envvars

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// OptionsFunc is a function used to configure ReadPasswordFromEnvVar
type OptionsFunc func(*ReadPasswordFromEnvVar)

type ReadPasswordFromEnvVar struct {
	envvar string
}

func NewReadPasswordFromEnvVar(options ...OptionsFunc) *ReadPasswordFromEnvVar {
	secret := &ReadPasswordFromEnvVar{}
	secret.Options(options...)

	return secret
}

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

func (s *ReadPasswordFromEnvVar) Read() (string, error) {
	if s == nil {
		return "", errors.New("Read password from an environment variable component has not been initialized.")
	}

	secret := os.Getenv(s.envvar)
	if len(secret) <= 0 {
		errors.New(fmt.Sprintf("The environment variable '%s' is not set.", s.envvar))
	}

	return secret, nil
}
