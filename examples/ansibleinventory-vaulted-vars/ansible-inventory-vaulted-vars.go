package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/inventory"
)

func main() {
	ansibleInventoryOptions := inventory.AnsibleInventoryOptions{
		Graph:             true,
		Inventory:         "inventory.yml",
		Vars:              true,
		Yaml:              true,
		VaultPasswordFile: "vault_password.cfg",
	}

	inventoryCmd := inventory.NewAnsibleInventoryCmd(
		inventory.WithPattern("all"),
		inventory.WithInventoryOptions(&ansibleInventoryOptions),
	)

	fmt.Println("Test strings:", inventoryCmd.String())

	exec := execute.NewDefaultExecute(
		execute.WithCmd(inventoryCmd),
	)

	err := exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
