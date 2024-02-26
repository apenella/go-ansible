package adhoc

import (
	"testing"

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
