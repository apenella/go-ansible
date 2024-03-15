package main

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("input.yml"),
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
