package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	ansibler "github.com/apenella/go-ansible"
	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/stdoutcallback/results"
)

func main() {

	var err error
	res := &results.AnsiblePlaybookJSONResults{}
	buff := new(bytes.Buffer)

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	execute := execute.NewDefaultExecute(
		execute.WithWrite(io.Writer(buff)),
	)

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
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