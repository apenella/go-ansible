package stdoutcallback

import (
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/result"
)

// ExecutorStdoutCallbackSetter extends the executor interface by adding methods to configure the Stdout Callback configuration
type ExecutorStdoutCallbackSetter interface {
	execute.Executor
	AddEnvVar(key, value string)
	WithOutput(output result.ResultsOutputer)
}
