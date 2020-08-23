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


	res, err := playbook.Run()
	if err != nil && assert.Error(t, err) {
		fmt.Println(err.Error())
		assert.Equal(t, nil, err)
	}
	err = res.PlaybookResultsChecks()
	if err == nil {
		fmt.Println(res.RawStdout)
	}


}
