package text

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	tests := []struct {
		desc     string
		reader   *ReadPasswordFromText
		expected string
		err      error
	}{
		{
			desc: "Testing reading a password from text",
			reader: NewReadPasswordFromText(
				WithText("ThatIsAPassword"),
			),
			expected: "ThatIsAPassword",
			err:      nil,
		},
		{
			desc:   "Testing error reading a password from text when text is an empty string",
			reader: NewReadPasswordFromText(),
			err:    errors.New("Text must be specified to use the password input from text."),
		},
		{
			desc:   "Testing error reading a password from text when ReadPasswordFromText has not been initialized",
			reader: nil,
			err:    errors.New("Password input from text has not been initialized."),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Run(test.desc, func(t *testing.T) {
				secret, err := test.reader.Read()
				if err != nil {
					assert.EqualError(t, err, test.err.Error())
				} else {
					assert.Equal(t, test.expected, secret)
				}
			})
		})
	}
}
