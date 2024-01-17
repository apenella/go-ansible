package inventory

import (
	"context"
	goerrors "errors"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	"github.com/apenella/go-ansible/pkg/vault"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	execerrors "os/exec"
	"testing"
)

// TestGenerateCommandOptions tests
func TestGenerateCommandOptions(t *testing.T) {
	tests := []struct {
		desc                   string
		ansiblePlaybookOptions *AnsibleInventoryOptions
		err                    error
		options                []string
	}{
		{
			desc:                   "Testing nil AnsiblePlaybookOptions definition",
			ansiblePlaybookOptions: nil,
			err:                    errors.New("(inventory::GenerateCommandOptions)", "AnsibleInventoryOptions is nil"),
			options:                nil,
		},
		{
			desc:                   "Testing an empty AnsibleInventoryOptions definition",
			ansiblePlaybookOptions: &AnsibleInventoryOptions{},
			err:                    nil,
			options:                []string{},
		},
		{
			desc: "Testing AnsibleInventoryOptions except vars",
			ansiblePlaybookOptions: &AnsibleInventoryOptions{
				Host:      "localhost",
				Inventory: "test/ansible/inventory/all",
				Output:    "/tmp/output.ini",
			},
			err:     nil,
			options: []string{"--host", "localhost", "--inventory", "test/ansible/inventory/all", "--output", "/tmp/output.ini"},
		},
		{
			desc: "Testing AnsibleInventoryOptions with vars",
			ansiblePlaybookOptions: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key":     "value",
					"key_int": 1,
				},
				VarsFile:  []string{"@test.yml"},
				Host:      "localhost",
				Inventory: "test/ansible/inventory/all",
				Output:    "/tmp/output.ini",
			},
			err:     nil,
			options: []string{"--host", "localhost", "--inventory", "test/ansible/inventory/all", "--output", "/tmp/output.ini", "--vars", "{\"key\":\"value\",\"key_int\":1}", "--vars", "@test.yml"},
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

// TestString tests
func TestString(t *testing.T) {
	tests := []struct {
		desc                string
		err                 error
		ansibleInventoryCmd *AnsibleInventoryCmd
		res                 string
	}{
		{
			desc: "Testing AnsiblePlaybookCmd to string",
			err:  nil,
			ansibleInventoryCmd: &AnsibleInventoryCmd{
				Binary:  "ansible-inventory",
				Exec:    execute.NewMockExecute(),
				Pattern: "all",
				Options: &AnsibleInventoryOptions{
					AskVaultPassword: true,
					Export:           true,
					Graph:            true,
					Host:             "localhost",
					Inventory:        "test/ansible/inventory/all",
					List:             true,
					Output:           "/tmp/output.ini",
					PlaybookDir:      "/playbook/",
					Toml:             true,
					Vars: map[string]interface{}{
						"array":  []string{"one", "two"},
						"bool":   true,
						"dict":   map[string]bool{"one": true, "two": false},
						"int":    10,
						"string": "testing an string",
					},
					VarsFile:          []string{"@test/ansible/vars.yml"},
					VaultID:           "asdf",
					VaultPasswordFile: "/vault/password/file",
					Verbose:           true,
					Version:           true,
					Yaml:              true,
				},
			},
			res: "ansible-inventory  --ask-vault-password --export --graph --host localhost --inventory test/ansible/inventory/all --list --output /tmp/output.ini --playbook-dir /playbook/ --toml --vars '{\"array\":[\"one\",\"two\"],\"bool\":true,\"dict\":{\"one\":true,\"two\":false},\"int\":10,\"string\":\"testing an string\"}' --vars @test/ansible/vars.yml --vault-id asdf --vault-password-file /vault/password/file -vvvv --version --yaml all",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			res := test.ansibleInventoryCmd.String()

			assert.Equal(t, test.res, res, "Unexpected value")
		})
	}
}

// TestCommand tests
func TestCommand(t *testing.T) {
	tests := []struct {
		desc                string
		err                 error
		AnsibleInventoryCmd *AnsibleInventoryCmd
		command             []string
	}{
		{
			desc: "Testing generate AnsibleInventoryCmd command",
			err:  nil,
			AnsibleInventoryCmd: &AnsibleInventoryCmd{
				Pattern: "all",
				Options: &AnsibleInventoryOptions{
					Host:      "localhost",
					Inventory: "test/ansible/inventory/all",
					Output:    "/tmp/output.ini",
				},
			},
			command: []string{"ansible-inventory", "--host", "localhost", "--inventory", "test/ansible/inventory/all", "--output", "/tmp/output.ini", "all"},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			command, err := test.AnsibleInventoryCmd.Command()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.command, command, "Unexpected value")
			}
		})
	}
}

// TestRun tests
func TestRun(t *testing.T) {
	tests := []struct {
		desc                string
		ansibleInventoryCmd *AnsibleInventoryCmd
		err                 error
		prepareAssertFunc   func(cmd *AnsibleInventoryCmd)
	}{
		{
			desc:                "Run nil ansiblePlaybookCmd",
			ansibleInventoryCmd: nil,
			err:                 errors.New("(inventory::Run)", "AnsibleInventoryCmd is nil"),
		},
		{
			desc: "Testing run a ansibleInventoryCmd with unexisting binary file",
			ansibleInventoryCmd: &AnsibleInventoryCmd{
				Binary: "unexisting",
			},
			err: errors.New("(inventory::Run)", "Binary file 'unexisting' does not exists", &execerrors.Error{Name: "unexisting", Err: goerrors.New("executable file not found in $PATH")}),
		},
		{
			desc: "Testing run a ansibleInventoryCmd",
			ansibleInventoryCmd: &AnsibleInventoryCmd{
				Binary:  "ansible-inventory",
				Exec:    execute.NewMockExecute(),
				Pattern: "all",
				Options: &AnsibleInventoryOptions{
					AskVaultPassword: true,
					Export:           true,
					Graph:            true,
					Host:             "localhost",
					Inventory:        "test/ansible/inventory/all",
					List:             true,
					Output:           "/tmp/output.ini",
					PlaybookDir:      "/playbook/",
					Toml:             true,
					Vars: map[string]interface{}{
						"array":  []string{"one", "two"},
						"bool":   true,
						"dict":   map[string]bool{"one": true, "two": false},
						"int":    10,
						"string": "testing an string",
					},
					VarsFile:          []string{"@test/ansible/vars.yml"},
					VaultID:           "asdf",
					VaultPasswordFile: "/vault/password/file",
					Verbose:           true,
					Version:           true,
					Yaml:              true,
				},
				StdoutCallback: stdoutcallback.JSONStdoutCallback,
			},
			prepareAssertFunc: func(inventory *AnsibleInventoryCmd) {
				inventory.Exec.(*execute.MockExecute).On(
					"Execute",
					context.TODO(),
					[]string{"ansible-inventory",
						"--ask-vault-password",
						"--export",
						"--graph",
						"--host",
						"localhost",
						"--inventory",
						"test/ansible/inventory/all",
						"--list",
						"--output",
						"/tmp/output.ini",
						"--playbook-dir",
						"/playbook/",
						"--toml",
						"--vars",
						"{\"array\":[\"one\",\"two\"],\"bool\":true,\"dict\":{\"one\":true,\"two\":false},\"int\":10,\"string\":\"testing an string\"}",
						"--vars",
						"@test/ansible/vars.yml",
						"--vault-id",
						"asdf",
						"--vault-password-file",
						"/vault/password/file",
						"--version",
						"-vvvv",
						"--yaml",
						"all",
					},
					mock.AnythingOfType("StdoutCallbackResultsFunc"),
					[]execute.ExecuteOptions{},
				).Return(nil)
			},
			err: nil,
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			if test.prepareAssertFunc != nil {
				test.prepareAssertFunc(test.ansibleInventoryCmd)
			}

			err := test.ansibleInventoryCmd.Run(context.TODO())
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				test.ansibleInventoryCmd.Exec.(*execute.MockExecute).AssertExpectations(t)
			}
		})
	}
}

// TestGenerateVarsCommand tests
func TestGenerateVarsCommand(t *testing.T) {

	tests := []struct {
		desc    string
		options *AnsibleInventoryOptions
		err     error
		vars    string
	}{
		{
			desc: "Testing vars map[string]string",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key": "value",
				},
			},
			err:  nil,
			vars: "{\"key\":\"value\"}",
		},
		{
			desc: "Testing vars map[string]bool",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key": true,
				},
			},
			err:  nil,
			vars: "{\"key\":true}",
		},
		{
			desc: "Testing vars map[string]int",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key": 10,
				},
			},
			err:  nil,
			vars: "{\"key\":10}",
		},
		{
			desc: "Testing vars map[string][]string",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key": []string{"value"},
				},
			},
			err:  nil,
			vars: "{\"key\":[\"value\"]}",
		},
		{
			desc: "Testing vars map[string]map[string]string",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key": map[string]string{
						"var": "value",
					},
				},
			},
			err:  nil,
			vars: "{\"key\":{\"var\":\"value\"}}",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			vars, err := test.options.generateVarsCommand()

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, vars, test.vars, "Unexpected options value")
			}
		})
	}
}

// TestAddVar tests
func TestAddVar(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsibleInventoryOptions
		err     error
		key     string
		value   interface{}
		res     map[string]interface{}
	}{
		{
			desc: "Testing add an vars to a nil data structure",
			options: &AnsibleInventoryOptions{
				Vars: nil,
			},
			err:   nil,
			key:   "key",
			value: "value",
			res: map[string]interface{}{
				"key": "value",
			},
		},
		{
			desc: "Testing add an vars",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key1": "value1",
				},
			},
			err:   nil,
			key:   "key",
			value: "value",
			res: map[string]interface{}{
				"key1": "value1",
				"key":  "value",
			},
		},
		{
			desc: "Testing add an existing vars",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"key": "value",
				},
			},
			err:   errors.New("(inventory::AddVar)", "Var 'key' already exist"),
			key:   "key",
			value: "value",
			res:   nil,
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.options.AddVar(test.key, test.value)

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, test.options.Vars, "Unexpected options value")
			}
		})
	}

}

// TestAddVarsFile tests
func TestAddVarsFile(t *testing.T) {

	tests := []struct {
		desc    string
		file    string
		options *AnsibleInventoryOptions
		res     []string
		err     error
	}{
		{
			desc:    "Testing add an vars file when VarsFile is nil",
			file:    "@test.yml",
			options: &AnsibleInventoryOptions{},
			res:     []string{"@test.yml"},
			err:     &errors.Error{},
		},
		{
			desc: "Testing add an vars file",
			file: "@test2.yml",
			options: &AnsibleInventoryOptions{
				VarsFile: []string{"@test1.yml"},
			},
			res: []string{"@test1.yml", "@test2.yml"},
			err: &errors.Error{},
		},
		{
			desc: "Testing add an vars file without file mark prefix @",
			file: "test.yml",
			options: &AnsibleInventoryOptions{
				VarsFile: []string{},
			},
			res: []string{"@test.yml"},
			err: &errors.Error{},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.options.AddVarsFile(test.file)

			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, test.options.VarsFile, "Unexpected options value")
			}
		})
	}
}

// TestAddVaultedVar tests
func TestAddVaultedVar(t *testing.T) {
	vaulter := vault.NewMockVariableVaulter()
	vaulter.On("Vault", "plain_text_value").Return(vault.NewVaultVariableValue("encrypted_value"), nil)

	tests := []struct {
		desc    string
		options *AnsibleInventoryOptions
		vaulter Vaulter
		name    string
		value   string
		res     map[string]interface{}
		err     error
	}{
		{
			desc:    "Testing add a vaulted extra-var",
			options: &AnsibleInventoryOptions{},
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
			options: &AnsibleInventoryOptions{},
			vaulter: nil,
			err:     errors.New("(inventory::AddVaultedVar)", "To define a vaulted var you need to initialize a vaulter"),
		},
		{
			desc: "Testing error adding a vaulted extra-var when variable already exist",
			options: &AnsibleInventoryOptions{
				Vars: map[string]interface{}{
					"variable_name": "{\"__ansible_vault\":\"encrypted_value\"}",
				},
			},
			vaulter: vaulter,
			name:    "variable_name",
			value:   "plain_text_value",
			res:     map[string]interface{}{},
			err:     errors.New("(inventory::AddVaultedVar)", "Var 'variable_name' already exist"),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			err := test.options.AddVaultedVar(test.vaulter, test.name, test.value)
			if err != nil {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, test.options.Vars, "Unexpected options value")
			}
		})
	}
}

// TestGenerateVerbosityFlag tests
func TestGenerateVerbosityFlag(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsibleInventoryOptions
		res     string
		err     error
	}{
		{
			desc: "Testing generate verbosity flag",
			options: &AnsibleInventoryOptions{
				Verbose: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag V",
			options: &AnsibleInventoryOptions{
				VerboseV: true,
			},
			res: "-v",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV",
			options: &AnsibleInventoryOptions{
				VerboseVV: true,
			},
			res: "-vv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVV",
			options: &AnsibleInventoryOptions{
				VerboseVVV: true,
			},
			res: "-vvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVVV",
			options: &AnsibleInventoryOptions{
				VerboseVVVV: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV has precedence over V",
			options: &AnsibleInventoryOptions{
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
