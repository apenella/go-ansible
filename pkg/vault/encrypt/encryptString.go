package encrypt

import (
	"github.com/pkg/errors"
	vault "github.com/sosedoff/ansible-vault-go"
)

// OptionsFunc is a function used to configure EncryptString
type OptionsFunc func(*EncryptString)

type EncryptString struct {
	reader PasswordReader
}

func NewEncryptString(options ...OptionsFunc) *EncryptString {
	EncryptString := &EncryptString{}
	EncryptString.Options(options...)

	return EncryptString
}

func WithReader(reader PasswordReader) OptionsFunc {
	return func(c *EncryptString) {
		c.reader = reader
	}
}

// Options configure the ReadSecretFromEnvVar
func (c *EncryptString) Options(opts ...OptionsFunc) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *EncryptString) Encrypt(plainText string) (string, error) {
	var pass, encryptedText string
	var err error

	pass, err = c.reader.Read()
	if err != nil {
		return "", errors.Wrap(err, "Error reading the password")
	}

	encryptedText, err = vault.Encrypt(plainText, pass)
	if err != nil {
		return "", errors.Wrap(err, "Error encrypting the password")
	}

	return encryptedText, nil
}
