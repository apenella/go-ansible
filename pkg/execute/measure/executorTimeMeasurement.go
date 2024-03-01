package measure

import (
	"context"
	"time"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	errors "github.com/apenella/go-common-utils/error"
)

// ExecuteOptionFunc is a function to set executor options
type ExecuteOptionFunc func(*ExecutorTimeMeasurement)

// ExecutorTimeMeasurement is a middleware that measure the execution time of a command
type ExecutorTimeMeasurement struct {
	// executor is the next executor to be executed and measured
	executor execute.Executor
	// duration is the duration of the command
	duration time.Duration
}

// NewExecutorTimeMeasurement returns a new ExecutorTimeMeasurement
func NewExecutorTimeMeasurement(executor execute.Executor, options ...ExecuteOptionFunc) *ExecutorTimeMeasurement {

	exec := &ExecutorTimeMeasurement{
		executor: executor,
	}

	for _, option := range options {
		option(exec)
	}

	return exec
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *ExecutorTimeMeasurement) Execute(ctx context.Context) error {

	if e.executor == nil {
		return errors.New("(ExecutorTimeMeasurement::Execute)", "Executor must be provided on ExecutorTimeMeasurement")
	}

	timeInit := time.Now()
	defer func() {
		e.duration = time.Since(timeInit)
	}()

	err := e.executor.Execute(ctx)
	if err != nil {
		return errors.New("(ExecutorTimeMeasurement::Execute)", err.Error())
	}

	return nil
}

// Duration returns the duration of the command
func (e *ExecutorTimeMeasurement) Duration() time.Duration {
	return e.duration
}
