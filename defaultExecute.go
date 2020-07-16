package ansibler

import (
	"fmt"
	"bufio"
	"bytes"
	"errors"
	"os"
	"os/exec"
	"time"
)

// DefaultExecute is a simple definition of an executor
type Executor struct {
	TimeElapsed string
	Stdout string
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *Executor) Execute(command string, args []string) error {

	var stdBuf string
	stderr := &bytes.Buffer{}

	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
    cmd.Env = append(cmd.Env, "ANSIBLE_STDOUT_CALLBACK=json")
    cmd.Env = append(cmd.Env, "ANSIBLE_HOST_KEY_CHECKING=False")
	cmd.Stderr = stderr

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return errors.New("(DefaultExecute::Execute) -> " + err.Error())
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			stdBuf = stdBuf +"\n"+ scanner.Text()
		}
	}()

	timeInit := time.Now()
	err = cmd.Start()
	if err != nil {
		return errors.New("(DefaultExecute::Execute) -> " + err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		return errors.New("(DefaultExecute::Execute) -> " + stderr.String())
	}

	e.TimeElapsed = time.Since(timeInit).String()
	e.Stdout = stdBuf
	fmt.Println(stdBuf)

	return nil
}
