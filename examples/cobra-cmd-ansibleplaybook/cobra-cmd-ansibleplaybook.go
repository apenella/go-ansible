package main

/*
 `go run cobra-cmd-ansibleplaybook.go -i 127.0.0.1, -p site.yml -L -e example=cobra-cmd-ansibleplaybook`
*/

import (
	"errors"
	"fmt"
	"os"
	"strings"

	ansibler "github.com/apenella/go-ansible"
	"github.com/spf13/cobra"
)

var inventory string
var playbook string
var connectionLocal bool
var extravars []string

const (
	extraVarsSplitToken = "="
)

func init() {
	rootCmd.Flags().StringVarP(&inventory, "inventory", "i", "", "Specify ansible playbook inventory")
	rootCmd.Flags().StringVarP(&playbook, "playbook", "p", "", "Main playbook to run")
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

	if len(playbook) < 1 {
		return errors.New("To run ansible-playbook playbook file path must be specified")
	}

	if len(inventory) < 1 {
		return errors.New("To run ansible-playbook an inventory must be specified")
	}

	vars, err := varListToMap(extravars)
	if err != nil {
		return errors.New("Error parsing extra variables. " + err.Error())
	}

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{}
	if connectionLocal {
		ansiblePlaybookConnectionOptions.Connection = "local"
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: inventory,
	}

	for keyVar, valueVar := range vars {
		ansiblePlaybookOptions.AddExtraVar(keyVar, valueVar)
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          playbook,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		ExecPrefix:        "Example cobra-cmd-ansibleplaybook",
	}

	ansibler.AnsibleForceColor()

	err = playbook.Run()
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
			return nil, errors.New("Invalid extra variable format on '" + v + "'")
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
