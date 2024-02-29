package workflow

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/pkg/execute"
)

type WorkflowExecute struct {
	// ExecutorList is a list of executors
	ExecutorList []execute.Executor
	// ContinueOnError is a flag to continue on error
	ContinueOnError bool
}

// NewWorkflowExecute creates a new WorkflowExecute
func NewWorkflowExecute(e ...execute.Executor) *WorkflowExecute {
	return &WorkflowExecute{
		ExecutorList: e,
	}
}

// AppendExecutor appends an executor to the list
func (e *WorkflowExecute) AppendExecutor(executor execute.Executor) *WorkflowExecute {
	e.ExecutorList = append(e.ExecutorList, executor)
	return e
}

// WithContinueOnError sets the continue on error flag to true
func (e *WorkflowExecute) WithContinueOnError() *WorkflowExecute {
	e.ContinueOnError = true
	return e
}

// Execute runs the executors
func (e *WorkflowExecute) Execute(ctx context.Context) error {
	var errList []error = make([]error, 0)

	for _, executor := range e.ExecutorList {
		err := executor.Execute(ctx)
		if err != nil {
			errList = append(errList, err)

			if !e.ContinueOnError {
				// leave the loop when the continue on error flag is false which is the default behaviour
				break
			}
		}
	}

	if len(errList) > 0 {
		errs := errList[0]
		for _, err := range errList[1:] {
			errs = fmt.Errorf("%s\n%s", errs, err)
		}

		return errs
	}

	return nil
}
