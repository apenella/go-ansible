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

type DefaultStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewDefaultStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *DefaultStdoutCallbackExecute {
	return &DefaultStdoutCallbackExecute{executor: executor}
}

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
