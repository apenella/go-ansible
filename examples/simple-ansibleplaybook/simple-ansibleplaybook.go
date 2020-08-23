package main

import (
	ansibler "github.com/apenella/go-ansible"
)

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.PlaybookConnectionOptions{
		Connection: "local",
		User:       "aleix",
	}

	ansiblePlaybookOptions := &ansibler.PlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.PlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		ExecPrefix:        "Go-ansible example",
	}

	err := playbook.Run()
	if err != nil {
		panic(err)
	}
}
