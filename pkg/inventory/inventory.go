package inventory

import (
	"context"
	"fmt"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	common "github.com/apenella/go-common-utils/data"
	errors "github.com/apenella/go-common-utils/error"
	"os/exec"
)

const (
	// DefaultAnsibleInventoryBinary is the default value for ansible binary file to run inventory modules
	DefaultAnsibleInventoryBinary = "ansible-inventory"

	// AskVaultPasswordFlag ask for vault password
	AskVaultPasswordFlag = "--ask-vault-password"

	// ExportFlag When doing an –list, represent in a way that is optimized for export, not as an accurate representation of how Ansible has processed it
	ExportFlag = "--export"

	// GraphFlag create inventory graph, if supplying pattern it must be a valid group name
	GraphFlag = "--graph"

	// HostFlag Output specific host info, works as inventory script
	HostFlag = "--host"

	// InventoryFlag is the inventory flag for ansible-inventory
	InventoryFlag = "--inventory"

	// ListFlag Output all hosts info, works as inventory script
	ListFlag = "--list"

	// OutputFlag When doing –list, send the inventory to a file instead of to the screen
	OutputFlag = "--output"

	// PlaybookDirFlag Since this tool does not use playbooks, use this as a substitute inventory directory. This sets the relative path for many features including roles/ group_vars/ etc.
	PlaybookDirFlag = "--playbook-dir"

	// TomlFlag Use TOML format instead of default JSON, ignored for –graph
	TomlFlag = "--toml"

	// VarsFlag Add vars to graph display, ignored unless used with –graph
	VarsFlag = "--vars"

	// ValutIdFlag the vault identity to use
	VaultIdFlag = "--vault-id"

	// ValutPasswordFileFlag vault password file
	VaultPasswordFileFlag = "--vault-password-file"

	// VerboseFlag verbose with -vvvv is enabled
	VerboseFlag = "-vvvv"

	// VerboseVFlag verbose with -v is enabled
	VerboseVFlag = "-v"

	// VerboseVVFlag verbose with -vv is enabled
	VerboseVVFlag = "-vv"

	// VerboseVVVFlag verbose with -vvv is enabled
	VerboseVVVFlag = "-vvv"

	// VerboseVVVVFlag verbose with -vvvv is enabled
	VerboseVVVVFlag = "-vvvv"

	// VersionFlag show program’s version number, config file location, configured module search path, module location, executable location and exit
	VersionFlag = "--version"

	// YamlFlag Use YAML format instead of default JSON, ignored for –graph
	YamlFlag = "--yaml"
)

// AnsibleInventoryOptionFunc is a function to set executor options
type AnsibleInventoryOptionFunc func(*AnsibleInventoryCmd)

// AnsibleInventoryCmd object is the main object which defines the `ansible-inventory` inventory command and how to execute it.
type AnsibleInventoryCmd struct {
	// Ansible-inventory binary file
	Binary string
	// Exec is the executor item
	Exec execute.Executor
	// Pattern is the ansible's group pattern
	Pattern string
	// Options are the ansible's inventory options
	Options *AnsibleInventoryOptions
	// StdoutCallback defines which is the stdout callback method. By default is used 'default' method. Supported stdout method by go-ansible are: debug, default, dense, json, minimal, null, oneline, stderr, timer, yaml
	StdoutCallback string
}

// Run method runs the ansible-inventory
func (p *AnsibleInventoryCmd) Run(ctx context.Context) error {
	var err error
	var command []string
	options := []execute.ExecuteOptions{}

	if p == nil {
		return errors.New("(inventory::Run)", "AnsibleInventoryCmd is nil")
	}

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsibleInventoryBinary
	}

	_, err = exec.LookPath(p.Binary)
	if err != nil {
		return errors.New("(inventory::Run)", fmt.Sprintf("Binary file '%s' does not exists", p.Binary), err)
	}

	// Define a default executor when it is not defined on AnsibleInventoryCmd
	if p.Exec == nil {
		p.Exec = execute.NewDefaultExecute()
	}

	// Configure StdoutCallback method. By default is used ansible's 'default' callback method
	stdoutcallback.AnsibleStdoutCallbackSetEnv(p.StdoutCallback)

	// Generate the command to be run
	command, err = p.Command()
	if err != nil {
		return errors.New("(inventory::Run)", fmt.Sprintf("Error running '%s'", p.String()), err)
	}

	// Execute the command an return
	return p.Exec.Execute(ctx, command, stdoutcallback.GetResultsFunc(p.StdoutCallback), options...)
}

// Command generate the ansible command which will be executed
func (p *AnsibleInventoryCmd) Command() ([]string, error) {
	cmd := []string{}

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsibleInventoryBinary
	}

	// Set the ansible-inventory binary file
	cmd = append(cmd, p.Binary)

	// Determine the options to be set
	if p.Options != nil {
		options, err := p.Options.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(inventory::Command)", "Error creating options", err)
		}
		cmd = append(cmd, options...)
	}

	cmd = append(cmd, p.Pattern)

	return cmd, nil
}

// String returns AnsibleInventoryCmd as string
func (p *AnsibleInventoryCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsibleInventoryBinary
	}

	str := p.Binary

	if p.Options != nil {
		str = fmt.Sprintf("%s %s", str, p.Options.String())
	}

	str = fmt.Sprintf("%s %s", str, p.Pattern)

	return str
}

// AnsibleInventoryOptions object has those parameters described on `Options` section within ansible-inventory's man page, and which defines which should be the ansible-inventory execution behavior.
type AnsibleInventoryOptions struct {
	// AskVaultPassword ask for vault password
	AskVaultPassword bool

	// Export When doing an –list, represent in a way that is optimized for export,not as an accurate representation of how Ansible has processed it
	Export bool

	// Graph create inventory graph, if supplying pattern it must be a valid group name
	Graph bool

	// Host Output specific host info, works as inventory script
	Host string

	// Inventory is the inventory flag for ansible-inventory
	Inventory string

	// List Output all hosts info, works as inventory script
	List bool

	// Output When doing –list, send the inventory to a file instead of to the screen
	Output string

	// PlaybookDir Since this tool does not use playbooks, use this as a substitute inventory directory.This sets the relative path for many features including roles/ group_vars/ etc.
	PlaybookDir string

	// Toml Use TOML format instead of default JSON, ignored for –graph
	Toml bool

	// Vars Add vars to graph display, ignored unless used with –graph
	Vars map[string]interface{}

	// VarsFile is a list of files used to load vars
	VarsFile []string

	// VaultID the vault identity to use
	VaultID string

	// VaultPasswordFile vault password file
	VaultPasswordFile string

	// Verbose verbose mode enabled
	Verbose bool

	// VerboseV verbose with -v is enabled
	VerboseV bool

	// VerboseVV verbose with -vv is enabled
	VerboseVV bool

	// VerboseVVV verbose with -vvv is enabled
	VerboseVVV bool

	// VerboseVVVV verbose with -vvvv is enabled
	VerboseVVVV bool

	// Version show program’s version number, config file location, configured module search path, module location, executable location and exit
	Version bool

	// Yaml Use YAML format instead of default JSON, ignored for –graph
	Yaml bool
}

// GenerateCommandOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleInventoryOptions) GenerateCommandOptions() ([]string, error) {
	errContext := "(inventory::GenerateCommandOptions)"

	cmd := []string{}

	if o == nil {
		return nil, errors.New(errContext, "AnsibleInventoryOptions is nil")
	}

	if o.AskVaultPassword {
		cmd = append(cmd, AskVaultPasswordFlag)
	}

	if o.Export {
		cmd = append(cmd, ExportFlag)
	}

	if o.Graph {
		cmd = append(cmd, GraphFlag)
	}

	if o.Host != "" {
		cmd = append(cmd, HostFlag)
		cmd = append(cmd, o.Host)
	}

	if o.Inventory != "" {
		cmd = append(cmd, InventoryFlag)
		cmd = append(cmd, o.Inventory)
	}

	if o.List {
		cmd = append(cmd, ListFlag)
	}

	if o.Output != "" {
		cmd = append(cmd, OutputFlag)
		cmd = append(cmd, o.Output)
	}

	if o.PlaybookDir != "" {
		cmd = append(cmd, PlaybookDirFlag)
		cmd = append(cmd, o.PlaybookDir)
	}

	if o.Toml {
		cmd = append(cmd, TomlFlag)
	}

	if len(o.Vars) > 0 {
		cmd = append(cmd, VarsFlag)
		vars, err := o.generateVarsCommand()
		if err != nil {
			return nil, errors.New(errContext, "Error generating vars", err)
		}
		cmd = append(cmd, vars)
	}

	for _, file := range o.VarsFile {
		cmd = append(cmd, VarsFlag)
		cmd = append(cmd, file)
	}

	if o.VaultID != "" {
		cmd = append(cmd, VaultIdFlag)
		cmd = append(cmd, o.VaultID)
	}

	if o.VaultPasswordFile != "" {
		cmd = append(cmd, VaultPasswordFileFlag)
		cmd = append(cmd, o.VaultPasswordFile)
	}

	if o.Version {
		cmd = append(cmd, VersionFlag)
	}

	if o.Verbose {
		// Assuming there is a method to generate the correct verbosity flag
		verboseFlag, err := o.generateVerbosityFlag()
		if err != nil {
			return nil, errors.New(errContext, "", err)
		}
		if verboseFlag != "" {
			cmd = append(cmd, verboseFlag)
		}
	}

	if o.Yaml {
		cmd = append(cmd, YamlFlag)
	}

	return cmd, nil
}

// generateVerbosityFlag return a string with the verbose flag. Higher verbosity (more v's) has precedence over lower
func (o *AnsibleInventoryOptions) generateVerbosityFlag() (string, error) {
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

// generateVarsCommand return a string which is a json structure having all the variable
func (o *AnsibleInventoryOptions) generateVarsCommand() (string, error) {

	vars, err := common.ObjectToJSONString(o.Vars)
	if err != nil {
		return "", errors.New("(inventory::generateExtraVarsCommand)", "Error creationg vars JSON object to string", err)
	}

	return vars, nil
}

// AddVar registers a new variable
func (o *AnsibleInventoryOptions) AddVar(name string, value interface{}) error {

	if o.Vars == nil {
		o.Vars = map[string]interface{}{}
	}
	_, exists := o.Vars[name]
	if exists {
		return errors.New("(inventory::AddVar)", fmt.Sprintf("Var '%s' already exist", name))
	}

	o.Vars[name] = value

	return nil
}

// AddVarsFile adds an vars file on ansible-inventory options item
func (o *AnsibleInventoryOptions) AddVarsFile(file string) error {
	if o.VarsFile == nil {
		o.VarsFile = []string{}
	}

	if file[0] != '@' {
		file = fmt.Sprintf("@%s", file)
	}

	o.VarsFile = append(o.VarsFile, file)

	return nil
}

// AddVaultedExtraVar registers a new variable on ansible-inventory options item vaulting its value
func (o *AnsibleInventoryOptions) AddVaultedVar(vaulter Vaulter, name string, value string) error {

	if vaulter == nil {
		return errors.New("(inventory::AddVaultedVar)", "To define a vaulted var you need to initialize a vaulter")
	}

	if o.Vars == nil {
		o.Vars = map[string]interface{}{}
	}

	_, exists := o.Vars[name]
	if exists {
		return errors.New("(inventory::AddVaultedVar)", fmt.Sprintf("Var '%s' already exist", name))
	}

	vaultedValue, err := vaulter.Vault(value)
	if err != nil {
		return errors.New("(inventory::AddVaultedVar)", fmt.Sprintf("Variable '%s' can not be vaulted", name), err)
	}

	o.Vars[name] = vaultedValue

	return nil
}

// GenerateCommandCommonOptions return a list of command options flags to be used on ansible execution
func (o *AnsibleInventoryOptions) String() string {
	str := ""

	if o.AskVaultPassword {
		str = fmt.Sprintf("%s %s", str, AskVaultPasswordFlag)
	}

	if o.Export {
		str = fmt.Sprintf("%s %s", str, ExportFlag)
	}

	if o.Graph {
		str = fmt.Sprintf("%s %s", str, GraphFlag)
	}

	if o.Host != "nil" {
		str = fmt.Sprintf("%s %s %s", str, HostFlag, o.Host)
	}

	if o.Inventory != "" {
		str = fmt.Sprintf("%s %s %s", str, InventoryFlag, o.Inventory)
	}

	if o.List {
		str = fmt.Sprintf("%s %s", str, ListFlag)
	}

	if o.Output != "nil" {
		str = fmt.Sprintf("%s %s %s", str, OutputFlag, o.Output)
	}

	if o.PlaybookDir != "nil" {
		str = fmt.Sprintf("%s %s %s", str, PlaybookDirFlag, o.PlaybookDir)
	}

	if o.Toml {
		str = fmt.Sprintf("%s %s", str, TomlFlag)
	}

	if len(o.Vars) > 0 {
		vars, _ := o.generateVarsCommand()
		str = fmt.Sprintf("%s %s %s", str, VarsFlag, fmt.Sprintf("'%s'", vars))
	}

	for _, varsFile := range o.VarsFile {
		str = fmt.Sprintf("%s %s %s", str, VarsFlag, varsFile)
	}

	if o.VaultID != "" {
		str = fmt.Sprintf("%s %s %s", str, VaultIdFlag, o.VaultID)
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

	if o.Yaml {
		str = fmt.Sprintf("%s %s", str, YamlFlag)
	}

	return str
}
