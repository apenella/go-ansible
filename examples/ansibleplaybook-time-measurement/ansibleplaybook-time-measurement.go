package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/apenella/go-ansible/v2/pkg/execute/measure"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := measure.NewExecutorTimeMeasurement(
		configuration.NewAnsibleWithConfigurationSettingsExecute(
			execute.NewDefaultExecute(
				execute.WithCmd(playbookCmd),
				execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			),
			configuration.WithAnsibleForceColor(),
		),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Duration: ", exec.Duration().String())
}
