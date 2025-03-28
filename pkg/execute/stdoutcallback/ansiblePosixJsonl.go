package stdoutcallback

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	jsonresults "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
)

const (
	// AnsiblePosixJsonlStdoutCallback ansible posix jsonl stdout output
	AnsiblePosixJsonlStdoutCallback = "ansible.posix.jsonl"
)

// AnsiblePosixJsonlStdoutCallbackExecute defines an executor to run an ansible command with a ansible posix jsonl stdout callback
type AnsiblePosixJsonlStdoutCallbackExecute struct {
	executor ExecutorQuietStdoutCallbackSetter
}

// NewAnsiblePosixJsonlStdoutCallbackExecute creates a AnsiblePosixJsonlStdoutCallbackExecute
func NewAnsiblePosixJsonlStdoutCallbackExecute(executor ExecutorQuietStdoutCallbackSetter) *AnsiblePosixJsonlStdoutCallbackExecute {
	return &AnsiblePosixJsonlStdoutCallbackExecute{executor: executor}
}

// WithExecutor sets the executor for the AnsiblePosixJsonlStdoutCallbackExecute
func (e *AnsiblePosixJsonlStdoutCallbackExecute) WithExecutor(exec ExecutorQuietStdoutCallbackSetter) *AnsiblePosixJsonlStdoutCallbackExecute {
	e.executor = exec
	return e
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *AnsiblePosixJsonlStdoutCallbackExecute) Execute(ctx context.Context) error {

	if e.executor == nil {
		return fmt.Errorf("AnsiblePosixJsonlStdoutCallbackExecute executor requires an executor")
	}

	e.executor.Quiet()
	e.executor.WithOutput(jsonresults.NewAnsiblePosixJSONLResults())

	return configuration.NewAnsibleWithConfigurationSettingsExecute(e.executor,
		configuration.WithAnsibleStdoutCallback(AnsiblePosixJsonlStdoutCallback),
	).Execute(ctx)
}
