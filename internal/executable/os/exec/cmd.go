package exec

import (
	"io"
	"os/exec"
)

// Cmd struct is a wrapper of exec.Cmd
type Cmd struct {
	cmd *exec.Cmd
}

// NewCmd return a wa
func NewCmd(cmd *exec.Cmd) *Cmd {
	return &Cmd{cmd}
}

// CombinedOutput is a wrapper of exec.Cmd CombinedOutput method
func (c *Cmd) CombinedOutput() ([]byte, error) {
	return c.cmd.CombinedOutput()
}

// Environ is a wrapper of exec.Cmd Environ method
func (c *Cmd) Environ() []string {
	return c.cmd.Environ()
}

// Output is a wrapper of exec.Cmd Output method
func (c *Cmd) Output() ([]byte, error) {
	return c.cmd.Output()
}

// Run is a wrapper of exec.Cmd Run method
func (c *Cmd) Run() error {
	return c.cmd.Run()
}

// Start is a wrapper of exec.Cmd Start method
func (c *Cmd) Start() error {
	return c.cmd.Start()
}

// StderrPipe is a wrapper of exec.Cmd StderrPipe method
func (c *Cmd) StderrPipe() (io.ReadCloser, error) {
	return c.cmd.StderrPipe()
}

// StdinPipe is a wrapper of exec.Cmd StdinPipe method
func (c *Cmd) StdinPipe() (io.WriteCloser, error) {
	return c.cmd.StdinPipe()
}

// StdoutPipe is a wrapper of exec.Cmd StdoutPipe method
func (c *Cmd) StdoutPipe() (io.ReadCloser, error) {
	return c.cmd.StdoutPipe()
}

// String is a wrapper of exec.Cmd String method
func (c *Cmd) String() string {
	return c.cmd.String()
}

// Wait is a wrapper of exec.Cmd Wait method
func (c *Cmd) Wait() error {
	return c.cmd.Wait()
}
