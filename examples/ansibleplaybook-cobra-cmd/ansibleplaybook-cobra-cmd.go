package main

/*
 `go run cobra-cmd-ansibleplaybook.go -i 127.0.0.1, -p site.yml -L -e example=cobra-cmd-ansibleplaybook`
*/

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
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
	Use:   "ansibleplaybook-cobra-cmd",
	Short: "ansibleplaybook-cobra-cmd",
	Long: `ansibleplaybook-cobra-cmd is an example which show how to use go-ansible library from cobra cli
	
 Run the example:
go run ansibleplaybook-cobra-cmd.go -L -i 127.0.0.1, -p site.yml -e example="hello go-ansible!"
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

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: inventory,
	}

	if connectionLocal {
		ansiblePlaybookOptions.Connection = "local"
	}

	for keyVar, valueVar := range vars {
		_ = ansiblePlaybookOptions.AddExtraVar(keyVar, valueVar)
	}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbookFiles...),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			execute.WithTransformers(
				transformer.Prepend("Go-ansible example with become"),
			),
		),
		configuration.WithAnsibleForceColor(),
	)

	err = exec.Execute(context.TODO())
	if err != nil {
		panic(err)
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
