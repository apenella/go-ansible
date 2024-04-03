package galaxycollectioninstall

import (
	"fmt"

	galaxy "github.com/apenella/go-ansible/v2/pkg/galaxy"
	galaxycollection "github.com/apenella/go-ansible/v2/pkg/galaxy/collection"
)

const (
	// AnsibleGalaxyCollectionInstallSubCommand is the ansible-galaxy collection install subcommand
	AnsibleGalaxyCollectionInstallSubCommand = "install"
)

// AnsibleGalaxyCollectionInstallOptionsFunc is a function to set executor options
type AnsibleGalaxyCollectionInstallOptionsFunc func(*AnsibleGalaxyCollectionInstallCmd)

// AnsibleGalaxyCollectionInstallCmd object is the main object which defines the `ansible-galaxy` command to install collections.
type AnsibleGalaxyCollectionInstallCmd struct {
	// Binary is the ansible-galaxy binary file
	Binary string

	// CollectionNames is the ansible-galaxy's collection names to be installed
	CollectionNames []string

	// GalaxyCollectionInstallOptions are the ansible-galaxy's collection install options
	GalaxyCollectionInstallOptions *AnsibleGalaxyCollectionInstallOptions
}

// NewAnsibleGalaxyCollectionInstallCmd creates a new AnsibleGalaxyCollectionInstallCmd instance
func NewAnsibleGalaxyCollectionInstallCmd(options ...AnsibleGalaxyCollectionInstallOptionsFunc) *AnsibleGalaxyCollectionInstallCmd {
	cmd := &AnsibleGalaxyCollectionInstallCmd{}

	for _, option := range options {
		option(cmd)
	}

	return cmd
}

// WithBinary set the ansible-galaxy binary file
func WithBinary(binary string) AnsibleGalaxyCollectionInstallOptionsFunc {
	return func(p *AnsibleGalaxyCollectionInstallCmd) {
		p.Binary = binary
	}
}

// WithGalaxyCollectionInstallOptions set the ansible-galaxy collection install options
func WithGalaxyCollectionInstallOptions(options *AnsibleGalaxyCollectionInstallOptions) AnsibleGalaxyCollectionInstallOptionsFunc {
	return func(p *AnsibleGalaxyCollectionInstallCmd) {
		p.GalaxyCollectionInstallOptions = options
	}
}

// WithCollectionNames set the ansible-galaxy role names
func WithCollectionNames(roleNames ...string) AnsibleGalaxyCollectionInstallOptionsFunc {
	return func(p *AnsibleGalaxyCollectionInstallCmd) {
		p.CollectionNames = append([]string{}, roleNames...)
	}
}

// Command generate the ansible-galaxy role install command which will be executed
func (p *AnsibleGalaxyCollectionInstallCmd) Command() ([]string, error) {
	cmd := []string{}

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = galaxy.DefaultAnsibleGalaxyBinary
	}

	cmd = append(cmd, p.Binary, galaxycollection.AnsibleGalaxyCollectionSubCommand, AnsibleGalaxyCollectionInstallSubCommand)

	// Add the options
	if p.GalaxyCollectionInstallOptions != nil {
		options, err := p.GalaxyCollectionInstallOptions.GenerateCommandOptions()
		if err != nil {
			return nil, err
		}
		cmd = append(cmd, options...)
	}

	// Add the role names
	cmd = append(cmd, p.CollectionNames...)

	return cmd, nil
}

// String returns the ansible-galaxy role install command as a string
func (p *AnsibleGalaxyCollectionInstallCmd) String() string {

	// Use default binary when it is not already defined
	if p.Binary == "" {
		p.Binary = galaxy.DefaultAnsibleGalaxyBinary
	}

	str := fmt.Sprintf("%s %s %s", p.Binary, galaxycollection.AnsibleGalaxyCollectionSubCommand, AnsibleGalaxyCollectionInstallSubCommand)

	if p.GalaxyCollectionInstallOptions != nil {
		str = fmt.Sprintf("%s %s", str, p.GalaxyCollectionInstallOptions.String())
	}

	// Include the role names
	for _, roleName := range p.CollectionNames {
		str = fmt.Sprintf("%s %s", str, roleName)
	}

	return str
}
