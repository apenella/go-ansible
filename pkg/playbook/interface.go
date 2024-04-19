package playbook

import "github.com/apenella/go-ansible/v2/pkg/vault"

type Vaulter interface {
	Vault(value string) (*vault.VaultVariableValue, error)
}

type ExitCodeErrorer interface {
	ExitCode() int
}
