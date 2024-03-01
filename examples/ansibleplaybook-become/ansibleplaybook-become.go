package main

import (
	"context"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		User:          "apenella",
		Inventory:     "127.0.0.1,",
		Become:        true,
		AskBecomePass: true,
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"site.yml"},
		PlaybookOptions: ansiblePlaybookOptions,
	}

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbook),
		execute.WithTransformers(
			transformer.Prepend("Go-ansible example with become"),
		),
		execute.WithEnvVars(
			map[string]string{
				"ANSIBLE_FORCE_COLOR": "true",
			},
		),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
