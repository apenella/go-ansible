package adhoc

import (
	"fmt"

	common "github.com/apenella/go-common-utils/data"
	errors "github.com/apenella/go-common-utils/error"
)

const (
	// ArgsFlag module arguments
	ArgsFlag = "--args"

	// AskVaultPasswordFlag ask for vault password
	AskVaultPasswordFlag = "--ask-vault-password"

	// BackgroundFlag un asynchronously, failing after X seconds (default=N/A)
	BackgroundFlag = "--background"

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

	// ModuleNameFlag module name to execute (default=command)
	ModuleNameFlag = "--module-name"

	// ModulePathFlag repend colon-separated path(s) to module library (default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)
	ModulePathFlag = "--module-path"

	// OneLineFlag condense output
	OneLineFlag = "--one-line"

	// PlaybookDirFlag since this tool does not use playbooks, use this as a substitute playbook directory.This sets the relative path for many features including roles/ group_vars/ etc.
	PlaybookDirFlag = "--playbook-dir"

	// PollFlag set the poll interval if using -B (default=15)
	PollFlag = "--poll"

	// SyntaxCheckFlag is the syntax check flag for ansible-playbook
	SyntaxCheckFlag = "--syntax-check"

	// TreeFlag log output to this directory
	TreeFlag = "--tree"

	// VaultIDFlag the vault identity to use
	VaultIDFlag = "--vault-id"

	// VaultPasswordFileFlag is the vault password file flag for ansible-playbook
	VaultPasswordFileFlag = "--vault-password-file"

	// VersionFlag show program's version number, config file location, configured module search path, module location, executable location and exit
	VersionFlag = "--version"

	// VerboseFlag verbose mode enabled to connection debugging
	VerboseFlag = "-vvvv"

	// VerboseVFlag verbose with -v is enabled
	VerboseVFlag = "-v"

	// VerboseVVFlag verbose with -vv is enabled
	VerboseVVFlag = "-vv"

	// VerboseVFlag verbose with -vvv is enabled
	VerboseVVVFlag = "-vvv"

	// VerboseVFlag verbose with -vvvv is enabled
	VerboseVVVVFlag = "-vvvv"

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
)

// AnsibleAdhocOptions object has those parameters described on `Options` section within ansible-playbook's man page, and which defines which should be the ansible-playbook execution behavior.
type AnsibleAdhocOptions struct {
	// Args module arguments
	Args string

	// AskVaultPassword ask for vault password
	AskVaultPassword bool

	// Background un asynchronously, failing after X seconds (default=N/A)
	Background int

	// Check don't make any changes; instead, try to predict some of the changes that may occur
	Check bool

	// Diff when changing (small) files and templates, show the differences in those files; works great with --check
	Diff bool

	// ExtraVars is a map of extra variables used on ansible-playbook execution
	ExtraVars map[string]interface{}

	// ExtraVarsFile is a list of files used to load extra-vars
	ExtraVarsFile []string

	// Forks specify number of parallel processes to use (default=50)
	Forks string

	// Inventory specify inventory host path
	Inventory string

	// Limit is selected hosts additional pattern
	Limit string

	// ListHosts outputs a list of matching hosts
	ListHosts bool

	// ModuleName module name to execute (default=command)
	ModuleName string

	// ModulePath repend colon-separated path(s) to module library (default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)
	ModulePath string

	// OneLine condense output
	OneLine bool

	// PlaybookDir since this tool does not use playbooks, use this as a substitute playbook directory.This sets the relative path for many features including roles/ group_vars/ etc.
	PlaybookDir string

	// Poll set the poll interval if using -B (default=15)
	Poll int

	// SyntaxCheck is the syntax check flag for ansible-playbook
	SyntaxCheck bool

	// Tree log output to this directory
	Tree string

	// VaultID the vault identity to use
	VaultID string

	// VaultPasswordFile path to the file holding vault decryption key
	VaultPasswordFile string

	// Verbose verbose mode enabled to connection debugging
	Verbose bool

	// Verbose verbose mode -v enabled
	VerboseV bool

	// Verbose verbose mode -vv enabled
	VerboseVV bool

	// Verbose verbose mode -vvv enabled
	VerboseVVV bool

	// Verbose verbose mode -vvvv enabled
	VerboseVVVV bool

	// Version show program's version number, config file location, configured module search path, module location, executable location and exit
	Version bool

	// Parameters defined on `Connections Options` section within ansible-playbook's man page, and which defines how to connect to hosts.

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
	Timeout int

	// User is the user to use to connect to a host
	User string

	// Parameters defined on `Privilege Escalation Options` section within ansible-playbook's man page, and which controls how and which user you become as on target hosts.

	// AskBecomePass is ansble-playbook's ask for become user password flag
	AskBecomePass bool

	// Become is ansble-playbook's become flag
	Become bool

	// BecomeMethod is ansble-playbook's become method. The accepted become methods are:
	// 	- ksu        Kerberos substitute user
	// 	- pbrun      PowerBroker run
	// 	- enable     Switch to elevated permissions on a network device
	// 	- sesu       CA Privileged Access Manager
	// 	- pmrun      Privilege Manager run
	// 	- runas      Run As user
	// 	- sudo       Substitute User DO
	// 	- su         Substitute User
	// 	- doas       Do As user
	// 	- pfexec     profile based execution
	// 	- machinectl Systemd's machinectl privilege escalation
	// 	- dzdo       Centrify's Direct Authorize
	BecomeMethod string

	// BecomeUser is ansble-playbook's become user
	BecomeUser string
}

// AddExtraVar registers a new extra variable
func (o *AnsibleAdhocOptions) AddExtraVar(name string, value interface{}) error {

	if o.ExtraVars == nil {
		o.ExtraVars = map[string]interface{}{}
	}
	_, exists := o.ExtraVars[name]
	if exists {
		return errors.New("(adhoc::AddExtraVar)", fmt.Sprintf("ExtraVar '%s' already exist", name))
	}

	o.ExtraVars[name] = value

	return nil
}

// AddExtraVarsFile adds an extra-vars file on ansible-playbook options item
func (o *AnsibleAdhocOptions) AddExtraVarsFile(file string) error {

	if o.ExtraVarsFile == nil {
		o.ExtraVarsFile = []string{}
	}

	if file[0] != '@' {
		file = fmt.Sprintf("@%s", file)
	}

	o.ExtraVarsFile = append(o.ExtraVarsFile, file)

	return nil
}

// GenerateAnsibleAdhocOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleAdhocOptions) GenerateAnsibleAdhocOptions() ([]string, error) {
	cmd := []string{}

	if o == nil {
		return nil, errors.New("(adhoc::GenerateAnsibleAdhocOptions)", "AnsibleAdhocOptions is nil")
	}

	if o.Args != "" {
		cmd = append(cmd, ArgsFlag)
		cmd = append(cmd, o.Args)
	}

	if o.AskVaultPassword {
		cmd = append(cmd, AskVaultPasswordFlag)
	}

	if o.Background > 0 {
		cmd = append(cmd, BackgroundFlag)
		cmd = append(cmd, fmt.Sprint(o.Background))
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
			return nil, errors.New("(adhoc::GenerateAnsibleAdhocOptions)", "Error generating extra-vars", err)
		}
		cmd = append(cmd, extraVars)
	}

	for _, extraVarsFile := range o.ExtraVarsFile {
		cmd = append(cmd, ExtraVarsFlag)
		cmd = append(cmd, extraVarsFile)
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

	if o.ModuleName != "" {
		cmd = append(cmd, ModuleNameFlag)
		cmd = append(cmd, o.ModuleName)
	}

	if o.ModulePath != "" {
		cmd = append(cmd, ModulePathFlag)
		cmd = append(cmd, o.ModulePath)
	}

	if o.OneLine {
		cmd = append(cmd, OneLineFlag)
	}

	if o.PlaybookDir != "" {
		cmd = append(cmd, PlaybookDirFlag)
		cmd = append(cmd, o.PlaybookDir)
	}

	if o.Poll > 0 {
		cmd = append(cmd, PollFlag)
		cmd = append(cmd, fmt.Sprint(o.Poll))
	}

	if o.SyntaxCheck {
		cmd = append(cmd, SyntaxCheckFlag)
	}

	if o.Tree != "" {
		cmd = append(cmd, TreeFlag)
		cmd = append(cmd, o.Tree)
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

	// Connection options

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
		cmd = append(cmd, o.SCPExtraArgs)
	}

	if o.SFTPExtraArgs != "" {
		cmd = append(cmd, SFTPExtraArgsFlag)
		cmd = append(cmd, o.SFTPExtraArgs)
	}

	if o.SSHCommonArgs != "" {
		cmd = append(cmd, SSHCommonArgsFlag)
		cmd = append(cmd, o.SSHCommonArgs)
	}

	if o.SSHExtraArgs != "" {
		cmd = append(cmd, SSHExtraArgsFlag)
		cmd = append(cmd, o.SSHExtraArgs)
	}

	if o.Timeout > 0 {
		cmd = append(cmd, TimeoutFlag)
		cmd = append(cmd, fmt.Sprint(o.Timeout))
	}

	if o.User != "" {
		cmd = append(cmd, UserFlag)
		cmd = append(cmd, o.User)
	}

	// Privilege escalation options

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

// generateExtraVarsCommand return a string which is a json structure having all the extra variable
func (o *AnsibleAdhocOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(adhoc::generateExtraVarsCommand)", "Error creationg extra-vars JSON object to string", err)
	}
	return extraVars, nil
}

// generateVerbosityFlag return a string with the verbose flag. Higher verbosity (more v's) has precedence over lower
func (o *AnsibleAdhocOptions) generateVerbosityFlag() (string, error) {
	if o.Verbose {
		return VerboseFlag, nil
	}

	if o.VerboseVVVV {
		return VerboseVVVVFlag, nil
	}

	if o.VerboseVVV {
		return VerboseVVVFlag, nil
	}

	if o.VerboseVV {
		return VerboseVVFlag, nil
	}

	if o.VerboseV {
		return VerboseVFlag, nil
	}

	return "", nil
}

// GenerateCommandCommonOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleAdhocOptions) String() string {
	str := ""

	if o.Args != "" {
		str = fmt.Sprintf("%s %s '%s'", str, ArgsFlag, o.Args)
	}

	if o.AskVaultPassword {
		str = fmt.Sprintf("%s %s", str, AskVaultPasswordFlag)
	}

	if o.Background > 0 {
		str = fmt.Sprintf("%s %s %d", str, BackgroundFlag, o.Background)
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

	for _, file := range o.ExtraVarsFile {
		str = fmt.Sprintf("%s %s %s", str, ExtraVarsFlag, file)
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

	if o.ModuleName != "" {
		str = fmt.Sprintf("%s %s %s", str, ModuleNameFlag, o.ModuleName)
	}

	if o.ModulePath != "" {
		str = fmt.Sprintf("%s %s %s", str, ModulePathFlag, o.ModulePath)
	}

	if o.OneLine {
		str = fmt.Sprintf("%s %s", str, OneLineFlag)
	}

	if o.PlaybookDir != "" {
		str = fmt.Sprintf("%s %s %s", str, PlaybookDirFlag, o.PlaybookDir)
	}

	if o.Poll > 0 {
		str = fmt.Sprintf("%s %s %d", str, PollFlag, o.Poll)
	}

	if o.SyntaxCheck {
		str = fmt.Sprintf("%s %s", str, SyntaxCheckFlag)
	}

	if o.Tree != "" {
		str = fmt.Sprintf("%s %s %s", str, TreeFlag, o.Tree)
	}

	if o.VaultID != "" {
		str = fmt.Sprintf("%s %s %s", str, VaultIDFlag, o.VaultID)
	}

	if o.VaultPasswordFile != "" {
		str = fmt.Sprintf("%s %s %s", str, VaultPasswordFileFlag, o.VaultPasswordFile)
	}

	verbosityString, _ := o.generateVerbosityFlag()
	if verbosityString != "" {
		str = fmt.Sprintf("%s %s", str, verbosityString)
	}

	if o.Version {
		str = fmt.Sprintf("%s %s", str, VersionFlag)
	}

	// Connection options

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
		str = fmt.Sprintf("%s %s '%s'", str, SCPExtraArgsFlag, o.SCPExtraArgs)
	}

	if o.SFTPExtraArgs != "" {
		str = fmt.Sprintf("%s %s '%s'", str, SFTPExtraArgsFlag, o.SFTPExtraArgs)
	}

	if o.SSHCommonArgs != "" {
		str = fmt.Sprintf("%s %s '%s'", str, SSHCommonArgsFlag, o.SSHCommonArgs)
	}

	if o.SSHExtraArgs != "" {
		str = fmt.Sprintf("%s %s '%s'", str, SSHExtraArgsFlag, o.SSHExtraArgs)
	}

	if o.Timeout > 0 {
		str = fmt.Sprintf("%s %s %d", str, TimeoutFlag, o.Timeout)
	}

	if o.User != "" {
		str = fmt.Sprintf("%s %s %s", str, UserFlag, o.User)
	}

	// Privilege escalation options

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
