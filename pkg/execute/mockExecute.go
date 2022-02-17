package execute

import (
	"context"

	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	"github.com/stretchr/testify/mock"
)

// MockExecute defines a simple executor for testing purposal
type MockExecute struct {
	mock.Mock
	//	Write io.Writer
}

// NewMockExecute returns a new instance of MockExecute
func NewMockExecute() *MockExecute {
	return &MockExecute{}
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *MockExecute) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error {
	args := e.Called(ctx, command, resultsFunc, options)
	return args.Error(0)

	// if e.Write == nil {
	// 	e.Write = os.Stdout
	// }
	// if command[0] == "error" {
	// 	return errors.New("(MockExecute::Execute) error")
	// }

	// fmt.Fprintf(e.Write, "%v", command)
	// return nil
}
