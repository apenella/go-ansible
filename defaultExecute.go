package ansibler

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// DefaultExecute is a simple definition of an executor
type DefaultExecute struct {
	Write io.Writer
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *DefaultExecute) Execute(command string, args []string, prefix string) error {

	stderr := &bytes.Buffer{}

	if e.Write == nil {
		e.Write = os.Stdout
	}

	cmd := exec.Command(command, args...)
	cmd.Stderr = stderr

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return errors.New("(DefaultExecute::Execute) -> " + err.Error())
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Fprintf(e.Write, "%s =>  %s\n", prefix, scanner.Text())
		}
	}()

	timeInit := time.Now()
	err = cmd.Start()
	if err != nil {
		return errors.New("(DefaultExecute::Execute) -> " + err.Error())
	}

	err = cmd.Wait()
	elapsedTime := time.Since(timeInit)
	if err != nil {
		return errors.New("(DefaultExecute::Execute) -> " + stderr.String())
	}

	fmt.Fprintf(e.Write, "Duration: %s\n", elapsedTime.String())

	return nil
}
