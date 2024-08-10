package exec

import (
	"context"
	"os/exec"
)

// Exec struct wrapps the OS Exec package
type Exec struct{}

func NewExec() *Exec {
	return &Exec{}
}

// Command is a wrapper of exec.Command
func (e *Exec) Command(name string, arg ...string) Cmder {
	return exec.Command(name, arg...)
}

// CommandContext is a wrapper of exec.CommandContext
func (e *Exec) CommandContext(ctx context.Context, name string, arg ...string) Cmder {
	return exec.CommandContext(ctx, name, arg...)
}
