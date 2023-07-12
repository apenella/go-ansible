package vault

import (
	"github.com/pkg/errors"
)

// OptionsFunc is a function used to configure ReadPasswordFromEnvVar
type OptionsFunc func(*VariableVaulter)

type VariableVaulter struct {
	encrypt Encrypter
}

func NewVariableVaulter(options ...OptionsFunc) *VariableVaulter {
	vault := &VariableVaulter{}
	vault.Options(options...)

	return vault
}

func WithEncrypt(e Encrypter) OptionsFunc {
	return func(v *VariableVaulter) {
		v.encrypt = e
	}
}

func (v *VariableVaulter) Options(opts ...OptionsFunc) {
	for _, opt := range opts {
		opt(v)
	}
}

func (v *VariableVaulter) Vault(value string) (*VaultVariableValue, error) {
	var err error
	var encryptedValue string

	if v == nil {
		return nil, errors.New("VariableVaulter must be initialized before vaulting a variable.")
	}

	if v.encrypt == nil {
		return nil, errors.New("Encrypter must be provided to encrypt a variable.")
	}

	encryptedValue, err = v.encrypt.Encrypt(value)
	if err != nil {
		return nil, errors.Wrap(err, "Error encrypting variable value.")
	}

	VariableVaulterValue := NewVaultVariableValue(encryptedValue)
	// encryptedValueJSON, err = VariableVaulterValue.ToJSON()
	// if err != nil {
	// 	return "", errors.Wrap(err, "Vault variable could not be converted to JSON.")
	// }

	return VariableVaulterValue, nil
}
