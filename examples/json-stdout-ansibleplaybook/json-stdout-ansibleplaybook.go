package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
)

func main() {
	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	output := io.ReadWriter(new(bytes.Buffer))
	execute := execute.NewDefaultExecute(
		execute.WithWrite(output),
	)

	playbookNames := []string{"site.yml"}
	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         playbookNames,
		Exec:              execute,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		StdoutCallback:    "json",
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err := results.JSONParse(playbookNames, output)
	if err != nil {
		panic(err)
	}

	for _, current := range res {
		fmt.Println(current.Playbook)
		fmt.Println(current.String())
	}
}
