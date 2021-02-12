package execute

import (
	"os"
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestDefaultExecute(t *testing.T) {
	tests := []struct {
		desc    string
		err     error
		execute *DefaultExecute
		command []string
		prefix  string
		res     string
	}{
		{
			desc: "Testing an execution",
			err:  nil,
			execute: &DefaultExecute{
				Write: os.Stdout,
			},
			command: []string{"echo", "hello"},
			prefix:  "test",
		},
		{
			desc: "Testing an ansible-playbook --version execution",
			err:  nil,
			execute: &DefaultExecute{
				Write: os.Stdout,
			},
			command: []string{"ansible-playbook", "--version"},
			prefix:  "test",
		},
		{
			desc: "Testing an ansible-playbook with local connection",
			err:  nil,
			execute: &DefaultExecute{
				Write: os.Stdout,
			},
			command: []string{"ansible-playbook", "--inventory", "127.0.0.1,", "test/site.yml", "-c", "local"},
			prefix:  "test",
		},
		{
			desc: "Testing an ansible-playbook forcing ",
			err:  errors.New("(DefaultExecute::Execute)", "Error during command execution: ansible-playbook error: parser error\n\nCommand executed: /home/aleix/.local/bin/ansible-playbook --inventory 127.0.0.1, test/site.yml --user apenella\n\nexit status 4"),
			execute: &DefaultExecute{
				Write: os.Stdout,
			},
			command: []string{"ansible-playbook", "--inventory", "127.0.0.1,", "test/site.yml", "--user", "apenella"},
			prefix:  "test",
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
