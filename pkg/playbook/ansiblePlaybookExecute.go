package playbook

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
)

// AnsiblePlaybookExecute is an executor for ansible-playbook command that runs the command using a DefaultExecute with default options
type AnsiblePlaybookExecute struct {
	cmd *AnsiblePlaybookCmd
}

// NewAnsiblePlaybookExecute returns a new AnsiblePlaybookExecute. It receives a list of playbooks to be executed
func NewAnsiblePlaybookExecute(playbook ...string) *AnsiblePlaybookExecute {

	ansiblePlaybookCmd := &AnsiblePlaybookCmd{
		Playbooks: append([]string{}, playbook...),
	}

	exec := &AnsiblePlaybookExecute{
		cmd: ansiblePlaybookCmd,
	}

	return exec
}

// WithAnsible return a AnsiblePlaybookExecute withthe binary file set
func (e *AnsiblePlaybookExecute) WithBinary(binary string) *AnsiblePlaybookExecute {
	e.cmd.Binary = binary

	return e
}

// WithPlaybookOptions returns an AnsiblePlaybookExecute with the ansible's playbook options set
func (e *AnsiblePlaybookExecute) WithPlaybookOptions(options *AnsiblePlaybookOptions) *AnsiblePlaybookExecute {
	e.cmd.PlaybookOptions = options

	return e
}

// Execute method runs the ansible-playbook command using a DefaultExecute with default options
func (e *AnsiblePlaybookExecute) Execute(ctx context.Context) error {

	exec := execute.NewDefaultExecute(
		execute.WithCmd(e.cmd),
		execute.WithErrorEnrich(NewAnsiblePlaybookErrorEnrich()),
	)

	err := exec.Execute(ctx)
	if err != nil {
		return err
	}

	return nil
}
