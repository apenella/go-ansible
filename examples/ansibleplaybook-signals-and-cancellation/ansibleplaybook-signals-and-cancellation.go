package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/configuration"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	signalChan := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		User:       "aleix",
		Inventory:  "127.0.0.1,",
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"site.yml"},
		PlaybookOptions: ansiblePlaybookOptions,
	}

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbook),
			execute.WithTransformers(
				transformer.Prepend("[ansibleplaybook-signals-and-cancellation]"),
			),
		),
	).WithAnsibleForceColor()

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
