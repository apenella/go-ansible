package main

import (
	"context"
	"github.com/apenella/go-ansible/pkg/inventory"
	"log"
)

func main() {
	ansibleInventoryOptions := inventory.AnsibleInventoryOptions{
		Graph:             true,
		Inventory:         "inventory.yml",
		Vars:              true,
		Yaml:              true,
		VaultPasswordFile: "vault_password.cfg",
	}

	inventoryCmd := inventory.AnsibleInventoryCmd{
		Pattern: "all",
		Options: &ansibleInventoryOptions,
	}

	results, _ := inventoryCmd.Command()
	log.Println("Test strings", inventoryCmd.String())
	for _, result := range results {
		log.Println("Command Data: ", result)
	}

	err := inventoryCmd.Run(context.TODO())
	if err != nil {
		panic(err)
	}

}
