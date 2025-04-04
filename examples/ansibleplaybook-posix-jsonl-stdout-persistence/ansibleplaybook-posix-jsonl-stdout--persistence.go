package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/execute/configuration"
	jsonresults "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

//
// Content for the EnrichedJSONLEvent
//

// EnrichedJSONLEvent is a struct that represents the enriched JSONL event
type EnrichedJSONLEvent struct {
	Hash string `json:"_hash"`
	Data string `json:"data"`
}

// NewEnrichedJSONLEvent creates a new EnrichedJSONLEvent
func NewEnrichedJSONLEvent(data string) EnrichedJSONLEvent {
	// Generate a SHA256 hash from the data
	hash := sha256.Sum256([]byte(data))

	return EnrichedJSONLEvent{
		Hash: fmt.Sprintf("%x", hash),
		Data: data,
	}
}

//
// Content for a transformer that generates a EnrichedJSONLEvent
//

// EnrichedJSONLEventTransformer is a function that transforms a JSONL event into an enriched JSONL event
func EnrichedJSONLEventTransformer(event string) string {

	enrichedEvent := NewEnrichedJSONLEvent(event)
	jsonData, err := json.Marshal(enrichedEvent)
	if err != nil {
		return event
	}

	return string(jsonData)
}

//
// Content for the JSONLEventWriter
//

// JSONLEventWriter is a custom writer that implements the io.Writer interface
type JSONLEventWriter struct {
	writer      io.Writer
	persistency Persistency
}

// NewJSONLEventWriter creates a new JSONLEventWriter
func NewJSONLEventWriter(p Persistency, w io.Writer) *JSONLEventWriter {
	return &JSONLEventWriter{
		persistency: p,
		writer:      w,
	}
}

// Write implements the io.Writer interface for JSONLEventWriter
func (e *JSONLEventWriter) Write(data []byte) (n int, err error) {

	if e.writer == nil {
		e.writer = os.Stdout
	}

	if e.persistency == nil {
		return len(data), fmt.Errorf("persistency is not initialized")
	}
	e.persistency.Add(data)

	return len(data), nil
}

//
// Content for the Persistency
//

// Persistency is a map that stores EnrichedJSONLEvent objects
type Persistency map[string][]byte

// NewPersistency creates a new database
func NewPersistency() Persistency {
	return make(Persistency)
}

// Add adds an event to the database
func (p Persistency) Add(event []byte) {
	hash := fmt.Sprintf("%x", sha256.Sum256(event))
	p[hash] = event
}

// Iterate iterates over the database and applies the function f to each event
func (p Persistency) Iterate() func(func(string, []byte) bool) {
	return func(yield func(string, []byte) bool) {
		for hash, event := range p {
			if !yield(hash, event) {
				return
			}
		}
	}
}

//
// main function
//

func main() {
	var err error

	persistency := NewPersistency()

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Connection: "local",
		Inventory:  "127.0.0.1,",
	}

	playbooksList := []string{"site1.yml", "site2.yml", "site3.yml"}

	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(playbooksList...),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)

	baseExec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		execute.WithWrite(NewJSONLEventWriter(persistency, os.Stdout)),
	)
	baseExec.Quiet()
	baseExec.WithOutput(
		jsonresults.NewJSONLEventStdoutCallbackResults(
			jsonresults.WithJSONLEventTransformers(
				EnrichedJSONLEventTransformer,
			),
		),
	)

	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(baseExec,
		configuration.WithAnsibleStdoutCallback(stdoutcallback.AnsiblePosixJsonlStdoutCallback),
	)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	err = exec.Execute(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, event := range persistency {

		enrichedEvent := new(EnrichedJSONLEvent)
		eventOriginal := new(jsonresults.AnsiblePlaybookJSONLEventResults)

		err = json.Unmarshal(event, &enrichedEvent)
		if err != nil {
			fmt.Printf("Error unmarshaling event: %v\n", err)
			continue
		}

		err = json.Unmarshal([]byte(enrichedEvent.Data), &eventOriginal)
		if err != nil {
			fmt.Printf("Error unmarshaling original event: %v\n", err)
			continue
		}

		fmt.Printf("(%s)\n%s\n", enrichedEvent.Hash, eventOriginal.String())
	}
}
