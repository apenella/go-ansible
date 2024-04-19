package playbook

import (
	"fmt"
	"testing"

	"github.com/apenella/go-ansible/v2/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEnrich(t *testing.T) {

	tests := []struct {
		desc     string
		err      error
		expected string
	}{
		{
			desc: "Testing enrich with a ansible-playbook general error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeGeneralError,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageGeneralError, "error cause"),
		},
		{
			desc: "Testing enrich with a ansible-playbook hosts failed error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeOneOrMoreHostFailed,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageOneOrMoreHostFailed, "error cause"),
		},
		{
			desc: "Testing enrich with a ansible-playbook hosts unreachable error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable, "error cause"),
		},
		{
			desc: "Testing enrich with a ansible-playbook parser error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeParserError,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageParserError, "error cause"),
		},
		{
			desc: "Testing enrich with a ansible-playbook bad or incomplete options error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeBadOrIncompleteOptions,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageBadOrIncompleteOptions, "error cause"),
		},
		{
			desc: "Testing enrich with a ansible-playbook user interrupted execution error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeUserInterruptedExecution,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageUserInterruptedExecution, "error cause"),
		},
		{
			desc: "Testing enrich with a ansible-playbook unexpected error",
			err: &mocks.MockExitCodeErr{
				Code:    AnsiblePlaybookErrorCodeUnexpectedError,
				Message: "error cause",
			},
			expected: fmt.Sprintf("%s: %s", AnsiblePlaybookErrorMessageUnexpectedError, "error cause"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := NewAnsiblePlaybookErrorEnrich()
			err := e.Enrich(test.err)
			assert.Equal(t, test.expected, err.Error())
		})
	}
}
