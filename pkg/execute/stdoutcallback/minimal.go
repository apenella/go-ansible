package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// MinimalStdoutCallback minmal ansible screen output
	MinimalStdoutCallback = "minimal"
)

type MinimalStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewMinimalStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *MinimalStdoutCallbackExecute {
	return &MinimalStdoutCallbackExecute{executor: executor}
}

func (e *MinimalStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *MinimalStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *MinimalStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("MinimalStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(MinimalStdoutCallback),
	).Execute(ctx)
}
