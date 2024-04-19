package playbook

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
)

const (
	// DefaultAnsiblePlaybookBinary is the ansible-playbook binary file default value
	DefaultAnsiblePlaybookBinary = "ansible-playbook"
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
}

// NewAnsiblePlaybookCmd creates a new AnsiblePlaybookCmd instance
func NewAnsiblePlaybookCmd(options ...AnsiblePlaybookOptionsFunc) *AnsiblePlaybookCmd {
	cmd := &AnsiblePlaybookCmd{}

	for _, option := range options {
		option(cmd)
	}

	return cmd
}

// WithBinary set the ansible-playbook binary file
func WithBinary(binary string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		p.Binary = binary
	}
}

// WithPlaybookOptions set the ansible-playbook options
func WithPlaybookOptions(options *AnsiblePlaybookOptions) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		p.PlaybookOptions = options
	}
}

// WithPlaybooks set the ansible-playbook playbooks
func WithPlaybooks(playbooks ...string) AnsiblePlaybookOptionsFunc {
	return func(p *AnsiblePlaybookCmd) {
		p.Playbooks = append([]string{}, playbooks...)
	}
}

// Command generate the ansible-playbook command which will be executed
func (p *AnsiblePlaybookCmd) Command() ([]string, error) {
	cmd := []string{}

	if len(p.Playbooks) == 0 {
		return nil, errors.New("(playbook::Command)", "No playbooks defined")
	}

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

	// Include the ansible playbook
	for _, playbook := range p.Playbooks {
		str = fmt.Sprintf("%s %s", str, playbook)
	}

	return str
}
