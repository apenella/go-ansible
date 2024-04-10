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

// ExecutorQuietStdoutCallbackSetter extends the ExecutorStdoutCallbackSetter interface by adding a method to force the non-verbose mode in the Stdout Callback configuration
type ExecutorQuietStdoutCallbackSetter interface {
	ExecutorStdoutCallbackSetter
	Quiet()
}
