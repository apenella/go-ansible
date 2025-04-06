/*

	This example shows how to configure the go-ansible library executor to use a custom writer that writes events to a persistence layer, as well as how to use a transformer to update the original events returned by the ansible.posix.jsonl callback plugin.

	Although the go-ansible library provides an executor that handles ansible.posix.jsonl stdout callback results, it does not allow the use of transformers to update the original events directly. In this case, you need to use the DefaultExecute executor, configure the transformer there, and then create an AnsibleWithConfigurationSettingsExecute executor with the ansible.posix.jsonl stdout callback plugin properly set.

*/

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
// Code for the EnrichedJSONLEvent struct. This struct is used to extend the original event data by adding a hash for enrichment purposes.
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
// Code for a transformer that generates a EnrichedJSONLEvent
//

// EnrichedJSONLEventTransformer is a function that transforms a JSONL event into an enriched JSONL event. The function receives the original event as a string and returns the enriched event as a string.
func EnrichedJSONLEventTransformer(event string) string {

	enrichedEvent := NewEnrichedJSONLEvent(event)
	jsonData, err := json.Marshal(enrichedEvent)
	if err != nil {
		return event
	}

	return string(jsonData)
}

//
// Code for the JSONLEventWriter. This component is responsible for writing the events to a persistence layer.
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
// Code for the Persistency layer. This component is responsible for storing the events in a map.
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

	// Create a new DefaultExecute executor using the playbook command. Set the writer to a custom JSONLEventWriter that persists events to a storage layer.
	baseExec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCmd),
		execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
		execute.WithWrite(NewJSONLEventWriter(persistency, os.Stdout)),
	)
	baseExec.Quiet()
	// Configure the DefaultExecute output to use JSONLEventStdoutCallbackResults and assign the EnrichedJSONLEventTransformer to enrich the event data.
	baseExec.WithOutput(
		jsonresults.NewJSONLEventStdoutCallbackResults(
			jsonresults.WithJSONLEventTransformers(
				EnrichedJSONLEventTransformer,
			),
		),
	)

	// Finally, create a new AnsibleWithConfigurationSettingsExecute executor using the previously defined base executor. This step manually sets the ansible.posix.jsonl stdout callback method.
	exec := configuration.NewAnsibleWithConfigurationSettingsExecute(baseExec,
		configuration.WithAnsibleStdoutCallback(stdoutcallback.AnsiblePosixJsonlStdoutCallback),
	)

	/*

		   If you don't need to apply a transformer to the original event, you can use the AnsiblePosixJsonlStdoutCallbackExecute executor directly, as is show below in that comment. Keep in mind that this will change the behavior of the example: when reading events from persistence, they will be the original events returned by the ansible.posix.jsonl callback plugin, rather than EnrichedJSONLEvent objects.

				  exec := stdoutcallback.NewAnsiblePosixJsonlStdoutCallbackExecute(
						execute.NewDefaultExecute(
							execute.WithCmd(playbookCmd),
							execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
							execute.WithWrite(NewJSONLEventWriter(persistency, os.Stdout)),
						),
				  )

	*/

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
			fmt.Printf("Error unmarshaling EnrichedJSONLEvent: %v. Received event data: %s\n", err, string(event))
			continue
		}

		if enrichedEvent.Data == "" {
			fmt.Printf("EnrichedJSONLEvent data is empty. Received event data: %s\n", string(event))
			continue
		}

		err = json.Unmarshal([]byte(enrichedEvent.Data), &eventOriginal)
		if err != nil {
			fmt.Printf("Error unmarshaling original event: %v. Received event data: %s\n", err, string(event))
			continue
		}

		fmt.Printf("(%s)\n%s\n", enrichedEvent.Hash, eventOriginal.String())
	}
}
