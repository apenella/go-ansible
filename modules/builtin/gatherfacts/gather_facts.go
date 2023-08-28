package gatherfacts

import (
	"bytes"
	"text/template"

	"github.com/apenella/go-ansible/modules/enumtipe"
	"github.com/apenella/go-ansible/modules/helper"
)

var Tmplt_ansible_gather_facts = `
    - name: {{ .Name }} 
      ansible.builtin.gather_facts:
        {{- if .Parallel }}
        parallel: {{ .Parallel }}
        {{- else }}
        parallel: false
        {{- end }}
`

type AnsibleBuiltinGatherFacts struct {
	Name     string              `json:"name"`
	Parallel enumtipe.CostomBool `json:"parallel"`
}

func (a *AnsibleBuiltinGatherFacts) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinGatherFacts) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	tmpl := template.Must(template.New("ansible_builtin_gather_facts").Parse(Tmplt_ansible_gather_facts))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
