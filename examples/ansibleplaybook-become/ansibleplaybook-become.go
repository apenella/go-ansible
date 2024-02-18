package main

import (
	"context"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		User: "apenella",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	ansiblePlaybookPrivilegeEscalationOptions := &options.AnsiblePrivilegeEscalationOptions{
		Become:        true,
		AskBecomePass: true,
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:                  []string{"site.yml"},
		ConnectionOptions:          ansiblePlaybookConnectionOptions,
		PrivilegeEscalationOptions: ansiblePlaybookPrivilegeEscalationOptions,
		Options:                    ansiblePlaybookOptions,
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
