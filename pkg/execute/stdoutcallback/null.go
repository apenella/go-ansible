package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// NullStdoutCallback don't display stuff to screen
	NullStdoutCallback = "null"
)

// NullStdoutCallbackExecute defines an executor to run an ansible command with a null stdout callback
type NullStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

// NewNullStdoutCallbackExecute creates a NullStdoutCallbackExecute
func NewNullStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *NullStdoutCallbackExecute {
	return &NullStdoutCallbackExecute{executor: executor}
}

// WithExecutor sets the executor for the NullStdoutCallbackExecute
func (e *NullStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *NullStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *NullStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("NullStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(NullStdoutCallback),
	).Execute(ctx)
}
