package playbook

import (
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/vault"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

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
				AskBecomePass:    true,
				AskPass:          true,
				AskVaultPassword: true,
				Become:           true,
				BecomeMethod:     "become-method",
				BecomeUser:       "become-user",
				Check:            true,
				Connection:       "local",
				Diff:             true,
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
				ExtraVarsFile:     []string{"@test.yml"},
				FlushCache:        true,
				ForceHandlers:     true,
				Forks:             "10",
				Inventory:         "inventory",
				Limit:             "limit",
				ListHosts:         true,
				ListTags:          true,
				ListTasks:         true,
				ModulePath:        "module-path",
				PrivateKey:        "private-key",
				SCPExtraArgs:      "scp-extra-args1 scp-extra-args2",
				SFTPExtraArgs:     "sftp-extra-args1 sftp-extra-args2",
				SkipTags:          "skip-tags",
				SSHCommonArgs:     "ssh-common-args1 ssh-common-args2",
				SSHExtraArgs:      "ssh-extra-args1 ssh-extra-args2",
				StartAtTask:       "start-at-task",
				Step:              true,
				SyntaxCheck:       true,
				Tags:              "tags",
				Timeout:           11,
				User:              "user",
				VaultID:           "vault-ID",
				VaultPasswordFile: "vault-password-file",
				Verbose:           true,
				Version:           true,
			},
			err:     nil,
			options: []string{"--ask-vault-password", "--check", "--diff", "--extra-vars", "{\"extra\":\"var\"}", "--extra-vars", "@test.yml", "--flush-cache", "--force-handlers", "--forks", "10", "--inventory", "inventory", "--limit", "limit", "--list-hosts", "--list-tags", "--list-tasks", "--module-path", "module-path", "--skip-tags", "skip-tags", "--start-at-task", "start-at-task", "--step", "--syntax-check", "--tags", "tags", "--vault-id", "vault-ID", "--vault-password-file", "vault-password-file", "-vvvv", "--version", "--ask-pass", "--connection", "local", "--private-key", "private-key", "--scp-extra-args", "scp-extra-args1 scp-extra-args2", "--sftp-extra-args", "sftp-extra-args1 sftp-extra-args2", "--ssh-common-args", "ssh-common-args1 ssh-common-args2", "--ssh-extra-args", "ssh-extra-args1 ssh-extra-args2", "--timeout", "11", "--user", "user", "--ask-become-pass", "--become", "--become-method", "become-method", "--become-user", "become-user"},
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

func TestAnsiblePlaybookOptionsString(t *testing.T) {
	tests := []struct {
		desc                   string
		ansiblePlaybookOptions *AnsiblePlaybookOptions
		res                    string
	}{
		{
			desc:                   "Testing an empty AnsiblePlaybookOptions definition",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{},
			res:                    "",
		},
		{
			desc: "Testing AnsiblePlaybookOptions except extra vars",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				AskBecomePass:    true,
				AskPass:          true,
				AskVaultPassword: true,
				Become:           true,
				BecomeMethod:     "become-method",
				BecomeUser:       "become-user",
				Check:            true,
				Connection:       "local",
				Diff:             true,
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
				ExtraVarsFile:     []string{"@test.yml"},
				FlushCache:        true,
				ForceHandlers:     true,
				Forks:             "10",
				Inventory:         "inventory",
				Limit:             "limit",
				ListHosts:         true,
				ListTags:          true,
				ListTasks:         true,
				ModulePath:        "module-path",
				PrivateKey:        "private-key",
				SCPExtraArgs:      "scp-extra-args",
				SFTPExtraArgs:     "sftp-extra-args",
				SkipTags:          "skip-tags",
				SSHCommonArgs:     "ssh-common-args",
				SSHExtraArgs:      "ssh-extra-args",
				StartAtTask:       "start-at-task",
				Step:              true,
				SyntaxCheck:       true,
				Tags:              "tags",
				Timeout:           11,
				User:              "user",
				VaultID:           "vault-ID",
				VaultPasswordFile: "vault-password-file",
				Verbose:           true,
				Version:           true,
			},
			res: " --ask-vault-password --check --diff --extra-vars '{\"extra\":\"var\"}' --extra-vars @test.yml --flush-cache --force-handlers --forks 10 --inventory inventory --limit limit --list-hosts --list-tags --list-tasks --module-path module-path --skip-tags skip-tags --start-at-task start-at-task --step --syntax-check --tags tags --vault-id vault-ID --vault-password-file vault-password-file -vvvv --version --ask-pass --connection local --private-key private-key --scp-extra-args 'scp-extra-args' --sftp-extra-args 'sftp-extra-args' --ssh-common-args 'ssh-common-args' --ssh-extra-args 'ssh-extra-args' --timeout 11 --user user --ask-become-pass --become --become-method become-method --become-user become-user",
		},
		{
			desc: "Testing AnsiblePlaybookOptions setting the VerboseV flag as true",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				VerboseV: true,
			},
			res: " -v",
		},
		{
			desc: "Testing AnsiblePlaybookOptions setting the VerboseVV flag as true",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				VerboseVV: true,
			},
			res: " -vv",
		},
		{
			desc: "Testing AnsiblePlaybookOptions setting the VerboseVVV flag as true",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				VerboseVVV: true,
			},
			res: " -vvv",
		},
		{
			desc: "Testing AnsiblePlaybookOptions setting the VerboseVVVV flag as true",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				VerboseVVVV: true,
			},
			res: " -vvvv",
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res := test.ansiblePlaybookOptions.String()

			assert.Equal(t, test.res, res, "Unexpected options value for AnsiblePlaybookOptions.String()")

		})
	}
}
