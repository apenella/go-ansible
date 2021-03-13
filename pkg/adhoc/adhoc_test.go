package adhoc

import (
	"bytes"
	"context"
	goerrors "errors"
	"io"
	execerrors "os/exec"
	"testing"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var w bytes.Buffer

	tests := []struct {
		desc            string
		ansibleAdhocCmd *AnsibleAdhocCmd
		res             string
		err             error
	}{
		{
			desc:            "Testing run an adhoc command with a nil AnsibleAdhocCmd",
			ansibleAdhocCmd: nil,
			err:             errors.New("(adhoc::Run)", "AnsibleAdhocCmd is nil"),
		},
		{
			desc: "Testing run an adhoc command with unexisting binary file",
			ansibleAdhocCmd: &AnsibleAdhocCmd{
				Binary: "unexisting",
			},
			err: errors.New("(adhoc::Run)", "Binary file 'unexisting' does not exists", &execerrors.Error{Name: "unexisting", Err: goerrors.New("executable file not found in $PATH")}),
		},
		{
			desc: "Testing run an adhoc command",
			ansibleAdhocCmd: &AnsibleAdhocCmd{
				Pattern: "all",
				Exec: execute.NewDefaultExecute(
					execute.WithWrite(io.Writer(&w)),
				),
				Options: &AnsibleAdhocOptions{
					Inventory:  "127.0.0.1,",
					ModuleName: "ping",
				},
				ConnectionOptions: &options.AnsibleConnectionOptions{
					Connection: "local",
				},
				StdoutCallback: "oneline",
			},
			res: `127.0.0.1 | SUCCESS => {"changed": false,"ping": "pong"}
Playbook run took 0 days, 0 hours, 0 minutes, 0 seconds
`,
		},
	}

	for _, test := range tests {
		w = bytes.Buffer{}
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)
			w.Reset()

			err := test.ansibleAdhocCmd.Run(context.TODO())
			if err != nil && assert.Error(t, err) {
				assert.Equal(t, test.err, err)
			} else {
				assert.Equal(t, test.res, w.String(), "Unexpected value")
			}
		})

	}

}

func TestCommand(t *testing.T) {
	t.Log("Testing generate ansible adhoc command string")

	adhoc := &AnsibleAdhocCmd{
		Binary:  "custom-binary",
		Pattern: "pattenr",
		Options: &AnsibleAdhocOptions{
			Args:        "args",
			Background:  11,
			ModuleName:  "module-name",
			OneLine:     true,
			PlaybookDir: "playbook-dir",
			Poll:        12,
			Tree:        "tree",
		},
		ConnectionOptions: &options.AnsibleConnectionOptions{
			Connection: "local",
		},
		PrivilegeEscalationOptions: &options.AnsiblePrivilegeEscalationOptions{},
		StdoutCallback:             "oneline",
	}

	expected := []string{
		"custom-binary",
		"pattenr",
		"--args",
		"args",
		"--background",
		"11",
		"--module-name",
		"module-name",
		"--one-line",
		"--playbook-dir",
		"playbook-dir",
		"--poll",
		"12",
		"--tree",
		"tree",
		"--connection",
		"local",
	}

	res, _ := adhoc.Command()

	assert.Equal(t, expected, res)
}

func TestString(t *testing.T) {

	t.Log("Testing generate ansible adhoc command string")

	adhoc := &AnsibleAdhocCmd{
		Binary:  "custom-binary",
		Pattern: "pattenr",
		Options: &AnsibleAdhocOptions{
			Args:        "args",
			Background:  11,
			ModuleName:  "module-name",
			OneLine:     true,
			PlaybookDir: "playbook-dir",
			Poll:        12,
			Tree:        "tree",
		},
		ConnectionOptions: &options.AnsibleConnectionOptions{
			Connection: "local",
		},
		PrivilegeEscalationOptions: &options.AnsiblePrivilegeEscalationOptions{},
		StdoutCallback:             "oneline",
	}

	expected := "custom-binary pattenr  --args args --background 11 --module-name module-name --one-line --playbook-dir playbook-dir --poll 12 --tree tree  --connection local "

	res := adhoc.String()

	assert.Equal(t, expected, res)
}

func TestGenerateAnsibleAdhocOptions(t *testing.T) {

	t.Log("Testing generate ansible command options")

	options := &AnsibleAdhocOptions{

		Args:        "args",
		Background:  11,
		ModuleName:  "module-name",
		OneLine:     true,
		PlaybookDir: "playbook-dir",
		Poll:        12,
		Tree:        "tree",

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

func TestAnsibleAdhocString(t *testing.T) {

	t.Log("Testing generate ansible command options string")

	options := &AnsibleAdhocOptions{

		Args:        "args",
		Background:  11,
		ModuleName:  "module-name",
		OneLine:     true,
		PlaybookDir: "playbook-dir",
		Poll:        12,
		Tree:        "tree",

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

	expected := " --args args --ask-vault-password --background 11 --check --diff --extra-vars '{\"var1\":\"value1\",\"var2\":false}' --forks 10 --inventory 127.0.0.1, --limit myhost --list-hosts --module-name module-name --module-path /dev/null --one-line --playbook-dir playbook-dir --poll 12 --syntax-check --tree tree --vault-id asdf --vault-password-file /dev/null -vvvv --version"

	assert.Equal(t, expected, cmd)
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
