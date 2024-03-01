package stdoutcallback

import (
	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/result"
)

// ExecutorStdoutCallbackSetter extends the executor interface by adding methods to configure the Stdout Callback configuration
type ExecutorStdoutCallbackSetter interface {
	execute.Executor
	AddEnvVar(key, value string)
	WithOutput(output result.ResultsOutputer)
}
