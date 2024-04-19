package execute

import (
	"context"

	"github.com/apenella/go-ansible/v2/internal/executable/os/exec"
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
	String() string
}

// ErrorEnricher interface to enrich and customize errors
type ErrorEnricher interface {
	Enrich(err error) error
}
