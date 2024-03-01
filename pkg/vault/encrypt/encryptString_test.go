package encrypt

import (
	"regexp"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/vault/password/text"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		desc           string
		text           string
		encrypt        *EncryptString
		expectedRegexp string
		expectedLen    uint
		err            error
	}{
		{
			desc: "Testing encrypting a message",
			text: "ThatIsASecretMessage",
			encrypt: NewEncryptString(
				WithReader(
					text.NewReadPasswordFromText(
						text.WithText("secret"),
					),
				),
			),
			expectedRegexp: "\\$ANSIBLE_VAULT;1.1;AES256",
			expectedLen:    418,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res, err := test.encrypt.Encrypt(test.text)
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Regexp(t, regexp.MustCompile(test.expectedRegexp), res)
				assert.Len(t, res, int(test.expectedLen))
			}
		})
	}
}
