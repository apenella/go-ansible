package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	"github.com/fatih/color"
)

// customTrasnformer
func outputColored() results.TransformerFunc {
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
		Exec: execute.NewDefaultExecute(
			execute.WithEnvVar("ANSIBLE_FORCE_COLOR", "true"),
			execute.WithTransformers(
				outputColored(),
				results.Prepend("Go-ansible example"),
				results.LogFormat(results.DefaultLogFormatLayout, results.Now),
			),
			execute.WithShowDuration(),
		),
	}

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

	err := playbook.Run(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
