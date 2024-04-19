package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	signalChan := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		User:       "aleix",
		Inventory:  "127.0.0.1,",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			execute.WithTransformers(
				transformer.Prepend("[ansibleplaybook-signals-and-cancellation]"),
			),
		),
		configuration.WithAnsibleForceColor(),
	)

	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			cancel()
		case <-ctx.Done():
		}
	}()

	err := exec.Execute(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
