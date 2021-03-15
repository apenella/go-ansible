package main

/*
 `go run cobra-cmd-ansibleplaybook.go -i 127.0.0.1, -p site.yml -L -e example=cobra-cmd-ansibleplaybook`
*/

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	errors "github.com/apenella/go-common-utils/error"
	"github.com/spf13/cobra"
)

var inventory string
var playbookFiles []string
var connectionLocal bool
var extravars []string

const (
	extraVarsSplitToken = "="
)

func init() {
	rootCmd.Flags().StringVarP(&inventory, "inventory", "i", "", "Specify ansible playbook inventory")
	rootCmd.Flags().StringSliceVarP(&playbookFiles, "playbook", "p", []string{}, "Playbook(s) to run")
	rootCmd.Flags().BoolVarP(&connectionLocal, "connection-local", "L", false, "Run playbook using local connection")
	rootCmd.Flags().StringSliceVarP(&extravars, "extra-var", "e", []string{}, "Set extra variables to use during the playbook execution. The format of each variable must be <key>=<value>")
}

var rootCmd = &cobra.Command{
	Use:   "cobra-cmd-ansibleplaybook",
	Short: "cobra-cmd-ansibleplaybook",
	Long: `cobra-cmd-ansibleplaybook is an example which show how to use go-ansible library from cobra cli
	
 Run the example:
go run cobra-cmd-ansibleplaybook.go -L -i 127.0.0.1, -p site.yml -e example="hello go-ansible!"
`,
	RunE: commandHandler,
}

func commandHandler(cmd *cobra.Command, args []string) error {

	if len(playbookFiles) < 1 {
		return errors.New("(commandHandler)", "To run ansible-playbook playbook file path must be specified")
	}

	if len(inventory) < 1 {
		return errors.New("(commandHandler)", "To run ansible-playbook an inventory must be specified")
	}

	vars, err := varListToMap(extravars)
	if err != nil {
		return errors.New("(commandHandler)", "Error parsing extra variables", err)
	}

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{}
	if connectionLocal {
		ansiblePlaybookConnectionOptions.Connection = "local"
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: inventory,
	}

	for keyVar, valueVar := range vars {
		ansiblePlaybookOptions.AddExtraVar(keyVar, valueVar)
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         playbookFiles,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec: execute.NewDefaultExecute(
			execute.WithTransformers(
				results.Prepend("cobra-cmd-ansibleplaybook example"),
			),
		),
	}

	options.AnsibleForceColor()

	err = playbook.Run(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func varListToMap(varsList []string) (map[string]interface{}, error) {

	vars := map[string]interface{}{}

	for _, v := range varsList {
		tokens := strings.Split(v, extraVarsSplitToken)

		if len(tokens) != 2 {
			return nil, errors.New("(varListToMap)", fmt.Sprintf("Invalid extra variable format on '%s'", v))
		}
		vars[tokens[0]] = tokens[1]
	}

	return vars, nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
