package ansibler

import (
	"fmt"
	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/stdoutcallback"
	common "github.com/apenella/go-common-utils/data"
	errors "github.com/apenella/go-common-utils/error"
	"io"
	"os/exec"
)

const (
	// ModuleNameFlag is the module-name flag for ansible adhoc
	ModuleNameFlag = "--module-name"

	// ModuleArgsFlag is module argument flag ansible adhoc
	ModuleArgsFlag = "--args"
)

// AnsibleAdhocCmd object is the main object which defines the `ansible` command and how to execute it.
type AnsibleAdhocCmd struct {
	// Ansible binary file
	Binary string
	// Exec is the executor item
	Exec Executor
	// ExecPrefix is a text that is set at the beginning of each execution line
	ExecPrefix string
	// Pattern is the ansible's host patterns
	Pattern string
	// Options are the ansible's playbook options
	Options *AnsibleAdhocOptions
	// ConnectionOptions are the ansible's playbook specific options for connection
	ConnectionOptions *AnsibleConnectionOptions
	// PrivilegeEscalationOptions are the ansible's playbook privilage escalation options
	PrivilegeEscalationOptions *AnsiblePrivilegeEscalationOptions
	// StdoutCallback defines which is the stdout callback method. By default is used 'default' method. Supported stdout method by go-ansible are: debug, default, dense, json, minimal, null, oneline, stderr, timer, yaml
	StdoutCallback string
	// Writer manages the output
	Writer io.Writer
}

// Run method runs the ansible-playbook
func (a *AnsibleAdhocCmd) Run() error {
	var err error
	var cmd []string

	if a == nil {
		return errors.New("(ansible:Run)", "AnsiblePlaybookCmd is nil")
	}

	// Use default binary when it is not already defined
	if a.Binary == "" {
		a.Binary = DefaultAnsibleBinary
	}

	_, err = exec.LookPath(a.Binary)
	if err != nil {
		return errors.New("(ansible:Run)", fmt.Sprintf("Binary file '%s' does not exists", a.Binary), err)
	}

	// Define a default executor when it is not defined on AnsibleAdhocCmd
	if a.Exec == nil {
		a.Exec = &execute.DefaultExecute{
			Write:       a.Writer,
			ResultsFunc: stdoutcallback.GetResultsFunc(a.StdoutCallback),
		}
	}

	// Generate the command to be run
	cmd, err = a.Command()
	if err != nil {
		return errors.New("(ansible:Run)", fmt.Sprintf("Error running '%s'", a.String()), err)
	}

	// Set default prefix
	if len(a.ExecPrefix) <= 0 {
		a.ExecPrefix = ""
	}

	// Configure StdoutCallback method. By default is used ansible's 'default' callback method
	stdoutcallback.AnsibleStdoutCallbackSetEnv(a.StdoutCallback)

	// Execute the command an return
	return a.Exec.Execute(cmd[0], cmd[1:], a.ExecPrefix)
}

// Command generate the ansible-playbook command which will be executed
func (a *AnsibleAdhocCmd) Command() ([]string, error) {
	cmd := []string{}

	// Use default binary when it is not already defined
	if a.Binary == "" {
		a.Binary = DefaultAnsiblePlaybookBinary
	}

	// Set the ansible-playbook binary file
	cmd = append(cmd, a.Binary)

	// Determine the options to be set
	if a.Options != nil {
		options, err := a.Options.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command)", "Error creating options", err)
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Determine the connection options to be set
	if a.ConnectionOptions != nil {
		options, err := a.ConnectionOptions.GenerateCommandConnectionOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command)", "Error creating connection options", err)
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Determine the privilege escalation options to be set
	if a.PrivilegeEscalationOptions != nil {
		options, err := a.PrivilegeEscalationOptions.GenerateCommandPrivilegeEscalationOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command)", "Error creating privilege escalation options", err)
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	// Include the ansible host pattern
	cmd = append(cmd, a.Pattern)

	return cmd, nil
}

// String returns AnsibleAdhocCmd as string
func (p *AnsibleAdhocCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsibleBinary
	}

	str := p.Binary

	if p.Options != nil {
		str = fmt.Sprintf("%s %s", str, p.Options.String())
	}
	if p.ConnectionOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.ConnectionOptions.String())
	}
	if p.PrivilegeEscalationOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.PrivilegeEscalationOptions.String())
	}

	str = fmt.Sprintf("%s %s", str, p.Pattern)

	return str
}

// AnsibleAdhocOptions object has those parameters described on `Options` section within ansible's man page, and which defines which should be the ansible execution behavior.
type AnsibleAdhocOptions struct {
	// Module name is the module name that will be executed
	ModuleName string
	// ModulesArgs are the arguments passed to the module
	ModulesArgs string
	// ExtraVars is a map of extra variables used on ansible-playbook execution
	ExtraVars map[string]interface{}
	// Inventory specify inventory host path
	Inventory string
	// Limit is selected hosts additional pattern
	Limit string
	// VaultPasswordFile path to the file holding vault decryption key
	VaultPasswordFile string
}

// GenerateCommandOptions return a list of options flags to be used on ansible execution
func (o *AnsibleAdhocOptions) GenerateCommandOptions() ([]string, error) {
	cmd := []string{}

	if o == nil {
		return nil, errors.New("(ansible::GenerateCommandOptions)", "AnsiblePlaybookOptions is nil")
	}

	if o.ModuleName != "" {
		cmd = append(cmd, ModuleNameFlag)
		cmd = append(cmd, o.ModuleName)
	}

	if o.ModulesArgs != "" {
		cmd = append(cmd, ModuleArgsFlag)
		cmd = append(cmd, o.ModulesArgs)
	}

	if o.Inventory != "" {
		cmd = append(cmd, InventoryFlag)
		cmd = append(cmd, o.Inventory)
	}

	if o.Limit != "" {
		cmd = append(cmd, LimitFlag)
		cmd = append(cmd, o.Limit)
	}

	if o.VaultPasswordFile != "" {
		cmd = append(cmd, VaultPasswordFileFlag)
		cmd = append(cmd, o.VaultPasswordFile)
	}

	if len(o.ExtraVars) > 0 {
		cmd = append(cmd, ExtraVarsFlag)
		extraVars, err := o.generateExtraVarsCommand()
		if err != nil {
			return nil, errors.New("(ansible::GenerateCommandOptions)", "Error generating extra-vars", err)
		}
		cmd = append(cmd, extraVars)
	}

	return cmd, nil
}

// generateExtraVarsCommand return an string which is a json structure having all the extra variable
func (o *AnsibleAdhocOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := common.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(ansible::generateExtraVarsCommand)", "Error creationg extra-vars JSON object to string", err)
	}
	return extraVars, nil
}

// AddExtraVar registers a new extra variable on ansible options item
func (o *AnsibleAdhocOptions) AddExtraVar(name string, value interface{}) error {

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

// String returns AnsibleAdhocOptions as string
func (o *AnsibleAdhocOptions) String() string {
	str := ""

	if o.Inventory != "" {
		str = fmt.Sprintf("%s %s %s", str, InventoryFlag, o.Inventory)
	}

	if o.Limit != "" {
		str = fmt.Sprintf("%s %s %s", str, LimitFlag, o.Limit)
	}

	if len(o.ExtraVars) > 0 {
		extraVars, _ := o.generateExtraVarsCommand()
		str = fmt.Sprintf("%s %s %s", str, ExtraVarsFlag, fmt.Sprintf("'%s'", extraVars))
	}

	return str
}
