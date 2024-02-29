package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/playbook"
)

type MyExecutor struct {
	Prefix string
	Cmd    execute.Commander
}

func NewMyExecutor(cmd execute.Commander) *MyExecutor {
	return &MyExecutor{
		Cmd: cmd,
	}
}

func (e *MyExecutor) WithPrefix(prefix string) {
	e.Prefix = prefix
}

func (e *MyExecutor) Execute(ctx context.Context) error {
	fmt.Println(fmt.Sprintf("%s %s\n", e.Prefix, "I am MyExecutor and I am doing nothing"))

	return nil
}

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"site.yml"},
		PlaybookOptions: ansiblePlaybookOptions,
	}

	exec := NewMyExecutor(playbook)
	exec.WithPrefix("[Go ansible example]")

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
