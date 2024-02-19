package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/pkg/inventory"
)

func main() {
	ansibleInventoryOptions := inventory.AnsibleInventoryOptions{
		Inventory: "inventory.yml",
		List:      true,
		Yaml:      true,
	}

	err := inventory.NewAnsibleInventoryExecute().
		WithOptions(&ansibleInventoryOptions).
		WithPattern("all").
		Execute(context.TODO())

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
