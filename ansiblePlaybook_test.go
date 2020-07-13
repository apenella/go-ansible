package ansibler

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateCommandConnectionOptions
func TestGenerateCommandConnectionOptions(t *testing.T) {
	tests := []struct {
		desc                             string
		ansiblePlaybookConnectionOptions *AnsiblePlaybookConnectionOptions
		err                              error
		options                          []string
	}{
		{
			desc: "Testing generate connection options",
			ansiblePlaybookConnectionOptions: &AnsiblePlaybookConnectionOptions{
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
		t.Log(test.desc)

		options, err := test.ansiblePlaybookConnectionOptions.GenerateCommandConnectionOptions()

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, options, test.options, "Unexpected options value")
		}
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
			err:                    errors.New("(ansible::GenerateCommandOptions) AnsiblePlaybookOptions is nil"),
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
				FlushCache: true,
				Inventory:  "inventory",
				Limit:      "limit",
				ListHosts:  true,
				ListTags:   true,
				ListTasks:  true,
				Tags:       "tags",
			},
			err:     nil,
			options: []string{"--flush-cache", "--inventory", "inventory", "--limit", "limit", "--list-hosts", "--list-tags", "--list-tasks", "--tags", "tags"},
		},
		{
			desc: "Testing AnsiblePlaybookOptions with extra vars",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
				FlushCache: true,
				Inventory:  "inventory",
				Limit:      "limit",
				ListHosts:  true,
				ListTags:   true,
				ListTasks:  true,
				Tags:       "tags",
			},
			err:     nil,
			options: []string{"--flush-cache", "--inventory", "inventory", "--limit", "limit", "--list-hosts", "--list-tags", "--list-tasks", "--tags", "tags", "--extra-vars", "{\"extra\":\"var\"}"},
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		options, err := test.ansiblePlaybookOptions.GenerateCommandOptions()

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, options, test.options, "Unexpected options value")
		}
	}
}

func TestGenerateExtraVarsCommand(t *testing.T) {

	tests := []struct {
		desc                   string
		ansiblePlaybookOptions *AnsiblePlaybookOptions
		err                    error
		extravars              string
	}{
		{
			desc: "Testing extra vars map[string]string",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:       nil,
			extravars: "{\"extra\":\"var\"}",
		},
		{
			desc: "Testing extra vars map[string]bool",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": true,
				},
			},
			err:       nil,
			extravars: "{\"extra\":true}",
		},
		{
			desc: "Testing extra vars map[string]int",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": 10,
				},
			},
			err:       nil,
			extravars: "{\"extra\":10}",
		},
		{
			desc: "Testing extra vars map[string][]string",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": []string{"var"},
				},
			},
			err:       nil,
			extravars: "{\"extra\":[\"var\"]}",
		},
		{
			desc: "Testing extra vars map[string]map[string]string",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
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
		t.Log(test.desc)

		extravars, err := test.ansiblePlaybookOptions.generateExtraVarsCommand()

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, extravars, test.extravars, "Unexpected options value")
		}
	}
}

func TestAddExtraVar(t *testing.T) {
	tests := []struct {
		desc                   string
		ansiblePlaybookOptions *AnsiblePlaybookOptions
		err                    error
		extraVarName           string
		extraVarValue          interface{}
		res                    map[string]interface{}
	}{
		{
			desc: "Testing add an extraVar to a nil data structure",
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
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
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
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
			ansiblePlaybookOptions: &AnsiblePlaybookOptions{
				ExtraVars: map[string]interface{}{
					"extra": "var",
				},
			},
			err:           errors.New("(ansible::AddExtraVar) ExtraVar 'extra' already exist"),
			extraVarName:  "extra",
			extraVarValue: "var",
			res:           nil,
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		err := test.ansiblePlaybookOptions.AddExtraVar(test.extraVarName, test.extraVarValue)

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, test.res, test.ansiblePlaybookOptions.ExtraVars, "Unexpected options value")
		}
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
			desc: "",
			err:  nil,
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Playbook: "test/ansible/site.yml",
				ConnectionOptions: &AnsiblePlaybookConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/ansible/inventory/all",
				},
			},
			command: []string{"ansible-playbook", "--inventory", "test/ansible/inventory/all", "--connection", "local", "test/ansible/site.yml"},
		},
	}

	for _, test := range tests {
		t.Log(test.desc)

		command, err := test.ansiblePlaybookCmd.Command()

		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, command, test.command, "Unexpected value")
		}
	}
}

func TestAnsibleForceColor(t *testing.T) {
	AnsibleForceColor()

	assert.Equal(t, "true", os.Getenv(AnsibleForceColorEnv), "Unexpected value")
}

func TestRun(t *testing.T) {

	var w bytes.Buffer

	tests := []struct {
		desc               string
		ansiblePlaybookCmd *AnsiblePlaybookCmd
		res                string
		err                error
	}{
		{
			desc:               "Run nil ansiblePlaybookCmd",
			ansiblePlaybookCmd: nil,
			res:                "",
			err:                errors.New("(ansible:Run) AnsiblePlaybookCmd is nil"),
		},
		{
			desc: "Run a ansiblePlaybookCmd",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Playbook: "test/test_site.yml",
				ConnectionOptions: &AnsiblePlaybookConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/ansible/inventory/all",
				},
			},
			res: "",
			err: nil,
		},
		{
			desc: "Run a ansiblePlaybookCmd without executor",
			ansiblePlaybookCmd: &AnsiblePlaybookCmd{
				Playbook: "test/test_site.yml",
				ConnectionOptions: &AnsiblePlaybookConnectionOptions{
					Connection: "local",
				},
				Options: &AnsiblePlaybookOptions{
					Inventory: "test/all",
				},
			},
			res: "",
			err: nil,
		},
	}

	for _, test := range tests {
		w = bytes.Buffer{}

		t.Log(test.desc)

		err := test.ansiblePlaybookCmd.Run()
		if err != nil && assert.Error(t, err) {
			assert.Equal(t, test.err, err)
		} else {
			assert.Equal(t, test.res, w.String(), "Unexpected value")
		}
	}
}
