package execute

import (
	"context"

	"github.com/apenella/go-ansible/internal/executable/os/exec"
)

// Executor interface to execute commands
type Executor interface {
	Execute(ctx context.Context) error
}

// Executabler is an interface to run commands
type Executabler interface {
	Command(name string, arg ...string) exec.Cmder
	CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}

// Commander generates commands to be executed
type Commander interface {
	Command() ([]string, error)
}
