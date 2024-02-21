package adhoc

import (
	"context"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
)

// AnsibleAdhocExecute is an executor for ansible command that runs the command using a DefaultExecute with default options
type AnsibleAdhocExecute struct {
	cmd *AnsibleAdhocCmd
}

// NewAnsibleAdhocExecute returns a new AnsibleAdhocExecute. It receives host pattern to be executed
func NewAnsibleAdhocExecute(pattern string) *AnsibleAdhocExecute {

	ansiblePlaybookCmd := &AnsibleAdhocCmd{
		Pattern: pattern,
	}

	exec := &AnsibleAdhocExecute{
		cmd: ansiblePlaybookCmd,
	}

	return exec
}

// WithBinary return a AnsibleAdhocExecute with the binary file set
func (e *AnsibleAdhocExecute) WithBinary(binary string) *AnsibleAdhocExecute {
	e.cmd.Binary = binary

	return e
}

// WithAdhocOptions returns an AnsibleAdhocExecute with the ansible's playbook options set
func (e *AnsibleAdhocExecute) WithAdhocOptions(options *AnsibleAdhocOptions) *AnsibleAdhocExecute {
	e.cmd.AdhocOptions = options

	return e
}

// WithConnectionOptions returns an AnsibleAdhocExecute with the ansible's playbook specific options for connection set
func (e *AnsibleAdhocExecute) WithConnectionOptions(options *options.AnsibleConnectionOptions) *AnsibleAdhocExecute {
	e.cmd.ConnectionOptions = options

	return e
}

// WithPrivilegeEscalationOptions returns an AnsibleAdhocExecute with the ansible's playbook privilege escalation options set
func (e *AnsibleAdhocExecute) WithPrivilegeEscalationOptions(options *options.AnsiblePrivilegeEscalationOptions) *AnsibleAdhocExecute {
	e.cmd.PrivilegeEscalationOptions = options

	return e
}

// Execute method runs the ansible command using a DefaultExecute with default options
func (e *AnsibleAdhocExecute) Execute(ctx context.Context) error {

	exec := execute.NewDefaultExecute(
		execute.WithCmd(e.cmd),
	)

	err := exec.Execute(ctx)
	if err != nil {
		return err
	}

	return nil
}
