package main

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/options"
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

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	err = ansiblePlaybookOptions.AddVaultedExtraVar(vaulter, "vaulted_extra_var", "That is a secret")
	if err != nil {
		panic(err)
	}

	executor := execute.NewDefaultExecute(
		execute.WithEnvVar("ANSIBLE_VAULT_PASSWORD_FILE", "./vault_password.cfg"),
	)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		Exec:              executor,
	}

	fmt.Printf("\n  Ansible playbook command:\n%s\n\n", playbook.String())

	err = playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
