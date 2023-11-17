package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/pkg/options"
)

func main() {

	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  " 127.0.0.1,",
		ModuleName: "command",
		Args:       "ping 127.0.0.1 -c 2",
	}

	adhoc := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
	}

	fmt.Println("Command: ", adhoc.String())

	onelineExecute := stdoutcallback.NewOnelineStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(adhoc),
		),
	)

	err := onelineExecute.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
