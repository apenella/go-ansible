package adhoc

import (
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAnsibleAdhocOptions(t *testing.T) {

	t.Log("Testing generate ansible adhoc options")

	options := &AnsibleAdhocOptions{
		Args:             "args",
		AskBecomePass:    true,
		AskPass:          true,
		AskVaultPassword: true,
		Background:       11,
		Become:           true,
		BecomeMethod:     "become-method",
		BecomeUser:       "become-user",
		Check:            true,
		Connection:       "local",
		Diff:             true,
		ExtraVars: map[string]interface{}{
			"var1": "value1",
			"var2": false,
		},
		ExtraVarsFile:     []string{"@extra-vars-file.yml"},
		Forks:             "10",
		Inventory:         "127.0.0.1,",
		Limit:             "myhost",
		ListHosts:         true,
		ModuleName:        "module-name",
		ModulePath:        "/dev/null",
		OneLine:           true,
		PlaybookDir:       "playbook-dir",
		Poll:              12,
		PrivateKey:        "pk",
		SCPExtraArgs:      "scp-extra-args1 scp-extra-args2",
		SFTPExtraArgs:     "sftp-extra-args1 sftp-extra-args2",
		SSHCommonArgs:     "ssh-common-args1 ssh-common-args2",
		SSHExtraArgs:      "ssh-extra-args1 ssh-extra-args2",
		SyntaxCheck:       true,
		Timeout:           10,
		Tree:              "tree",
		User:              "user",
		VaultID:           "asdf",
		VaultPasswordFile: "/dev/null",
		Verbose:           true,
		Version:           true,
	}

	opts, _ := options.GenerateAnsibleAdhocOptions()

	expected := []string{
		"--args", "args",
		"--ask-vault-password",
		"--background",
		"11",
		"--check",
		"--diff",
		"--extra-vars",
		"{\"var1\":\"value1\",\"var2\":false}",
		"--extra-vars",
		"@extra-vars-file.yml",
		"--forks",
		"10",
		"--inventory",
		"127.0.0.1,",
		"--limit",
		"myhost",
		"--list-hosts",
		"--module-name",
		"module-name",
		"--module-path",
		"/dev/null",
		"--one-line",
		"--playbook-dir",
		"playbook-dir",
		"--poll",
		"12",
		"--syntax-check",
		"--tree",
		"tree",
		"--vault-id",
		"asdf",
		"--vault-password-file",
		"/dev/null",
		"-vvvv",
		"--version",
		"--ask-pass",
		"--connection",
		"local",
		"--private-key",
		"pk",
		"--scp-extra-args",
		"scp-extra-args1 scp-extra-args2",
		"--sftp-extra-args",
		"sftp-extra-args1 sftp-extra-args2",
		"--ssh-common-args",
		"ssh-common-args1 ssh-common-args2",
		"--ssh-extra-args",
		"ssh-extra-args1 ssh-extra-args2",
		"--timeout",
		"10",
		"--user",
		"user",
		"--ask-become-pass",
		"--become",
		"--become-method",
		"become-method",
		"--become-user",
		"become-user",
	}

	assert.Equal(t, expected, opts)
}

func TestGenerateExtraVarsCommand(t *testing.T) {

	tests := []struct {
		desc      string
		options   *AnsibleAdhocOptions
		err       error
		extravars string
	}{
		{
			desc: "Testing extra vars map[string]string",
			options: &AnsibleAdhocOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:       nil,
			extravars: "{\"extra\":\"var\"}",
		},
		{
			desc: "Testing extra vars map[string]bool",
			options: &AnsibleAdhocOptions{
				ExtraVars: map[string]interface{}{
					"extra": true,
				},
			},
			err:       nil,
			extravars: "{\"extra\":true}",
		},
		{
			desc: "Testing extra vars map[string]int",
			options: &AnsibleAdhocOptions{
				ExtraVars: map[string]interface{}{
					"extra": 10,
				},
			},
			err:       nil,
			extravars: "{\"extra\":10}",
		},
		{
			desc: "Testing extra vars map[string][]string",
			options: &AnsibleAdhocOptions{
				ExtraVars: map[string]interface{}{
					"extra": []string{"var"},
				},
			},
			err:       nil,
			extravars: "{\"extra\":[\"var\"]}",
		},
		{
			desc: "Testing extra vars map[string]map[string]string",
			options: &AnsibleAdhocOptions{
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

func TestAnsibleAdhocOptionsString(t *testing.T) {

	tests := []struct {
		desc    string
		options *AnsibleAdhocOptions
		res     string
	}{
		{
			desc: "Testing generate ansible adhoc options string",
			options: &AnsibleAdhocOptions{
				Args:             "args",
				AskBecomePass:    true,
				AskPass:          true,
				AskVaultPassword: true,
				Background:       11,
				Become:           true,
				BecomeMethod:     "become-method",
				BecomeUser:       "become-user",
				Check:            true,
				Connection:       "local",
				Diff:             true,
				ExtraVars: map[string]interface{}{
					"var1": "value1",
					"var2": false,
				},
				ExtraVarsFile:     []string{"@test/ansible/extra_vars.yml"},
				Forks:             "10",
				Inventory:         "127.0.0.1,",
				Limit:             "myhost",
				ListHosts:         true,
				ModuleName:        "module-name",
				ModulePath:        "/dev/null",
				OneLine:           true,
				PlaybookDir:       "playbook-dir",
				Poll:              12,
				PrivateKey:        "pk",
				SCPExtraArgs:      "scp-extra-args",
				SFTPExtraArgs:     "sftp-extra-args",
				SSHCommonArgs:     "ssh-common-args",
				SSHExtraArgs:      "ssh-extra-args",
				SyntaxCheck:       true,
				Timeout:           10,
				Tree:              "tree",
				User:              "user",
				VaultID:           "asdf",
				VaultPasswordFile: "/dev/null",
				Verbose:           true,
				Version:           true,
			},
			res: " --args 'args' --ask-vault-password --background 11 --check --diff --extra-vars '{\"var1\":\"value1\",\"var2\":false}' --extra-vars @test/ansible/extra_vars.yml --forks 10 --inventory 127.0.0.1, --limit myhost --list-hosts --module-name module-name --module-path /dev/null --one-line --playbook-dir playbook-dir --poll 12 --syntax-check --tree tree --vault-id asdf --vault-password-file /dev/null -vvvv --version --ask-pass --connection local --private-key pk --scp-extra-args 'scp-extra-args' --sftp-extra-args 'sftp-extra-args' --ssh-common-args 'ssh-common-args' --ssh-extra-args 'ssh-extra-args' --timeout 10 --user user --ask-become-pass --become --become-method become-method --become-user become-user",
		},
		{
			desc: "Testing AnsibleAdhocOptions setting the VerboseV flag as true",
			options: &AnsibleAdhocOptions{
				VerboseV: true,
			},
			res: " -v",
		},
		{
			desc: "Testing AnsibleAdhocOptions setting the VerboseVV flag as true",
			options: &AnsibleAdhocOptions{
				VerboseVV: true,
			},
			res: " -vv",
		},
		{
			desc: "Testing AnsibleAdhocOptions setting the VerboseVVV flag as true",
			options: &AnsibleAdhocOptions{
				VerboseVVV: true,
			},
			res: " -vvv",
		},
		{
			desc: "Testing AnsibleAdhocOptions setting the VerboseVVVV flag as true",
			options: &AnsibleAdhocOptions{
				VerboseVVVV: true,
			},
			res: " -vvvv",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			res := test.options.String()
			assert.Equal(t, test.res, res)
		})
	}

}

func TestAddExtraVar(t *testing.T) {
	tests := []struct {
		desc          string
		options       *AnsibleAdhocOptions
		err           error
		extraVarName  string
		extraVarValue interface{}
		res           map[string]interface{}
	}{
		{
			desc: "Testing add an extraVar to a nil data structure",
			options: &AnsibleAdhocOptions{
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
			options: &AnsibleAdhocOptions{
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
			options: &AnsibleAdhocOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:           errors.New("(adhoc::AddExtraVar)", "ExtraVar 'extra' already exist"),
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
		options *AnsibleAdhocOptions
		res     []string
		err     error
	}{
		{
			desc:    "Testing add an extra-vars file when ExtraVarsFile is nil",
			file:    "@test.yml",
			options: &AnsibleAdhocOptions{},
			res:     []string{"@test.yml"},
			err:     &errors.Error{},
		},
		{
			desc: "Testing add an extra-vars file",
			file: "@test2.yml",
			options: &AnsibleAdhocOptions{
				ExtraVarsFile: []string{"@test1.yml"},
			},
			res: []string{"@test1.yml", "@test2.yml"},
			err: &errors.Error{},
		},
		{
			desc: "Testing add an extra-vars file without file mark prefix @",
			file: "test.yml",
			options: &AnsibleAdhocOptions{
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

func TestGenerateVerbosityFlag(t *testing.T) {
	tests := []struct {
		desc    string
		options *AnsibleAdhocOptions
		res     string
		err     error
	}{
		{
			desc: "Testing generate verbosity flag",
			options: &AnsibleAdhocOptions{
				Verbose: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag V",
			options: &AnsibleAdhocOptions{
				VerboseV: true,
			},
			res: "-v",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV",
			options: &AnsibleAdhocOptions{
				VerboseVV: true,
			},
			res: "-vv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVV",
			options: &AnsibleAdhocOptions{
				VerboseVVV: true,
			},
			res: "-vvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VVVV",
			options: &AnsibleAdhocOptions{
				VerboseVVVV: true,
			},
			res: "-vvvv",
			err: &errors.Error{},
		},
		{
			desc: "Testing generate verbosity flag VV has precedence over V",
			options: &AnsibleAdhocOptions{
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
