package exec

import (
	"context"
	"os/exec"
)

// OsExec struct wrapps the OS Exec package
type OsExec struct{}

func NewOsExec() *OsExec {
	return &OsExec{}
}

// Command is a wrapper of exec.Command
func (e *OsExec) Command(name string, arg ...string) Cmder {
	return exec.Command(name, arg...)
}

// CommandContext is a wrapper of exec.CommandContext
func (e *OsExec) CommandContext(ctx context.Context, name string, arg ...string) Cmder {
	return exec.CommandContext(ctx, name, arg...)
}
