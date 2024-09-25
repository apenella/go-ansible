package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Become:       true,
		BecomeMethod: "sudo",
		BecomeUser:   "root",
		ExtraVars: map[string]interface{}{
			"ansible_ssh_private_key_file": "/ssh/id_rsa",
			"ansible_sudo_pass":            "12345",
		},
		Inventory:     "inventory.yml",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
		Timeout:       300,
		User:          "aleix",
	}

	cmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(cmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		),
		configuration.WithAnsibleForceColor(),
		configuration.WithAnsibleStdoutCallback(stdoutcallback.YAMLStdoutCallback),
	)

	fmt.Printf("Executing command: '%s'", cmd.String())

	err := exec.Execute(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
