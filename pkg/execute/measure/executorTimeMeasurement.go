package measure

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/apenella/go-ansible/pkg/execute"
	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	errors "github.com/apenella/go-common-utils/error"
)

// ExecuteOptionFunc is a function to set executor options
type ExecuteOptionFunc func(*ExecutorTimeMeasurement)

// ExecutorTimeMeasurement is a middleware that measure the execution time of a command
type ExecutorTimeMeasurement struct {
	// executor is the next executor to be executed and measured
	executor execute.Executor
	// write is the writer to be used to print the duration message
	write io.Writer
	// showDuration is a flag to show the duration once the command is executed
	showDuration bool
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
func (e *ExecutorTimeMeasurement) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...execute.ExecuteOptions) error {

	if e.executor == nil {
		return errors.New("(ExecutorTimeMeasurement::Execute)", "Executor must be provided on ExecutorTimeMeasurement")
	}

	if e.write == nil {
		e.write = os.Stdout
	}

	timeInit := time.Now()
	defer func() {
		e.duration = time.Since(timeInit)
		if e.showDuration {
			fmt.Fprintln(e.write, fmt.Sprintf("Duration: %s", e.duration))
		}
	}()

	err := e.executor.Execute(ctx, command, resultsFunc, options...)
	if err != nil {
		return errors.New("(ExecutorTimeMeasurement::Execute)",
			fmt.Sprintf("%s",
				err.Error(),
			))
	}

	return nil
}

// Duration returns the duration of the command
func (e *ExecutorTimeMeasurement) Duration() time.Duration {
	return e.duration
}

// WithWrite set the writer to be used by DefaultExecutor
func WithWrite(w io.Writer) ExecuteOptionFunc {
	return func(e *ExecutorTimeMeasurement) {
		e.write = w
	}
}

// WithShowDuration enables to show command duration
func WithShowDuration() ExecuteOptionFunc {
	return func(e *ExecutorTimeMeasurement) {
		e.showDuration = true
	}
}
