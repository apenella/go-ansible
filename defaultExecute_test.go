package ansibler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultExecute(t *testing.T) {
	tests := []struct {
		desc    string
		err     error
		execute *Executor
		command []string
		res     string
	}{
		{
			desc: "Testing an execution",
			err:  nil,
			execute: &Executor{},
			command: []string{"echo", "hello"},
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.execute.Execute(test.command[0], test.command[1:])
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		}
	}

}
