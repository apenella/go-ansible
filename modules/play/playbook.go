package play

import (
	"bytes"
	"context"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/apenella/go-ansible/modules/enumtipe"
	"github.com/apenella/go-ansible/modules/helper"
	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/measure"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
)

type ITaskMaker interface {
	MakeAnsibleTask() (string, error)
}

var Tmplt_ansible_playbook = `
---
{{ if .Name }}
- name: {{ .Name }}
  hosts: all
{{ else }}
- hosts: all
{{- end }}
  {{- if .RemoteUser}}
  remote_user: {{ .RemoteUser }}
  {{- end }}
  gather_facts: {{ .GatherFacts }}

  {{- if .PrevTasks }}
  pre_tasks:
    {{- range .PreTasks }}
    {{ .MakeAnsibleTask }}
    {{- end }}
  {{- end }}
  tasks:
    {{- range .Tasks }}
    {{- .MakeAnsibleTask }}
    {{- end }}
  {{- if .PostTasks }}
  post_tasks:
    {{- range .PostTasks }}
    {{- .MakeAnsibleTask }}
    {{- end }}
  {{- end }}
`

type Playbook struct {
	Name       string   `json:"name"`
	Hosts      []string `json:"hosts"`
	RemoteUser string   `json:"remote_user"`
	Password   string   `json:"password"`
	PrivateKey string   `json:"private_key"`

	GatherFacts enumtipe.CostomBool `json:"gather_facts"`

	PrevTasks []ITaskMaker `json:"prev_tasks"`
	Tasks     []ITaskMaker `json:"tasks"`
	PostTasks []ITaskMaker `json:"post_tasks"`
}

func (a *Playbook) String() string {
	return helper.MarshalToString(a)
}

func (a *Playbook) MakeAnsibleTask() (string, error) {

	tmpl := template.Must(template.New("ansible_playbook").Parse(Tmplt_ansible_playbook))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}

func (a *Playbook) ExecPlaybook(ctx context.Context) (*results.AnsiblePlaybookJSONResults, string, int64, error) {

	var err error
	var res *results.AnsiblePlaybookJSONResults
	buff := new(bytes.Buffer)

	tempDir, err := os.MkdirTemp("", "octopus_ansibleplaybook")
	if err != nil {
		return nil, "", -50, err
	}

	connectionOptions := &options.AnsibleConnectionOptions{
		User:          a.RemoteUser,
		SSHCommonArgs: "'-o StrictHostKeyChecking=no'",
	}

	var pk_file_path string
	if a.PrivateKey != "" {
		pk_file_path = filepath.Join(tempDir, "priv_key")
		err = os.WriteFile(pk_file_path, []byte(a.PrivateKey), 0400)
		if err == nil {
			connectionOptions.PrivateKey = pk_file_path
			defer func() {
				os.Remove(pk_file_path)
			}()
		}

	}

	hosts := strings.Join(a.Hosts, ",")
	playbookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: hosts,
	}

	executorTimeMeasurement := measure.NewExecutorTimeMeasurement(
		execute.NewDefaultExecute(
			execute.WithWrite(io.Writer(buff)),
		),
	)

	// prepare playbook
	playbooks := make([]string, 0)

	pb_content, err := a.MakeAnsibleTask()
	if err != nil {
		return nil, "", -10, err
	}

	pb_file_path := filepath.Join(tempDir, "playbook.yaml")

	err = os.WriteFile(pb_file_path, []byte(pb_content), 0644)
	if err != nil {
		return nil, pb_content, -60, err
	}

	playbooks = append(playbooks, pb_file_path)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         playbooks,
		ConnectionOptions: connectionOptions,
		Options:           playbookOptions,
		Exec:              executorTimeMeasurement,
		StdoutCallback:    "json",
	}

	err = playbook.Run(ctx)
	duration := executorTimeMeasurement.Duration()
	if err != nil {
		return nil, pb_content, duration.Milliseconds(), err
	}

	res, err = results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {

		return nil, pb_content, duration.Milliseconds(), err
	}

	return res, pb_content, duration.Milliseconds(), nil
}
