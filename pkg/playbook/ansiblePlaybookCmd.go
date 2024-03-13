package playbook

import (
	"fmt"

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
