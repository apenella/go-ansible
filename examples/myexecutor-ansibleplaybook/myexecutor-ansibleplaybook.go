package main

import (
	"fmt"

	ansibler "github.com/apenella/go-ansible"
)

type MyExecutor struct{}

func (e *MyExecutor) Execute(command string, args []string, prefix string) error {
	fmt.Println("I am doing nothing")

	return nil
}

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.PlaybookConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &ansibler.PlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.PlaybookCmd{
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
