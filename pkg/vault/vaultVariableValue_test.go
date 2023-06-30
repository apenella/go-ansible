package vault

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJSON(t *testing.T) {
	tests := []struct {
		desc     string
		variable *VaultVariableValue
		expected string
		err      error
	}{
		{
			desc:     "Testing converting vaulted variable value to JSON",
			variable: NewVaultVariableValue("encrypted_variable_value"),
			expected: "{\"__ansible_vault\":\"encrypted_variable_value\"}",
			err:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res, err := test.variable.ToJSON()
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, test.expected, res)
			}
		})
	}
}
