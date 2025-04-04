package json

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"

	"github.com/apenella/go-ansible/v2/pkg/execute/result"
	errors "github.com/apenella/go-common-utils/error"
)

// JSONLEventStdoutCallbackResults handles the ansible.posix.jsonl callback plugin output
type JSONLEventStdoutCallbackResults struct{}

// NewJSONLEventStdoutCallbackResults creates a new JSONLEventStdoutCallbackResults instance
func NewJSONLEventStdoutCallbackResults() *JSONLEventStdoutCallbackResults {
	return &JSONLEventStdoutCallbackResults{}
}

// Print handles the ansible.posix.jsonl callback plugin output
func (r *JSONLEventStdoutCallbackResults) Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...result.OptionsFunc) error {
	printChan := make(chan []byte)
	errChan := make(chan error)
	done := make(chan struct{})

	errContext := "(result::json::JSONLEventStdoutCallbackResults::Print)"

	if reader == nil {
		return errors.New(errContext, "JSONLEventStdoutCallbackResults requires a reader to print the output of the execution")
	}

	if writer == nil {
		return errors.New(errContext, "JSONLEventStdoutCallbackResults requires a writer to print the output of the execution")
	}

	go func() {
		defer close(printChan)
		defer close(errChan)
		defer close(done)

		errs := []error{}

		for data, err := range readResultsStream(reader) {
			if err != nil {
				errs = append(errs, err)
				continue
			}
			printChan <- data
		}

		if len(errs) > 0 {
			errChan <- errors.New(errContext, "Error processing the execution output", errs...)
		}

		done <- struct{}{}
	}()

	for {
		select {
		case data := <-printChan:
			_, err := writer.Write(data)
			if err != nil {
				return errors.New(errContext, "Error writing to writer", err)
			}
		case err := <-errChan:
			if err != nil {
				return errors.New(errContext, "Error reading the results stream", err)
			}
		case <-done:
			return nil
		case <-ctx.Done():
			return nil
		}
	}
}

func readResultsStream(reader io.Reader) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			line := scanner.Text()

			// Validate if the line is a properly formed JSON object
			var event any
			err := json.Unmarshal([]byte(line), &event)
			if err != nil {
				if !yield(nil, fmt.Errorf("error decoding JSON: %w", err)) {
					return
				}
				continue
			}

			// Convert the JSON object back to a byte array
			eventByteArray, err := json.Marshal(event)
			if err != nil {
				if !yield(nil, fmt.Errorf("error converting event to string: %w", err)) {
					return
				}
				continue
			}

			// Yield the valid JSON byte array
			if !yield(eventByteArray, nil) {
				return
			}
		}

		// Handle any errors that occurred during scanning
		if err := scanner.Err(); err != nil {
			yield(nil, fmt.Errorf("error reading input: %w", err))
		}
	}
}
