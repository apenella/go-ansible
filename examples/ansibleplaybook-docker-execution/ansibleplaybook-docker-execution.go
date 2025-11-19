package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/docker/docker/client"
)

func main() {

	// ctx := context.Background()
	apiClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	executable := NewDockerExec(apiClient)

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	executor := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithExecutable(executable),
		execute.WithTransformers(
			transformer.Prepend("ansibleplaybook-docker-executor example"),
		),
	)

	err = executor.Execute(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
