package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/adhoc"
)

func main() {

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Args: `msg="
		{{ arg1 }}
		{{ arg2 }}
		{{ arg3 }}
		"`,
		Connection: "local",
		ExtraVars: map[string]interface{}{
			"arg1": map[string]interface{}{"subargument": "subargument_value"},
			"arg2": "arg2_value",
			"arg3": "arg3_value",
		},
		Inventory:  "127.0.0.1,",
		ModuleName: "debug",
	}

	err := adhoc.NewAnsibleAdhocExecute("all").
		WithAdhocOptions(ansibleAdhocOptions).
		Execute(context.TODO())

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
