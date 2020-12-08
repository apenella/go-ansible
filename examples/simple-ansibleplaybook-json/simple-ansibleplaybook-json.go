package main

import (
	"bytes"
	"fmt"
	"io"

	ansibler "github.com/apenella/go-ansible"
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

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		ExecPrefix:        "Go-ansible example",
		StdoutCallback:    "json",
		Writer:            io.Writer(buff),
	}

	err = playbook.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err = results.JSONParse(buff.Bytes())
	if err != nil {
		panic(err)
	}

	fmt.Println(res.String())
}
