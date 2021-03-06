package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	ansibler "github.com/apenella/go-ansible"
	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/stdoutcallback/results"
)

func main() {

	var timeout int
	flag.IntVar(&timeout, "timeout", 15, "Timeout in seconds")
	flag.Parse()

	fmt.Printf("Timeout: %d seconds\n", timeout)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithTransformers(
				results.Prepend("Go-ansible example"),
			),
		),
		StdoutCallback: "json",
	}

	err := playbook.Run(ctx)
	if err != nil {
		panic(err)
	}
}
