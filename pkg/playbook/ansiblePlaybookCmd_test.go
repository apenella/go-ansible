package playbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewAnsiblePlaybookCmd tests
func TestNewAnsiblePlaybookCmd(t *testing.T) {
	tests := []struct {
		desc                       string
		expectedAnsiblePlaybookCmd *AnsiblePlaybookCmd
		binary                     string
		playbooks                  []string
		playbookOptions            *AnsiblePlaybookOptions
	}{
		{
			desc:      "Testing create new AnsiblePlaybookCmd",
			binary:    "custom-ansible-playbook",
			playbooks: []string{"test/ansible/site.yml", "test/ansible/site2.yml"},
			playbookOptions: &AnsiblePlaybookOptions{
				AskBecomePass:    true,
				AskPass:          true,
				AskVaultPassword: true,
				Become:           true,
			},
			expectedAnsiblePlaybookCmd: &AnsiblePlaybookCmd{
				Binary:    "custom-ansible-playbook",
				Playbooks: []string{"test/ansible/site.yml", "test/ansible/site2.yml"},
				PlaybookOptions: &AnsiblePlaybookOptions{
					AskBecomePass:    true,
					AskPass:          true,
					AskVaultPassword: true,
					Become:           true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			ansiblePlaybookCmd := NewAnsiblePlaybookCmd(
				WithBinary(test.binary),
				WithPlaybooks(test.playbooks...),
				WithPlaybookOptions(test.playbookOptions),
			)

			assert.Equal(t, test.expectedAnsiblePlaybookCmd, ansiblePlaybookCmd, "Unexpected value")
		})
	}
}

// TestCommand tests
func TestCommand(t *testing.T) {
	tests := []struct {
		desc               string
		err                error
		ansiblePlaybookCmd *AnsiblePlaybookCmd
		command            []string
	}{
		{
			desc: "Testing generate AnsiblePlaybookCmd command",
			err:  nil,
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Playbooks: []string{"test/ansible/site.yml"},

				PlaybookOptions: &AnsiblePlaybookOptions{
					AskBecomePass:    true,
					AskPass:          true,
					AskVaultPassword: true,
					Become:           true,
					BecomeMethod:     "sudo",
					BecomeUser:       "apenella",
					Check:            true,
					Connection:       "local",
					ExtraVars: map[string]interface{}{
						"var1": "value1",
					},
					Diff:              true,
					FlushCache:        true,
					Forks:             "10",
					Inventory:         "test/ansible/inventory/all",
					Limit:             "myhost",
					ListHosts:         true,
					ModulePath:        "/dev/null",
					PrivateKey:        "pk",
					SyntaxCheck:       true,
					Tags:              "tag1",
					Timeout:           10,
					User:              "apenella",
					VaultID:           "asdf",
					VaultPasswordFile: "/dev/null",
					Verbose:           true,
					Version:           true,
				},
			},
			command: []string{"ansible-playbook", "--ask-vault-password", "--check", "--diff", "--extra-vars", "{\"var1\":\"value1\"}", "--flush-cache", "--forks", "10", "--inventory", "test/ansible/inventory/all", "--limit", "myhost", "--list-hosts", "--module-path", "/dev/null", "--syntax-check", "--tags", "tag1", "--vault-id", "asdf", "--vault-password-file", "/dev/null", "-vvvv", "--version", "--ask-pass", "--connection", "local", "--private-key", "pk", "--timeout", "10", "--user", "apenella", "--ask-become-pass", "--become", "--become-method", "sudo", "--become-user", "apenella", "test/ansible/site.yml"},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			command, err := test.ansiblePlaybookCmd.Command()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.command, command, "Unexpected value")
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		desc               string
		err                error
		ansiblePlaybookCmd *AnsiblePlaybookCmd
		res                string
	}{
		{
			desc: "Testing AnsiblePlaybookCmd to string",
			err:  nil,
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Playbooks: []string{"test/ansible/site.yml", "test/ansible/site2.yml"},
				PlaybookOptions: &AnsiblePlaybookOptions{
					AskBecomePass:    true,
					AskPass:          true,
					AskVaultPassword: true,
					Become:           true,
					BecomeMethod:     "sudo",
					BecomeUser:       "apenella",
					Check:            true,
					Connection:       "local",
					Diff:             true,
					ExtraVars: map[string]interface{}{
						"var1": "value1",
					},
					ExtraVarsFile:     []string{"@test/ansible/extra_vars.yml"},
					FlushCache:        true,
					ForceHandlers:     true,
					Forks:             "10",
					Inventory:         "test/ansible/inventory/all",
					Limit:             "myhost",
					ListHosts:         true,
					ListTags:          true,
					ListTasks:         true,
					ModulePath:        "/dev/null",
					PrivateKey:        "pk",
					SkipTags:          "tagN",
					StartAtTask:       "task1",
					Step:              true,
					SyntaxCheck:       true,
					Tags:              "tag1",
					Timeout:           10,
					User:              "apenella",
					VaultID:           "asdf",
					VaultPasswordFile: "/dev/null",
					Verbose:           true,
					Version:           true,
				},
			},
			res: "ansible-playbook  --ask-vault-password --check --diff --extra-vars '{\"var1\":\"value1\"}' --extra-vars @test/ansible/extra_vars.yml --flush-cache --force-handlers --forks 10 --inventory test/ansible/inventory/all --limit myhost --list-hosts --list-tags --list-tasks --module-path /dev/null --skip-tags tagN --start-at-task task1 --step --syntax-check --tags tag1 --vault-id asdf --vault-password-file /dev/null -vvvv --version --ask-pass --connection local --private-key pk --timeout 10 --user apenella --ask-become-pass --become --become-method sudo --become-user apenella test/ansible/site.yml test/ansible/site2.yml",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res := test.ansiblePlaybookCmd.String()

			assert.Equal(t, test.res, res, "Unexpected value")
		})
	}

}
