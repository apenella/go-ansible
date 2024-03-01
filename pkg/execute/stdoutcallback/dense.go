package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// DenseStdoutCallback dense stdout output
	DenseStdoutCallback = "dense"
)

type DenseStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewDenseStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *DenseStdoutCallbackExecute {
	return &DenseStdoutCallbackExecute{executor: executor}
}

func (e *DenseStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *DenseStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DenseStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("DenseStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(DenseStdoutCallback),
	).Execute(ctx)
}
