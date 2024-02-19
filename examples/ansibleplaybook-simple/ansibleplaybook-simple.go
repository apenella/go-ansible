package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
)

func main() {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	err := playbook.NewAnsiblePlaybookExecute("site.yml", "site2.yml").
		WithOptions(ansiblePlaybookOptions).
		WithConnectionOptions(ansiblePlaybookConnectionOptions).
		Execute(context.TODO())

	if err != nil {
		fmt.Print()
		os.Exit(1)
	}
}
