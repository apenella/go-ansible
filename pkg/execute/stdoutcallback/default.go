package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// DefaultStdoutCallback default ansible screen output
	DefaultStdoutCallback = "default"
)

// DefaultStdoutCallbackExecute defines an executor to run an ansible command with a default stdout callback
type DefaultStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

// NewDefaultStdoutCallbackExecute creates a DefaultStdoutCallbackExecute
func NewDefaultStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *DefaultStdoutCallbackExecute {
	return &DefaultStdoutCallbackExecute{executor: executor}
}

// WithExecutor sets the executor for the DefaultStdoutCallbackExecute
func (e *DefaultStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *DefaultStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DefaultStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("DefaultStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(DefaultStdoutCallback),
	).Execute(ctx)

}
