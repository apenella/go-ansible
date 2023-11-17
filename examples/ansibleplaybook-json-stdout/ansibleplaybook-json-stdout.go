package main

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	results "github.com/apenella/go-ansible/pkg/execute/result/json"
	"github.com/apenella/go-ansible/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	var err error
	var res *results.AnsiblePlaybookJSONResults

	buff := new(bytes.Buffer)

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbooksList := []string{"site1.yml", "site2.yml", "site3.yml"}
	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         playbooksList,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	exec := measure.NewExecutorTimeMeasurement(
		stdoutcallback.NewJSONStdoutCallbackExecute(
			execute.NewDefaultExecute(
				execute.WithCmd(playbook),
				execute.WithWrite(io.Writer(buff)),
			),
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

	fmt.Println(res.String())
	fmt.Println("Duration: ", exec.Duration().String())

}
