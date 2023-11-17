package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/pkg/execute/result/default"
)

const (
	// StderrStdoutCallback splits output, sending failed tasks to stderr
	StderrStdoutCallback = "stderr"
)

type StderrStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewStderrStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *StderrStdoutCallbackExecute {
	return &StderrStdoutCallbackExecute{executor: executor}
}

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

	return configuration.NewExecutorWithAnsibleConfigurationSettings(e.executor).
		WithAnsibleStdoutCallback(StderrStdoutCallback).Execute(ctx)
}
