package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/moby/moby/client"
)

func main() {

	// ctx := context.Background()
	apiClient, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	_ = NewDockerCmd(apiClient)

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	err = playbook.NewAnsiblePlaybookExecute("site.yml").
		WithPlaybookOptions(ansiblePlaybookOptions).
		Execute(context.TODO())

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
