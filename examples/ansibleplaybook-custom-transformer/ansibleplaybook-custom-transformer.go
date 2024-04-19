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
	"github.com/fatih/color"
)

// customTrasnformer
func outputColored() transformer.TransformerFunc {
	return func(message string) string {
		yellow := color.New(color.FgYellow).SprintFunc()
		return fmt.Sprintf("%s", yellow(message))
	}
}

func main() {

	signalChan := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
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
				outputColored(),
				transformer.Prepend("Go-ansible example"),
				transformer.LogFormat(transformer.DefaultLogFormatLayout, transformer.Now),
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

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}

}
