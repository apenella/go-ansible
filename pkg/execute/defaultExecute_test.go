package execute

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/apenella/go-ansible/v2/mocks"
	"github.com/apenella/go-ansible/v2/pkg/execute/exec"
	defaultresults "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
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

	errContext := "(execute::DefaultExecute::Execute)"

	tests := []struct {
		desc              string
		err               error
		execute           *DefaultExecute
		exec              *exec.MockCmd
		prepareAssertFunc func(*exec.MockExec, *exec.MockCmd)
		assertFunc        func(*exec.MockExec, *exec.MockCmd)
	}{
		{
			desc: "Testing error executing a command when Command is not defiend",
			err:  errors.New(errContext, "Command is not defined"),
			execute: NewDefaultExecute(
				WithWrite(io.Writer(&stdout)),
				WithWriteError(io.Writer(&stderr)),
			),
		},
		{
			desc: "Testing execute a command",
			err:  &errors.Error{},
			exec: exec.NewMockCmd(),
			execute: NewDefaultExecute(
				WithExecutable(exec.NewMockExec()),
				WithWrite(io.Writer(&stdout)),
				WithWriteError(io.Writer(&stderr)),
				WithCmd(
					mocks.NewMockAnsibleCmd([]string{"ansible-playbook", "--connection", "local", "../../test/test_site.yml"}, nil),
				),
			),
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

				e.On("CommandContext", context.TODO(), "ansible-playbook", []string{"--connection", "local", "../../test/test_site.yml"}).Return(cmd)
			},
			assertFunc: func(e *exec.MockExec, cmd *exec.MockCmd) {
				cmd.AssertExpectations(t)
				e.AssertExpectations(t)
			},
		},
		{
			desc: "Testing execute a command in a quiet mode",
			err:  &errors.Error{},
			exec: exec.NewMockCmd(),
			execute: &DefaultExecute{
				Exec:        exec.NewMockExec(),
				Cmd:         mocks.NewMockAnsibleCmd([]string{"ansible-playbook", "--connection", "local", "../../test/test_site.yml"}, nil),
				Write:       io.Writer(&stdout),
				WriterError: io.Writer(&stderr),
				quiet:       true,
			},

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

				e.On("CommandContext", context.TODO(), "ansible-playbook", []string{"--connection", "local", "../../test/test_site.yml"}).Return(cmd)
			},
			assertFunc: func(e *exec.MockExec, cmd *exec.MockCmd) {
				cmd.AssertExpectations(t)
				e.AssertExpectations(t)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			stdout.Reset()
			stderr.Reset()

			if test.prepareAssertFunc != nil {
				test.prepareAssertFunc(test.execute.Exec.(*exec.MockExec), test.exec)
			}

			err := test.execute.Execute(context.TODO())
			if err != nil {
				assert.Equal(t, test.err, err)
			}

			if test.assertFunc != nil {
				test.assertFunc(test.execute.Exec.(*exec.MockExec), test.exec)
			}
		})
	}
}

// func TestExecuteFunctional(t *testing.T) {

// 	var stdout, stderr bytes.Buffer

// 	binary := "ansible-playbook"
// 	_, err := osexec.LookPath(binary)
// 	if err != nil {
// 		t.Skip("")
// 		t.Fatal(err)
// 	}

// 	tests := []struct {
// 		desc           string
// 		err            error
// 		execute        *DefaultExecute
// 		command        Commander
// 		expectedStdout string
// 	}{
// 		{
// 			desc: "Testing an ansible-playbook with local connection",
// 			err:  &errors.Error{},
// 			execute: NewDefaultExecute(
// 				WithWrite(io.Writer(&stdout)),
// 				WithWriteError(io.Writer(&stderr)),
// 				WithEnvVars(
// 					// It forces to use always the same stdout callback
// 					map[string]string{
// 						"ANSIBLE_STDOUT_CALLBACK": "yaml",
// 					},
// 				),
// 				WithCmd(
// 					// playbook.NewAnsiblePlaybookCmdBuilder().
// 					// WithBinary(binary).
// 					// WithPlaybooks([]string{"../../test/test_site.yml"}).
// 					// WithOptions(&playbook.AnsiblePlaybookOptions{
// 					// 	Inventory: ",127.0.0.1",
// 					// }).
// 					// WithConnectionOptions(&options.AnsibleConnectionOptions{
// 					// 	Connection: "local",
// 					// }).Build()),
// 					mocks.NewMockAnsibleCmd(),
// 				),
// 			),

// 			expectedStdout: `
// PLAY [all] *********************************************************************

// TASK [Print test message] ******************************************************
// ok: [127.0.0.1] =>
//   msg: That's a message to test

// PLAY RECAP *********************************************************************
// 127.0.0.1                  : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

// `,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.desc, func(t *testing.T) {
// 			t.Log(test.desc)

// 			stdout.Reset()
// 			stderr.Reset()

// 			err := test.execute.Execute(context.TODO())
// 			if err != nil && assert.Error(t, err) {
// 				assert.Equal(t, test.err, err)
// 			}

// 			assert.Equal(t, test.expectedStdout, stdout.String())
// 		})
// 	}
// }

// TestQuiet tests the function WithQuiet
func TestQuiet(t *testing.T) {
	execute := &DefaultExecute{}
	execute.Quiet()

	assert.Equal(t, execute.quiet, true)
}

func TestQuietCommand(t *testing.T) {
	tests := []struct {
		desc     string
		execute  *DefaultExecute
		expected []string
		err      error
	}{
		{
			desc: "Testing execute a command with verbose flags",
			err:  &errors.Error{},
			execute: &DefaultExecute{
				Exec: exec.NewMockExec(),
				Cmd: mocks.NewMockAnsibleCmd(
					[]string{
						"ansible-playbook",
						"--connection",
						"local",
						"site.yml",
						"-v",
						"-vv",
						"-vvv",
						"-vvvv",
						"--verbose",
					},
					nil),
				// The test executes the quietCommand method so it is not necessary to set the quiet flag
				// quiet: true,
			},

			expected: []string{
				"ansible-playbook",
				"--connection",
				"local",
				"site.yml",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			command, err := test.execute.quietCommand()
			if err != nil {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.expected, command)
			}
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

func TestAddEnvVar(t *testing.T) {
	tests := []struct {
		desc     string
		execute  *DefaultExecute
		key      string
		value    string
		expected string
	}{
		{
			desc:     "Testing add new environment variable to DefaultExecute",
			execute:  &DefaultExecute{},
			key:      "key",
			value:    "value",
			expected: "value",
		},
		{
			desc: "Testing add new environment variable to DefaultExecute",
			execute: &DefaultExecute{
				EnvVars: map[string]string{
					"key": "oldvalue",
				},
			},
			key:      "key",
			value:    "newvalue",
			expected: "newvalue",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			test.execute.AddEnvVar(test.key, test.value)
			assert.Equal(t, test.execute.EnvVars[test.key], test.expected)
		})
	}
}

func TestAddEnvVarSafe(t *testing.T) {

	errContext := "execute::DefaultExecute:AddEnvVarSafe"

	tests := []struct {
		desc     string
		execute  *DefaultExecute
		key      string
		value    string
		expected string
		err      error
	}{
		{
			desc:     "Testing add new environment variable to DefaultExecute",
			execute:  &DefaultExecute{},
			key:      "key",
			value:    "value",
			expected: "value",
			err:      &errors.Error{},
		},
		{
			desc: "Testing add new environment variable to DefaultExecute",
			execute: &DefaultExecute{
				EnvVars: map[string]string{
					"key": "oldvalue",
				},
			},
			key:      "key",
			value:    "newvalue",
			expected: "newvalue",
			err:      errors.New(errContext, "Environment variable 'key' already exists"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.execute.AddEnvVarSafe(test.key, test.value)
			if err != nil {
				assert.Equal(t, err, test.err)
			} else {
				assert.Equal(t, test.execute.EnvVars[test.key], test.expected)
			}
		})
	}
}

// TestWithCmd tests the function WithCmd
func TestWithCmd(t *testing.T) {
	cmd := mocks.NewMockAnsibleCmd([]string{"ansible-playbook", "--connection", "local", "../../test/test_site.yml"}, nil)

	execute := NewDefaultExecute(
		WithCmd(cmd),
	)

	assert.Equal(t, execute.Cmd, cmd)
}

// TestWithExecutable tests the function WithExecutable
func TestWithExecutable(t *testing.T) {
	e := exec.NewOsExec()

	execute := NewDefaultExecute(
		WithExecutable(e),
	)

	assert.Equal(t, execute.Exec, e)
}

// TestWithWrite tests the function WithWrite
func TestWithWrite(t *testing.T) {
	write := os.Stdout

	execute := NewDefaultExecute(
		WithWrite(write),
	)

	assert.Equal(t, execute.Write, write)
}

// TestWithWriteError tests the function WithWriteError
func TestWithWriteError(t *testing.T) {
	write := os.Stderr

	execute := NewDefaultExecute(
		WithWriteError(write),
	)

	assert.Equal(t, execute.WriterError, write)
}

// TestWithCmdRunDir tests the function WithCmdRunDir
func TestWithCmdRunDir(t *testing.T) {
	cmdRunDir := "/tmp"

	execute := NewDefaultExecute(
		WithCmdRunDir(cmdRunDir),
	)

	assert.Equal(t, execute.CmdRunDir, cmdRunDir)
}

// TestWithTransformers tests the function WithTransformers
func TestWithTransformers(t *testing.T) {
	trans := []transformer.TransformerFunc{
		transformer.Prepend("prepend"),
		transformer.Append("append"),
	}

	execute := NewDefaultExecute(
		WithTransformers(trans...),
	)

	assert.Equal(t, execute.Transformers, trans)
}

// TestWithEnvVars tests the func WithEnvVars
func TestWithEnvVars(t *testing.T) {
	envvars := EnvVars{
		"var1": "value1",
	}

	execute := NewDefaultExecute(
		WithEnvVars(envvars),
	)

	assert.Equal(t, execute.EnvVars, envvars)
}

// TestWithOutput tests the function WithOutput
func TestWithOutput(t *testing.T) {
	output := defaultresults.NewDefaultResults()

	execute := NewDefaultExecute(
		WithOutput(output),
	)

	assert.Equal(t, execute.Output, output)
}
