package stdoutcallback

import (
	"context"
	"errors"
	"testing"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNullStdoutCallbackExecute(t *testing.T) {
	t.Parallel()
	t.Run("Testing Null stdout callback execution", func(t *testing.T) {
		exec := execute.NewMockExecute()

		exec.On("WithOutput", mock.Anything).Return(exec)
		exec.On("AddEnvVar", configuration.AnsibleStdoutCallback, NullStdoutCallback)
		exec.On("Execute", mock.Anything).Return(nil)

		e := NewNullStdoutCallbackExecute(nil).WithExecutor(exec)
		err := e.Execute(context.TODO())

		assert.Nil(t, err)
		exec.AssertExpectations(t)
	})

	t.Run("Testing error on Null stdout callback when execute function returns an error", func(t *testing.T) {
		exec := execute.NewMockExecute()

		exec.On("WithOutput", mock.Anything).Return(exec)
		exec.On("AddEnvVar", configuration.AnsibleStdoutCallback, NullStdoutCallback)
		exec.On("Execute", mock.Anything).Return(errors.New("some error"))

		e := NewNullStdoutCallbackExecute(exec)
		err := e.Execute(context.TODO())

		assert.ErrorContains(t, err, "some error")
	})

	t.Run("Testing error on Null stdout callback when executor is not provided", func(t *testing.T) {
		e := NewNullStdoutCallbackExecute(nil)
		err := e.Execute(context.TODO())

		assert.ErrorContains(t, err, "NullStdoutCallbackExecute executor requires an executor")
	})
}
