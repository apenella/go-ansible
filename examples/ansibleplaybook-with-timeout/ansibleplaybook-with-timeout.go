package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "Timeout in seconds")
	flag.Parse()

	fmt.Printf("Timeout: %d seconds\n", timeout)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		User:       "apenella",
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		execute.WithTransformers(
			transformer.Prepend("Go-ansible example"),
		),
	)

	err := exec.Execute(ctx)
	if err != nil {
		panic(err)
	}
}
