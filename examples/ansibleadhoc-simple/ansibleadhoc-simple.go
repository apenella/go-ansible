package main

import (
	"context"
	"fmt"
	"os"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
)

func main() {

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  "127.0.0.1,",
		ModuleName: "debug",
		Args: `msg="
		{{ arg1 }}
		{{ arg2 }}
		{{ arg3 }}
		"`,
		ExtraVars: map[string]interface{}{
			"arg1": map[string]interface{}{"subargument": "subargument_value"},
			"arg2": "arg2_value",
			"arg3": "arg3_value",
		},
	}

	err := adhoc.NewAnsibleAdhocExecute("all").
		WithAdhocOptions(ansibleAdhocOptions).
		WithConnectionOptions(ansibleConnectionOptions).
		Execute(context.TODO())

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
