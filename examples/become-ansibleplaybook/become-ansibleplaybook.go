package main

import (
	ansibler "github.com/apenella/go-ansible"
)

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		User: "apenella",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
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
