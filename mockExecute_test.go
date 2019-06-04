package ansibler

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockExecute(t *testing.T) {
	tests := []struct {
		desc    string
		err     error
		execute *MockExecute
		command []string
		prefix  string
		res     string
	}{
		{
			desc:    "Testing a dummy error",
			err:     errors.New("(MockExecute::Execute) error"),
			execute: &MockExecute{},
			command: []string{"error"},
			prefix:  "",
			res:     "",
		},
		{
			desc:    "Testing an execution",
			err:     nil,
			execute: &MockExecute{},
			command: []string{"command", "arg1", "arg2"},
			prefix:  "prefix",
			res:     "prefix command arg1 arg2",
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.execute.Execute(test.command[0], test.command[1:], test.prefix)
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		}
	}
}
