package vault

import (
	"regexp"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/vault/encrypt"
	"github.com/apenella/go-ansible/v2/pkg/vault/password/text"
	"github.com/go-errors/errors"
	"github.com/stretchr/testify/assert"
)

func TestVault(t *testing.T) {

	value := "Secret text"

	encryptString := encrypt.NewMockEncryptString()
	encryptString.Mock.On("Encrypt", value).Return("encrypted_value", nil)

	tests := []struct {
		desc     string
		vaulter  *VariableVaulter
		value    string
		expected *VaultVariableValue
		err      error
	}{
		{
			desc: "Testing vaulting a variable",
			vaulter: NewVariableVaulter(
				WithEncrypt(
					encryptString,
				),
			),
			value:    value,
			expected: &VaultVariableValue{Value: "encrypted_value"},
			err:      errors.New(""),
		},
		{
			desc:    "Testing error vaulting a text when the VariableVaulter is not initialized ",
			vaulter: nil,
			err:     errors.New("VariableVaulter must be initialized before vaulting a variable."),
		},
		{
			desc:    "Testing error vaulting a text when the Encrypter is not initialized ",
			vaulter: NewVariableVaulter(),
			err:     errors.New("Encrypter must be provided to encrypt a variable."),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res, err := test.vaulter.Vault(test.value)
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, test.expected, res)
			}
		})
	}
}

func TestVaultIntegration(t *testing.T) {

	var err error
	var vaultedVariable *VaultVariableValue
	var VaultedVariableJSONString string

	// arrange
	encrypter := encrypt.NewEncryptString(
		encrypt.WithReader(
			text.NewReadPasswordFromText(
				text.WithText("s3cr3t"),
			),
		),
	)

	vaulter := NewVariableVaulter(
		WithEncrypt(encrypter),
	)

	text := "That is a plain text"
	expectedRegexp := "{\"__ansible_vault\":\"\\$ANSIBLE_VAULT;1.1;AES256\\\\n[0-9a-zA-Z]{80}\\\\n[0-9a-zA-Z]{80}\\\\n[0-9a-zA-Z]{80}\\\\n[0-9a-zA-Z]{80}\\\\n[0-9a-zA-Z]{68}"

	// act
	vaultedVariable, err = vaulter.Vault(text)
	assert.NoError(t, err)

	VaultedVariableJSONString, err = vaultedVariable.ToJSON()
	assert.NoError(t, err)

	// assert
	assert.Regexp(t, regexp.MustCompile(expectedRegexp), VaultedVariableJSONString)
}
