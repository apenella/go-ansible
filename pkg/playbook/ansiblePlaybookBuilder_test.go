package playbook

import (
	"testing"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/stretchr/testify/assert"
)

// TestNewAnsiblePlaybookCmdBuilder tests NewAnsiblePlaybookCmdBuilder function
func TestNewAnsiblePlaybookCmdBuilder(t *testing.T) {
	b := NewAnsiblePlaybookCmdBuilder()
	assert.NotNil(t, b)
}

func TestAnsiblePlaybookCmdBuilderAllOptions(t *testing.T) {

	playbookOptions := &AnsiblePlaybookOptions{}
	connectionOptions := &options.AnsibleConnectionOptions{}
	privilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{}

	b := NewAnsiblePlaybookCmdBuilder().
		WithBinary("ansible-playbook").
		WithPlaybooks([]string{"playbook.yml"}).
		WithOptions(playbookOptions).
		WithConnectionOptions(connectionOptions).
		WithPrivilegeEscalationOptions(privilegeEscalationOptions).Build()

	assert.Equal(t, "ansible-playbook", b.Binary)
	assert.Equal(t, []string{"playbook.yml"}, b.Playbooks)
	assert.Equal(t, playbookOptions, b.Options)
	assert.Equal(t, connectionOptions, b.ConnectionOptions)
	assert.Equal(t, privilegeEscalationOptions, b.PrivilegeEscalationOptions)

}
