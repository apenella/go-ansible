package execute

import "context"

// Executor interface is satisfied by those types which has a Execute(context.Context,string,[]string)error method
type Executor interface {
	Execute(ctx context.Context, command []string, options ...ExecuteOptions) error
}

// ExecuteOptions
type ExecuteOptions func(Executor)
