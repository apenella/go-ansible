package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	ansibler "github.com/apenella/go-ansible"
	"github.com/apenella/go-ansible/execute"
	"github.com/apenella/go-ansible/stdoutcallback/results"
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

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
		User:       "aleix",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithPrefix("Go-ansible examples"),
			execute.WithOutputFormat(execute.OutputFormatLogFormat),
			execute.WithTransformers(
				results.Prepend("before"),
				outputColored(),
				results.Prepend("after"),
			),
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
