package exec

import "io"

// Cmder is an interface to run a command
type Cmder interface {
	CombinedOutput() ([]byte, error)
	Environ() []string
	Output() ([]byte, error)
	Run() error
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	String() string
	Wait() error
}
