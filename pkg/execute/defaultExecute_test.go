package execute

import (
	"bytes"
	"context"
	"io"
	osexec "os/exec"
	"testing"

	"github.com/apenella/go-ansible/pkg/execute/executable/os/exec"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultExecute(t *testing.T) {
	wr := &bytes.Buffer{}
	runDir := "rundir"

	t.Log("Testing NewDefaultExecute and WithXXX methods")

	trans := func() transformer.TransformerFunc {
		return func(message string) string {
			return message
		}
	}

	exe := NewDefaultExecute(
		WithCmdRunDir(runDir),
		WithWrite(io.Writer(wr)),
		WithWriteError(io.Writer(wr)),
		WithTransformers(trans()),
	)

	assert.Equal(t, runDir, exe.CmdRunDir, "CmdRunDir does not match")
	assert.Equal(t, wr, exe.Write, "Write does not match")
	assert.Equal(t, wr, exe.WriterError, "WriteError does not match")
}

func TestExecute(t *testing.T) {
	var stdout, stderr bytes.Buffer
	var cmdRead bytes.Buffer

	tests := []struct {
		desc              string
		err               error
		execute           *DefaultExecute
		cmd               *exec.MockCmd
		command           []string
		options           []ExecuteOptions
		prepareAssertFunc func(e *exec.MockExec, cmd *exec.MockCmd)
		assertFunc        func(e *exec.MockExec, cmd *exec.MockCmd)
	}{
		{
			desc: "Testing execute a command",
			err:  &errors.Error{},
			cmd:  exec.NewMockCmd(),
			execute: NewDefaultExecute(
				WithExecutable(exec.NewMockExec()),
				WithWrite(io.Writer(&stdout)),
				WithWriteError(io.Writer(&stderr)),
			),
			command: []string{"command", "-flag"},
			options: []ExecuteOptions{},
			prepareAssertFunc: func(e *exec.MockExec, cmd *exec.MockCmd) {
				if e == nil {
					t.Fatal("prepareAssertFunc requires a *exec.MockExec")
				}

				if cmd == nil {
					t.Fatal("prepareAssertFunc requires a *exec.MockCmd")
				}

				cmd.On("StdoutPipe").Return(io.NopCloser(io.Reader(&cmdRead)), nil)
				cmd.On("StderrPipe").Return(io.NopCloser(io.Reader(&cmdRead)), nil)
				cmd.On("Start").Return(nil)
				cmd.On("Wait").Return(nil)

				e.On("CommandContext", context.TODO(), "command", []string{"-flag"}).Return(cmd)
			},
			assertFunc: func(e *exec.MockExec, cmd *exec.MockCmd) {
				cmd.AssertExpectations(t)
				e.AssertExpectations(t)
			},
		},

		// {
		// 	desc: "Testing error when the command fails with AnsiblePlaybookErrorCodeGeneralError",
		// 	err:  &errors.Error{},
		// 	cmd:  exec.NewMockCmd(),
		// 	execute: NewDefaultExecute(
		// 		WithExecutable(exec.NewMockExec()),
		// 		WithWrite(io.Writer(&stdout)),
		// 		WithWriteError(io.Writer(&stderr)),
		// 	),
		// 	command: []string{"command", "-flag"},
		// 	options: []ExecuteOptions{},
		// 	prepareAssertFunc: func(e *exec.MockExec, cmd *exec.MockCmd) {
		// 		if e == nil {
		// 			t.Fatal("prepareAssertFunc requires a *exec.MockExec")
		// 		}

		// 		if cmd == nil {
		// 			t.Fatal("prepareAssertFunc requires a *exec.MockCmd")
		// 		}

		// 		cmd.On("StdoutPipe").Return(io.NopCloser(io.Reader(&cmdRead)), nil)
		// 		cmd.On("StderrPipe").Return(io.NopCloser(io.Reader(&cmdRead)), nil)
		// 		cmd.On("Start").Return(nil)
		// 		cmd.On("String").Return("That is an error")
		// 		cmd.On("Wait").Return(nil)

		// 		e.On("CommandContext", context.TODO(), "command", []string{"-flag"}).Return(cmd)
		// 	},
		// 	assertFunc: func(e *exec.MockExec, cmd *exec.MockCmd) {
		// 		cmd.AssertExpectations(t)
		// 		e.AssertExpectations(t)
		// 	},
		// },
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			stdout.Reset()
			stderr.Reset()

			if test.prepareAssertFunc != nil {
				test.prepareAssertFunc(test.execute.Exec.(*exec.MockExec), test.cmd)
			}

			err := test.execute.Execute(context.TODO(), test.command, test.options...)
			if err != nil {
				assert.Equal(t, test.err, err)
			}

			if test.assertFunc != nil {
				test.assertFunc(test.execute.Exec.(*exec.MockExec), test.cmd)
			}
		})
	}
}

func TestExecuteFunctional(t *testing.T) {

	var stdout, stderr bytes.Buffer

	binary := "ansible-playbook"
	_, err := osexec.LookPath(binary)
	if err != nil {
		t.Skip("")
		t.Fatal(err)
	}

	tests := []struct {
		desc           string
		err            error
		execute        *DefaultExecute
		command        []string
		options        []ExecuteOptions
		expectedStdout string
	}{
		{
			desc: "Testing an ansible-playbook with local connection",
			err:  &errors.Error{},
			execute: NewDefaultExecute(
				WithWrite(io.Writer(&stdout)),
				WithWriteError(io.Writer(&stderr)),
			),
			command: []string{binary, "-i", "127.0.0.1,", "../../test/test_site.yml", "-c", "local"},
			expectedStdout: `
PLAY [all] *********************************************************************

TASK [Print test message] ******************************************************
ok: [127.0.0.1] => 
  msg: That's a message to test

PLAY RECAP *********************************************************************
127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

`,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			stdout.Reset()
			stderr.Reset()

			err := test.execute.Execute(context.TODO(), test.command, test.options...)
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			}

			assert.Equal(t, test.expectedStdout, stdout.String())
		})
	}
}

func TestEnviron(t *testing.T) {
	tests := []struct {
		desc           string
		envvars        EnvVars
		expectedResult []string
	}{
		{
			desc: "Testing basic test case",
			envvars: EnvVars{
				"KEY": "VALUE",
			},
			expectedResult: []string{"KEY=VALUE"},
		},
		{
			desc:           "Testing an empty env",
			envvars:        EnvVars{},
			expectedResult: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			assert.Equal(t, test.expectedResult, test.envvars.Environ())
		})
	}
}
