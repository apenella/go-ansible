package execute

import (
	"context"

	"github.com/apenella/go-ansible/pkg/execute/executable/os/exec"
)

// Executor interface is satisfied by those types which has a Execute(context.Context,[]string,stdoutcallback.StdoutCallbackResultsFunc,...ExecuteOptions)error method
type Executor interface {
	Execute(ctx context.Context, command []string, options ...ExecuteOptions) error
}

// Executabler is an interface to run commands
type Executabler interface {
	Command(name string, arg ...string) exec.Cmder
	CommandContext(ctx context.Context, name string, arg ...string) exec.Cmder
}
