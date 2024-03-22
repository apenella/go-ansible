package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/adhoc"
	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
)

func main() {

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Args:       "ping 127.0.0.1 -c 2",
		Connection: "local",
		Inventory:  " 127.0.0.1,",
		ModuleName: "command",
	}

	adhocCmd := adhoc.NewAnsibleAdhocCmd(
		adhoc.WithPattern("all"),
		adhoc.WithAdhocOptions(ansibleAdhocOptions),
	)

	fmt.Println("Command: ", adhocCmd.String())

	onelineExecute := stdoutcallback.NewOnelineStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(adhocCmd),
		),
	)

	err := onelineExecute.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
