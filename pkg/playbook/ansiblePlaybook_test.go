package playbook

import (
	"testing"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/vault"
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
			err:                    errors.New("(playbook::GenerateCommandOptions)", "AnsiblePlaybookOptions is nil"),
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
		{
			desc: "Testing AnsiblePlaybookOptions with extra vars",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
				ExtraVarsFile: []string{"@test.yml"},
				FlushCache:    true,
				Inventory:     "inventory",
				Limit:         "limit",
				ListHosts:     true,
				ListTags:      true,
				ListTasks:     true,
				Tags:          "tags",
			},
			err:     nil,
			options: []string{"--extra-vars", "{\"extra\":\"var\"}", "--extra-vars", "@test.yml", "--flush-cache", "--inventory", "inventory", "--limit", "limit", "--list-hosts", "--list-tags", "--list-tasks", "--tags", "tags"},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			options, err := test.ansiblePlaybookOptions.GenerateCommandOptions()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.options, options, "Unexpected options value")
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
				Playbooks: []string{"test/ansible/site.yml"},
				ConnectionOptions: &options.AnsibleConnectionOptions{
					AskPass:    true,
					Connection: "local",
					PrivateKey: "pk",
					Timeout:    10,
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
				ConnectionOptions: &options.AnsibleConnectionOptions{
					AskPass:    true,
					Connection: "local",
					PrivateKey: "pk",
					Timeout:    10,
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
					ExtraVarsFile: []string{"@test/ansible/extra_vars.yml"},
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
			res: "ansible-playbook  --ask-vault-password --check --diff --extra-vars '{\"var1\":\"value1\"}' --extra-vars @test/ansible/extra_vars.yml --flush-cache --force-handlers --forks 10 --inventory test/ansible/inventory/all --limit myhost --list-hosts --list-tags --list-tasks --module-path /dev/null --skip-tags tagN --start-at-task task1 --step --syntax-check --tags tag1 --vault-id asdf --vault-password-file /dev/null -vvvv --version  --ask-pass --connection local --private-key pk --timeout 10 --user apenella  --ask-become-pass --become --become-method sudo --become-user apenella test/ansible/site.yml test/ansible/site2.yml",
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

func TestGenerateExtraVarsCommand(t *testing.T) {

	tests := []struct {
		desc      string
		options   *AnsiblePlaybookOptions
		err       error
		extravars string
	}{
		{
			desc: "Testing extra vars map[string]string",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:       nil,
			extravars: "{\"extra\":\"var\"}",
		},
		{
			desc: "Testing extra vars map[string]bool",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": true,
				},
			},
			err:       nil,
			extravars: "{\"extra\":true}",
		},
		{
			desc: "Testing extra vars map[string]int",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": 10,
				},
			},
			err:       nil,
			extravars: "{\"extra\":10}",
		},
		{
			desc: "Testing extra vars map[string][]string",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": []string{"var"},
				},
			},
			err:       nil,
			extravars: "{\"extra\":[\"var\"]}",
		},
		{
			desc: "Testing extra vars map[string]map[string]string",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": map[string]string{
						"var": "value",
					},
				},
			},
			err:       nil,
			extravars: "{\"extra\":{\"var\":\"value\"}}",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			extravars, err := test.options.generateExtraVarsCommand()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, extravars, test.extravars, "Unexpected options value")
			}
		})
	}
}

func TestAddExtraVar(t *testing.T) {
	tests := []struct {
		desc          string
		options       *AnsiblePlaybookOptions
		err           error
		extraVarName  string
		extraVarValue interface{}
		res           map[string]interface{}
	}{
		{
			desc: "Testing add an extraVar to a nil data structure",
			options: &AnsiblePlaybookOptions{
				ExtraVars: nil,
			},
			err:           nil,
			extraVarName:  "extra",
			extraVarValue: "var",
			res: map[string]interface{}{
				"extra": "var",
			},
		},
		{
			desc: "Testing add an extraVar",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra1": "var1",
				},
			},
			err:           nil,
			extraVarName:  "extra",
			extraVarValue: "var",
			res: map[string]interface{}{
				"extra1": "var1",
				"extra":  "var",
			},
		},
		{
			desc: "Testing add an existing extraVar",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:           errors.New("(playbook::AddExtraVar)", "ExtraVar 'extra' already exist"),
			extraVarName:  "extra",
			extraVarValue: "var",
			res:           nil,
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.options.AddExtraVar(test.extraVarName, test.extraVarValue)

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, test.options.ExtraVars, "Unexpected options value")
			}
		})
	}

}

func TestAddExtraVarsFile(t *testing.T) {

	tests := []struct {
		desc    string
		file    string
		options *AnsiblePlaybookOptions
		res     []string
		err     error
	}{
		{
			desc:    "Testing add an extra-vars file when ExtraVarsFile is nil",
			file:    "@test.yml",
			options: &AnsiblePlaybookOptions{},
			res:     []string{"@test.yml"},
			err:     &errors.Error{},
		},
		{
			desc: "Testing add an extra-vars file",
			file: "@test2.yml",
			options: &AnsiblePlaybookOptions{
				ExtraVarsFile: []string{"@test1.yml"},
			},
			res: []string{"@test1.yml", "@test2.yml"},
			err: &errors.Error{},
		},
		{
			desc: "Testing add an extra-vars file without file mark prefix @",
			file: "test.yml",
			options: &AnsiblePlaybookOptions{
				ExtraVarsFile: []string{},
			},
			res: []string{"@test.yml"},
			err: &errors.Error{},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.options.AddExtraVarsFile(test.file)

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, test.options.ExtraVarsFile, "Unexpected options value")
			}
		})
	}
}

// AddVaultedExtraVar(vaulter Vaulter, name string, value string)
func TestAddVaultedExtraVar(t *testing.T) {
	vaulter := vault.NewMockVariableVaulter()
	vaulter.On("Vault", "plain_text_value").Return(vault.NewVaultVariableValue("encrypted_value"), nil)

	tests := []struct {
		desc    string
		options *AnsiblePlaybookOptions
		vaulter Vaulter
		name    string
		value   string
		res     map[string]interface{}
		err     error
	}{
		{
			desc:    "Testing add a vaulted extra-var",
			options: &AnsiblePlaybookOptions{},
			vaulter: vaulter,
			name:    "variable_name",
			value:   "plain_text_value",
			res: map[string]interface{}{
				"variable_name": vault.NewVaultVariableValue("encrypted_value"),
			},
			err: &errors.Error{},
		},
		{
			desc:    "Testing error adding a vaulted extra-var when vaulter is nil",
			options: &AnsiblePlaybookOptions{},
			vaulter: nil,
			err:     errors.New("(playbook::AddVaultedExtraVar)", "To define a vaulted extra-var you need to initialize a vaulter"),
		},
		{
			desc: "Testing error adding a vaulted extra-var when variable already exist",
			options: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"variable_name": "{\"__ansible_vault\":\"encrypted_value\"}",
				},
			},
			vaulter: vaulter,
			name:    "variable_name",
			value:   "plain_text_value",
			res:     map[string]interface{}{},
			err:     errors.New("(playbook::AddVaultedExtraVar)", "ExtraVar 'variable_name' already exist"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.options.AddVaultedExtraVar(test.vaulter, test.name, test.value)
			if err != nil {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, test.options.ExtraVars, "Unexpected options value")
			}
		})
	}
}

func TestGenerateVerbosityFlag(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsiblePlaybookOptions
		res     string
		err     error
	}{
		{
			desc: "Testing generate verbosity flag",
			options: &AnsiblePlaybookOptions{
				Verbose: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag V",
			options: &AnsiblePlaybookOptions{
				VerboseV: true,
			},
			res: "-v",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV",
			options: &AnsiblePlaybookOptions{
				VerboseVV: true,
			},
			res: "-vv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVV",
			options: &AnsiblePlaybookOptions{
				VerboseVVV: true,
			},
			res: "-vvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVVV",
			options: &AnsiblePlaybookOptions{
				VerboseVVVV: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV has precedence over V",
			options: &AnsiblePlaybookOptions{
				VerboseVV: true,
				VerboseV:  true,
			},
			res: "-vv",
			err: &errors.Error{},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res, err := test.options.generateVerbosityFlag()
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, res)
			}
		})
	}
}
