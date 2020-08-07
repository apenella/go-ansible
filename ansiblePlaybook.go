package ansibler

import (
	"errors"
	"io"
	"os"

	common "github.com/apenella/go-common-utils/data"
)

const (
	// AnsiblePlaybookBin is the ansible-playbook binary file value
	AnsiblePlaybookBin = "ansible-playbook"

	// AskBecomePassFlag is ansble-playbook's ask for become user password flag
	AskBecomePassFlag = "--ask-become-pass"

	// AskPassFlag is ansble-playbook's ask for connection password flag
	AskPassFlag = "--ask-pass"

	// BecomeFlag is ansble-playbook's become flag
	BecomeFlag = "--become"

	// BecomeMethodFlag is ansble-playbook's become method flag
	BecomeMethodFlag = "--become-method"

	// BecomeUserFlag is ansble-playbook's become user flag
	BecomeUserFlag = "--become-user"

	// ConnectionFlag is the connection flag for ansible-playbook
	ConnectionFlag = "--connection"

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

	// PrivateKeyFlag is the private key file flag for ansible-playbook
	PrivateKeyFlag = "--private-key"

	// TagsFlag is the tags flag for ansible-playbook
	TagsFlag = "--tags"

	// SyntaxCheckFlag is the syntax check flag for ansible-playbook
	SyntaxCheckFlag = "--syntax-check"

	// TimeoutFlag is the timeout flag for ansible-playbook
	TimeoutFlag = "--timeout"

	// UserFlag is the user flag for ansible-playbook
	UserFlag = "--user"

	// VaultPasswordFileFlag is the vault password file flag for ansible-playbook
	VaultPasswordFileFlag = "--vault-password-file"

	// AnsibleForceColorEnv is the environment variable which forces color mode
	AnsibleForceColorEnv = "ANSIBLE_FORCE_COLOR"
)

// Executor is and interface that should be implemented for those item which could run ansible playbooks
type Executor interface {
	Execute(command string, args []string, prefix string) error
}

// AnsibleForceColor change to a forced color mode
func AnsibleForceColor() {
	os.Setenv(AnsibleForceColorEnv, "true")
}

// AnsiblePlaybookCmd object is the main object which defines the `ansible-playbook` command and how to execute it.
type AnsiblePlaybookCmd struct {
	// Exec is the executor item
	Exec Executor
	// ExecPrefix is a text that is set at the beginning of each execution line
	ExecPrefix string
	// Playbook is the ansible's playbook name to be used
	Playbook string
	// Options are the ansible's playbook options
	Options *AnsiblePlaybookOptions
	// ConnectionOptions are the ansible's playbook specific options for connection
	ConnectionOptions *AnsiblePlaybookConnectionOptions
	// PrivilegeEscalationOptions are the ansible's playbook privilage escalation options
	PrivilegeEscalationOptions *AnsiblePlaybookPrivilegeEscalationOptions
	// Writer manages the output
	Writer io.Writer
}

// Run method runs the ansible-playbook
func (p *AnsiblePlaybookCmd) Run() error {
	if p == nil {
		return errors.New("(ansible:Run) AnsiblePlaybookCmd is nil")
	}

	// Define a default executor when it is not defined on AnsiblePlaybookCmd
	if p.Exec == nil {
		p.Exec = &DefaultExecute{
			Write: p.Writer,
		}
	}

	// Generate the command to be run
	cmd, err := p.Command()
	if err != nil {
		return errors.New("(ansible:Run) -> " + err.Error())
	}

	// Set default prefix
	if len(p.ExecPrefix) <= 0 {
		p.ExecPrefix = ""
	}

	// Execute the command an return
	return p.Exec.Execute(cmd[0], cmd[1:], p.ExecPrefix)
}

// Command generate the ansible-playbook command which will be executed
func (p *AnsiblePlaybookCmd) Command() ([]string, error) {
	cmd := []string{}
	// Set the ansible-playbook binary file
	cmd = append(cmd, AnsiblePlaybookBin)

	// Determine the options to be set
	if p.Options != nil {
		options, err := p.Options.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command) -> " + err.Error())
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Determine the connection options to be set
	if p.ConnectionOptions != nil {
		options, err := p.ConnectionOptions.GenerateCommandConnectionOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command) -> " + err.Error())
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Determine the privilege escalation options to be set
	if p.PrivilegeEscalationOptions != nil {
		options, err := p.PrivilegeEscalationOptions.GenerateCommandPrivilegeEscalationOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command) -> " + err.Error())
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Include the ansible playbook
	cmd = append(cmd, p.Playbook)

	return cmd, nil
}

// AnsiblePlaybookOptions object has those parameters described on `Options` section within ansible-playbook's man page, and which defines which should be the ansible-playbook execution behavior.
type AnsiblePlaybookOptions struct {
	// ExtraVars is a map of extra variables used on ansible-playbook execution
	ExtraVars map[string]interface{}
	// FlushCache clear the fact cache for every host in inventory
	FlushCache bool
	// Inventory specify inventory host path
	Inventory string
	// Limit is selected hosts additional pattern
	Limit string
	// ListHosts outputs a list of matching hosts
	ListHosts bool
	// ListTags list all available tags
	ListTags bool
	// ListTasks
	ListTasks bool
	// Tags list all tasks that would be executed
	Tags string
}

// GenerateCommandOptions return a list of options flags to be used on ansible-playbook execution
func (o *AnsiblePlaybookOptions) GenerateCommandOptions() ([]string, error) {
	cmd := []string{}

	if o == nil {
		return nil, errors.New("(ansible::GenerateCommandOptions) AnsiblePlaybookOptions is nil")
	}

	if o.FlushCache {
		cmd = append(cmd, FlushCacheFlag)
	}

	if o.Inventory != "" {
		cmd = append(cmd, InventoryFlag)
		cmd = append(cmd, o.Inventory)
	}

	if o.Limit != "" {
		cmd = append(cmd, LimitFlag)
		cmd = append(cmd, o.Limit)
	}

	if o.ListHosts {
		cmd = append(cmd, ListHostsFlag)
	}

	if o.ListTags {
		cmd = append(cmd, ListTagsFlag)
	}

	if o.ListTasks {
		cmd = append(cmd, ListTasksFlag)
	}

	if o.Tags != "" {
		cmd = append(cmd, TagsFlag)
		cmd = append(cmd, o.Tags)
	}

	if len(o.ExtraVars) > 0 {
		cmd = append(cmd, ExtraVarsFlag)
		extraVars, err := o.generateExtraVarsCommand()
		if err != nil {
			return nil, errors.New("(ansible::GenerateCommandOptions) -> " + err.Error())
		}
		cmd = append(cmd, extraVars)
	}

	return cmd, nil
}

// generateExtraVarsCommand return an string which is a json structure having all the extra variable
func (o *AnsiblePlaybookOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(ansible::generateExtraVarsCommand) -> " + err.Error())
	}
	return extraVars, nil
}

// AddExtraVar registers a new extra variable on ansible-playbook options item
func (o *AnsiblePlaybookOptions) AddExtraVar(name string, value interface{}) error {

	if o.ExtraVars == nil {
		o.ExtraVars = map[string]interface{}{}
	}
	_, exists := o.ExtraVars[name]
	if exists {
		return errors.New("(ansible::AddExtraVar) ExtraVar '" + name + "' already exist")
	}

	o.ExtraVars[name] = value

	return nil
}

// AnsiblePlaybookConnectionOptions object has those parameters described on `Connections Options` section within ansible-playbook's man page, and which defines how to connect to hosts.
type AnsiblePlaybookConnectionOptions struct {
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
func (o *AnsiblePlaybookConnectionOptions) GenerateCommandConnectionOptions() ([]string, error) {
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

// AnsiblePlaybookPrivilegeEscalationOptions object has those parameters described on `Privilege Escalation Options` section within ansible-playbook's man page, and which controls how and which user you become as on target hosts.
type AnsiblePlaybookPrivilegeEscalationOptions struct {
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
func (o *AnsiblePlaybookPrivilegeEscalationOptions) GenerateCommandPrivilegeEscalationOptions() ([]string, error) {
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
