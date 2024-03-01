package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// TimerStdoutCallback adds time to play stats
	TimerStdoutCallback = "timer"
)

type TimerStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewTimerStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *TimerStdoutCallbackExecute {
	return &TimerStdoutCallbackExecute{executor: executor}
}

func (e *TimerStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *TimerStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *TimerStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("TimerStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(TimerStdoutCallback),
	).Execute(ctx)
}
