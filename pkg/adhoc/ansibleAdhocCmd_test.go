package adhoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewAnsibleAdhocCmd tests
func TestNewAnsibleAdhocCmd(t *testing.T) {
	t.Log("Testing create new AnsibleAdhocCmd")

	tests := []struct {
		desc                    string
		expectedAnsibleAdhocCmd *AnsibleAdhocCmd
		binary                  string
		pattern                 string
		adhocOptions            *AnsibleAdhocOptions
	}{
		{
			desc:    "Testing create new AnsibleAdhocCmd",
			binary:  "custom-binary",
			pattern: "pattern",
			adhocOptions: &AnsibleAdhocOptions{
				Args:       "args",
				Background: 11,
				ModuleName: "module-name",
				OneLine:    true,
			},
			expectedAnsibleAdhocCmd: &AnsibleAdhocCmd{
				Binary:  "custom-binary",
				Pattern: "pattern",
				AdhocOptions: &AnsibleAdhocOptions{
					Args:       "args",
					Background: 11,
					ModuleName: "module-name",
					OneLine:    true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			ansibleAdhocCmd := NewAnsibleAdhocCmd(
				WithBinary(test.binary),
				WithPattern(test.pattern),
				WithAdhocOptions(test.adhocOptions),
			)

			assert.Equal(t, test.expectedAnsibleAdhocCmd, ansibleAdhocCmd, "Unexpected value")
		})
	}

}

func TestCommand(t *testing.T) {
	t.Log("Testing generate ansible adhoc command string")

	adhoc := &AnsibleAdhocCmd{
		Binary:  "custom-binary",
		Pattern: "pattenr",
		AdhocOptions: &AnsibleAdhocOptions{
			Args:        "args",
			Background:  11,
			ModuleName:  "module-name",
			OneLine:     true,
			PlaybookDir: "playbook-dir",
			Poll:        12,
			Tree:        "tree",
			Connection:  "local",
		},
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
		Pattern: "pattern",
		AdhocOptions: &AnsibleAdhocOptions{
			Args:        "args",
			Background:  11,
			ModuleName:  "module-name",
			OneLine:     true,
			PlaybookDir: "playbook-dir",
			Poll:        12,
			Tree:        "tree",
			Connection:  "local",
		},
	}

	expected := "custom-binary pattern  --args 'args' --background 11 --module-name module-name --one-line --playbook-dir playbook-dir --poll 12 --tree tree --connection local"

	res := adhoc.String()

	assert.Equal(t, expected, res)
}
