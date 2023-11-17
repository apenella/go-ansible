package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "Timeout in seconds")
	flag.Parse()

	fmt.Printf("Timeout: %d seconds\n", timeout)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbook),
		execute.WithTransformers(
			transformer.Prepend("Go-ansible example"),
		),
	)

	err := exec.Execute(ctx)
	if err != nil {
		panic(err)
	}
}
