package vault

import (
	common "github.com/apenella/go-common-utils/data"
	"github.com/pkg/errors"
)

type VaultVariableValue struct {
	Value interface{} `json:"__ansible_vault"`
}

func NewVaultVariableValue(value interface{}) *VaultVariableValue {
	return &VaultVariableValue{
		Value: value,
	}
}

func (v *VaultVariableValue) ToJSON() (string, error) {

	// jsonValue, err := json.Marshal(v)
	jsonValue, err := common.ObjectToJSONString(v)
	if err != nil {
		return "", errors.Wrap(err, "Error converting the vault variable value to JSON.")
	}

	return string(jsonValue), nil
}
