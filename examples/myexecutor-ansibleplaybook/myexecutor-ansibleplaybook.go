package main

import (
	"fmt"

	ansibler "github.com/apenella/go-ansible"
)

type MyExecutor struct{}

func (e *MyExecutor) SetCmdRunDir(string) {}

func (e *MyExecutor) Execute(command string, args []string, prefix string) error {
	fmt.Println("I am doing nothing")

	return nil
}

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              &MyExecutor{},
	}

	err := playbook.Run()
	if err != nil {
		panic(err)
	}
}
