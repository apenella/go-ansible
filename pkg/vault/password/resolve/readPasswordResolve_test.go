package resolve

import (
	"errors"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/vault/password/envvars"
	"github.com/apenella/go-ansible/v2/pkg/vault/password/file"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	testFs := afero.NewMemMapFs()

	t.Setenv("VAULT_PASSWORD", "ThatIsAPasswordFromEnvVar")
	err := afero.WriteFile(testFs, "/password", []byte("ThatIsAPasswordFromFile"), 0666)
	if err != nil {
		t.Log(err)
	}

	tests := []struct {
		desc     string
		reader   *ReadPasswordResolve
		expected string
		err      error
	}{
		{
			desc: "Testing resolve the password from file password reader",
			reader: NewReadPasswordResolve(
				WithReader(
					file.NewReadPasswordFromFile(
						file.WithFs(testFs),
						file.WithFile("/password"),
					),
					envvars.NewReadPasswordFromEnvVar(
						envvars.WithEnvVar("VAULT_PASSWORD"),
					),
				),
			),
			expected: "ThatIsAPasswordFromFile",
			err:      nil,
		},
		{
			desc: "Testing resolve the password from file password reader",
			reader: NewReadPasswordResolve(
				WithReader(
					envvars.NewReadPasswordFromEnvVar(
						envvars.WithEnvVar("VAULT_PASSWORD"),
					),
					file.NewReadPasswordFromFile(
						file.WithFs(testFs),
						file.WithFile("/password"),
					),
				),
			),
			expected: "ThatIsAPasswordFromEnvVar",
			err:      nil,
		},
		{
			desc:   "Testing error resolving the password reader when ReadPasswordResolve is not initialized",
			reader: nil,
			err:    errors.New("The component to resolve read password mechanism has not been initialized."),
		},
		{
			desc:   "Testing error resolve the password reader when no reader is found",
			reader: NewReadPasswordResolve(),
			err:    errors.New("The component to resolve read password does not found a password."),
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			password, err := test.reader.Read()
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, test.expected, password)
			}
		})
	}
}
