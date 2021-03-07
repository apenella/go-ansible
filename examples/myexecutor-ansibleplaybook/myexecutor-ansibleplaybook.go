package main

import (
	"context"
	"fmt"

	ansibler "github.com/apenella/go-ansible"
	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/stdoutcallback"
)

type MyExecutor struct {
	Prefix string
}

func (e *MyExecutor) Options(options ...execute.ExecuteOptions) {
	// apply all options to the executor
	for _, opt := range options {
		opt(e)
	}
}

func WithPrefix(prefix string) execute.ExecuteOptions {
	return func(e execute.Executor) {
		e.(*MyExecutor).Prefix = prefix
	}
}

func (e *MyExecutor) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...execute.ExecuteOptions) error {

	// apply all options to the executor
	for _, opt := range options {
		opt(e)
	}

	fmt.Println(fmt.Sprintf("%s %s\n", e.Prefix, "I am MyExecutor and I am doing nothing"))

	return nil
}

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	exe := &MyExecutor{}
	exe.Options(
		WithPrefix("[Go ansible example]"),
	)

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              exe,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
