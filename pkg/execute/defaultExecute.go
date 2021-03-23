package execute

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/apenella/go-ansible/pkg/stdoutcallback"
	"github.com/apenella/go-ansible/pkg/stdoutcallback/results"
	errors "github.com/apenella/go-common-utils/error"
)

const (
	// AnsiblePlaybookErrorCodeGeneralError
	AnsiblePlaybookErrorCodeGeneralError = 1
	// AnsiblePlaybookErrorCodeOneOrMoreHostFailed
	AnsiblePlaybookErrorCodeOneOrMoreHostFailed = 2
	// AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable
	AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable = 3
	// AnsiblePlaybookErrorCodeParserError
	AnsiblePlaybookErrorCodeParserError = 4
	// AnsiblePlaybookErrorCodeBadOrIncompleteOptions
	AnsiblePlaybookErrorCodeBadOrIncompleteOptions = 5
	// AnsiblePlaybookErrorCodeUserInterruptedExecution
	AnsiblePlaybookErrorCodeUserInterruptedExecution = 99
	// AnsiblePlaybookErrorCodeUnexpectedError
	AnsiblePlaybookErrorCodeUnexpectedError = 250

	// AnsiblePlaybookErrorMessageGeneralError
	AnsiblePlaybookErrorMessageGeneralError = "ansible-playbook error: general error"
	// AnsiblePlaybookErrorMessageOneOrMoreHostFailed
	AnsiblePlaybookErrorMessageOneOrMoreHostFailed = "ansible-playbook error: one or more host failed"
	// AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable
	AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable = "ansible-playbook error: one or more host unreachable"
	// AnsiblePlaybookErrorMessageParserError
	AnsiblePlaybookErrorMessageParserError = "ansible-playbook error: parser error"
	// AnsiblePlaybookErrorMessageBadOrIncompleteOptions
	AnsiblePlaybookErrorMessageBadOrIncompleteOptions = "ansible-playbook error: bad or incomplete options"
	// AnsiblePlaybookErrorMessageUserInterruptedExecution
	AnsiblePlaybookErrorMessageUserInterruptedExecution = "ansible-playbook error: user interrupted execution"
	// AnsiblePlaybookErrorMessageUnexpectedError
	AnsiblePlaybookErrorMessageUnexpectedError = "ansible-playbook error: unexpected error"
)

// DefaultExecute is a simple definition of an executor
type DefaultExecute struct {
	// Writer is where is written the command stdout
	Write io.Writer
	// WriterError is where is written the command stderr
	WriterError io.Writer
	// ShowDuration enables to show the execution duration time after the command finishes
	ShowDuration bool
	// CmdRunDir specifies the working directory of the command.
	CmdRunDir string
	// OutputFormat
	Transformers []results.TransformerFunc
}

// NewDefaultExecute return a new DefaultExecute instance with all options
func NewDefaultExecute(options ...ExecuteOptions) *DefaultExecute {
	execute := &DefaultExecute{}

	for _, opt := range options {
		opt(execute)
	}

	return execute
}

// WithWrite set the writer to be used by DefaultExecutor
func WithWrite(w io.Writer) ExecuteOptions {
	return func(e Executor) {
		e.(*DefaultExecute).Write = w
	}
}

// WithWriteError set the error writer to be used by DefaultExecutor
func WithWriteError(w io.Writer) ExecuteOptions {
	return func(e Executor) {
		e.(*DefaultExecute).WriterError = w
	}
}

// WithCmdRunDir set the command run directory to be used by DefaultExecutor
func WithCmdRunDir(cmdRunDir string) ExecuteOptions {
	return func(e Executor) {
		e.(*DefaultExecute).CmdRunDir = cmdRunDir
	}
}

// WithShowDuration enables to show command duration
func WithShowDuration() ExecuteOptions {
	return func(e Executor) {
		e.(*DefaultExecute).ShowDuration = true
	}
}

// WithTransformers add trasformes
func WithTransformers(trans ...results.TransformerFunc) ExecuteOptions {
	return func(e Executor) {
		e.(*DefaultExecute).Transformers = trans
	}
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DefaultExecute) Execute(ctx context.Context, command []string, resultsFunc stdoutcallback.StdoutCallbackResultsFunc, options ...ExecuteOptions) error {

	var (
		err error
		cmdStderr, cmdStdout io.ReadCloser
		wg sync.WaitGroup
	)

	execErrChan := make(chan error)

	// apply all options to the executor
	for _, opt := range options {
		opt(e)
	}

	if resultsFunc == nil {
		resultsFunc = results.DefaultStdoutCallbackResults
	}

	// default stdout and stderr for the main process
	if e.Write == nil {
		e.Write = os.Stdout
	}

	if e.WriterError == nil {
		e.WriterError = os.Stderr
	}

	cmd := exec.CommandContext(ctx, command[0], command[1:]...)

	if len(e.CmdRunDir) > 0 {
		cmd.Dir = e.CmdRunDir
	}

	cmd.Stdin = os.Stdin	// connects the main process' stdin to ansible's stdin

	cmdStdout, err = cmd.StdoutPipe()
	defer cmdStdout.Close()
	if err != nil {
		return errors.New("(DefaultExecute::Execute)", "Error creating stdout pipe", err)
	}

	cmdStderr, err = cmd.StderrPipe()
	defer cmdStderr.Close()
	if err != nil {
		return errors.New("(DefaultExecute::Execute)", "Error creating stderr pipe", err)
	}

	timeInit := time.Now()
	err = cmd.Start()
	if err != nil {
		return errors.New("(DefaultExecute::Execute)", "Error starting command", err)
	}

	// Waig for stdout and stderr
	wg.Add(2)

	// stdout management
	go func() {
		defer close(execErrChan)

		trans := []results.TransformerFunc{}

		for _, t := range e.Transformers {
			trans = append(trans, t)
		}

		// when using the default results func DefaultStdoutCallbackResults,
		// reads from ansible's stdout and writes to main process' stdout
		err := resultsFunc(ctx, cmdStdout, e.Write, trans...)
		wg.Done()
		execErrChan <- err
	}()

	// stderr management
	go func() {
		// show stderr messages using default stdout callback results
		results.DefaultStdoutCallbackResults(ctx, cmdStderr, e.WriterError, []results.TransformerFunc{}...)
		wg.Done()
	}()

	wg.Wait()

	if err := <-execErrChan; err != nil {
		return errors.New("(DefaultExecute::Execute)", "Error managing results output", err)
	}

	err = cmd.Wait()
	if err != nil {

		if ctx.Err() != nil {
			fmt.Fprintf(e.Write, "%s\n", fmt.Sprintf("\nWhoops! %s\n", ctx.Err()))
		} else {
			errorMessage := string(err.(*exec.ExitError).Stderr)
			errorMessage = fmt.Sprintf("Command executed: %s\n%s\n%s", cmd.String(), errorMessage, err.Error())

			exitError, exists := err.(*exec.ExitError)
			if exists {
				ws := exitError.Sys().(syscall.WaitStatus)
				switch ws.ExitStatus() {
				case AnsiblePlaybookErrorCodeGeneralError:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageGeneralError, errorMessage)
				case AnsiblePlaybookErrorCodeOneOrMoreHostFailed:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageOneOrMoreHostFailed, errorMessage)
				case AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable, errorMessage)
				case AnsiblePlaybookErrorCodeParserError:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageParserError, errorMessage)
				case AnsiblePlaybookErrorCodeBadOrIncompleteOptions:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageBadOrIncompleteOptions, errorMessage)
				case AnsiblePlaybookErrorCodeUserInterruptedExecution:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageUserInterruptedExecution, errorMessage)
				case AnsiblePlaybookErrorCodeUnexpectedError:
					errorMessage = fmt.Sprintf("%s\n\n%s", AnsiblePlaybookErrorMessageUnexpectedError, errorMessage)
				}
			}
			return errors.New("(DefaultExecute::Execute)", fmt.Sprintf("Error during command execution: %s", errorMessage))
		}
	}

	elapsedTime := time.Since(timeInit)

	if e.ShowDuration {
		fmt.Fprintf(e.Write, "Duration: %s\n", elapsedTime.String())
	}

	return nil
}
