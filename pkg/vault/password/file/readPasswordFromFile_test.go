package file

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	testFs := afero.NewMemMapFs()

	err := afero.WriteFile(testFs, "/password", []byte("ThatIsAPassword"), 0666)
	if err != nil {
		t.Log(err)
	}

	tests := []struct {
		desc     string
		reader   *ReadPasswordFromFile
		expected string
		err      error
	}{
		{
			desc: "Testing reading a password from a file",
			reader: NewReadPasswordFromFile(
				WithFs(testFs),
				WithFile("/password"),
			),
			expected: "ThatIsAPassword",
			err:      nil,
		},
		{
			desc:   "Testing error reading a password from file when ReadPasswordFromFile is not initialized",
			reader: nil,
			err:    errors.New("Read password from a file component has not been initialized."),
		},
		{
			desc: "Testing error reading a password from file when ReadPasswordFromFile file is not specified",
			reader: NewReadPasswordFromFile(
				WithFs(testFs),
				WithFile(""),
			),
			err: errors.New("File path must be specified to read the password from a file."),
		},
		{
			desc: "Testing error reading a password from file when ReadPasswordFromFile file does not exists",
			reader: NewReadPasswordFromFile(
				WithFs(testFs),
				WithFile("/unexisting"),
			),
			err: errors.Wrap(errors.New("open /unexisting: file does not exist"), "Error describing the file '/unexisting'."),
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
