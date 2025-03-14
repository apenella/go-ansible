package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbooksList := []string{"site1.yml", "site2.yml", "site3.yml"}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbooksList...),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := stdoutcallback.NewAnsiblePosixJsonlStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		),
	)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	err := exec.Execute(ctx)
	if err != nil {
		fmt.Println(err)
	}
}
