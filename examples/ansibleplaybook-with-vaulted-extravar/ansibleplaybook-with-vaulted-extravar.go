package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/execute/configuration"
	"github.com/apenella/go-ansible/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/pkg/playbook"
	"github.com/apenella/go-ansible/pkg/vault"
	"github.com/apenella/go-ansible/pkg/vault/encrypt"
	"github.com/apenella/go-ansible/pkg/vault/password/file"
	"github.com/apenella/go-ansible/pkg/vault/password/resolve"
	"github.com/spf13/afero"
)

func main() {
	var err error

	vaulter := vault.NewVariableVaulter(
		vault.WithEncrypt(
			encrypt.NewEncryptString(
				encrypt.WithReader(
					resolve.NewReadPasswordResolve(
						resolve.WithReader(
							file.NewReadPasswordFromFile(
								file.WithFs(afero.NewOsFs()),
								file.WithFile("./vault_password.cfg"),
							),
							//
							// Uncomment this lines to use other password readers
							//
							// text.NewReadPasswordFromText(
							// 	text.WithText("s3cr3t"),
							// ),
							// envvars.NewReadPasswordFromEnvVar(
							// 	envvars.WithEnvVar("ANSIBLE_VAULT_PASSWORD"),
							// ),
						),
					),
				),
			),
		),
	)

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	err = ansiblePlaybookOptions.AddVaultedExtraVar(vaulter, "vaulted_extra_var", "That is a secret")
	if err != nil {
		panic(err)
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:       []string{"site.yml"},
		PlaybookOptions: ansiblePlaybookOptions,
	}

	fmt.Printf("\n  Ansible playbook command:\n%s\n\n", playbook.String())

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbook),
			execute.WithTransformers(
				transformer.Prepend("Go-ansible example with become"),
			),
		),
	).WithAnsibleVaultPasswordFile("./vault_password.cfg").WithAnsibleForceColor()

	err = exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
