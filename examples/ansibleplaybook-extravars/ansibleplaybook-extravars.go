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
		ExtraVars: map[string]interface{}{
			"extravar1":    "value11",
			"extravar2":    "value12",
			"ansible_port": "22225",
		},
		Inventory:     "inventory.yml",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	fmt.Println("Command: ", playbookCmd.String())

	yamlexec := stdoutcallback.NewYAMLStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		),
	)

	err := yamlexec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
