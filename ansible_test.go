package ansibler

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestAnsible(t *testing.T) {

	playbook := &PlaybookCmd{
		Playbook: "test/test_site.yml",
		ConnectionOptions: &PlaybookConnectionOptions{
			Connection: "local",
		},
		Options: &PlaybookOptions{
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

	res := &PlaybookResults{}
	res, err := playbook.Run()
	fmt.Println(res.Changed)
	if err != nil && assert.Error(t, err) {
		assert.Equal(t, nil, err)
	}
}
