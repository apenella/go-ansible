package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	defaultresult "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
)

const (
	// YAMLStdoutCallback yamlized ansible screen output
	YAMLStdoutCallback = "yaml"
)

type YAMLStdoutCallbackExecute struct {
	executor ExecutorStdoutCallbackSetter
}

func NewYAMLStdoutCallbackExecute(executor ExecutorStdoutCallbackSetter) *YAMLStdoutCallbackExecute {
	return &YAMLStdoutCallbackExecute{executor: executor}
}

func (e *YAMLStdoutCallbackExecute) WithExecutor(exec ExecutorStdoutCallbackSetter) *YAMLStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *YAMLStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("YAMLStdoutCallbackExecute executor requires an executor")
	}

	e.executor.WithOutput(defaultresult.NewDefaultResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(YAMLStdoutCallback),
	).Execute(ctx)
}
