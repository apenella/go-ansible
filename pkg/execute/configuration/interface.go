package configuration

import "github.com/apenella/go-ansible/v2/pkg/execute"

type ExecutorEnvVarSetter interface {
	execute.Executor
	AddEnvVar(key, value string)
}
