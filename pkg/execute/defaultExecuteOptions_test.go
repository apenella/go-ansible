package execute

// import (
// 	"os"
// 	"testing"

// 	"github.com/apenella/go-ansible/pkg/execute/executable/os/exec"
// 	defaultresults "github.com/apenella/go-ansible/pkg/execute/result/default"
// 	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
// 	"github.com/apenella/go-ansible/pkg/playbook"
// 	"github.com/stretchr/testify/assert"
// )

// // TestWithCmd tests the function WithCmd
// func TestWithCmd(t *testing.T) {
// 	cmd := &playbook.AnsiblePlaybookCmd{}

// 	execute := NewDefaultExecute(
// 		WithCmd(cmd),
// 	)

// 	assert.Equal(t, execute.Cmd, cmd)
// }

// // TestWithExecutable tests the function WithExecutable
// func TestWithExecutable(t *testing.T) {
// 	e := exec.NewExec()

// 	execute := NewDefaultExecute(
// 		WithExecutable(e),
// 	)

// 	assert.Equal(t, execute.Exec, e)
// }

// // TestWithWrite tests the function WithWrite
// func TestWithWrite(t *testing.T) {
// 	write := os.Stdout

// 	execute := NewDefaultExecute(
// 		WithWrite(write),
// 	)

// 	assert.Equal(t, execute.Write, write)
// }

// // TestWithWriteError tests the function WithWriteError
// func TestWithWriteError(t *testing.T) {
// 	write := os.Stderr

// 	execute := NewDefaultExecute(
// 		WithWriteError(write),
// 	)

// 	assert.Equal(t, execute.WriterError, write)
// }

// // TestWithCmdRunDir tests the function WithCmdRunDir
// func TestWithCmdRunDir(t *testing.T) {
// 	cmdRunDir := "/tmp"

// 	execute := NewDefaultExecute(
// 		WithCmdRunDir(cmdRunDir),
// 	)

// 	assert.Equal(t, execute.CmdRunDir, cmdRunDir)
// }

// // TestWithTransformers tests the function WithTransformers
// func TestWithTransformers(t *testing.T) {
// 	trans := []transformer.TransformerFunc{
// 		transformer.Prepend("prepend"),
// 		transformer.Append("append"),
// 	}

// 	execute := NewDefaultExecute(
// 		WithTransformers(trans...),
// 	)

// 	assert.Equal(t, execute.Transformers, trans)
// }

// // TestWithEnvVars tests the func WithEnvVars
// func TestWithEnvVars(t *testing.T) {
// 	envvars := EnvVars{
// 		"var1": "value1",
// 	}

// 	execute := NewDefaultExecute(
// 		WithEnvVars(envvars),
// 	)

// 	assert.Equal(t, execute.EnvVars, envvars)
// }

// // TestWithOutput tests the function WithOutput
// func TestWithOutput(t *testing.T) {
// 	output := defaultresults.NewDefaultResults()

// 	execute := NewDefaultExecute(
// 		WithOutput(output),
// 	)

// 	assert.Equal(t, execute.Output, output)
// }
