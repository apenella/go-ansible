package configuration

import "github.com/apenella/go-ansible/pkg/execute"

type ExecutorEnvVarSetter interface {
	execute.Executor
	AddEnvVar(key, value string)
}
