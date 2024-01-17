package inventory

import "github.com/apenella/go-ansible/pkg/vault"

type Vaulter interface {
	Vault(value string) (*vault.VaultVariableValue, error)
}
