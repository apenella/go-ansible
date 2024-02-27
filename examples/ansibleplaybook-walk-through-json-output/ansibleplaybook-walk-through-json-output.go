package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/apenella/go-ansible/pkg/execute"
	results "github.com/apenella/go-ansible/pkg/execute/result/json"
	"github.com/apenella/go-ansible/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	var err error
	var res *results.AnsiblePlaybookJSONResults

	buff := new(bytes.Buffer)

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		User:       "apenella",
		Inventory:  "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"site1.yml", "site2.yml"},
		PlaybookOptions: ansiblePlaybookOptions,
	}

	exec := stdoutcallback.NewJSONStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbook),
			execute.WithWrite(io.Writer(buff)),
		),
	)

	err = exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err = results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {
		panic(err)
	}

	msgOutput := struct {
		Host    string `json:"host"`
		Message string `json:"message"`
	}{}

	for _, play := range res.Plays {
		for _, task := range play.Tasks {
			if task.Task.Name == "ansibleplaybook-walk-through-json-output" {
				for _, content := range task.Hosts {

					err = json.Unmarshal([]byte(fmt.Sprint(content.Stdout)), &msgOutput)
					if err != nil {
						panic(err)
					}

					fmt.Printf("[%s] %s\n", msgOutput.Host, msgOutput.Message)
				}
			}
		}
	}
}
