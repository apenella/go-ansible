package playbook

import (
	"fmt"

	"github.com/apenella/go-ansible/pkg/options"
	common "github.com/apenella/go-common-utils/data"
	errors "github.com/apenella/go-common-utils/error"
)

const (
	// TODO: error management
	// // AnsiblePlaybookErrorCodeGeneralError is the error code for a general error
	// AnsiblePlaybookErrorCodeGeneralError = 1
	// // AnsiblePlaybookErrorCodeOneOrMoreHostFailed is the error code for a one or more host failed
	// AnsiblePlaybookErrorCodeOneOrMoreHostFailed = 2
	// // AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable is the error code for a one or more host unreachable
	// AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable = 3
	// // AnsiblePlaybookErrorCodeParserError is the error code for a parser error
	// AnsiblePlaybookErrorCodeParserError = 4
	// // AnsiblePlaybookErrorCodeBadOrIncompleteOptions is the error code for a bad or incomplete options
	// AnsiblePlaybookErrorCodeBadOrIncompleteOptions = 5
	// // AnsiblePlaybookErrorCodeUserInterruptedExecution is the error code for a user interrupted execution
	// AnsiblePlaybookErrorCodeUserInterruptedExecution = 99
	// // AnsiblePlaybookErrorCodeUnexpectedError is the error code for a unexpected error
	// AnsiblePlaybookErrorCodeUnexpectedError = 250

	// // AnsiblePlaybookErrorMessageGeneralError is the error message for a general error
	// AnsiblePlaybookErrorMessageGeneralError = "ansible-playbook error: general error"
	// // AnsiblePlaybookErrorMessageOneOrMoreHostFailed is the error message for a one or more host failed
	// AnsiblePlaybookErrorMessageOneOrMoreHostFailed = "ansible-playbook error: one or more host failed"
	// // AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable is the error message for a one or more host unreachable
	// AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable = "ansible-playbook error: one or more host unreachable"
	// // AnsiblePlaybookErrorMessageParserError is the error message for a parser error
	// AnsiblePlaybookErrorMessageParserError = "ansible-playbook error: parser error"
	// // AnsiblePlaybookErrorMessageBadOrIncompleteOptions is the error message for a bad or incomplete options
	// AnsiblePlaybookErrorMessageBadOrIncompleteOptions = "ansible-playbook error: bad or incomplete options"
	// // AnsiblePlaybookErrorMessageUserInterruptedExecution is the error message for a user interrupted execution
	// AnsiblePlaybookErrorMessageUserInterruptedExecution = "ansible-playbook error: user interrupted execution"
	// // AnsiblePlaybookErrorMessageUnexpectedError is the error message for a unexpected error
	// AnsiblePlaybookErrorMessageUnexpectedError = "ansible-playbook error: unexpected error"

	// DefaultAnsiblePlaybookBinary is the ansible-playbook binary file default value
	DefaultAnsiblePlaybookBinary = "ansible-playbook"

	// AskVaultPasswordFlag ask for vault password
	AskVaultPasswordFlag = "--ask-vault-password"

	// CheckFlag don't make any changes; instead, try to predict some of the changes that may occur
	CheckFlag = "--check"

	// DiffFlag when changing (small) files and templates, show the differences in those files; works great with --check
	DiffFlag = "--diff"

	// ExtraVarsFlag is the extra variables flag for ansible-playbook
	ExtraVarsFlag = "--extra-vars"

	// FlushCacheFlag is the flush cache flag for ansible-playbook
	FlushCacheFlag = "--flush-cache"

	// ForceHandlersFlag run handlers even if a task fails
	ForceHandlersFlag = "--force-handlers"

	// ForksFlag specify number of parallel processes to use (default=50)
	ForksFlag = "--forks"

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

	// ModulePathFlag repend colon-separated path(s) to module library (default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)
	ModulePathFlag = "--module-path"

	// SkipTagsFlag only run plays and tasks whose tags do not match these values
	SkipTagsFlag = "--skip-tags"

	// StartAtTaskFlag start the playbook at the task matching this name
	StartAtTaskFlag = "--start-at-task"

	// StepFlag one-step-at-a-time: confirm each task before running
	StepFlag = "--step"

	// SyntaxCheckFlag is the syntax check flag for ansible-playbook
	SyntaxCheckFlag = "--syntax-check"

	// TagsFlag is the tags flag for ansible-playbook
	TagsFlag = "--tags"

	// VaultIDFlag the vault identity to use
	VaultIDFlag = "--vault-id"

	// VaultPasswordFileFlag is the vault password file flag for ansible-playbook
	VaultPasswordFileFlag = "--vault-password-file"

	// VersionFlag show program's version number, config file location, configured module search path, module location, executable location and exit
	VersionFlag = "--version"

	// VerboseFlag verbose mode enabled
	VerboseFlag = "-vvvv"

	// VerboseVFlag verbose with -v is enabled
	VerboseVFlag = "-v"

	// VerboseVVFlag verbose with -vv is enabled
	VerboseVVFlag = "-vv"

	// VerboseVFlag verbose with -vvv is enabled
	VerboseVVVFlag = "-vvv"

	// VerboseVFlag verbose with -vvvv is enabled
	VerboseVVVVFlag = "-vvvv"
)

// AnsiblePlaybookOptionsFunc is a function to set executor options
type AnsiblePlaybookOptionsFunc func(*AnsiblePlaybookCmd)

// AnsiblePlaybookCmd object is the main object which defines the `ansible-playbook` command and how to execute it.
type AnsiblePlaybookCmd struct {
	// Ansible binary file
	Binary string
	// Playbooks is the ansible's playbooks list to be used
	Playbooks []string
	// PlaybookOptions are the ansible's playbook options
	PlaybookOptions *AnsiblePlaybookOptions
	// ConnectionOptions are the ansible's playbook specific options for connection
	ConnectionOptions *options.AnsibleConnectionOptions
	// PrivilegeEscalationOptions are the ansible's playbook privilege escalation options
	PrivilegeEscalationOptions *options.AnsiblePrivilegeEscalationOptions
}

// Command generate the ansible-playbook command which will be executed
func (p *AnsiblePlaybookCmd) Command() ([]string, error) {
	cmd := []string{}

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsiblePlaybookBinary
	}

	// Set the ansible-playbook binary file
	cmd = append(cmd, p.Binary)

	// Determine the options to be set
	if p.PlaybookOptions != nil {
		options, err := p.PlaybookOptions.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(playbook::Command)", "Error creating options", err)
		}
		cmd = append(cmd, options...)
	}

	// Determine the connection options to be set
	if p.ConnectionOptions != nil {
		options, err := p.ConnectionOptions.GenerateCommandConnectionOptions()
		if err != nil {
			return nil, errors.New("(playbook::Command)", "Error creating connection options", err)
		}
		cmd = append(cmd, options...)
	}

	// Determine the privilege escalation options to be set
	if p.PrivilegeEscalationOptions != nil {
		options, err := p.PrivilegeEscalationOptions.GenerateCommandPrivilegeEscalationOptions()
		if err != nil {
			return nil, errors.New("(playbook::Command)", "Error creating privilege escalation options", err)
		}
		cmd = append(cmd, options...)
	}

	// Include the ansible playbook
	cmd = append(cmd, p.Playbooks...)

	return cmd, nil
}

// String returns AnsiblePlaybookCmd as string
func (p *AnsiblePlaybookCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsiblePlaybookBinary
	}

	str := p.Binary

	if p.PlaybookOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.PlaybookOptions.String())
	}
	if p.ConnectionOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.ConnectionOptions.String())
	}
	if p.PrivilegeEscalationOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.PrivilegeEscalationOptions.String())
	}

	// Include the ansible playbook
	for _, playbook := range p.Playbooks {
		str = fmt.Sprintf("%s %s", str, playbook)
	}

	return str
}

// TODO: error management for Ansible Playbook
// func (p *AnsiblePlaybookCmd) Error(ctx context.Context, err error) error {

// 	if err != nil {
// 		if ctx.Err() != nil {
// 			goerrors.Wrap(err, fmt.Sprintf("\nWhoops! %s\n", ctx.Err()))

// 			fmt.Fprintf(e.Write, "%s\n", fmt.Sprintf("\nWhoops! %s\n", ctx.Err()))
// 		} else {
// 			errorMessage := fmt.Sprintf("Command executed:\n%s\n", cmd.String())
// 			if len(e.EnvVars) > 0 {
// 				errorMessage = fmt.Sprintf("%s\nEnvironment variables:\n%s\n", errorMessage, strings.Join(e.EnvVars.Environ(), "\n"))
// 			}
// 			errorMessage = fmt.Sprintf("%s\nError:\n%s\n", errorMessage, err.Error())
// 			stderrErrorMessage := string(err.(*osexec.ExitError).Stderr)
// 			if len(stderrErrorMessage) > 0 {
// 				errorMessage = fmt.Sprintf("%s\n'%s'\n", errorMessage, stderrErrorMessage)
// 			}

// 			exitError, exists := err.(*osexec.ExitError)
// 			if exists {
// 				ws := exitError.Sys().(syscall.WaitStatus)
// 				switch ws.ExitStatus() {
// 				case AnsiblePlaybookErrorCodeGeneralError:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageGeneralError, errorMessage)
// 				case AnsiblePlaybookErrorCodeOneOrMoreHostFailed:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageOneOrMoreHostFailed, errorMessage)
// 				case AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable, errorMessage)
// 				case AnsiblePlaybookErrorCodeParserError:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageParserError, errorMessage)
// 				case AnsiblePlaybookErrorCodeBadOrIncompleteOptions:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageBadOrIncompleteOptions, errorMessage)
// 				case AnsiblePlaybookErrorCodeUserInterruptedExecution:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageUserInterruptedExecution, errorMessage)
// 				case AnsiblePlaybookErrorCodeUnexpectedError:
// 					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageUnexpectedError, errorMessage)
// 				}
// 			}
// 			return errors.New("(DefaultExecute::Execute)", fmt.Sprintf("Error during command execution: %s", errorMessage))
// 		}
// 	}
// }

// AnsiblePlaybookOptions object has those parameters described on `Options` section within ansible-playbook's man page, and which defines which should be the ansible-playbook execution behavior.
type AnsiblePlaybookOptions struct {

	// AskVaultPassword ask for vault password
	AskVaultPassword bool

	// Check don't make any changes; instead, try to predict some of the changes that may occur
	Check bool

	// Diff when changing (small) files and templates, show the differences in those files; works great with --check
	Diff bool

	// ExtraVars is a map of extra variables used on ansible-playbook execution
	ExtraVars map[string]interface{}

	// ExtraVarsFile is a list of files used to load extra-vars
	ExtraVarsFile []string

	// FlushCache is the flush cache flag for ansible-playbook
	FlushCache bool

	// ForceHandlers run handlers even if a task fails
	ForceHandlers bool

	// Forks specify number of parallel processes to use (default=50)
	Forks string

	// Inventory specify inventory host path
	Inventory string

	// Limit is selected hosts additional pattern
	Limit string

	// ListHosts outputs a list of matching hosts
	ListHosts bool

	// ListTags is the list tags flag for ansible-playbook
	ListTags bool

	// ListTasks is the list tasks flag for ansible-playbook
	ListTasks bool

	// ModulePath repend colon-separated path(s) to module library (default=~/.ansible/plugins/modules:/usr/share/ansible/plugins/modules)
	ModulePath string

	// SkipTags only run plays and tasks whose tags do not match these values
	SkipTags string

	// StartAtTask start the playbook at the task matching this name
	StartAtTask string

	// Step one-step-at-a-time: confirm each task before running
	Step bool

	// SyntaxCheck is the syntax check flag for ansible-playbook
	SyntaxCheck bool

	// Tags is the tags flag for ansible-playbook
	Tags string

	// VaultID the vault identity to use
	VaultID string

	// VaultPasswordFile path to the file holding vault decryption key
	VaultPasswordFile string

	// Verbose verbose mode enabled
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
}

// GenerateCommandOptions return a list of options flags to be used on ansible-playbook execution
func (o *AnsiblePlaybookOptions) GenerateCommandOptions() ([]string, error) {

	errContext := "(playbook::GenerateCommandOptions)"

	cmd := []string{}

	if o == nil {
		return nil, errors.New(errContext, "AnsiblePlaybookOptions is nil")
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
			return nil, errors.New(errContext, "Error generating extra-vars", err)
		}
		cmd = append(cmd, extraVars)
	}

	for _, file := range o.ExtraVarsFile {
		cmd = append(cmd, ExtraVarsFlag)
		cmd = append(cmd, file)
	}

	if o.FlushCache {
		cmd = append(cmd, FlushCacheFlag)
	}

	if o.ForceHandlers {
		cmd = append(cmd, ForceHandlersFlag)
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

	if o.ListTags {
		cmd = append(cmd, ListTagsFlag)
	}

	if o.ListTasks {
		cmd = append(cmd, ListTasksFlag)
	}

	if o.ModulePath != "" {
		cmd = append(cmd, ModulePathFlag)
		cmd = append(cmd, o.ModulePath)
	}

	if o.SkipTags != "" {
		cmd = append(cmd, SkipTagsFlag)
		cmd = append(cmd, o.SkipTags)
	}

	if o.StartAtTask != "" {
		cmd = append(cmd, StartAtTaskFlag)
		cmd = append(cmd, o.StartAtTask)
	}

	if o.Step {
		cmd = append(cmd, StepFlag)
	}

	if o.SyntaxCheck {
		cmd = append(cmd, SyntaxCheckFlag)
	}

	if o.Tags != "" {
		cmd = append(cmd, TagsFlag)
		cmd = append(cmd, o.Tags)
	}

	if o.VaultID != "" {
		cmd = append(cmd, VaultIDFlag)
		cmd = append(cmd, o.VaultID)
	}

	if o.VaultPasswordFile != "" {
		cmd = append(cmd, VaultPasswordFileFlag)
		cmd = append(cmd, o.VaultPasswordFile)
	}

	verboseFlag, err := o.generateVerbosityFlag()
	if err != nil {
		return nil, errors.New(errContext, "", err)
	}
	if verboseFlag != "" {
		cmd = append(cmd, verboseFlag)
	}

	if o.Version {
		cmd = append(cmd, VersionFlag)
	}

	return cmd, nil
}

// generateVerbosityFlag return a string with the verbose flag. Higher verbosity (more v's) has precedence over lower
func (o *AnsiblePlaybookOptions) generateVerbosityFlag() (string, error) {
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

// generateExtraVarsCommand return a string which is a json structure having all the extra variable
func (o *AnsiblePlaybookOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(playbook::generateExtraVarsCommand)", "Error creationg extra-vars JSON object to string", err)
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
		return errors.New("(playbook::AddExtraVar)", fmt.Sprintf("ExtraVar '%s' already exist", name))
	}

	o.ExtraVars[name] = value

	return nil
}

// AddExtraVarsFile adds an extra-vars file on ansible-playbook options item
func (o *AnsiblePlaybookOptions) AddExtraVarsFile(file string) error {

	if o.ExtraVarsFile == nil {
		o.ExtraVarsFile = []string{}
	}

	if file[0] != '@' {
		file = fmt.Sprintf("@%s", file)
	}

	o.ExtraVarsFile = append(o.ExtraVarsFile, file)

	return nil
}

// AddVaultedExtraVar registers a new extra variable on ansible-playbook options item vaulting its value
func (o *AnsiblePlaybookOptions) AddVaultedExtraVar(vaulter Vaulter, name string, value string) error {

	if vaulter == nil {
		return errors.New("(playbook::AddVaultedExtraVar)", "To define a vaulted extra-var you need to initialize a vaulter")
	}

	if o.ExtraVars == nil {
		o.ExtraVars = map[string]interface{}{}
	}

	_, exists := o.ExtraVars[name]
	if exists {
		return errors.New("(playbook::AddVaultedExtraVar)", fmt.Sprintf("ExtraVar '%s' already exist", name))
	}

	vaultedValue, err := vaulter.Vault(value)
	if err != nil {
		return errors.New("(playbook::AddVaultedExtraVar)", fmt.Sprintf("Variable '%s' can not be vaulted", name), err)
	}

	o.ExtraVars[name] = vaultedValue

	return nil
}

// String returns AnsiblePlaybookOptions as string
func (o *AnsiblePlaybookOptions) String() string {

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

	for _, extraVarsFile := range o.ExtraVarsFile {
		str = fmt.Sprintf("%s %s %s", str, ExtraVarsFlag, extraVarsFile)
	}

	if o.FlushCache {
		str = fmt.Sprintf("%s %s", str, FlushCacheFlag)
	}

	if o.ForceHandlers {
		str = fmt.Sprintf("%s %s", str, ForceHandlersFlag)
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

	if o.ListTags {
		str = fmt.Sprintf("%s %s", str, ListTagsFlag)
	}

	if o.ListTasks {
		str = fmt.Sprintf("%s %s", str, ListTasksFlag)
	}

	if o.ModulePath != "" {
		str = fmt.Sprintf("%s %s %s", str, ModulePathFlag, o.ModulePath)
	}

	if o.SkipTags != "" {
		str = fmt.Sprintf("%s %s %s", str, SkipTagsFlag, o.SkipTags)
	}

	if o.StartAtTask != "" {
		str = fmt.Sprintf("%s %s %s", str, StartAtTaskFlag, o.StartAtTask)
	}

	if o.Step {
		str = fmt.Sprintf("%s %s", str, StepFlag)
	}

	if o.SyntaxCheck {
		str = fmt.Sprintf("%s %s", str, SyntaxCheckFlag)
	}

	if o.Tags != "" {
		str = fmt.Sprintf("%s %s %s", str, TagsFlag, o.Tags)
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
