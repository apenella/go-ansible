package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// DebugStdoutCallback formatted stdout/stderr output
	DebugStdoutCallback = "debug"
)

type DebugStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewDebugStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *DebugStdoutCallbackExecute {
	return &DebugStdoutCallbackExecute{executor: executor}
}

func (e *DebugStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *DebugStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DebugStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("DebugStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(DebugStdoutCallback),
	).Execute(ctx)
}
