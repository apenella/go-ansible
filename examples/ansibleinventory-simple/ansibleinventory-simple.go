package main

import (
	"context"
	"github.com/apenella/go-ansible/pkg/inventory"
	"log"
)

func main() {
	ansibleInventoryOptions := inventory.AnsibleInventoryOptions{
		Inventory: "inventory.yml",
		List:      true,
		Yaml:      true,
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
