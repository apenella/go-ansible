package execute

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestDefaultExecute(t *testing.T) {

	var stdout, stderr bytes.Buffer

	binary, err := exec.LookPath("ansible-playbook")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		desc           string
		err            error
		execute        *DefaultExecute
		command        []string
		options        []ExecuteOptions
		res            string
		stdout         io.Writer
		stderr         io.Writer
		expectedStderr string
		expectedStdout string
		ctx            context.Context
	}{
		{
			desc: "Testing an ansible-playbook with local connection",
			err:  nil,
			execute: &DefaultExecute{
				Write:       io.Writer(&stdout),
				WriterError: io.Writer(&stderr),
			},
			ctx:     context.TODO(),
			command: []string{binary, "--inventory", "127.0.0.1,", "test/site.yml", "-c", "local"},
			options: []ExecuteOptions{
				WithPrefix("test"),
			},
			expectedStdout: ``,
		},
		{
			desc: "Testing an ansible-playbook forcing an invalid charaters warning message",
			err:  errors.New("(DefaultExecute::Execute)", "Error during command execution: ansible-playbook error: parser error\n\nCommand executed: "+binary+" --inventory test/all test/site.yml --user apenella\n\nexit status 4"),
			execute: &DefaultExecute{
				Write:       io.Writer(&stdout),
				WriterError: io.Writer(&stderr),
			},
			ctx:     context.TODO(),
			command: []string{binary, "--inventory", "test/all", "test/site.yml", "--user", "apenella"},
			options: []ExecuteOptions{
				WithPrefix("test"),
			},
			expectedStderr: `test ── [WARNING]: Invalid characters were found in group names but not replaced, use
test ── -vvvv to see details
`,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		stdout.Reset()
		stderr.Reset()

		err := test.execute.Execute(test.ctx, test.command, test.options...)

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		}
		assert.Equal(t, test.expectedStderr, stderr.String())
	}

}
