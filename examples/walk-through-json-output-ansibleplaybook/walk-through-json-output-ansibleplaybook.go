package main

import (
	"bytes"
	"context"
	"encoding/json"
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
		execute.WithWrite(io.Writer(output)),
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

	msgOutput := struct {
		Host    string `json:"host"`
		Message string `json:"message"`
	}{}

	for _, current := range res {
		fmt.Println(current.Playbook)
		for _, play := range current.Plays {
			for _, task := range play.Tasks {
				if task.Task.Name == "walk-through-json-output-ansibleplaybook" {
					for _, content := range task.Hosts {

						err = json.Unmarshal([]byte(content.Stdout), &msgOutput)
						if err != nil {
							panic(err)
						}

						fmt.Printf("[%s] %s\n", msgOutput.Host, msgOutput.Message)
					}
				}
			}
		}
	}
}
