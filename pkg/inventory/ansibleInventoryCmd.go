package inventory

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
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

	// LimitFlag further limit selected hosts to an additional pattern
	LimitFlag = "--limit"

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
	// Pattern is the ansible's host or group pattern
	Pattern string
	// Options are the ansible's inventory options
	InventoryOptions *AnsibleInventoryOptions
}

// NewAnsibleInventoryCmd creates a new AnsibleInventoryCmd instance
func NewAnsibleInventoryCmd(options ...AnsibleInventoryOptionFunc) *AnsibleInventoryCmd {
	cmd := &AnsibleInventoryCmd{}

	for _, option := range options {
		option(cmd)
	}

	return cmd
}

// WithBinary set the ansible-inventory binary file
func WithBinary(binary string) AnsibleInventoryOptionFunc {
	return func(p *AnsibleInventoryCmd) {
		p.Binary = binary
	}
}

// WithPattern set the adhoc pattern
func WithPattern(pattern string) AnsibleInventoryOptionFunc {
	return func(p *AnsibleInventoryCmd) {
		p.Pattern = pattern
	}
}

// WithInventoryOptions set the ansible-inventory options
func WithInventoryOptions(options *AnsibleInventoryOptions) AnsibleInventoryOptionFunc {
	return func(p *AnsibleInventoryCmd) {
		p.InventoryOptions = options
	}
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

	// set the pattern
	cmd = append(cmd, p.Pattern)

	// Determine the options to be set
	if p.InventoryOptions != nil {
		options, err := p.InventoryOptions.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(inventory::Command)", "Error creating options", err)
		}
		cmd = append(cmd, options...)
	}

	return cmd, nil
}

// String returns AnsibleInventoryCmd as string
func (p *AnsibleInventoryCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = DefaultAnsibleInventoryBinary
	}

	str := p.Binary

	str = fmt.Sprintf("%s %s", str, p.Pattern)

	if p.InventoryOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.InventoryOptions.String())
	}

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

	// Limit further limit selected hosts to an additional pattern
	Limit string

	// List Output all hosts info, works as inventory script
	List bool

	// Output When doing –list, send the inventory to a file instead of to the screen
	Output string

	// PlaybookDir Since this tool does not use playbooks, use this as a substitute inventory directory.This sets the relative path for many features including roles/ group_vars/ etc.
	PlaybookDir string

	// Toml Use TOML format instead of default JSON, ignored for –graph
	Toml bool

	// Vars Add vars to graph display, ignored unless used with –graph
	Vars bool

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

	if o.Limit != "" {
		cmd = append(cmd, LimitFlag)
		cmd = append(cmd, o.Limit)
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

	if o.Vars {
		cmd = append(cmd, VarsFlag)
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

	if o.Host != "" {
		str = fmt.Sprintf("%s %s %s", str, HostFlag, o.Host)
	}

	if o.Inventory != "" {
		str = fmt.Sprintf("%s %s %s", str, InventoryFlag, o.Inventory)
	}

	if o.Limit != "" {
		str = fmt.Sprintf("%s %s %s", str, LimitFlag, o.Limit)
	}

	if o.List {
		str = fmt.Sprintf("%s %s", str, ListFlag)
	}

	if o.Output != "" {
		str = fmt.Sprintf("%s %s %s", str, OutputFlag, o.Output)
	}

	if o.PlaybookDir != "" {
		str = fmt.Sprintf("%s %s %s", str, PlaybookDirFlag, o.PlaybookDir)
	}

	if o.Toml {
		str = fmt.Sprintf("%s %s", str, TomlFlag)
	}

	if o.Vars {
		str = fmt.Sprintf("%s %s", str, VarsFlag)
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
