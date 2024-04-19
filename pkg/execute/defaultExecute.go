package execute

import (
	"context"
	"fmt"
	"io"
	"os"
	osexec "os/exec"
	"strings"
	"sync"

	"github.com/apenella/go-ansible/v2/internal/executable/os/exec"
	"github.com/apenella/go-ansible/v2/pkg/execute/result"
	defaultresults "github.com/apenella/go-ansible/v2/pkg/execute/result/default"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
)

const (
	// AnsiblePlaybookErrorCodeGeneralError is the error code for a general error
	AnsiblePlaybookErrorCodeGeneralError = 1
	// AnsiblePlaybookErrorCodeOneOrMoreHostFailed is the error code for a one or more host failed
	AnsiblePlaybookErrorCodeOneOrMoreHostFailed = 2
	// AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable is the error code for a one or more host unreachable
	AnsiblePlaybookErrorCodeOneOrMoreHostUnreachable = 3
	// AnsiblePlaybookErrorCodeParserError is the error code for a parser error
	AnsiblePlaybookErrorCodeParserError = 4
	// AnsiblePlaybookErrorCodeBadOrIncompleteOptions is the error code for a bad or incomplete options
	AnsiblePlaybookErrorCodeBadOrIncompleteOptions = 5
	// AnsiblePlaybookErrorCodeUserInterruptedExecution is the error code for a user interrupted execution
	AnsiblePlaybookErrorCodeUserInterruptedExecution = 99
	// AnsiblePlaybookErrorCodeUnexpectedError is the error code for a unexpected error
	AnsiblePlaybookErrorCodeUnexpectedError = 250

	// AnsiblePlaybookErrorMessageGeneralError is the error message for a general error
	AnsiblePlaybookErrorMessageGeneralError = "ansible-playbook error: general error"
	// AnsiblePlaybookErrorMessageOneOrMoreHostFailed is the error message for a one or more host failed
	AnsiblePlaybookErrorMessageOneOrMoreHostFailed = "ansible-playbook error: one or more host failed"
	// AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable is the error message for a one or more host unreachable
	AnsiblePlaybookErrorMessageOneOrMoreHostUnreachable = "ansible-playbook error: one or more host unreachable"
	// AnsiblePlaybookErrorMessageParserError is the error message for a parser error
	AnsiblePlaybookErrorMessageParserError = "ansible-playbook error: parser error"
	// AnsiblePlaybookErrorMessageBadOrIncompleteOptions is the error message for a bad or incomplete options
	AnsiblePlaybookErrorMessageBadOrIncompleteOptions = "ansible-playbook error: bad or incomplete options"
	// AnsiblePlaybookErrorMessageUserInterruptedExecution is the error message for a user interrupted execution
	AnsiblePlaybookErrorMessageUserInterruptedExecution = "ansible-playbook error: user interrupted execution"
	// AnsiblePlaybookErrorMessageUnexpectedError is the error message for a unexpected error
	AnsiblePlaybookErrorMessageUnexpectedError = "ansible-playbook error: unexpected error"
)

// EnvVars represents a custom environment for an ansible playbook execution.
type EnvVars map[string]string

// Environ returns a copy of strings representing the custom environment, in the form "key=value".
func (e EnvVars) Environ() []string {
	result := make([]string, 0, len(e))
	for k, v := range e {
		result = append(result, fmt.Sprintf("%s=%s", k, v))
	}
	return result
}

// DefaultExecute is a simple definition of an executor
type DefaultExecute struct {
	// Cmd is the command generator
	Cmd Commander
	// CmdRunDir specifies the working directory of the command.
	CmdRunDir string
	// EnvVars specifies env vars of the command.
	EnvVars EnvVars
	// ErrContext is the error context
	ErrorEnrich ErrorEnricher
	// Exec is the executor
	Exec Executabler
	// Output manages the output of the command
	Output result.ResultsOutputer
	// quiet is a flag to set the executor in quiet mode
	quiet bool
	// Transformers is the list of transformers func for the output
	Transformers []transformer.TransformerFunc
	// Writer is where is written the command stdout
	Write io.Writer
	// WriterError is where is written the command stderr
	WriterError io.Writer
}

// NewDefaultExecute return a new DefaultExecute instance with all options
func NewDefaultExecute(options ...ExecuteOptions) *DefaultExecute {
	execute := &DefaultExecute{
		EnvVars: make(map[string]string),
	}

	for _, opt := range options {
		opt(execute)
	}

	return execute
}

// WithOutput sets the output mechanism to DefaultExecutor
func (e *DefaultExecute) WithOutput(output result.ResultsOutputer) {
	e.Output = output
}

// AddEnvVar add the provided environment variable. It overwrites the variable when it already exists.
func (e *DefaultExecute) AddEnvVar(key, value string) {
	if e.EnvVars == nil {
		e.EnvVars = make(EnvVars)
	}

	e.EnvVars[key] = value
}

// AddEnvVarSafe add the provided environment variable. It returns an error when the variable already exists
func (e *DefaultExecute) AddEnvVarSafe(key, value string) error {

	errContext := "execute::DefaultExecute:AddEnvVarSafe"

	if e.EnvVars == nil {
		e.EnvVars = make(EnvVars)
	}

	_, exists := e.EnvVars[key]
	if exists {
		return errors.New(errContext, fmt.Sprintf("Environment variable '%s' already exists", key))
	}

	e.EnvVars[key] = value
	return nil
}

// Quiet sets the executor in quiet mode
func (e *DefaultExecute) Quiet() {
	e.quiet = true
}

// quietCommand returns the command without the verbose flags -v, -vv, -vvv, -vvvv and --verbose
func (e *DefaultExecute) quietCommand() ([]string, error) {

	errContext := "(execute::DefaultExecute:quietCommand)"

	command, err := e.Cmd.Command()
	if err != nil {
		return nil, errors.New(errContext, "Error creating command", err)
	}

	quietCommand := make([]string, 0)
	for _, cmd := range command {
		if cmd == "-v" || cmd == "-vv" || cmd == "-vvv" || cmd == "-vvvv" || cmd == "--verbose" {
			continue
		}
		quietCommand = append(quietCommand, cmd)
	}

	return quietCommand, nil
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DefaultExecute) Execute(ctx context.Context) (err error) {

	var errCmd error
	var cmdStderr, cmdStdout io.ReadCloser
	var wg sync.WaitGroup

	errContext := "(execute::DefaultExecute::Execute)"

	defer e.checkCompatibility()

	execErrChan := make(chan error)

	// default stdout and stderr for the main process
	if e.Write == nil {
		e.Write = os.Stdout
	}

	if e.WriterError == nil {
		e.WriterError = os.Stderr
	}

	if e.Exec == nil {
		e.Exec = exec.NewExec()
	}

	if e.Cmd == nil {
		return errors.New(errContext, "Command is not defined")
	}

	command, err := e.Cmd.Command()
	if err != nil {
		return errors.New(errContext, "Error creating command", err)
	}

	if e.quiet {
		command, err = e.quietCommand()
		if err != nil {
			return errors.New(errContext, "Error creating quiet command", err)
		}
	}

	cmd := e.Exec.CommandContext(ctx, command[0], command[1:]...)

	// Assert if cmd's type is the Golang's exec.Cmd as set the desired values for that case
	_, isOsExecCmd := cmd.(*osexec.Cmd)
	if isOsExecCmd {
		if len(e.CmdRunDir) > 0 {
			cmd.(*osexec.Cmd).Dir = e.CmdRunDir
		}

		if len(e.EnvVars) > 0 {
			cmd.(*osexec.Cmd).Env = append(os.Environ(), e.EnvVars.Environ()...)
		}

		// connects the main process' stdin to ansible's stdin
		cmd.(*osexec.Cmd).Stdin = os.Stdin
	}

	trans := make([]transformer.TransformerFunc, 0)
	trans = append(trans, e.Transformers...)

	cmdStdout, err = cmd.StdoutPipe()
	defer func() {
		_ = cmdStdout.Close()
	}()
	if err != nil {
		return errors.New(errContext, "Error creating stdout pipe", err)
	}

	cmdStderr, err = cmd.StderrPipe()
	defer func() {
		_ = cmdStderr.Close()
	}()
	if err != nil {
		return errors.New(errContext, "Error creating stderr pipe", err)
	}

	if e.Output == nil {

		e.Output = defaultresults.NewDefaultResults(
			defaultresults.WithTransformers(trans...),
		)
	}

	err = cmd.Start()
	if err != nil {
		return errors.New(errContext, "Error starting command", err)
	}

	// Waig for stdout and stderr
	wg.Add(2)

	// stdout management
	go func() {
		defer close(execErrChan)

		// when using the default results func DefaultStdoutCallbackResults,
		// reads from ansible's stdout and writes to main process' stdout
		e.Output.Print(ctx, cmdStdout, e.Write)

		wg.Done()
		execErrChan <- err
	}()

	// stderr management
	go func() {
		// show stderr messages using default stdout callback results
		e.Output.Print(ctx, cmdStderr, e.WriterError)
		wg.Done()
	}()

	wg.Wait()

	if err := <-execErrChan; err != nil {
		return errors.New(errContext, "Error managing results output", err)
	}

	err = cmd.Wait()
	if err != nil {

		if ctx.Err() != nil {
			fmt.Fprintf(e.Write, "%s\n", fmt.Sprintf("\nWhoops! %s\n", ctx.Err()))
		} else {

			if e.ErrorEnrich != nil {
				errCmd = e.ErrorEnrich.Enrich(err)
			} else {
				errCmd = err
			}

			errorMessage := fmt.Sprintf(" Command executed: %s\n", e.Cmd.String())
			if len(e.EnvVars) > 0 {
				errorMessage = fmt.Sprintf("%s\n Environment variables:\n%s\n", errorMessage, strings.Join(e.EnvVars.Environ(), "\n"))
			}

			stderrErrorMessage := string(err.(*osexec.ExitError).Stderr)
			if len(stderrErrorMessage) > 0 {
				errorMessage = fmt.Sprintf("%s\n'%s'\n", errorMessage, stderrErrorMessage)
			}

			return errors.New(errContext, fmt.Sprintf("Error during command execution.\n%s", errorMessage), errCmd)
		}
	}

	return nil
}

func (e *DefaultExecute) checkCompatibility() {}
