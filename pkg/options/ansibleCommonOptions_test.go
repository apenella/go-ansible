package options

import (
	"testing"

	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestGenerateCommandCommonOptions(t *testing.T) {

	t.Log("Testing generate command common options")

	options := &AnsibleCommonOptions{
		AskVaultPassword: true,
		Check:            true,
		Diff:             true,
		ExtraVars: map[string]interface{}{
			"var1": "value1",
			"var2": false,
		},
		Forks:             "10",
		Inventory:         "127.0.0.1,",
		Limit:             "myhost",
		ListHosts:         true,
		ModulePath:        "/dev/null",
		SyntaxCheck:       true,
		VaultID:           "asdf",
		VaultPasswordFile: "/dev/null",
		Verbose:           true,
		Version:           true,
	}

	opts, _ := options.GenerateCommandCommonOptions()

	expected := []string{
		"--ask-vault-password",
		"--check",
		"--diff",
		"--extra-vars",
		"{\"var1\":\"value1\",\"var2\":false}",
		"--forks",
		"10",
		"--inventory",
		"127.0.0.1,",
		"--limit",
		"myhost",
		"--list-hosts",
		"--module-path",
		"/dev/null",
		"--syntax-check",
		"--vault-id",
		"asdf",
		"--vault-password-file",
		"/dev/null",
		"-vvvv",
		"--version",
	}

	assert.Equal(t, expected, opts)

}

func TestGenerateExtraVarsCommand(t *testing.T) {

	tests := []struct {
		desc      string
		options   *AnsibleCommonOptions
		err       error
		extravars string
	}{
		{
			desc: "Testing extra vars map[string]string",
			options: &AnsibleCommonOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:       nil,
			extravars: "{\"extra\":\"var\"}",
		},
		{
			desc: "Testing extra vars map[string]bool",
			options: &AnsibleCommonOptions{
				ExtraVars: map[string]interface{}{
					"extra": true,
				},
			},
			err:       nil,
			extravars: "{\"extra\":true}",
		},
		{
			desc: "Testing extra vars map[string]int",
			options: &AnsibleCommonOptions{
				ExtraVars: map[string]interface{}{
					"extra": 10,
				},
			},
			err:       nil,
			extravars: "{\"extra\":10}",
		},
		{
			desc: "Testing extra vars map[string][]string",
			options: &AnsibleCommonOptions{
				ExtraVars: map[string]interface{}{
					"extra": []string{"var"},
				},
			},
			err:       nil,
			extravars: "{\"extra\":[\"var\"]}",
		},
		{
			desc: "Testing extra vars map[string]map[string]string",
			options: &AnsibleCommonOptions{
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

func TestString(t *testing.T) {

	t.Log("Testing generate command common options string")

	options := &AnsibleCommonOptions{
		AskVaultPassword: true,
		Check:            true,
		Diff:             true,
		ExtraVars: map[string]interface{}{
			"var1": "value1",
			"var2": false,
		},
		Forks:             "10",
		Inventory:         "127.0.0.1,",
		Limit:             "myhost",
		ListHosts:         true,
		ModulePath:        "/dev/null",
		SyntaxCheck:       true,
		VaultID:           "asdf",
		VaultPasswordFile: "/dev/null",
		Verbose:           true,
		Version:           true,
	}

	cmd := options.String()

	expected := " --ask-vault-password --check --diff --extra-vars '{\"var1\":\"value1\",\"var2\":false}' --forks 10 --inventory 127.0.0.1, --limit myhost --list-hosts --module-path /dev/null --syntax-check --vault-id asdf --vault-password-file /dev/null -vvvv --version"

	assert.Equal(t, expected, cmd)
}

func TestAddExtraVar(t *testing.T) {
	tests := []struct {
		desc          string
		options       *AnsibleCommonOptions
		err           error
		extraVarName  string
		extraVarValue interface{}
		res           map[string]interface{}
	}{
		{
			desc: "Testing add an extraVar to a nil data structure",
			options: &AnsibleCommonOptions{
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
			options: &AnsibleCommonOptions{
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
			options: &AnsibleCommonOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:           errors.New("(ansible::AddExtraVar)", "ExtraVar 'extra' already exist"),
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
