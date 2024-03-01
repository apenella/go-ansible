package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/inventory"
)

func main() {
	ansibleInventoryOptions := inventory.AnsibleInventoryOptions{
		Inventory: "inventory.yml",
		List:      true,
		Yaml:      true,
	}

	err := inventory.NewAnsibleInventoryExecute().
		WithInventoryOptions(&ansibleInventoryOptions).
		WithPattern("all").
		Execute(context.TODO())

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
