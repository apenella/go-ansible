package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	jsonresults "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func main() {
	var err error

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbooksList := []string{"site1.yml", "site2.yml", "site3.yml"}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbooksList...),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	pipeReader, pipeWriter := io.Pipe()

	exec := stdoutcallback.NewAnsiblePosixJsonlStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			execute.WithWrite(pipeWriter),
		),
	)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	result := make(chan error, 1)
	go func() {
		result <- exec.Execute(ctx)
		pipeWriter.Close()
	}()

	for data := range parseAnsiblePlaybookJSONLEventResultsStream(pipeReader) {
		fmt.Println(data)
	}

	err = <-result
	if err != nil {
		fmt.Println(err)
	}
}

func parseAnsiblePlaybookJSONLEventResultsStream(reader io.Reader) iter.Seq[jsonresults.AnsiblePlaybookJSONLEventResults] {
	return func(yield func(jsonresults.AnsiblePlaybookJSONLEventResults) bool) {
		var event jsonresults.AnsiblePlaybookJSONLEventResults
		bufReader := bufio.NewReader(reader)

		for {
			line, _, err := bufReader.ReadLine()
			if err == io.EOF {
				return
			}
			if err != nil {
				fmt.Println("Error reading line:", err)
				continue
			}
			err = json.Unmarshal(line, &event)
			if err != nil {
				fmt.Println("Error unmarshaling line:", err)
				continue
			}

			if !yield(event) {
				return
			}
		}
	}
}
