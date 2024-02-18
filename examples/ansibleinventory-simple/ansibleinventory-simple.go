package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/inventory"
)

func main() {
	ansibleInventoryOptions := inventory.AnsibleInventoryOptions{
		Inventory: "inventory.yml",
		List:      true,
		Yaml:      true,
	}

	inventoryCmd := &inventory.AnsibleInventoryCmd{
		Pattern: "all",
		Options: &ansibleInventoryOptions,
	}

	fmt.Println("Test strings:", inventoryCmd.String())

	exec := execute.NewDefaultExecute(
		execute.WithCmd(inventoryCmd),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
