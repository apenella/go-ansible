package ansible

import (
	"oc-images-utils/ansible"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnsible(t *testing.T) {

	playbook := &ansible.AnsiblePlaybookCmd{
		Playbook: "test/ansible/test_site.yml",
		ConnectionOptions: &ansible.AnsiblePlaybookConnectionOptions{
			Connection: "local",
		},
		Options: &ansible.AnsiblePlaybookOptions{
			Inventory: "test/ansible/inventory/all",
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
