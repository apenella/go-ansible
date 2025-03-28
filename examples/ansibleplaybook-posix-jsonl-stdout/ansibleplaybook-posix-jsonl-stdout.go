package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

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

	exec := stdoutcallback.NewAnsiblePosixJsonlStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			execute.WithWrite(NewJSONLEventWriter(os.Stdout)),
		),
	)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	err = exec.Execute(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

// JSONLEventWriter is a custom writer that implements the io.Writer interface
type JSONLEventWriter struct {
	writer io.Writer
}

// NewJSONLEventWriter creates a new JSONLEventWriter
func NewJSONLEventWriter(w io.Writer) *JSONLEventWriter {
	return &JSONLEventWriter{
		writer: w,
	}
}

// Write implements the io.Writer interface for JSONLEventWriter
func (e *JSONLEventWriter) Write(p []byte) (n int, err error) {

	var event jsonresults.AnsiblePlaybookJSONLEventResults

	if e.writer == nil {
		e.writer = os.Stdout
	}

	err = json.Unmarshal(p, &event)
	if err != nil {
		return len(p), fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	fmt.Fprintf(e.writer, "%s\n", event)

	return len(p), nil
}
