package playbook

import (
	"github.com/apenella/go-ansible/pkg/options"
)

// AnsiblePlaybookCmdBuilder is a builder for the AnsiblePlaybookCmd struct
type AnsiblePlaybookCmdBuilder struct {
	cmd *AnsiblePlaybookCmd
}

// NewAnsiblePlaybookCmdBuilder return a new AnsiblePlaybookCmdBuilder
func NewAnsiblePlaybookCmdBuilder() *AnsiblePlaybookCmdBuilder {
	return &AnsiblePlaybookCmdBuilder{
		cmd: &AnsiblePlaybookCmd{},
	}
}

// Build return a AnsiblePlaybookCmd struct
func (b *AnsiblePlaybookCmdBuilder) Build() *AnsiblePlaybookCmd {
	return b.cmd
}

// WithAnsible return a AnsiblePlaybookCmdBuilder withthe binary file set
func (b *AnsiblePlaybookCmdBuilder) WithBinary(binary string) *AnsiblePlaybookCmdBuilder {
	b.cmd.Binary = binary

	return b
}

// WithPlaybooks returns an AnsiblePlaybookCmdBuilder with the ansible's playbooks list set
func (b *AnsiblePlaybookCmdBuilder) WithPlaybooks(playbooks []string) *AnsiblePlaybookCmdBuilder {
	b.cmd.Playbooks = playbooks

	return b
}

// WithOptions returns an AnsiblePlaybookCmdBuilder with the ansible's playbook options set
func (b *AnsiblePlaybookCmdBuilder) WithOptions(options *AnsiblePlaybookOptions) *AnsiblePlaybookCmdBuilder {
	b.cmd.Options = options

	return b
}

// WithConnectionOptions returns an AnsiblePlaybookCmdBuilder with the ansible's playbook specific options for connection set
func (b *AnsiblePlaybookCmdBuilder) WithConnectionOptions(options *options.AnsibleConnectionOptions) *AnsiblePlaybookCmdBuilder {
	b.cmd.ConnectionOptions = options

	return b
}

// WithPrivilegeEscalationOptions returns an AnsiblePlaybookCmdBuilder with the ansible's playbook privilege escalation options set
func (b *AnsiblePlaybookCmdBuilder) WithPrivilegeEscalationOptions(options *options.AnsiblePrivilegeEscalationOptions) *AnsiblePlaybookCmdBuilder {
	b.cmd.PrivilegeEscalationOptions = options

	return b
}
