package adhoc

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	common "github.com/apenella/go-common-utils/data"
	errors "github.com/apenella/go-common-utils/error"
)

const (

	// DefaultAnsibleAdhocBinary is the default value for ansible binary file to run adhoc modules
	DefaultAnsibleAdhocBinary = "ansible"

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
)

// AnsibleAdhocOptionsFunc is a function to set executor options
type AnsibleAdhocOptionsFunc func(*AnsibleAdhocCmd)

// AnsibleAdhocCmd object is the main object which defines the `ansible` adhoc command and how to execute it.
type AnsibleAdhocCmd struct {
	// Ansible binary file
	Binary string
	// Exec is the executor item
	Exec execute.Executor
	// Pattern is the ansible's host pattern
	Pattern string
	// Options are the ansible's playbook options
	Options *AnsibleAdhocOptions
	// ConnectionOptions are the ansible's playbook specific options for connection
	ConnectionOptions *options.AnsibleConnectionOptions
	// PrivilegeEscalationOptions are the ansible's playbook privilage escalation options
	PrivilegeEscalationOptions *options.AnsiblePrivilegeEscalationOptions
	// StdoutCallback defines which is the stdout callback method. By default is used 'default' method. Supported stdout method by go-ansible are: debug, default, dense, json, minimal, null, oneline, stderr, timer, yaml
	StdoutCallback string
}

// Run method runs the ansible-playbook
func (a *AnsibleAdhocCmd) Run(ctx context.Context) error {
	var err error
	var command []string
	options := []execute.ExecuteOptions{}

	if a == nil {
		return errors.New("(adhoc::Run)", "AnsibleAdhocCmd is nil")
	}

	// Use default binary when it is not already defined
	if a.Binary == "" {
		a.Binary = DefaultAnsibleAdhocBinary
	}

	_, err = exec.LookPath(a.Binary)
	if err != nil {
		return errors.New("(adhoc::Run)", fmt.Sprintf("Binary file '%s' does not exists", a.Binary), err)
	}

	// Define a default executor when it is not defined on AnsibleAdhocCmd
	if a.Exec == nil {
		a.Exec = execute.NewDefaultExecute()
	}

	// Configure StdoutCallback method. By default is used ansible's 'default' callback method
	stdoutcallback.AnsibleStdoutCallbackSetEnv(a.StdoutCallback)

	// Generate the command to be run
	command, err = a.Command()
	if err != nil {
		return errors.New("(adhoc::Run)", fmt.Sprintf("Error running '%s'", a.String()), err)
	}

	// Execute the command an return
	return a.Exec.Execute(ctx, command, stdoutcallback.GetResultsFunc(a.StdoutCallback), options...)
}

// Command generate the ansible command which will be executed
func (a *AnsibleAdhocCmd) Command() ([]string, error) {
	cmd := []string{}

	// Use default binary when it is not already defined
	if a.Binary == "" {
		a.Binary = DefaultAnsibleAdhocBinary
	}

	// Set the ansible-playbook binary file
	cmd = append(cmd, a.Binary)

	// Include the ansible playbook
	cmd = append(cmd, a.Pattern)

	// Determine the options to be set
	if a.Options != nil {
		options, err := a.Options.GenerateAnsibleAdhocOptions()
		if err != nil {
			return nil, errors.New("(adhoc::Command)", "Error creating options", err)
		}
		cmd = append(cmd, options...)

	}

	// Determine the connection options to be set
	if a.ConnectionOptions != nil {
		options, err := a.ConnectionOptions.GenerateCommandConnectionOptions()
		if err != nil {
			return nil, errors.New("(adhoc::Command)", "Error creating connection options", err)
		}

		cmd = append(cmd, options...)
	}

	// Determine the privilege escalation options to be set
	if a.PrivilegeEscalationOptions != nil {
		options, err := a.PrivilegeEscalationOptions.GenerateCommandPrivilegeEscalationOptions()
		if err != nil {
			return nil, errors.New("(adhoc::Command)", "Error creating privilege escalation options", err)
		}

		cmd = append(cmd, options...)
	}

	return cmd, nil
}

// String returns AnsibleAdhocCmd as string
func (a *AnsibleAdhocCmd) String() string {

	// Use default binary when it is not already defined
	if a.Binary == "" {
		a.Binary = DefaultAnsibleAdhocBinary
	}

	str := a.Binary

	str = fmt.Sprintf("%s %s", str, a.Pattern)

	if a.Options != nil {
		str = fmt.Sprintf("%s %s", str, a.Options.String())
	}
	if a.ConnectionOptions != nil {
		str = fmt.Sprintf("%s %s", str, a.ConnectionOptions.String())
	}
	if a.PrivilegeEscalationOptions != nil {
		str = fmt.Sprintf("%s %s", str, a.PrivilegeEscalationOptions.String())
	}

	return str
}

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
	// ModuleName module name to execute (default=command)
	ModuleName string
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
	// Version show program's version number, config file location, configured module search path, module location, executable location and exit
	Version bool
}

// GenerateCommandAdhocOptions return a list of command options flags to be used on ansible execution
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

	return cmd, nil
}

// generateExtraVarsCommand return an string which is a json structure having all the extra variable
func (o *AnsibleAdhocOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(adhoc::generateExtraVarsCommand)", "Error creationg extra-vars JSON object to string", err)
	}
	return extraVars, nil
}

// GenerateCommandCommonOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleAdhocOptions) String() string {
	str := ""

	if o.Args != "" {
		str = fmt.Sprintf("%s %s %s", str, ArgsFlag, o.Args)
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

	if o.Verbose {
		str = fmt.Sprintf("%s %s", str, VerboseFlag)
	}

	if o.Version {
		str = fmt.Sprintf("%s %s", str, VersionFlag)
	}

	return str
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
