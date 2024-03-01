package inventory

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
)

// AnsibleInventoryExecute is an executor for ansible-inventory command that runs the command using a DefaultExecute with default options
type AnsibleInventoryExecute struct {
	cmd *AnsibleInventoryCmd
}

// NewAnsibleInventoryExecute returns a new AnsibleInventoryExecute. It receives a pattern to be used to filter the inventory
func NewAnsibleInventoryExecute() *AnsibleInventoryExecute {

	ansibleInventoryCmd := &AnsibleInventoryCmd{}

	exec := &AnsibleInventoryExecute{
		cmd: ansibleInventoryCmd,
	}

	return exec
}

// WithBinary return a AnsibleInventoryExecute withthe binary file set
func (e *AnsibleInventoryExecute) WithBinary(binary string) *AnsibleInventoryExecute {
	e.cmd.Binary = binary

	return e
}

// WithInventoryOptions returns an AnsibleInventoryExecute with the ansible's inventory options set
func (e *AnsibleInventoryExecute) WithInventoryOptions(options *AnsibleInventoryOptions) *AnsibleInventoryExecute {
	e.cmd.InventoryOptions = options

	return e
}

// WithPattern return a AnsibleInventoryExecute withthe binary file set
func (e *AnsibleInventoryExecute) WithPattern(pattern string) *AnsibleInventoryExecute {
	e.cmd.Pattern = pattern

	return e
}

// Execute method runs the ansible-inventory command using a DefaultExecute with default options
func (e *AnsibleInventoryExecute) Execute(ctx context.Context) error {

	exec := execute.NewDefaultExecute(
		execute.WithCmd(e.cmd),
	)

	err := exec.Execute(ctx)
	if err != nil {
		return err
	}

	return nil
}
