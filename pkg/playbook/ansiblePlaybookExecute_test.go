package playbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAnsiblePlaybookExecute(t *testing.T) {
	tests := []struct {
		desc     string
		playbook []string
		expected *AnsiblePlaybookExecute
	}{
		{
			desc:     "Testing creating a new AnsiblePlaybookExecute",
			playbook: []string{"test1"},
			expected: &AnsiblePlaybookExecute{
				cmd: &AnsiblePlaybookCmd{
					Playbooks: []string{"test1"},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := NewAnsiblePlaybookExecute(test.playbook...)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithBinary(t *testing.T) {
	tests := []struct {
		desc     string
		binary   string
		expected *AnsiblePlaybookExecute
	}{
		{
			desc:   "Testing setting binary to AnsiblePlaybookExecute",
			binary: "test1",
			expected: &AnsiblePlaybookExecute{
				cmd: &AnsiblePlaybookCmd{
					Binary: "test1",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsiblePlaybookExecute{
				cmd: &AnsiblePlaybookCmd{},
			}

			res := e.WithBinary(test.binary)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithPlaybookOptions(t *testing.T) {
	tests := []struct {
		desc     string
		options  *AnsiblePlaybookOptions
		expected *AnsiblePlaybookExecute
	}{
		{
			desc: "Testing setting playbook options to AnsiblePlaybookExecute",
			options: &AnsiblePlaybookOptions{
				Inventory: "test1",
			},
			expected: &AnsiblePlaybookExecute{
				cmd: &AnsiblePlaybookCmd{
					PlaybookOptions: &AnsiblePlaybookOptions{
						Inventory: "test1",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsiblePlaybookExecute{
				cmd: &AnsiblePlaybookCmd{},
			}

			res := e.WithPlaybookOptions(test.options)

			assert.Equal(t, test.expected, res)
		})
	}
}
