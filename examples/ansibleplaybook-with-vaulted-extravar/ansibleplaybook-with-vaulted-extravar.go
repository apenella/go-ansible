package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
	"github.com/apenella/go-ansible/v2/pkg/vault"
	"github.com/apenella/go-ansible/v2/pkg/vault/encrypt"
	"github.com/apenella/go-ansible/v2/pkg/vault/password/file"
	"github.com/apenella/go-ansible/v2/pkg/vault/password/resolve"
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

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks("site.yml"),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	fmt.Printf("\n  Ansible playbook command:\n%s\n\n", playbookCmd.String())

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			execute.WithTransformers(
				transformer.Prepend("Go-ansible example with become"),
			),
		),
		configuration.WithAnsibleVaultPasswordFile("./vault_password.cfg"),
		configuration.WithAnsibleForceColor(),
	)

	err = exec.Execute(context.TODO())
	if err != nil {
		panic(err)
	}
}
