package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/configuration"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
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

	exec := configuration.NewExecutorWithAnsibleConfigurationSettings(
		execute.NewDefaultExecute(
			execute.WithCmd(playbook),
			execute.WithTransformers(
				outputColored(),
				transformer.Prepend("Go-ansible example"),
				transformer.LogFormat(transformer.DefaultLogFormatLayout, transformer.Now),
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

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}

}
