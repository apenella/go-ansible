package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/workflow"
	galaxy "github.com/apenella/go-ansible/v2/pkg/galaxy/role/install"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory:     "inventory.yml",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
	}

	galaxyInstallRolesCmd := galaxy.NewAnsibleGalaxyRoleInstallCmd(
		// galaxy.WithRoleNames("geerlingguy.go"),
		galaxy.WithGalaxyRoleInstallOptions(&galaxy.AnsibleGalaxyRoleInstallOptions{
			Force:    true,
			RoleFile: "requirements.yml",
		}),
	)

	galaxyInstallRolesExec := execute.NewDefaultExecute(
		execute.WithCmd(galaxyInstallRolesCmd),
	)

	playbookCmd := playbook.NewAnsiblePlaybookExecute("site.yml").
		WithPlaybookOptions(ansiblePlaybookOptions)

	err := workflow.NewWorkflowExecute(galaxyInstallRolesExec, playbookCmd).
		WithTrace().
		Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
