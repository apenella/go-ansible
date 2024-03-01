package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// OnelineStdoutCallback oneline ansible screen output
	OnelineStdoutCallback = "oneline"
)

type OnelineStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewOnelineStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *OnelineStdoutCallbackExecute {
	return &OnelineStdoutCallbackExecute{executor: executor}
}

func (e *OnelineStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *OnelineStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *OnelineStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("OnelineStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(OnelineStdoutCallback),
	).Execute(ctx)
}
