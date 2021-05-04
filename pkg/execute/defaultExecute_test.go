package execute

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"testing"

	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultExecute(t *testing.T) {
	wr := &bytes.Buffer{}
	runDir := "rundir"

	t.Log("Testing NewDefaultExecute and WithXXX methods")

	trans := func() results.TransformerFunc {
		return func(message string) string {
			return message
		}
	}

	exe := NewDefaultExecute(
		WithCmdRunDir(runDir),
		WithWrite(io.Writer(wr)),
		WithWriteError(io.Writer(wr)),
		WithShowDuration(),
		WithTransformers(trans()),
	)

	assert.Equal(t, runDir, exe.CmdRunDir, "CmdRunDir does not match")
	assert.True(t, exe.ShowDuration, "ShowDuration does not match")
	assert.Equal(t, wr, exe.Write, "Write does not match")
	assert.Equal(t, wr, exe.WriterError, "WriteError does not match")
}

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
			execute: NewDefaultExecute(
				WithWrite(io.Writer(&stdout)),
				WithWriteError(io.Writer(&stderr)),
			),
			ctx:            context.TODO(),
			command:        []string{binary, "--inventory", "127.0.0.1,", "test/site.yml", "-c", "local"},
			expectedStdout: ``,
		},
		{
			desc: "Testing an ansible-playbook forcing an invalid charaters warning message",
			err:  errors.New("(DefaultExecute::Execute)", "Error during command execution: ansible-playbook error: parser error\n\nCommand executed: "+binary+" --inventory test/all test/site.yml --user apenella\n\nexit status 4"),
			execute: NewDefaultExecute(
				WithWrite(io.Writer(&stdout)),
				WithWriteError(io.Writer(&stderr)),
			),
			ctx:     context.TODO(),
			command: []string{binary, "--inventory", "test/all", "test/site.yml", "--user", "apenella"},
			expectedStderr: `[WARNING]: Invalid characters were found in group names but not replaced, use
-vvvv to see details
`,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		stdout.Reset()
		stderr.Reset()

		err := test.execute.Execute(test.ctx, test.command, nil, test.options...)

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		}
		assert.Equal(t, test.expectedStderr, stderr.String())
	}

}
