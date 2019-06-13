package ansibler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnsible(t *testing.T) {

	playbook := &AnsiblePlaybookCmd{
		Playbook: "test/test_site.yml",
		ConnectionOptions: &AnsiblePlaybookConnectionOptions{
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
	}

	err := playbook.Run()
	if err != nil && assert.Error(t, err) {
		assert.Equal(t, nil, err)
	}
}
