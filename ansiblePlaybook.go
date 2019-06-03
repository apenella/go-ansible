package ansible

import (
	"errors"
	"os"

	"oc-images-utils/helpers"
)

const (
	AnsiblePlaybookBin    = "ansible-playbook"
	ConnectionFlag        = "--connection"
	ExtraVarsFlag         = "--extra-vars"
	FlushCacheFlag        = "--flush-cache"
	InventoryFlag         = "--inventory"
	LimitFlag             = "--limit"
	ListHostsFlag         = "--list-hosts"
	ListTagsFlag          = "--list-tags"
	ListTasksFlag         = "--list-tasks"
	TagsFlag              = "--tags"
	SyntaxCheckFlag       = "--syntax-check"
	VaultPasswordFileFlag = "--vault-password-file"

	AnsibleForceColorEnv = "ANSIBLE_FORCE_COLOR"
)

// Executor interface of executors
type Executor interface {
	Execute(command string, args []string, prefix string) error
}

type AnsiblePlaybookCmd struct {
	Exec              Executor
	ExecPrefix        string
	Playbook          string
	Options           *AnsiblePlaybookOptions
	ConnectionOptions *AnsiblePlaybookConnectionOptions
}

type AnsiblePlaybookOptions struct {
	ExtraVars  map[string]interface{}
	FlushCache bool
	Inventory  string
	Limit      string
	ListHosts  bool
	ListTags   bool
	ListTasks  bool
	Tags       string
}

type AnsiblePlaybookConnectionOptions struct {
	Connection string
}

func AnsibleForceColor() {
	os.Setenv(AnsibleForceColorEnv, "true")
}

// Run method runs the ansible-playbook
func (p *AnsiblePlaybookCmd) Run() error {
	if p == nil {
		return errors.New("(ansible:Run) AnsiblePlaybookCmd is nil")
	}

	if p.Exec == nil {
		p.Exec = &DefaultExecute{
			Write: os.Stdout,
		}
	}

	cmd, err := p.Command()
	if err != nil {
		return errors.New("(ansible:Run) -> " + err.Error())
	}

	if len(p.ExecPrefix) <= 0 {
		p.ExecPrefix = ""
	}

	return p.Exec.Execute(cmd[0], cmd[1:], p.ExecPrefix)
}

func (p *AnsiblePlaybookCmd) Command() ([]string, error) {
	cmd := []string{}
	cmd = append(cmd, AnsiblePlaybookBin)

	if p.Options != nil {
		options, err := p.Options.GenerateCommandOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command) -> " + err.Error())
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	if p.ConnectionOptions != nil {
		options, err := p.ConnectionOptions.GenerateCommandConnectionOptions()
		if err != nil {
			return nil, errors.New("(ansible::Command) -> " + err.Error())
		}
		for _, option := range options {
			cmd = append(cmd, option)
		}
	}

	cmd = append(cmd, p.Playbook)

	return cmd, nil
}

func (o *AnsiblePlaybookOptions) GenerateCommandOptions() ([]string, error) {
	cmd := []string{}

	if o == nil {
		return nil, errors.New("(ansible::GenerateCommandOptions) AnsiblePlaybookOptions is nil")
	}

	if o.FlushCache {
		cmd = append(cmd, FlushCacheFlag)
	}

	if o.Inventory != "" {
		cmd = append(cmd, InventoryFlag)
		cmd = append(cmd, o.Inventory)
	}

	if o.Limit != "" {
		cmd = append(cmd, LimitFlag)
		cmd = append(cmd, o.Limit)
	}

	if o.ListHosts {
		cmd = append(cmd, ListHostsFlag)
	}

	if o.ListTags {
		cmd = append(cmd, ListTagsFlag)
	}

	if o.ListTasks {
		cmd = append(cmd, ListTasksFlag)
	}

	if o.Tags != "" {
		cmd = append(cmd, TagsFlag)
		cmd = append(cmd, o.Tags)
	}

	if len(o.ExtraVars) > 0 {
		cmd = append(cmd, ExtraVarsFlag)
		extraVars, err := o.generateExtraVarsCommand()
		if err != nil {
			return nil, errors.New("(ansible::GenerateCommandOptions) -> " + err.Error())
		}
		cmd = append(cmd, extraVars)
	}

	return cmd, nil
}

func (o *AnsiblePlaybookOptions) generateExtraVarsCommand() (string, error) {

	extraVars, err := helpers.ObjectToJSONString(o.ExtraVars)
	if err != nil {
		return "", errors.New("(ansible::generateExtraVarsCommand) -> " + err.Error())
	}
	return extraVars, nil
}

func (o *AnsiblePlaybookOptions) AddExtraVar(name string, value interface{}) error {

	if o.ExtraVars == nil {
		o.ExtraVars = map[string]interface{}{}
	}
	_, exists := o.ExtraVars[name]
	if exists {
		return errors.New("(ansible::AddExtraVar) ExtraVar '" + name + "' already exist.")
	}

	o.ExtraVars[name] = value

	return nil
}

func (o *AnsiblePlaybookConnectionOptions) GenerateCommandConnectionOptions() ([]string, error) {
	cmd := []string{}

	if o.Connection != "" {
		cmd = append(cmd, ConnectionFlag)
		cmd = append(cmd, o.Connection)
	}

	return cmd, nil
}
