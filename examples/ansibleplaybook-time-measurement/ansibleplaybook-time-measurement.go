package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/configuration"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	exec := measure.NewExecutorTimeMeasurement(
		configuration.NewExecutorWithAnsibleConfigurationSettings(
			execute.NewDefaultExecute(
				execute.WithCmd(playbook),
			),
		).WithAnsibleForceColor(),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Duration: ", exec.Duration().String())
}
