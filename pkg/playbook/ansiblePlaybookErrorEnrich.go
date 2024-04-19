package playbook

import (
	"github.com/pkg/errors"
)

const (
	// TODO: error management
	// AnsiblePlaybookErrorCodeGeneralError is the error code for a general error
	AnsiblePlaybookErrorCodeGeneralError = 1
	// AnsiblePlaybookErrorCodeOneOrMoreHostFailed is the error code for a one or more host failed
	AnsiblePlaybookErrorCodeOneOrMoreHostFailed = 2
	// AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable is the error code for a one or more host unreachable
	AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable = 3
	// AnsiblePlaybookErrorCodeParserError is the error code for a parser error
	AnsiblePlaybookErrorCodeParserError = 4
	// AnsiblePlaybookErrorCodeBadOrIncompleteOptions is the error code for a bad or incomplete options
	AnsiblePlaybookErrorCodeBadOrIncompleteOptions = 5
	// AnsiblePlaybookErrorCodeUserInterruptedExecution is the error code for a user interrupted execution
	AnsiblePlaybookErrorCodeUserInterruptedExecution = 99
	// AnsiblePlaybookErrorCodeUnexpectedError is the error code for a unexpected error
	AnsiblePlaybookErrorCodeUnexpectedError = 250

	// AnsiblePlaybookErrorMessageGeneralError is the error message for a general error
	AnsiblePlaybookErrorMessageGeneralError = "ansible-playbook error: general error"
	// AnsiblePlaybookErrorMessageOneOrMoreHostFailed is the error message for a one or more host failed
	AnsiblePlaybookErrorMessageOneOrMoreHostFailed = "ansible-playbook error: one or more host failed"
	// AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable is the error message for a one or more host unreachable
	AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable = "ansible-playbook error: one or more host unreachable"
	// AnsiblePlaybookErrorMessageParserError is the error message for a parser error
	AnsiblePlaybookErrorMessageParserError = "ansible-playbook error: parser error"
	// AnsiblePlaybookErrorMessageBadOrIncompleteOptions is the error message for a bad or incomplete options
	AnsiblePlaybookErrorMessageBadOrIncompleteOptions = "ansible-playbook error: bad or incomplete options"
	// AnsiblePlaybookErrorMessageUserInterruptedExecution is the error message for a user interrupted execution
	AnsiblePlaybookErrorMessageUserInterruptedExecution = "ansible-playbook error: user interrupted execution"
	// AnsiblePlaybookErrorMessageUnexpectedError is the error message for a unexpected error
	AnsiblePlaybookErrorMessageUnexpectedError = "ansible-playbook error: unexpected error"
)

// AnsiblePlaybookErrorEnrich is an error enricher for ansible-playbook errors
type AnsiblePlaybookErrorEnrich struct{}

// NewAnsiblePlaybookErrorEnrich creates a new AnsiblePlaybookErrorEnrich instance
func NewAnsiblePlaybookErrorEnrich() *AnsiblePlaybookErrorEnrich {
	return &AnsiblePlaybookErrorEnrich{}
}

// Enrich return an error enriched with ansible-playbook error information
func (e *AnsiblePlaybookErrorEnrich) Enrich(err error) error {

	var errorMessage string

	_, hasExitCode := err.(ExitCodeErrorer)

	if hasExitCode {
		switch err.(ExitCodeErrorer).ExitCode() {
		case AnsiblePlaybookErrorCodeGeneralError:
			errorMessage = AnsiblePlaybookErrorMessageGeneralError
		case AnsiblePlaybookErrorCodeOneOrMoreHostFailed:
			errorMessage = AnsiblePlaybookErrorMessageOneOrMoreHostFailed
		case AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable:
			errorMessage = AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable
		case AnsiblePlaybookErrorCodeParserError:
			errorMessage = AnsiblePlaybookErrorMessageParserError
		case AnsiblePlaybookErrorCodeBadOrIncompleteOptions:
			errorMessage = AnsiblePlaybookErrorMessageBadOrIncompleteOptions
		case AnsiblePlaybookErrorCodeUserInterruptedExecution:
			errorMessage = AnsiblePlaybookErrorMessageUserInterruptedExecution
		case AnsiblePlaybookErrorCodeUnexpectedError:
			errorMessage = AnsiblePlaybookErrorMessageUnexpectedError
		}
	}

	return errors.Wrap(err, errorMessage)
}
