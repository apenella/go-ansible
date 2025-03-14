package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// StderrStdoutCallback splits output, sending failed tasks to stderr
	StderrStdoutCallback = "stderr"
)

// StderrStdoutCallbackExecute defines an executor to run an ansible command with a stderr stdout callback
type StderrStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

// NewStderrStdoutCallbackExecute creates a StderrStdoutCallbackExecute
func NewStderrStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *StderrStdoutCallbackExecute {
	return &StderrStdoutCallbackExecute{executor: executor}
}

// WithExecutor sets the executor for the StderrStdoutCallbackExecute
func (e *StderrStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *StderrStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *StderrStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("StderrStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(StderrStdoutCallback),
	).Execute(ctx)
}
