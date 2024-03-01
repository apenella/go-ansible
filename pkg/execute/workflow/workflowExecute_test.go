package workflow

import (
	"context"
	"errors"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/stretchr/testify/assert"
)

// TestNewWorkflowExecute tests NewWorkflowExecute function
func TestNewWorkflowExecute(t *testing.T) {
	t.Parallel()

	desc := "Testing create new WorkflowExecute using NewWorkflowExecute function"
	t.Run(desc, func(t *testing.T) {
		t.Log(desc)

		e := NewWorkflowExecute()
		assert.Equal(t, e, &WorkflowExecute{})
	})
}

func TestAppendExecutor(t *testing.T) {
	t.Parallel()

	desc := "Testing append executor to WorkflowExecute"

	t.Run(desc, func(t *testing.T) {
		t.Log(desc)

		e := NewWorkflowExecute()
		executor := &execute.MockExecute{}

		e.AppendExecutor(executor)

		assert.Equal(t, e, &WorkflowExecute{
			ExecutorList: []execute.Executor{executor}},
		)
	})

}

func TestWithContinueOnError(t *testing.T) {
	t.Parallel()

	desc := "Testing set continue on error flag to true in WorkFlowExecute"

	t.Run(desc, func(t *testing.T) {
		t.Log(desc)

		e := NewWorkflowExecute()

		e.WithContinueOnError()

		assert.Equal(t, e, &WorkflowExecute{
			ContinueOnError: true,
		})
	})
}

func TestExecute(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		desc              string
		err               error
		workflow          *WorkflowExecute
		expectedError     error
		prepareAssertFunc func(t *testing.T, e *WorkflowExecute)
		assertFunc        func(t *testing.T, e *WorkflowExecute)
	}{
		{
			desc:          "Testing execute workflow with multiple executor",
			workflow:      NewWorkflowExecute(),
			expectedError: errors.New("some error"),
			prepareAssertFunc: func(t *testing.T, e *WorkflowExecute) {
				executor1 := execute.NewMockExecute()
				executor2 := execute.NewMockExecute()
				executor3 := execute.NewMockExecute()

				executor1.On("Execute", context.TODO()).Return(nil)
				executor2.On("Execute", context.TODO()).Return(nil)
				executor3.On("Execute", context.TODO()).Return(nil)

				e.AppendExecutor(executor1)
				e.AppendExecutor(executor2)
				e.AppendExecutor(executor3)
			},
			assertFunc: func(t *testing.T, e *WorkflowExecute) {
				for _, executor := range e.ExecutorList {
					executor.(*execute.MockExecute).AssertExpectations(t)
				}
			},
		},
		{
			desc:          "Testing error when executing a workflow with a failing execution",
			workflow:      NewWorkflowExecute(),
			expectedError: errors.New("some error"),
			prepareAssertFunc: func(t *testing.T, e *WorkflowExecute) {
				executor1 := execute.NewMockExecute()
				executor2 := execute.NewMockExecute()
				// That will not be executed because of the error in executor2
				executor3 := execute.NewMockExecute()

				executor1.On("Execute", context.TODO()).Return(nil)
				executor2.On("Execute", context.TODO()).Return(errors.New("some error"))

				e.AppendExecutor(executor1)
				e.AppendExecutor(executor2)
				e.AppendExecutor(executor3)
			},
			assertFunc: func(t *testing.T, e *WorkflowExecute) {
				for _, executor := range e.ExecutorList {
					executor.(*execute.MockExecute).AssertExpectations(t)
				}
			},
		},
		{
			desc:          "Testing error when executing a workflow with a failing execution but continue on error flag is set to true",
			workflow:      NewWorkflowExecute().WithContinueOnError(),
			expectedError: errors.New("some error in executor2\nsome error in executor3"),
			prepareAssertFunc: func(t *testing.T, e *WorkflowExecute) {
				executor1 := execute.NewMockExecute()
				executor2 := execute.NewMockExecute()
				executor3 := execute.NewMockExecute()

				executor1.On("Execute", context.TODO()).Return(nil)
				executor2.On("Execute", context.TODO()).Return(errors.New("some error in executor2"))
				executor3.On("Execute", context.TODO()).Return(errors.New("some error in executor3"))

				e.AppendExecutor(executor1)
				e.AppendExecutor(executor2)
				e.AppendExecutor(executor3)
			},
			assertFunc: func(t *testing.T, e *WorkflowExecute) {
				for _, executor := range e.ExecutorList {
					executor.(*execute.MockExecute).AssertExpectations(t)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Log(test.desc)

			if test.prepareAssertFunc != nil {
				test.prepareAssertFunc(t, test.workflow)
			}

			err := test.workflow.Execute(context.TODO())

			if err != nil {
				assert.Equal(t, test.expectedError, err)

				if test.assertFunc != nil {
					test.assertFunc(t, test.workflow)
				}

			} else {
				if test.assertFunc != nil {
					test.assertFunc(t, test.workflow)
				} else {
					// If no assert function is provided, we assume that the test is wrong
					assert.True(t, false, "No assert function provided")
				}
			}
		})
	}
}
