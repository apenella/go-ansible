package execute

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"

	"github.com/apenella/go-ansible/stdoutcallback"
	"github.com/apenella/go-ansible/stdoutcallback/results"
	errors "github.com/apenella/go-common-utils/error"
)

// DefaultExecuteOptions
type DefaultExecuteOptions struct {
	// Prefix is a text that is set at the beginning of each execution line
	Prefix string
	// CmdRunDir specifies the working directory of the command.
	CmdRunDir string
}

// DefaultExecute is a simple definition of an executor
type DefaultExecute struct {
	// Writer is where is written the command stdout
	Write io.Writer
	// WriterError is where is written the command stderr
	WriterError io.Writer
	// ResultsFunc is the function that manages execution output
	ResultsFunc stdoutcallback.StdoutCallbackResultsFunc
	// ShowDuration enables to show the execution duration time after the command finishes
	ShowDuration bool
}

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

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DefaultExecute) Execute(command string, args []string, prefix string) error {

	var err error
	var cmdStderr, cmdStdout io.ReadCloser
	var wg sync.WaitGroup

	execErrChan := make(chan error)

	if e.Write == nil {
		e.Write = os.Stdout
	}

	cmd := exec.Command(command, args...)

	if e.CmdRunDir != "" {
		cmd.Dir = e.CmdRunDir
	}

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

		if e.ResultsFunc == nil {
			e.ResultsFunc = results.DefaultStdoutCallbackResults
		}

		err := e.ResultsFunc(prefix, cmdStdout, e.Write)
		wg.Done()
		execErrChan <- err
	}()

	// stderr management
	go func() {
		if e.WriterError == nil {
			e.WriterError = os.Stderr
		}

		// show stderr messages using default stdout callback results
		results.DefaultStdoutCallbackResults(prefix, cmdStderr, e.WriterError)
		wg.Done()
	}()

	wg.Wait()

	if err := <-execErrChan; err != nil {
		return errors.New("(DefaultExecute::Execute)", "Error managing results output", err)
	}

	err = cmd.Wait()
	if err != nil {
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

	elapsedTime := time.Since(timeInit)

	if e.ShowDuration {
		fmt.Fprintf(e.Write, "Duration: %s\n", elapsedTime.String())
	}

	return nil
}
