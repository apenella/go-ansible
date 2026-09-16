package execute

import (
	"io"

	"github.com/apenella/go-ansible/v2/pkg/execute/result"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
)

// ExecuteOptions is a function to set executor options
type ExecuteOptions func(*DefaultExecute)

// WithCmd set the execuctable parameter
func WithCmd(cmd Commander) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.Cmd = cmd
	}
}

// WithExecutable set the execuctable parameter
func WithExecutable(executable Executabler) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.Exec = executable
	}
}

// WithWrite set the writer to be used by DefaultExecutor
func WithWrite(w io.Writer) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.Write = w
	}
}

// WithWriteError set the error writer to be used by DefaultExecutor
func WithWriteError(w io.Writer) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.WriterError = w
	}
}

// WithStdin sets the reader to be connected to the command standard input. It defaults to os.Stdin.
// An *os.File is passed to the command as is. Any other reader is copied to the command through a
// pipe, so a reader that never returns EOF (such as an interactive terminal) keeps Execute blocked
// after the command exits.
func WithStdin(r io.Reader) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.Stdin = r
	}
}

// WithCmdRunDir set the command run directory to be used by DefaultExecutor
func WithCmdRunDir(cmdRunDir string) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.CmdRunDir = cmdRunDir
	}
}

// WithTransformers set trasformes functions
func WithTransformers(trans ...transformer.TransformerFunc) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.Transformers = trans
	}
}

// WithEnvVars adds the provided env var to the command
func WithEnvVars(vars map[string]string) ExecuteOptions {
	return func(e *DefaultExecute) {
		for key, value := range vars {
			e.EnvVars[key] = value
		}
	}
}

// WithOutput sets the output mechanism to DefaultExecutor
func WithOutput(output result.ResultsOutputer) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.WithOutput(output)
	}
}

// WithErrorEnrich sets the error context mechanism to DefaultExecutor
func WithErrorEnrich(enricher ErrorEnricher) ExecuteOptions {
	return func(e *DefaultExecute) {
		e.ErrorEnrich = enricher
	}
}
