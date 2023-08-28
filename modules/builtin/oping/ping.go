package oping

import (
	"bytes"
	"text/template"

	"github.com/apenella/go-ansible/modules/helper"
)

var Tmplt_ansible_ping = `
    - name: {{ .Name }}  
      ansible.builtin.ping:
`

type AnsibleBuiltinPing struct {
	Name string `json:"name"`
}

func (a *AnsibleBuiltinPing) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinPing) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	tmpl := template.Must(template.New("ansible_builtin_ping").Parse(Tmplt_ansible_ping))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
