package inventory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAnsibleInventoryExecute(t *testing.T) {
	tests := []struct {
		desc     string
		expected *AnsibleInventoryExecute
	}{
		{
			desc: "Testing creating a new AnsibleInventoryExecute",
			expected: &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := NewAnsibleInventoryExecute()

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithBinary(t *testing.T) {
	tests := []struct {
		desc     string
		binary   string
		expected *AnsibleInventoryExecute
	}{
		{
			desc:   "Testing setting binary to AnsibleInventoryExecute",
			binary: "test1",
			expected: &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{
					Binary: "test1",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{},
			}

			res := e.WithBinary(test.binary)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithInventoryOptions(t *testing.T) {
	tests := []struct {
		desc     string
		options  *AnsibleInventoryOptions
		expected *AnsibleInventoryExecute
	}{
		{
			desc: "Testing setting inventory options to AnsibleInventoryExecute",
			options: &AnsibleInventoryOptions{
				Host: "test1",
			},
			expected: &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{
					InventoryOptions: &AnsibleInventoryOptions{
						Host: "test1",
					},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{},
			}

			res := e.WithInventoryOptions(test.options)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestWithPattern(t *testing.T) {
	tests := []struct {
		desc     string
		pattern  string
		expected *AnsibleInventoryExecute
	}{
		{
			desc:    "Testing setting pattern to AnsibleInventoryExecute",
			pattern: "test1",
			expected: &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{
					Pattern: "test1",
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := &AnsibleInventoryExecute{
				cmd: &AnsibleInventoryCmd{},
			}

			res := e.WithPattern(test.pattern)

			assert.Equal(t, test.expected, res)
		})
	}
}
