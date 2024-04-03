package galaxyroleinstall

import (
	"fmt"

	galaxy "github.com/apenella/go-ansible/v2/pkg/galaxy"
	galaxyrole "github.com/apenella/go-ansible/v2/pkg/galaxy/role"
)

const (
	// // DefaultAnsibleGalaxyRoleInstallBinary is the ansible-galaxy binary file default value
	// DefaultAnsibleGalaxyRoleInstallBinary = "ansible-galaxy"

	// // AnsibleGalaxyRoleSubCommand is the ansible-galaxy role subcommand
	// AnsibleGalaxyRoleSubCommand = "role"

	// AnsibleGalaxyRoleInstallSubCommand is the ansible-galaxy role install subcommand
	AnsibleGalaxyRoleInstallSubCommand = "install"
)

// AnsibleGalaxyRoleInstallOptionsFunc is a function to set executor options
type AnsibleGalaxyRoleInstallOptionsFunc func(*AnsibleGalaxyRoleInstallCmd)

// AnsibleGalaxyRoleInstallCmd object is the main object which defines the `ansible-galaxy` command to install roles.
type AnsibleGalaxyRoleInstallCmd struct {
	// Binary is the ansible-galaxy binary file
	Binary string

	// RoleNames is the ansible-galaxy's role names to be installed
	RoleNames []string

	// GalaxyRoleInstallOptions are the ansible-galaxy's role install options
	GalaxyRoleInstallOptions *AnsibleGalaxyRoleInstallOptions
}

// NewAnsibleGalaxyRoleInstallCmd creates a new AnsibleGalaxyRoleInstallCmd instance
func NewAnsibleGalaxyRoleInstallCmd(options ...AnsibleGalaxyRoleInstallOptionsFunc) *AnsibleGalaxyRoleInstallCmd {
	cmd := &AnsibleGalaxyRoleInstallCmd{}

	for _, option := range options {
		option(cmd)
	}

	return cmd
}

// WithBinary set the ansible-galaxy binary file
func WithBinary(binary string) AnsibleGalaxyRoleInstallOptionsFunc {
	return func(p *AnsibleGalaxyRoleInstallCmd) {
		p.Binary = binary
	}
}

// WithGalaxyRoleInstallOptions set the ansible-galaxy role install options
func WithGalaxyRoleInstallOptions(options *AnsibleGalaxyRoleInstallOptions) AnsibleGalaxyRoleInstallOptionsFunc {
	return func(p *AnsibleGalaxyRoleInstallCmd) {
		p.GalaxyRoleInstallOptions = options
	}
}

// WithRoleNames set the ansible-galaxy role names
func WithRoleNames(roleNames ...string) AnsibleGalaxyRoleInstallOptionsFunc {
	return func(p *AnsibleGalaxyRoleInstallCmd) {
		p.RoleNames = append([]string{}, roleNames...)
	}
}

// Command generate the ansible-galaxy role install command which will be executed
func (p *AnsibleGalaxyRoleInstallCmd) Command() ([]string, error) {
	cmd := []string{}

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = galaxy.DefaultAnsibleGalaxyBinary
	}

	cmd = append(cmd, p.Binary, galaxyrole.AnsibleGalaxyRoleSubCommand, AnsibleGalaxyRoleInstallSubCommand)

	// Add the options
	if p.GalaxyRoleInstallOptions != nil {
		options, err := p.GalaxyRoleInstallOptions.GenerateCommandOptions()
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, options...)
	}

	// Add the role names
	cmd = append(cmd, p.RoleNames...)

	return cmd, nil
}

// String returns the ansible-galaxy role install command as a string
func (p *AnsibleGalaxyRoleInstallCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = galaxy.DefaultAnsibleGalaxyBinary
	}

	str := fmt.Sprintf("%s %s %s", p.Binary, galaxyrole.AnsibleGalaxyRoleSubCommand, AnsibleGalaxyRoleInstallSubCommand)

	if p.GalaxyRoleInstallOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.GalaxyRoleInstallOptions.String())
	}

	// Include the role names
	for _, roleName := range p.RoleNames {
		str = fmt.Sprintf("%s %s", str, roleName)
	}

	return str
}
