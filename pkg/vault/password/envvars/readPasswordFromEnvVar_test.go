package envvars

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {

	t.Setenv("VAULT_PASSWORD", "ThatIsAPassword")

	tests := []struct {
		desc     string
		reader   *ReadPasswordFromEnvVar
		expected string
		err      error
	}{
		{
			desc: "Testing reading a secret from an environment variable",
			reader: NewReadPasswordFromEnvVar(
				WithEnvVar("VAULT_PASSWORD"),
			),
			expected: "ThatIsAPassword",
			err:      nil,
		},
		{
			desc:   "Testing error reading a secret from environment variable when ReadPasswordFromEnvVar is not initialized",
			reader: nil,
			err:    errors.New("Read password from an environment variable component has not been initialized."),
		},
		{
			desc: "Testing error reading a secret from an environment variable when it is not set",
			reader: NewReadPasswordFromEnvVar(
				WithEnvVar("UNSET_VARIALBE"),
			),
			err: errors.New("The environment variable 'UNSET_VARIALBE' is not set"),
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			secret, err := test.reader.Read()
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, test.expected, secret)
			}
		})
	}
}
