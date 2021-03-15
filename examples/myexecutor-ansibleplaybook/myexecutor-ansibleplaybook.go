package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback"
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

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	exe := &MyExecutor{}
	exe.Options(
		WithPrefix("[Go ansible example]"),
	)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              exe,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
