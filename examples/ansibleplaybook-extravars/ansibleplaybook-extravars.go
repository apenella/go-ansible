package main

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		User:       "apenella",
		Inventory:  "127.0.0.1,",
		ExtraVars: map[string]interface{}{
			"extravar1": "value11",
			"extravar2": "value12",
		},
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
