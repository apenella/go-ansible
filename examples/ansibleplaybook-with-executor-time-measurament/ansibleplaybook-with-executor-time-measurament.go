package main

import (
	"context"
	"os"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
)

func main() {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	exec := execute.NewDefaultExecute(
		execute.WithEnvVar("ANSIBLE_FORCE_COLOR", "true"),
		execute.WithTransformers(
			results.Prepend("Go-ansible example"),
			results.LogFormat(results.DefaultLogFormatLayout, results.Now),
		),
		execute.WithShowDuration(),
	)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              execute.NewExecutorTimeMeasurement(os.Stdout, exec),
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
