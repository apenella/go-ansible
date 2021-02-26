package ansibler

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// MockExecute defines a simple executor for testing purposal
type MockExecute struct {
	Write io.Writer
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *MockExecute) Execute(command string, args []string, prefix string) error {
	if e.Write == nil {
		e.Write = os.Stdout
	}
	if command == "error" {
		return errors.New("(MockExecute::Execute) error")
	}

	fmt.Fprintf(e.Write, "%s %v", command, args)
	return nil
}
