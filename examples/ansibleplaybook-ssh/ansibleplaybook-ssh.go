package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory:     "inventory.yml",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
	}

	cmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(cmd),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(cmd.String())

}
