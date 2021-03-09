package playbook

import (
	"bytes"
	"context"
	goerrors "errors"
	execerrors "os/exec"
	"testing"

	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/stdoutcallback"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

// TestGenerateCommandConnectionOptions
func TestGenerateCommandConnectionOptions(t *testing.T) {
	tests := []struct {
		desc                             string
		ansiblePlaybookConnectionOptions *options.AnsibleConnectionOptions
		err                              error
		options                          []string
	}{
		{
			desc: "Testing generate connection options",
			ansiblePlaybookConnectionOptions: &options.AnsibleConnectionOptions{
				Connection: "local",
			},
			err: nil,
			options: []string{
				"--connection",
				"local",
			},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			options, err := test.ansiblePlaybookConnectionOptions.GenerateCommandConnectionOptions()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, options, test.options, "Unexpected options value")
			}
		})
	}

}

// TestGenerateCommandOptions tests
func TestGenerateCommandOptions(t *testing.T) {
	tests := []struct {
		desc                   string
		ansiblePlaybookOptions *AnsiblePlaybookOptions
		err                    error
		options                []string
	}{
		{
			desc:                   "Testing nil AnsiblePlaybookOptions definition",
			ansiblePlaybookOptions: nil,
			err:                    errors.New("(ansible::GenerateCommandOptions)", "AnsiblePlaybookOptions is nil"),
			options:                nil,
		},
		{
			desc:                   "Testing an empty AnsiblePlaybookOptions definition",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{},
			err:                    nil,
			options:                []string{},
		},
		{
			desc: "Testing AnsiblePlaybookOptions except extra vars",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				FlushCache:    true,
				ForceHandlers: true,
				ListTags:      true,
				ListTasks:     true,
				SkipTags:      "tagN",
				StartAtTask:   "second",
				Step:          true,
				Tags:          "tags",
			},
			err:     nil,
			options: []string{"--flush-cache", "--force-handlers", "--list-tags", "--list-tasks", "--skip-tags", "tagN", "--start-at-task", "second", "--step", "--tags", "tags"},
		},
		// TODO
		// {
		// 	desc: "Testing AnsiblePlaybookOptions with extra vars",
		// 	Options: &AnsiblePlaybookOptions{
		// 		ExtraVars: map[string]interface{}{
		// 			"extra": "var",
		// 		},
		// 		FlushCache: true,
		// 		Inventory:  "inventory",
		// 		Limit:      "limit",
		// 		ListHosts:  true,
		// 		ListTags:   true,
		// 		ListTasks:  true,
		// 		Tags:       "tags",
		// 	},
		// 	err:     nil,
		// 	options: []string{"--flush-cache", "--inventory", "inventory", "--limit", "limit", "--list-hosts", "--list-tags", "--list-tasks", "--tags", "tags", "--extra-vars", "{\"extra\":\"var\"}"},
		// },
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			options, err := test.ansiblePlaybookOptions.GenerateCommandOptions()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, options, test.options, "Unexpected options value")
			}
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
				Playbook: "test/ansible/site.yml",
				ConnectionOptions: &options.AnsibleConnectionOptions{
					AskPass:    true,
					Connection: "local",
					PrivateKey: "pk",
					Timeout:    "10",
					User:       "apenella",
				},
				Options: &AnsiblePlaybookOptions{
					AskVaultPassword:  true,
					Check:             true,
					Diff:              true,
					Forks:             "10",
					ListHosts:         true,
					ModulePath:        "/dev/null",
					SyntaxCheck:       true,
					VaultID:           "asdf",
					VaultPasswordFile: "/dev/null",
					Verbose:           true,
					Version:           true,

					Inventory: "test/ansible/inventory/all",
					Limit:     "myhost",
					ExtraVars: map[string]interface{}{
						"var1": "value1",
					},
					FlushCache: true,
					Tags:       "tag1",
				},
				PrivilegeEscalationOptions: &options.AnsiblePrivilegeEscalationOptions{
					Become:        true,
					BecomeMethod:  "sudo",
					BecomeUser:    "apenella",
					AskBecomePass: true,
				},
			},
			command: []string{"ansible-playbook", "--extra-vars", "{\"var1\":\"value1\"}", "--inventory", "test/ansible/inventory/all", "--limit", "myhost", "--flush-cache", "--tags", "tag1", "--ask-pass", "--connection", "local", "--private-key", "pk", "--user", "apenella", "--timeout", "10", "--ask-become-pass", "--become", "--become-method", "sudo", "--become-user", "apenella", "test/ansible/site.yml"},
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
				Playbook: "test/ansible/site.yml",
				ConnectionOptions: &options.AnsibleConnectionOptions{
					AskPass:    true,
					Connection: "local",
					PrivateKey: "pk",
					Timeout:    "10",
					User:       "apenella",
				},
				Options: &AnsiblePlaybookOptions{
					AskVaultPassword:  true,
					Check:             true,
					Diff:              true,
					Forks:             "10",
					ListHosts:         true,
					ModulePath:        "/dev/null",
					SyntaxCheck:       true,
					VaultID:           "asdf",
					VaultPasswordFile: "/dev/null",
					Verbose:           true,
					Version:           true,
					Inventory:         "test/ansible/inventory/all",
					Limit:             "myhost",
					ExtraVars: map[string]interface{}{
						"var1": "value1",
					},
					FlushCache:    true,
					ForceHandlers: true,
					ListTags:      true,
					ListTasks:     true,
					SkipTags:      "tagN",
					StartAtTask:   "task1",
					Step:          true,
					Tags:          "tag1",
				},
				PrivilegeEscalationOptions: &options.AnsiblePrivilegeEscalationOptions{
					Become:        true,
					BecomeMethod:  "sudo",
					BecomeUser:    "apenella",
					AskBecomePass: true,
				},
			},
			res: "ansible-playbook  --extra-vars '{\"var1\":\"value1\"}' --inventory test/ansible/inventory/all --limit myhost --flush-cache --force-handlers --list-tags --list-tasks --skip-tags tagN --start-at-task task1 --step --tags tag1  --ask-pass --connection local --private-key pk --user apenella --timeout 10  --ask-become-pass --become --become-method sudo --become-user apenella test/ansible/site.yml",
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

func TestRun(t *testing.T) {

	var w bytes.Buffer

	tests := []struct {
		desc               string
		ansiblePlaybookCmd *AnsiblePlaybookCmd
		res                string
		ctx                context.Context
		err                error
	}{
		{
			desc:               "Run nil ansiblePlaybookCmd",
			ansiblePlaybookCmd: nil,
			res:                "",
			err:                errors.New("(ansible:Run)", "AnsiblePlaybookCmd is nil"),
		},
		{
			desc: "Testing run a ansiblePlaybookCmd with unexisting binary file",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Binary: "unexisting",
			},
			res: "",
			ctx: context.TODO(),
			err: errors.New("(ansible:Run)", "Binary file 'unexisting' does not exists", &execerrors.Error{Name: "unexisting", Err: goerrors.New("executable file not found in $PATH")}),
		},
		{
			desc: "Testing run a ansiblePlaybookCmd",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Exec: &execute.MockExecute{
					Write: &w,
				},
				Playbook: "test/ansible/site.yml",
				ConnectionOptions: &options.AnsibleConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/ansible/inventory/all",
				},
			},
			ctx: context.TODO(),
			res: "[ansible-playbook --inventory test/ansible/inventory/all --connection local test/ansible/site.yml]",
			err: nil,
		},
		{
			desc: "Testing run a ansiblePlaybookCmd without executor",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Exec:     nil,
				Playbook: "test/test_site.yml",
				ConnectionOptions: &options.AnsibleConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/ansible/inventory/all",
				},
			},
			res: "",
			ctx: context.TODO(),
			err: nil,
		},
		{
			desc: "Testing run a ansiblePlaybookCmd with JSON stdout callback",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				StdoutCallback: stdoutcallback.JSONStdoutCallback,
				Playbook:       "test/test_site.yml",
				ConnectionOptions: &options.AnsibleConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/all",
				},
			},
			res: "",
			ctx: context.TODO(),
			err: nil,
		},
		{
			desc: "Testing run a ansiblePlaybookCmd with multiple extravars",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Exec: &execute.MockExecute{
					Write: &w,
				},
				Playbook: "test/test_site.yml",
				ConnectionOptions: &options.AnsibleConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/all",
					ExtraVars: map[string]interface{}{
						"string": "testing an string",
						"bool":   true,
						"int":    10,
						"array":  []string{"one", "two"},
						"dict": map[string]bool{
							"one": true,
							"two": false,
						},
					},
				},
			},
			res: "[ansible-playbook --extra-vars {\"array\":[\"one\",\"two\"],\"bool\":true,\"dict\":{\"one\":true,\"two\":false},\"int\":10,\"string\":\"testing an string\"} --inventory test/all --connection local test/test_site.yml]",
			ctx: context.TODO(),
			err: nil,
		},
	}

	for _, test := range tests {
		w = bytes.Buffer{}

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)
			w.Reset()

			err := test.ansiblePlaybookCmd.Run(test.ctx)
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, w.String(), "Unexpected value")
			}
		})
	}
}
