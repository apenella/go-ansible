package main

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/apenella/go-ansible/v2/pkg/execute/measure"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/fatih/color"
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
