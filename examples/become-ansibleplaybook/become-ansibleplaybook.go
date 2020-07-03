package main

import (
	ansibler "github.com/apenella/go-ansible"
)

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		User: "aleix",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "192.168.0.19,",
	}

	ansiblePlaybookPrivilegeEscalationOptions := &ansibler.AnsiblePlaybookPrivilegeEscalationOptions{
		Become:        true,
		AskBecomePass: true,
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:                   "site.yml",
		ConnectionOptions:          ansiblePlaybookConnectionOptions,
		PrivilegeEscalationOptions: ansiblePlaybookPrivilegeEscalationOptions,
		Options:                    ansiblePlaybookOptions,
		ExecPrefix:                 "Go-ansible example with become",
	}

	err := playbook.Run()
	if err != nil {
		panic(err)
	}
}
