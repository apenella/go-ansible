package execute

import (
	"os"
	"testing"

	"github.com/apenella/go-ansible/internal/executable/os/exec"
	defaultresults "github.com/apenella/go-ansible/pkg/execute/result/default"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/stretchr/testify/assert"
)

// TestOptionsWithCmd tests the function WithCmd
func TestOptionsWithCmd(t *testing.T) {
	cmd := &playbook.AnsiblePlaybookCmd{}

	execute := NewDefaultExecute(
		WithCmd(cmd),
	)

	assert.Equal(t, execute.Cmd, cmd)
}

// TestOptionsWithExecutable tests the function WithExecutable
func TestOptionsWithExecutable(t *testing.T) {
	e := exec.NewExec()

	execute := NewDefaultExecute(
		WithExecutable(e),
	)

	assert.Equal(t, execute.Exec, e)
}

// TestOptionsWithWrite tests the function WithWrite
func TestOptionsWithWrite(t *testing.T) {
	write := os.Stdout

	execute := NewDefaultExecute(
		WithWrite(write),
	)

	assert.Equal(t, execute.Write, write)
}

// TestOptionsWithWriteError tests the function WithWriteError
func TestOptionsWithWriteError(t *testing.T) {
	write := os.Stderr

	execute := NewDefaultExecute(
		WithWriteError(write),
	)

	assert.Equal(t, execute.WriterError, write)
}

// TestOptionsWithCmdRunDir tests the function WithCmdRunDir
func TestOptionsWithCmdRunDir(t *testing.T) {
	cmdRunDir := "/tmp"

	execute := NewDefaultExecute(
		WithCmdRunDir(cmdRunDir),
	)

	assert.Equal(t, execute.CmdRunDir, cmdRunDir)
}

// TestOptionsWithTransformers tests the function WithTransformers
func TestOptionsWithTransformers(t *testing.T) {
	trans := []transformer.TransformerFunc{
		transformer.Prepend("prepend"),
		transformer.Append("append"),
	}

	execute := NewDefaultExecute(
		WithTransformers(trans...),
	)

	assert.Equal(t, execute.Transformers, trans)
}

// TestOptionsWithEnvVars tests the func WithEnvVars
func TestOptionsWithEnvVars(t *testing.T) {
	envvars := EnvVars{
		"var1": "value1",
	}

	execute := NewDefaultExecute(
		WithEnvVars(envvars),
	)

	assert.Equal(t, execute.EnvVars, envvars)
}

// TestOptionsWithOutput tests the function WithOutput
func TestOptionsWithOutput(t *testing.T) {
	output := defaultresults.NewDefaultResults()

	execute := NewDefaultExecute(
		WithOutput(output),
	)

	assert.Equal(t, execute.Output, output)
}
