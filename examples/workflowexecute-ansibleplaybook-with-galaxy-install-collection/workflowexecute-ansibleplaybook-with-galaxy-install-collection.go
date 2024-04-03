package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/workflow"
	galaxy "github.com/apenella/go-ansible/v2/pkg/galaxy/collection/install"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory:     "inventory.yml",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
	}

	galaxyInstallCollectionCmd := galaxy.NewAnsibleGalaxyCollectionInstallCmd(
		galaxy.WithGalaxyCollectionInstallOptions(&galaxy.AnsibleGalaxyCollectionInstallOptions{
			Force:            true,
			Upgrade:          true,
			RequirementsFile: "requirements.yml",
		}),
	)

	galaxyInstallCollectionExec := execute.NewDefaultExecute(
		execute.WithCmd(galaxyInstallCollectionCmd),
	)

	playbookCmd := playbook.NewAnsiblePlaybookExecute("site.yml").
		WithPlaybookOptions(ansiblePlaybookOptions)

	err := workflow.NewWorkflowExecute(galaxyInstallCollectionExec, playbookCmd).
		WithTrace().
		Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
