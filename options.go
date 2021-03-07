package ansibler

import (
	"fmt"
	"os"
)

// Binaries
const (
	// DefaultAnsiblePlaybookBinary is the ansible-playbook binary file default value
	DefaultAnsiblePlaybookBinary = "ansible-playbook"

	// DefaultAnsibleBinary is the ansible binary file default value
	DefaultAnsibleBinary = "ansible"
)

// Optional arguments
const (

	// ExtraVarsFlag is the extra variables flag for ansible-playbook
	ExtraVarsFlag = "--extra-vars"

	// FlushCacheFlag is the flush cache flag for ansible-playbook
	FlushCacheFlag = "--flush-cache"

	// InventoryFlag is the inventory flag for ansible-playbook
	InventoryFlag = "--inventory"

	// LimitFlag is the limit flag for ansible-playbook
	LimitFlag = "--limit"

	// ListHostsFlag is the list hosts flag for ansible-playbook
	ListHostsFlag = "--list-hosts"

	// ListTagsFlag is the list tags flag for ansible-playbook
	ListTagsFlag = "--list-tags"

	// ListTasksFlag is the list tasks flag for ansible-playbook
	ListTasksFlag = "--list-tasks"

	// TagsFlag is the tags flag for ansible-playbook
	TagsFlag = "--tags"

	// SyntaxCheckFlag is the syntax check flag for ansible-playbook
	SyntaxCheckFlag = "--syntax-check"

	// VaultPasswordFileFlag is the vault password file flag for ansible-playbook
	VaultPasswordFileFlag = "--vault-password-file"

	// ansible configuration consts

	// AnsibleForceColorEnv is the environment variable which forces color mode
	AnsibleForceColorEnv = "ANSIBLE_FORCE_COLOR"

	// AnsibleHostKeyCheckingEnv
	AnsibleHostKeyCheckingEnv = "ANSIBLE_HOST_KEY_CHECKING"
)

// Connection Options
const (
	// PrivateKeyFlag is the private key file flag for ansible-playbook
	PrivateKeyFlag = "--private-key"

	// TimeoutFlag is the timeout flag for ansible-playbook
	TimeoutFlag = "--timeout"

	// ConnectionFlag is the connection flag for ansible-playbook
	ConnectionFlag = "--connection"

	// AskPassFlag is ansble-playbook's ask for connection password flag
	AskPassFlag = "--ask-pass"

	// UserFlag is the user flag for ansible-playbook
	UserFlag = "--user"
)

// Privilege Escalation Options
const (
	// BecomeMethodFlag is ansble-playbook's become method flag
	BecomeMethodFlag = "--become-method"

	// BecomeUserFlag is ansble-playbook's become user flag
	BecomeUserFlag = "--become-user"

	// AskBecomePassFlag is ansble-playbook's ask for become user password flag
	AskBecomePassFlag = "--ask-become-pass"

	// BecomeFlag is ansble-playbook's become flag
	BecomeFlag = "--become"
)

// Executor is and interface that should be implemented for those item which could run ansible playbooks
type Executor interface {
	Execute(command string, args []string, prefix string) error
}

// AnsibleForceColor changes to a forced color mode
func AnsibleForceColor() {
	os.Setenv(AnsibleForceColorEnv, "true")
}

// AnsibleAvoidHostKeyChecking sets the hosts key checking to false
func AnsibleAvoidHostKeyChecking() {
	os.Setenv(AnsibleHostKeyCheckingEnv, "false")
}

// AnsibleSetEnv set any configuration by environment variables. Check ansible configuration at https://docs.ansible.com/ansible/latest/reference_appendices/config.html
func AnsibleSetEnv(key, value string) {
	os.Setenv(key, value)
}

// AnsibleConnectionOptions object has those parameters described on `Connections Options` section within ansible-playbook's man page, and which defines how to connect to hosts.
type AnsibleConnectionOptions struct {
	// AskPass defines whether user's password should be asked to connect to host
	AskPass bool
	// Connection is the type of connection used by ansible-playbook
	Connection string
	// PrivateKey is the user's private key file used to connect to a host
	PrivateKey string
	// Timeout is the connection timeout on ansible-playbook. Take care because Timeout is defined ad string
	Timeout string
	// User is the user to use to connect to a host
	User string
}

// GenerateCommandConnectionOptions return a list of connection options flags to be used on ansible-playbook execution
func (o *AnsibleConnectionOptions) GenerateCommandConnectionOptions() ([]string, error) {
	cmd := []string{}

	if o.AskPass {
		cmd = append(cmd, AskPassFlag)
	}

	if o.Connection != "" {
		cmd = append(cmd, ConnectionFlag)
		cmd = append(cmd, o.Connection)
	}

	if o.PrivateKey != "" {
		cmd = append(cmd, PrivateKeyFlag)
		cmd = append(cmd, o.PrivateKey)
	}

	if o.User != "" {
		cmd = append(cmd, UserFlag)
		cmd = append(cmd, o.User)
	}

	if o.Timeout != "" {
		cmd = append(cmd, TimeoutFlag)
		cmd = append(cmd, o.Timeout)
	}

	return cmd, nil
}

// String return a list of connection options flags to be used on ansible-playbook execution
func (o *AnsibleConnectionOptions) String() string {
	str := ""

	if o.AskPass {
		str = fmt.Sprintf("%s %s", str, AskPassFlag)
	}

	if o.Connection != "" {
		str = fmt.Sprintf("%s %s %s", str, ConnectionFlag, o.Connection)
	}

	if o.PrivateKey != "" {
		str = fmt.Sprintf("%s %s %s", str, PrivateKeyFlag, o.PrivateKey)
	}

	if o.User != "" {
		str = fmt.Sprintf("%s %s %s", str, UserFlag, o.User)
	}

	if o.Timeout != "" {
		str = fmt.Sprintf("%s %s %s", str, TimeoutFlag, o.Timeout)
	}

	return str
}

/* become methods
ksu        Kerberos substitute user
pbrun      PowerBroker run
enable     Switch to elevated permissions on a network device
sesu       CA Privileged Access Manager
pmrun      Privilege Manager run
runas      Run As user
sudo       Substitute User DO
su         Substitute User
doas       Do As user
pfexec     profile based execution
machinectl Systemd's machinectl privilege escalation
dzdo       Centrify's Direct Authorize
*/

// AnsiblePrivilegeEscalationOptions object has those parameters described on `Privilege Escalation Options` section within ansible-playbook's man page, and which controls how and which user you become as on target hosts.
type AnsiblePrivilegeEscalationOptions struct {
	// Become
	Become bool
	// BecomeMethod
	BecomeMethod string
	// BecomeUser
	BecomeUser string
	// AskBecomePass
	AskBecomePass bool
}

// GenerateCommandPrivilegeEscalationOptions return a list of privilege escalation options flags to be used on ansible-playbook execution
func (o *AnsiblePrivilegeEscalationOptions) GenerateCommandPrivilegeEscalationOptions() ([]string, error) {
	cmd := []string{}

	if o.AskBecomePass {
		cmd = append(cmd, AskBecomePassFlag)
	}

	if o.Become {
		cmd = append(cmd, BecomeFlag)
	}

	if o.BecomeMethod != "" {
		cmd = append(cmd, BecomeMethodFlag)
		cmd = append(cmd, o.BecomeMethod)
	}

	if o.BecomeUser != "" {
		cmd = append(cmd, BecomeUserFlag)
		cmd = append(cmd, o.BecomeUser)
	}

	return cmd, nil
}

// String return an string
func (o *AnsiblePrivilegeEscalationOptions) String() string {
	str := ""

	if o.AskBecomePass {
		str = fmt.Sprintf("%s %s", str, AskBecomePassFlag)
	}

	if o.Become {
		str = fmt.Sprintf("%s %s", str, BecomeFlag)
	}

	if o.BecomeMethod != "" {
		str = fmt.Sprintf("%s %s %s", str, BecomeMethodFlag, o.BecomeMethod)
	}

	if o.BecomeUser != "" {
		str = fmt.Sprintf("%s %s %s", str, BecomeUserFlag, o.BecomeUser)
	}

	return str
}
