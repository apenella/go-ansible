package main

import (
	"context"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/configuration"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/fatih/color"
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
				execute.WithTransformers(
					transformer.Prepend("Go-ansible example"),
					transformer.LogFormat(transformer.DefaultLogFormatLayout, transformer.Now),
				),
			),
		).WithAnsibleForceColor(),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}

	color.Cyan("\n\tDuration: %s\n\n", exec.Duration())
}
