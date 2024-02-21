package adhoc

// generate test for func (e *AnsibleAdhocExecute) WithBinary(binary string) *AnsibleAdhocExecute, func (e *AnsibleAdhocExecute) WithAdhocOptions(options *AnsibleAdhocOptions) *AnsibleAdhocExecute, func (e *AnsibleAdhocExecute) WithConnectionOptions(options *options.AnsibleConnectionOptions) *AnsibleAdhocExecute, func (e *AnsibleAdhocExecute) WithPrivilegeEscalationOptions(options *options.AnsiblePrivilegeEscalationOptions) *AnsibleAdhocExecute

import (
	"testing"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/stretchr/testify/assert"
)

func TestNewAnsibleAdhocExecute(t *testing.T) {
	tests := []struct {
		desc     string
		pattern  string
		expected *AnsibleAdhocExecute
	}{
		{
			desc:    "Testing creating a new AnsibleAdhocExecute",
			pattern: "test1",
			expected: &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{
					Pattern: "test1",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := NewAnsibleAdhocExecute(test.pattern)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithBinary(t *testing.T) {
	tests := []struct {
		desc     string
		binary   string
		expected *AnsibleAdhocExecute
	}{
		{

			desc:   "Testing setting binary to AnsibleAdhocExecute",
			binary: "test1",
			expected: &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{
					Binary: "test1",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{},
			}

			res := e.WithBinary(test.binary)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithAdhocOptions(t *testing.T) {
	tests := []struct {
		desc     string
		options  *AnsibleAdhocOptions
		expected *AnsibleAdhocExecute
	}{
		{

			desc: "Testing setting adhoc options to AnsibleAdhocExecute",
			options: &AnsibleAdhocOptions{
				Inventory: "test1",
			},
			expected: &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{
					AdhocOptions: &AnsibleAdhocOptions{
						Inventory: "test1",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{},
			}

			res := e.WithAdhocOptions(test.options)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithConnectionOptions(t *testing.T) {
	tests := []struct {
		desc     string
		options  *options.AnsibleConnectionOptions
		expected *AnsibleAdhocExecute
	}{
		{

			desc: "Testing setting connection options to AnsibleAdhocExecute",
			options: &options.AnsibleConnectionOptions{
				Connection: "test1",
			},
			expected: &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{
					ConnectionOptions: &options.AnsibleConnectionOptions{
						Connection: "test1",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{},
			}

			res := e.WithConnectionOptions(test.options)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithPrivilegeEscalationOptions(t *testing.T) {
	tests := []struct {
		desc     string
		options  *options.AnsiblePrivilegeEscalationOptions
		expected *AnsibleAdhocExecute
	}{
		{

			desc: "Testing setting privilege escalation options to AnsibleAdhocExecute",
			options: &options.AnsiblePrivilegeEscalationOptions{
				Become: true,
			},
			expected: &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{
					PrivilegeEscalationOptions: &options.AnsiblePrivilegeEscalationOptions{
						Become: true,
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleAdhocExecute{
				cmd: &AnsibleAdhocCmd{},
			}

			res := e.WithPrivilegeEscalationOptions(test.options)

			assert.Equal(t, test.expected, res)
		})
	}
}
