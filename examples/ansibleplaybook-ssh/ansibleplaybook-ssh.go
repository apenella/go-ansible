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
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
	)

	fmt.Println("Executing command: ", cmd.String())

	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
