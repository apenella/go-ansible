package playbook

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/stdoutcallback"
	errors "github.com/apenella/go-common-utils/error"
)

const (

	// DefaultAnsiblePlaybookBinary is the ansible-playbook binary file default value
	DefaultAnsiblePlaybookBinary = "ansible-playbook"

	// FlushCacheFlag is the flush cache flag for ansible-playbook
	FlushCacheFlag = "--flush-cache"

	// ForceHandlersFlag run handlers even if a task fails
	ForceHandlersFlag = "--force-handlers"

	// ListTagsFlag is the list tags flag for ansible-playbook
	ListTagsFlag = "--list-tags"

	// ListTasksFlag is the list tasks flag for ansible-playbook
	ListTasksFlag = "--list-tasks"

	// SkipTagsFlag only run plays and tasks whose tags do not match these values
	SkipTagsFlag = "--skip-tags"

	// StartAtTaskFlag start the playbook at the task matching this name
	StartAtTaskFlag = "--start-at-task"

	// StepFlag one-step-at-a-time: confirm each task before running
	StepFlag = "--step"

	// TagsFlag is the tags flag for ansible-playbook
	TagsFlag = "--tags"
)

// AnsiblePlaybookOptionsFunc is a function to set executor options
type AnsiblePlaybookOptionsFunc func(*AnsiblePlaybookCmd)

// AnsiblePlaybookCmd object is the main object which defines the `ansible-playbook` command and how to execute it.
type AnsiblePlaybookCmd struct {
	// Ansible binary file
	Binary string
	// Exec is the executor item
	Exec execute.Executor
	// Playbook is the ansible's playbook name to be used
	Playbook string
	// Options are the ansible's playbook options
	AnsiblePlaybookOptions *AnsiblePlaybookOptions
	// ConnectionOptions are the ansible's playbook specific options for connection
	ConnectionOptions *options.AnsibleConnectionOptions
	// PrivilegeEscalationOptions are the ansible's playbook privilage escalation options
	PrivilegeEscalationOptions *options.AnsiblePrivilegeEscalationOptions
	// StdoutCallback defines which is the stdout callback method. By default is used 'default' method. Supported stdout method by go-ansible are: debug, default, dense, json, minimal, null, oneline, stderr, timer, yaml
	StdoutCallback string
}

// Run method runs the ansible-playbook
func (p *AnsiblePlaybookCmd) Run(ctx context.Context) error {
	var err error
	var command []string
	options := []execute.ExecuteOptions{}

	if p == nil {
		return errors.New("(ansible:Run)", "AnsiblePlaybookCmd is nil")
	}

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsiblePlaybookBinary
	}

	_, err = exec.LookPath(p.Binary)
	if err != nil {
		return errors.New("(ansible:Run)", fmt.Sprintf("Binary file '%s' does not exists", p.Binary), err)
	}

	// Define a default executor when it is not defined on AnsiblePlaybookCmd
	if p.Exec == nil {
		p.Exec = execute.NewDefaultExecute()
	}

	// Configure StdoutCallback method. By default is used ansible's 'default' callback method
	stdoutcallback.AnsibleStdoutCallbackSetEnv(p.StdoutCallback)

	// Generate the command to be run
	command, err = p.Command()
	if err != nil {
		return errors.New("(ansible:Run)", fmt.Sprintf("Error running '%s'", p.String()), err)
	}

	// Execute the command an return
	return p.Exec.Execute(ctx, command, stdoutcallback.GetResultsFunc(p.StdoutCallback), options...)
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
	if p.AnsiblePlaybookOptions != nil {
		options, err := p.AnsiblePlaybookOptions.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command)", "Error creating options", err)
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Determine the connection options to be set
	if p.ConnectionOptions != nil {
		options, err := p.ConnectionOptions.GenerateCommandConnectionOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command)", "Error creating connection options", err)
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Determine the privilege escalation options to be set
	if p.PrivilegeEscalationOptions != nil {
		options, err := p.PrivilegeEscalationOptions.GenerateCommandPrivilegeEscalationOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command)", "Error creating privilege escalation options", err)
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Include the ansible playbook
	cmd = append(cmd, p.Playbook)

	return cmd, nil
}

// String returns AnsiblePlaybookCmd as string
func (p *AnsiblePlaybookCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsiblePlaybookBinary
	}

	str := p.Binary

	if p.AnsiblePlaybookOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.AnsiblePlaybookOptions.String())
	}
	if p.ConnectionOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.ConnectionOptions.String())
	}
	if p.PrivilegeEscalationOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.PrivilegeEscalationOptions.String())
	}

	str = fmt.Sprintf("%s %s", str, p.Playbook)

	return str
}

// AnsiblePlaybookOptions object has those parameters described on `Options` section within ansible-playbook's man page, and which defines which should be the ansible-playbook execution behavior.
type AnsiblePlaybookOptions struct {
	*options.AnsibleCommonOptions

	// FlushCache is the flush cache flag for ansible-playbook
	FlushCache bool

	// ForceHandlers run handlers even if a task fails
	ForceHandlers bool

	// ListTags is the list tags flag for ansible-playbook
	ListTags bool

	// ListTasks is the list tasks flag for ansible-playbook
	ListTasks bool

	// SkipTags only run plays and tasks whose tags do not match these values
	SkipTags string

	// StartAtTask start the playbook at the task matching this name
	StartAtTask string

	// Step one-step-at-a-time: confirm each task before running
	Step bool

	// Tags is the tags flag for ansible-playbook
	Tags string
}

// GenerateCommandOptions return a list of options flags to be used on ansible-playbook execution
func (o *AnsiblePlaybookOptions) GenerateCommandOptions() ([]string, error) {

	var err error
	cmd := []string{}

	if o == nil {
		return nil, errors.New("(ansible::GenerateCommandOptions)", "AnsiblePlaybookOptions is nil")
	}

	if o.AnsibleCommonOptions != nil {
		cmd, err = o.AnsibleCommonOptions.GenerateCommandCommonOptions()
		if err != nil {
			fmt.Println("!!!!!!", err.Error())
			return nil, errors.New("(ansible::GenerateCommandOptions)", "Error generating command commond options", err)
		}
	}

	if o.FlushCache {
		cmd = append(cmd, FlushCacheFlag)
	}

	if o.ForceHandlers {
		cmd = append(cmd, ForceHandlersFlag)
	}

	if o.ListTags {
		cmd = append(cmd, ListTagsFlag)
	}

	if o.ListTasks {
		cmd = append(cmd, ListTasksFlag)
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

	if o.Tags != "" {
		cmd = append(cmd, TagsFlag)
		cmd = append(cmd, o.Tags)
	}

	return cmd, nil
}

// TODO: remove or options facade
// generateExtraVarsCommand return an string which is a json structure having all the extra variable
// func (o *AnsiblePlaybookOptions) generateExtraVarsCommand() (string, error) {

// 	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
// 	if err != nil {
// 		return "", errors.New("(ansible::generateExtraVarsCommand)", "Error creationg extra-vars JSON object to string", err)
// 	}
// 	return extraVars, nil
// }

// TODO: remove or options facade
// AddExtraVar registers a new extra variable on ansible-playbook options item
// func (o *AnsiblePlaybookOptions) AddExtraVar(name string, value interface{}) error {

// 	if o.ExtraVars == nil {
// 		o.ExtraVars = map[string]interface{}{}
// 	}
// 	_, exists := o.ExtraVars[name]
// 	if exists {
// 		return errors.New("(ansible::AddExtraVar)", fmt.Sprintf("ExtraVar '%s' already exist", name))
// 	}

// 	o.ExtraVars[name] = value

// 	return nil
// }

// String returns AnsiblePlaybookOptions as string
func (o *AnsiblePlaybookOptions) String() string {

	str := ""

	if o.AnsibleCommonOptions != nil {
		str = o.AnsibleCommonOptions.String()
	}

	if o.FlushCache {
		str = fmt.Sprintf("%s %s", str, FlushCacheFlag)
	}

	if o.ForceHandlers {
		str = fmt.Sprintf("%s %s", str, ForceHandlersFlag)
	}

	if o.ListTags {
		str = fmt.Sprintf("%s %s", str, ListTagsFlag)
	}

	if o.ListTasks {
		str = fmt.Sprintf("%s %s", str, ListTasksFlag)
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

	if o.Tags != "" {
		str = fmt.Sprintf("%s %s %s", str, TagsFlag, o.Tags)
	}

	return str
}

// Options set the command options to ansible-playbook
func (p *AnsiblePlaybookCmd) Options(options ...AnsiblePlaybookOptionsFunc) {

	if p.AnsiblePlaybookOptions == nil {
		p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
	}

	for _, opt := range options {
		opt(p)
	}
}

// WithFlushCache set FlushCache ansible-playbook options value
func WithFlushCache(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.FlushCache = value
	}
}

// WithForceHandlers set ForceHandlers ansible-playbook options value
func WithForceHandlers(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.ForceHandlers = value
	}
}

// WithListTags set ListTags ansible-playbook options value
func WithListTags(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.ListTags = value
	}
}

// WithListTasks set ListTasks ansible-playbook options value
func WithListTasks(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.ListTasks = value
	}
}

// WithSkipTags set SkipTags ansible-playbook options value
func WithSkipTags(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.SkipTags = value
	}
}

// WithStartAtTask set StartAtTask ansible-playbook options value
func WithStartAtTask(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.StartAtTask = value
	}
}

// WithStep set Step ansible-playbook options value
func WithStep(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.Step = value
	}
}

// WithTags set Tags ansible-playbook options value
func WithTags(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		p.AnsiblePlaybookOptions.Tags = value
	}
}

// WithAskVaultPassword set AskVaultPassword ansible-playbook common options value
func WithAskVaultPassword(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.AskVaultPassword = value
	}
}

// WithCheck set Check ansible-playbook common options value
func WithCheck(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Check = value
	}
}

// WithDiff set Diff ansible-playbook common options value
func WithDiff(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Diff = value
	}
}

// WithExtraVars set ExtraVars ansible-playbook common options value
func WithExtraVars(value map[string]interface{}) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.ExtraVars = value
	}
}

// WithForks set Forks ansible-playbook common options value
func WithForks(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Forks = value
	}
}

// WithInventory set Inventory ansible-playbook common options value
func WithInventory(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Inventory = value
	}
}

// WithLimit set Limit ansible-playbook common options value
func WithLimit(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Limit = value
	}
}

// WithListHosts set ListHosts ansible-playbook common options value
func WithListHosts(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.ListHosts = value
	}
}

// WithModulePath set ModulePath ansible-playbook common options value
func WithModulePath(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.ModulePath = value
	}
}

// WithSyntaxCheck set SyntaxCheck ansible-playbook common options value
func WithSyntaxCheck(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.SyntaxCheck = value
	}
}

// WithVaultID set VaultID ansible-playbook common options value
func WithVaultID(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.VaultID = value
	}
}

// WithVaultPasswordFile set VaultPasswordFile ansible-playbook common options value
func WithVaultPasswordFile(value string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.VaultPasswordFile = value
	}
}

// WithVerbose set Verbose ansible-playbook common options value
func WithVerbose(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Verbose = value
	}
}

// WithVersion set Version ansible-playbook common options value
func WithVersion(value bool) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		if p.AnsiblePlaybookOptions == nil {
			p.AnsiblePlaybookOptions = &AnsiblePlaybookOptions{}
		}
		if p.AnsiblePlaybookOptions.AnsibleCommonOptions == nil {
			p.AnsiblePlaybookOptions.AnsibleCommonOptions = &options.AnsibleCommonOptions{}
		}

		p.AnsiblePlaybookOptions.AnsibleCommonOptions.Version = value
	}
}
