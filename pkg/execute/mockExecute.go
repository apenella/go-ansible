package execute

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/apenella/go-ansible/pkg/stdoutcallback"
)

// MockExecute defines a simple executor for testing purposal
type MockExecute struct {
	Write io.Writer
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *MockExecute) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error {
	if e.Write == nil {
		e.Write = os.Stdout
	}
	if command[0] == "error" {
		return errors.New("(MockExecute::Execute) error")
	}

	fmt.Fprintf(e.Write, "%v", command)
	return nil
}
