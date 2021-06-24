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

	var err error
	var res *results.AnsiblePlaybookJSONResults

	buff := new(bytes.Buffer)

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	execute := execute.NewDefaultExecute(
		execute.WithWrite(io.Writer(buff)),
	)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		Exec:              execute,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		StdoutCallback:    "json",
	}

	err = playbook.Run(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err = results.JSONParse(buff.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println(res.String())
}
