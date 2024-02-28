package main

import (
	"context"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/configuration"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/fatih/color"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"site.yml"},
		PlaybookOptions: ansiblePlaybookOptions,
	}

	exec := measure.NewExecutorTimeMeasurement(
		configuration.NewAnsibleWithConfigurationSettingsExecute(
			execute.NewDefaultExecute(
				execute.WithCmd(playbook),
				execute.WithTransformers(
					transformer.Prepend("Go-ansible example"),
					transformer.LogFormat(transformer.DefaultLogFormatLayout, transformer.Now),
				),
			),
			configuration.WithAnsibleForceColor(),
		),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}

	color.Cyan("\n\tDuration: %s\n\n", exec.Duration())
}
