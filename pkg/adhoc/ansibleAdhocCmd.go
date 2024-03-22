package adhoc

import (
	"fmt"

	errors "github.com/apenella/go-common-utils/error"
)

const (

	// DefaultAnsibleAdhocBinary is the default value for ansible binary file to run adhoc modules
	DefaultAnsibleAdhocBinary = "ansible"
)

// AnsibleAdhocOptionsFunc is a function to set executor options
type AnsibleAdhocOptionsFunc func(*AnsibleAdhocCmd)

// AnsibleAdhocCmd object is the main object which defines the `ansible` adhoc command and how to execute it.
type AnsibleAdhocCmd struct {
	// Ansible binary file
	Binary string
	// Pattern is the ansible's host pattern
	Pattern string
	// AdhocOptions are the ansible's playbook options
	AdhocOptions *AnsibleAdhocOptions
}

// NewAnsibleAdhocCmd creates a new AnsibleAdhocCmd instance
func NewAnsibleAdhocCmd(options ...AnsibleAdhocOptionsFunc) *AnsibleAdhocCmd {
	cmd := &AnsibleAdhocCmd{}

	for _, option := range options {
		option(cmd)
	}

	return cmd
}

// WithBinary set the ansible binary file
func WithBinary(binary string) AnsibleAdhocOptionsFunc {
	return func(a *AnsibleAdhocCmd) {
		a.Binary = binary
	}
}

// WithPattern set the adhoc pattern
func WithPattern(pattern string) AnsibleAdhocOptionsFunc {
	return func(a *AnsibleAdhocCmd) {
		a.Pattern = pattern
	}
}

// WithAdhocOptions set the adhoc options
func WithAdhocOptions(options *AnsibleAdhocOptions) AnsibleAdhocOptionsFunc {
	return func(a *AnsibleAdhocCmd) {
		a.AdhocOptions = options
	}
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
	if a.AdhocOptions != nil {
		options, err := a.AdhocOptions.GenerateAnsibleAdhocOptions()
		if err != nil {
			return nil, errors.New("(adhoc::Command)", "Error creating options", err)
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

	if a.AdhocOptions != nil {
		str = fmt.Sprintf("%s %s", str, a.AdhocOptions.String())
	}

	return str
}
