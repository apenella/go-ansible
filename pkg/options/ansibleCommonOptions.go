package options

import (
	"fmt"
	"os"

	common "github.com/apenella/go-common-utils/data"
	errors "github.com/apenella/go-common-utils/error"
)

const (
	//
	// Optional arguments

	// AskVaultPasswordFlag ask for vault password
	AskVaultPasswordFlag = "--ask-vault-password"

	// CheckFlag don't make any changes; instead, try to predict some of the changes that may occur
	CheckFlag = "--check"

	// DiffFlag when changing (small) files and templates, show the differences in those files; works great with --check
	DiffFlag = "--diff"

	// ExtraVarsFlag is the extra variables flag for ansible-playbook
	ExtraVarsFlag = "--extra-vars"

	// ForksFlag specify number of parallel processes to use (default=50)
	ForksFlag = "--forks"

	// InventoryFlag is the inventory flag for ansible-playbook
	InventoryFlag = "--inventory"

	// LimitFlag is the limit flag for ansible-playbook
	LimitFlag = "--limit"

	// ListHostsFlag is the list hosts flag for ansible-playbook
	ListHostsFlag = "--list-hosts"

	// ModulePathFlag repend colon-separated path(s) to module library (default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)
	ModulePathFlag = "--module-path"

	// SyntaxCheckFlag is the syntax check flag for ansible-playbook
	SyntaxCheckFlag = "--syntax-check"

	// VaultIDFlag the vault identity to use
	VaultIDFlag = "--vault-id"

	// VaultPasswordFileFlag is the vault password file flag for ansible-playbook
	VaultPasswordFileFlag = "--vault-password-file"

	// VersionFlag show program's version number, config file location, configured module search path, module location, executable location and exit
	VersionFlag = "--version"

	// VerboseFlag verbose mode enabled to connection debugging
	VerboseFlag = "-vvvv"

	//
	// Connection Options

	// AskPassFlag is ansble-playbook's ask for connection password flag
	AskPassFlag = "--ask-pass"

	// ConnectionFlag is the connection flag for ansible-playbook
	ConnectionFlag = "--connection"

	// PrivateKeyFlag is the private key file flag for ansible-playbook
	PrivateKeyFlag = "--private-key"

	// SCPExtraArgsFlag specify extra arguments to pass to scp only
	SCPExtraArgsFlag = "--scp-extra-args"

	// SFTPExtraArgsFlag specify extra arguments to pass to sftp only
	SFTPExtraArgsFlag = "--sftp-extra-args"

	// SSHCommonArgsFlag specify common arguments to pass to sftp/scp/ssh
	SSHCommonArgsFlag = "--ssh-common-args"

	// SSHExtraArgsFlag specify extra arguments to pass to ssh only
	SSHExtraArgsFlag = "--ssh-extra-args"

	// TimeoutFlag is the timeout flag for ansible-playbook
	TimeoutFlag = "--timeout"

	// UserFlag is the user flag for ansible-playbook
	UserFlag = "--user"

	//
	// Privilege Escalation Options

	// BecomeMethodFlag is ansble-playbook's become method flag
	BecomeMethodFlag = "--become-method"

	// BecomeUserFlag is ansble-playbook's become user flag
	BecomeUserFlag = "--become-user"

	// AskBecomePassFlag is ansble-playbook's ask for become user password flag
	AskBecomePassFlag = "--ask-become-pass"

	// BecomeFlag is ansble-playbook's become flag
	BecomeFlag = "--become"

	//
	// ansible configuration consts

	// AnsibleForceColorEnv is the environment variable which forces color mode
	AnsibleForceColorEnv = "ANSIBLE_FORCE_COLOR"

	// AnsibleHostKeyCheckingEnv
	AnsibleHostKeyCheckingEnv = "ANSIBLE_HOST_KEY_CHECKING"
)

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

// AnsibleCommonOptions object has those parameters described on `Options` section within ansible-playbook's man page, and which defines which should be the ansible-playbook execution behavior.
type AnsibleCommonOptions struct {
	// AskVaultPassword ask for vault password
	AskVaultPassword bool
	// Check don't make any changes; instead, try to predict some of the changes that may occur
	Check bool
	// Diff when changing (small) files and templates, show the differences in those files; works great with --check
	Diff bool
	// ExtraVars is a map of extra variables used on ansible-playbook execution
	ExtraVars map[string]interface{}
	// Forks specify number of parallel processes to use (default=50)
	Forks string
	// Inventory specify inventory host path
	Inventory string
	// Limit is selected hosts additional pattern
	Limit string
	// ListHosts outputs a list of matching hosts
	ListHosts bool
	// ModulePath repend colon-separated path(s) to module library (default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)
	ModulePath string
	// SyntaxCheck is the syntax check flag for ansible-playbook
	SyntaxCheck bool
	// VaultID the vault identity to use
	VaultID string
	// VaultPasswordFile path to the file holding vault decryption key
	VaultPasswordFile string
	// Verbose verbose mode enabled to connection debugging
	Verbose bool
	// Version show program's version number, config file location, configured module search path, module location, executable location and exit
	Version bool
}

// GenerateCommandCommonOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleCommonOptions) GenerateCommandCommonOptions() ([]string, error) {
	cmd := []string{}

	if o == nil {
		return nil, errors.New("(ansible::GenerateCommandCommonOptions)", "AnsibleCommonOptions is nil")
	}

	if o.AskVaultPassword {
		cmd = append(cmd, AskVaultPasswordFlag)
	}

	if o.Check {
		cmd = append(cmd, CheckFlag)
	}

	if o.Diff {
		cmd = append(cmd, DiffFlag)
	}

	if len(o.ExtraVars) > 0 {
		cmd = append(cmd, ExtraVarsFlag)
		extraVars, err := o.generateExtraVarsCommand()
		if err != nil {
			return nil, errors.New("(ansible::GenerateCommandCommonOptions)", "Error generating extra-vars", err)
		}
		cmd = append(cmd, extraVars)
	}

	if o.Forks != "" {
		cmd = append(cmd, ForksFlag)
		cmd = append(cmd, o.Forks)
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

	if o.ModulePath != "" {
		cmd = append(cmd, ModulePathFlag)
		cmd = append(cmd, o.ModulePath)
	}

	if o.SyntaxCheck {
		cmd = append(cmd, SyntaxCheckFlag)
	}

	if o.VaultID != "" {
		cmd = append(cmd, VaultIDFlag)
		cmd = append(cmd, o.VaultID)
	}

	if o.VaultPasswordFile != "" {
		cmd = append(cmd, VaultPasswordFileFlag)
		cmd = append(cmd, o.VaultPasswordFile)
	}

	if o.Verbose {
		cmd = append(cmd, VerboseFlag)
	}

	if o.Version {
		cmd = append(cmd, VersionFlag)
	}

	return cmd, nil
}

// generateExtraVarsCommand return an string which is a json structure having all the extra variable
func (o *AnsibleCommonOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(ansible::generateExtraVarsCommand)", "Error creationg extra-vars JSON object to string", err)
	}
	return extraVars, nil
}

// GenerateCommandCommonOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleCommonOptions) String() string {
	str := ""

	if o.AskVaultPassword {
		str = fmt.Sprintf("%s %s", str, AskVaultPasswordFlag)
	}

	if o.Check {
		str = fmt.Sprintf("%s %s", str, CheckFlag)
	}

	if o.Diff {
		str = fmt.Sprintf("%s %s", str, DiffFlag)
	}

	if len(o.ExtraVars) > 0 {
		extraVars, _ := o.generateExtraVarsCommand()
		str = fmt.Sprintf("%s %s %s", str, ExtraVarsFlag, fmt.Sprintf("'%s'", extraVars))
	}

	if o.Forks != "" {
		str = fmt.Sprintf("%s %s %s", str, ForksFlag, o.Forks)
	}

	if o.Inventory != "" {
		str = fmt.Sprintf("%s %s %s", str, InventoryFlag, o.Inventory)
	}

	if o.Limit != "" {
		str = fmt.Sprintf("%s %s %s", str, LimitFlag, o.Limit)
	}

	if o.ListHosts {
		str = fmt.Sprintf("%s %s", str, ListHostsFlag)
	}

	if o.ModulePath != "" {
		str = fmt.Sprintf("%s %s %s", str, ModulePathFlag, o.ModulePath)
	}

	if o.SyntaxCheck {
		str = fmt.Sprintf("%s %s", str, SyntaxCheckFlag)
	}

	if o.VaultID != "" {
		str = fmt.Sprintf("%s %s %s", str, VaultIDFlag, o.VaultID)
	}

	if o.VaultPasswordFile != "" {
		str = fmt.Sprintf("%s %s %s", str, VaultPasswordFileFlag, o.VaultPasswordFile)
	}

	if o.Verbose {
		str = fmt.Sprintf("%s %s", str, VerboseFlag)
	}

	if o.Version {
		str = fmt.Sprintf("%s %s", str, VersionFlag)
	}

	return str
}

// AddExtraVar registers a new extra variable
func (o *AnsibleCommonOptions) AddExtraVar(name string, value interface{}) error {

	if o.ExtraVars == nil {
		o.ExtraVars = map[string]interface{}{}
	}
	_, exists := o.ExtraVars[name]
	if exists {
		return errors.New("(ansible::AddExtraVar)", fmt.Sprintf("ExtraVar '%s' already exist", name))
	}

	o.ExtraVars[name] = value

	return nil
}

// AnsibleConnectionOptions object has those parameters described on `Connections Options` section within ansible-playbook's man page, and which defines how to connect to hosts.
type AnsibleConnectionOptions struct {
	// AskPass defines whether user's password should be asked to connect to host
	AskPass bool
	// Connection is the type of connection used by ansible-playbook
	Connection string
	// PrivateKey is the user's private key file used to connect to a host
	PrivateKey string
	// SCPExtraArgs specify extra arguments to pass to scp only
	SCPExtraArgs string
	// SFTPExtraArgs specify extra arguments to pass to sftp only
	SFTPExtraArgs string
	// SSHCommonArgs specify common arguments to pass to sftp/scp/ssh
	SSHCommonArgs string
	// SSHExtraArgs specify extra arguments to pass to ssh only
	SSHExtraArgs string
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

	if o.SCPExtraArgs != "" {
		cmd = append(cmd, SCPExtraArgsFlag)
	}

	if o.SFTPExtraArgs != "" {
		cmd = append(cmd, SFTPExtraArgsFlag)
	}

	if o.SSHCommonArgs != "" {
		cmd = append(cmd, o.SSHCommonArgs)
	}

	if o.SSHExtraArgs != "" {
		cmd = append(cmd, SSHExtraArgsFlag)
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

	if o.SCPExtraArgs != "" {
		str = fmt.Sprintf("%s %s %s", str, SCPExtraArgsFlag, o.SCPExtraArgs)
	}

	if o.SFTPExtraArgs != "" {
		str = fmt.Sprintf("%s %s %s", str, SFTPExtraArgsFlag, o.SFTPExtraArgs)
	}

	if o.SSHCommonArgs != "" {
		str = fmt.Sprintf("%s %s %s", str, SSHCommonArgsFlag, o.SSHCommonArgs)
	}

	if o.SSHExtraArgs != "" {
		str = fmt.Sprintf("%s %s %s", str, SSHExtraArgsFlag, o.SSHExtraArgs)
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
