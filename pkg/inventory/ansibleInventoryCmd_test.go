package inventory

import (
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

// TestNewAnsibleInventoryCmd tests
func TestNewAnsibleInventoryCmd(t *testing.T) {

	tests := []struct {
		desc                        string
		expectedAnsibleInventoryCmd *AnsibleInventoryCmd
		binary                      string
		pattern                     string
		ansibleInventoryOptions     *AnsibleInventoryOptions
	}{
		{
			desc:    "Testing create new AnsibleInventoryCmd",
			binary:  "custom-binary",
			pattern: "pattern",
			ansibleInventoryOptions: &AnsibleInventoryOptions{
				AskVaultPassword:  true,
				Export:            true,
				Graph:             true,
				Host:              "localhost",
				Inventory:         "test/ansible/inventory/all",
				Limit:             "limit",
				List:              true,
				Output:            "/tmp/output.ini",
				PlaybookDir:       "playbook-dir",
				Toml:              true,
				Vars:              true,
				VaultID:           "vault-id",
				VaultPasswordFile: "vault-password-file",
				Verbose:           true,
				Version:           true,
				Yaml:              true,
			},
			expectedAnsibleInventoryCmd: &AnsibleInventoryCmd{
				Binary:  "custom-binary",
				Pattern: "pattern",
				InventoryOptions: &AnsibleInventoryOptions{
					AskVaultPassword:  true,
					Export:            true,
					Graph:             true,
					Host:              "localhost",
					Inventory:         "test/ansible/inventory/all",
					Limit:             "limit",
					List:              true,
					Output:            "/tmp/output.ini",
					PlaybookDir:       "playbook-dir",
					Toml:              true,
					Vars:              true,
					VaultID:           "vault-id",
					VaultPasswordFile: "vault-password-file",
					Verbose:           true,
					Version:           true,
					Yaml:              true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			ansibleInventoryCmd := NewAnsibleInventoryCmd(
				WithBinary(test.binary),
				WithPattern(test.pattern),
				WithInventoryOptions(test.ansibleInventoryOptions),
			)

			assert.Equal(t, test.expectedAnsibleInventoryCmd, ansibleInventoryCmd, "Unexpected value")
		})
	}
}

// TestGenerateCommandOptions tests
func TestGenerateCommandOptions(t *testing.T) {
	tests := []struct {
		desc                    string
		ansibleInventoryOptions *AnsibleInventoryOptions
		err                     error
		options                 []string
	}{
		{
			desc:                    "Testing nil AnsibleInventoryOptions definition",
			ansibleInventoryOptions: nil,
			err:                     errors.New("(inventory::GenerateCommandOptions)", "AnsibleInventoryOptions is nil"),
			options:                 nil,
		},
		{
			desc:                    "Testing an empty AnsibleInventoryOptions definition",
			ansibleInventoryOptions: &AnsibleInventoryOptions{},
			err:                     nil,
			options:                 []string{},
		},
		{
			desc: "Testing AnsibleInventoryOptions except vars",
			ansibleInventoryOptions: &AnsibleInventoryOptions{
				Host:      "localhost",
				Inventory: "test/ansible/inventory/all",
				Output:    "/tmp/output.ini",
			},
			err:     nil,
			options: []string{"--host", "localhost", "--inventory", "test/ansible/inventory/all", "--output", "/tmp/output.ini"},
		},
		{
			desc: "Testing AnsibleInventoryOptions with vars",
			ansibleInventoryOptions: &AnsibleInventoryOptions{
				Vars:      true,
				Host:      "localhost",
				Inventory: "test/ansible/inventory/all",
				Output:    "/tmp/output.ini",
			},
			err:     nil,
			options: []string{"--host", "localhost", "--inventory", "test/ansible/inventory/all", "--output", "/tmp/output.ini", "--vars"},
		},
	}

	for _, test := range tests {

		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			options, err := test.ansibleInventoryOptions.GenerateCommandOptions()

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
			desc: "Testing AnsibleInventoryCmd to string",
			err:  nil,
			ansibleInventoryCmd: &AnsibleInventoryCmd{
				Binary:  "ansible-inventory",
				Pattern: "all",
				InventoryOptions: &AnsibleInventoryOptions{
					AskVaultPassword:  true,
					Export:            true,
					Graph:             true,
					Host:              "localhost",
					Inventory:         "test/ansible/inventory/all",
					Limit:             "myhost",
					List:              true,
					Output:            "/tmp/output.ini",
					PlaybookDir:       "/playbook/",
					Toml:              true,
					Vars:              true,
					VaultID:           "asdf",
					VaultPasswordFile: "/vault/password/file",
					Verbose:           true,
					Version:           true,
					Yaml:              true,
				},
			},
			res: "ansible-inventory all  --ask-vault-password --export --graph --host localhost --inventory test/ansible/inventory/all --limit myhost --list --output /tmp/output.ini --playbook-dir /playbook/ --toml --vars --vault-id asdf --vault-password-file /vault/password/file -vvvv --version --yaml",
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
				InventoryOptions: &AnsibleInventoryOptions{
					AskVaultPassword:  true,
					Export:            true,
					Graph:             true,
					Host:              "localhost",
					Inventory:         "test/ansible/inventory/all",
					Limit:             "limit",
					List:              true,
					Output:            "/tmp/output.ini",
					PlaybookDir:       "playbook-dir",
					Toml:              true,
					Vars:              true,
					VaultID:           "vault-id",
					VaultPasswordFile: "vault-password-file",
					Verbose:           true,
					Version:           true,
					Yaml:              true,
				},
			},
			command: []string{
				"ansible-inventory",
				"all",
				"--ask-vault-password",
				"--export",
				"--graph",
				"--host",
				"localhost",
				"--inventory",
				"test/ansible/inventory/all",
				"--limit",
				"limit",
				"--list",
				"--output",
				"/tmp/output.ini",
				"--playbook-dir",
				"playbook-dir",
				"--toml",
				"--vars",
				"--vault-id",
				"vault-id",
				"--vault-password-file",
				"vault-password-file",
				"--version",
				"-vvvv",
				"--yaml",
			},
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
