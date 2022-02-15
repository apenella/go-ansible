package execute

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	errors "github.com/apenella/go-common-utils/error"
)

// ExecutorTimeMeasurement is a middleware that measure the execution time of a command
type ExecutorTimeMeasurement struct {
	executor Executor
	writer   io.Writer
}

// NewExecutorTimeMeasurement returns a new ExecutorTimeMeasurement
func NewExecutorTimeMeasurement(w io.Writer, executor Executor) *ExecutorTimeMeasurement {
	if w == nil {
		w = os.Stdout
	}

	return &ExecutorTimeMeasurement{
		executor: executor,
		writer:   w,
	}
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *ExecutorTimeMeasurement) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error {

	timeInit := time.Now()
	err := e.executor.Execute(ctx, command, resultsFunc, options...)
	if err != nil {
		return errors.New("(ExecutorTimeMeasurement::Execute)",
			fmt.Sprintf("%s\n%s",
				err.Error(),
				durationMessage(time.Since(timeInit)),
			))
	}

	fmt.Fprintln(e.writer, durationMessage(time.Since(timeInit)))

	return nil
}

func durationMessage(d time.Duration) string {
	return fmt.Sprintf("Duration: %s", d)
}
