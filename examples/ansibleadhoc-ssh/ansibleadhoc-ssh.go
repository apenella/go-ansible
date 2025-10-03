package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/adhoc"
	"github.com/apenella/go-ansible/v2/pkg/execute"
)

func main() {
	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:     "inventory.yml",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -T",
		ModuleName:    "shell",
		Args:          "hostname",
	}

	cmd := adhoc.NewAnsibleAdhocCmd(
		adhoc.WithPattern("all"),
		adhoc.WithAdhocOptions(ansibleAdhocOptions),
	)

	c, _ := cmd.Command()
	fmt.Println("Command: ", c)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(cmd),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
