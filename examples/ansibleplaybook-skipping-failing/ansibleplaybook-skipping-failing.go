package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	results "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	var err error
	var res *results.AnsiblePlaybookJSONResults

	buff := new(bytes.Buffer)
	timeBuff := new(bytes.Buffer)

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbooksList := []string{"site1.yml"}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbooksList...),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	log.Println("Command: ", playbookCmd.String())

	exec := stdoutcallback.NewJSONStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
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

	for _, play := range res.Plays {
		for _, task := range play.Tasks {
			for _, content := range task.Hosts {
				if task.Task.Name == "skipping-task" {
					fmt.Printf("Task [%s] skipped [%t] with skip reason [%s]\n", task.Task.Name, content.Skipped, content.SkipReason)
				} else {
					fmt.Printf("Task [%s] failed [%t] with condition [%t]. Executed cmd: %v\n",
						task.Task.Name, content.Failed, content.FailedWhenResult, content.Cmd)
				}

			}
		}
	}

	fmt.Println(timeBuff.String())
}
