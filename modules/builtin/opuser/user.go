package opuser

import (
	"bytes"
	"html/template"

	"github.com/apenella/go-ansible/modules/enumtipe"
	"github.com/apenella/go-ansible/modules/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/user_module.html
var Tmplt_ansible_create_user = `
    - name: Add Systemuser {{ .Username }} with uid {{ .UID }}
      ansible.builtin.user:
        name: {{ .Username }}
        {{- if .Password }}
        password: {{ .Password }}
        {{- end }}
        {{- if .Home }}
        home: {{ .Home }}
        {{- end }}
        {{- if .CreateHome }}
        create_home: true
        {{- end }}
        {{- if .MoveHome}}
        move_home: true
        {{- end }}
        {{- if .Comment }}
        comment: {{ .Comment }}
        {{- end }}
        {{- if .UID }}
        uid: {{ .UID }}
        {{- end }}
        {{- if .Shell }}
        shell: {{ .Shell }}
        {{- end }}
        {{- if .Groups }}
        groups: {{ .Groups }}
        append: yes
        {{- end }}
        {{- if .State }}
        state: {{ .State }}
        {{- end }}
        {{- if .Remove }}
        remove: {{ .Remove }}
        {{- end }}
`

type AnsibleCreateUser struct {
	Username   string              `json:"username"`
	Password   string              `json:"password"`
	Home       string              `json:"home"`
	CreateHome enumtipe.CostomBool `json:"create_home"`
	MoveHome   enumtipe.CostomBool `json:"move_home"`
	Comment    string              `json:"comment"`
	Shell      string              `json:"shell"`
	Groups     string              `json:"groups"`
	UID        uint32              `json:"uid"`
	State      string              `json:"state"`
	Remove     enumtipe.CostomBool `json:"remove"`
}

func (a *AnsibleCreateUser) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleCreateUser) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	tmpl := template.Must(template.New("ansible_user").Parse(Tmplt_ansible_create_user))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", nil
	}

	return buff.String(), nil
}
