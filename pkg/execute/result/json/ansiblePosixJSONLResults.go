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

// AnsiblePosixJSONLResults handles the ansible.posix.jsonl callback plugin output
type AnsiblePosixJSONLResults struct{}

// NewAnsiblePosixJSONLResults creates a new AnsiblePosixJSONLResults instance
func NewAnsiblePosixJSONLResults() *AnsiblePosixJSONLResults {
	return &AnsiblePosixJSONLResults{}
}

// Print handles the ansible.posix.jsonl callback plugin output
func (r *AnsiblePosixJSONLResults) Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...result.OptionsFunc) error {
	printChan := make(chan []byte)
	errChan := make(chan error)
	done := make(chan struct{})

	errContext := "(result::json::AnsiblePosixJSONLResults::Print)"

	if reader == nil {
		return errors.New(errContext, "AnsiblePosixJSONLResults requires a reader to print the output of the execution")
	}

	if writer == nil {
		return errors.New(errContext, "AnsiblePosixJSONLResults requires a writer to print the output of the execution")
	}

	go func() {
		errs := []error{}

		for data, err := range parseAnsiblePlaybookJSONLEventResultsStream(reader) {
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
				return errors.New(errContext, "Error processing the execution output", err)
			}
		case <-done:
			return nil
		case <-ctx.Done():
			return nil
		}
	}
}

func parseAnsiblePlaybookJSONLEventResultsStream(reader io.Reader) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		buff := bufio.NewReader(reader)
		decoder := json.NewDecoder(buff)

		for {
			var event any
			var err error

			err = decoder.Decode(&event)
			if err == io.EOF {
				return
			}

			if err != nil {
				if !yield(nil, fmt.Errorf("error decoding JSON: %w", err)) {
					return
				}
				continue
			}

			eventByteArray, err := json.Marshal(event)
			if err != nil {
				if !yield(nil, fmt.Errorf("error converting event to string: %w", err)) {
					return
				}
				continue
			}

			if !yield(eventByteArray, nil) {
				return
			}
		}
	}
}
