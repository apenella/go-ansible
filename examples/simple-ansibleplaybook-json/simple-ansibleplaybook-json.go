package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	ansibler "github.com/apenella/go-ansible"
	"github.com/apenella/go-ansible/stdoutcallback"
	"github.com/apenella/go-ansible/stdoutcallback/results"
)

// JSONExecute is a simple definition of an executor
type JSONExecute struct {
	Write       io.Writer
	ResultsFunc stdoutcallback.StdoutCallbackResultsFunc
}

// Execute takes a command and args and runs it, streaming output to stdout
func (e *JSONExecute) Execute(command string, args []string, prefix string) error {

	var buff bytes.Buffer

	execDoneChan := make(chan int8)
	defer close(execDoneChan)
	execErrChan := make(chan error)
	defer close(execErrChan)

	if e.Write == nil {
		e.Write = os.Stdout
	}

	cmd := exec.Command(command, args...)
	cmd.Stderr = e.Write

	cmdReader, err := cmd.StdoutPipe()
	defer cmdReader.Close()

	if err != nil {
		return errors.New("(JSONExecute::Execute) -> " + err.Error())
	}

	go func() {

		if e.ResultsFunc == nil {
			e.ResultsFunc = results.DefaultStdoutCallbackResults
		}
		err := e.ResultsFunc(prefix, cmdReader, io.Writer(&buff))
		if err != nil {
			execErrChan <- err
			return
		}

		execDoneChan <- int8(0)
	}()

	timeInit := time.Now()
	err = cmd.Start()
	if err != nil {
		return errors.New("(JSONExecute::Execute) -> " + err.Error())
	}

	select {
	case <-execDoneChan:

		ansibleJSONResult, err := results.JSONParse(buff.Bytes())
		if err != nil {
			return errors.New("(JSONExecute::Execute) Error parsing JSON output. " + err.Error())
		}

		err = ansibleJSONResult.CheckStats()
		if err != nil {
			fmt.Fprintln(e.Write, err.Error())
		} else {
			fmt.Fprintln(e.Write, ansibleJSONResult.String())
		}

		elapsedTime := time.Since(timeInit)

		fmt.Fprintf(e.Write, "Duration: %s\n", elapsedTime.String())

		err = cmd.Wait()
		if err != nil {
			return errors.New("(JSONExecute::Execute) " + err.Error())
		}

	case err := <-execErrChan:
		return errors.New("(JSONExecute::Execute) " + err.Error())
	}

	return nil
}

func main() {

	ansiblePlaybookConnectionOptions := &ansibler.AnsiblePlaybookConnectionOptions{
		Connection: "local",
		User:       "aleix",
	}

	ansiblePlaybookOptions := &ansibler.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	playbook := &ansibler.AnsiblePlaybookCmd{
		Playbook:          "site.yml",
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		ExecPrefix:        "Go-ansible example",
		StdoutCallback:    "json",
		Exec: &JSONExecute{
			ResultsFunc: results.JSONStdoutCallbackResults,
		},
	}

	err := playbook.Run()
	if err != nil {
		panic(err)
	}

}
