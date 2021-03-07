package main

import (
	"context"

	ansibler "github.com/apenella/go-ansible"
	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/stdoutcallback/results"
)

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.AnsibleConnectionOptions{
		User: "apenella",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	ansiblePlaybookPrivilegeEscalationOptions := &ansibler.AnsiblePrivilegeEscalationOptions{
		Become:        true,
		AskBecomePass: true,
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:                   "site.yml",
		ConnectionOptions:          ansiblePlaybookConnectionOptions,
		PrivilegeEscalationOptions: ansiblePlaybookPrivilegeEscalationOptions,
		Options:                    ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithTransformers(
				results.Prepend("Go-ansible example with become"),
			),
		),
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
