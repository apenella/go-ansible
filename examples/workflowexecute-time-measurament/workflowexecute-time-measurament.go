package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute/measure"
	"github.com/apenella/go-ansible/v2/pkg/execute/workflow"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	first := playbook.NewAnsiblePlaybookExecute("first.yml").
		WithPlaybookOptions(ansiblePlaybookOptions)

	second := playbook.NewAnsiblePlaybookExecute("second.yml").
		WithPlaybookOptions(ansiblePlaybookOptions)

	exec := measure.NewExecutorTimeMeasurement(
		workflow.NewWorkflowExecute(first, second),
	)
	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Duration: ", exec.Duration().String())
}
