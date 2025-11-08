package json

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"iter"

	"github.com/apenella/go-ansible/v2/pkg/execute/result"
	"github.com/apenella/go-ansible/v2/pkg/execute/result/transformer"
	errors "github.com/apenella/go-common-utils/error"
	"golang.org/x/sync/errgroup"
)

// JSONLEventStdoutCallbackResults handles the ansible.posix.jsonl callback plugin output
type JSONLEventStdoutCallbackResults struct {
	trans []transformer.TransformerFunc
}

// NewJSONLEventStdoutCallbackResults creates a new JSONLEventStdoutCallbackResults instance
func NewJSONLEventStdoutCallbackResults(options ...result.OptionsFunc) *JSONLEventStdoutCallbackResults {
	results := &JSONLEventStdoutCallbackResults{}
	results.Options(options...)
	return results
}

// WithJSONLEventTransformers sets a transformers list to JSONLEventStdoutCallbackResults
func WithJSONLEventTransformers(trans ...transformer.TransformerFunc) result.OptionsFunc {
	return func(r result.ResultsOutputer) {
		r.(*JSONLEventStdoutCallbackResults).trans = append(r.(*JSONLEventStdoutCallbackResults).trans, trans...)
	}
}

// Options executes the options functions received as a parameters to set the JSONLEventStdoutCallbackResults attributes
func (r *JSONLEventStdoutCallbackResults) Options(options ...result.OptionsFunc) {
	for _, opt := range options {
		opt(r)
	}
}

// Print handles the ansible.posix.jsonl callback plugin output
func (r *JSONLEventStdoutCallbackResults) Print(ctx context.Context, reader io.Reader, writer io.Writer, options ...result.OptionsFunc) error {

	goroutine, ctx := errgroup.WithContext(ctx)
	dataChan := make(chan []byte)
	errContext := "(result::json::JSONLEventStdoutCallbackResults::Print)"

	if reader == nil {
		return errors.New(errContext, "JSONLEventStdoutCallbackResults requires a reader to print the output of the execution")
	}

	if writer == nil {
		return errors.New(errContext, "JSONLEventStdoutCallbackResults requires a writer to print the output of the execution")
	}

	r.Options(options...)

	goroutine.Go(func() error {
		defer close(dataChan)

		errs := []error{}
		for data, err := range readResultsStream(reader) {
			if err != nil {
				errs = append(errs, err)
				continue
			}

			// transformerFunc expects and returns a string so we need to convert the byte array to a string and back
			if len(r.trans) > 0 {
				dataString := string(data)
				for _, t := range r.trans {
					dataString = t(dataString)
				}
				data = []byte(dataString)
			}

			// while there is data it is written to the print channel
			dataChan <- data
		}

		if len(errs) > 0 {
			return errors.New(errContext, "error processing the execution output", errs...)
		}

		return nil
	})

	goroutine.Go(func() error {
		for {
			select {
			case data, ok := <-dataChan:
				if !ok {
					return nil
				}

				_, err := writer.Write(data)
				if err != nil {
					return errors.New(errContext, "error writing to writer", err)
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	// wait for both goroutines to complete or fail. The expected behaviour is that when there are no more lines to read, the first goroutine finishes, and the errgroup cancels the second goroutine, which handles the writing operations.
	err := goroutine.Wait()
	if err != nil {
		return errors.New(errContext, "error handling the results stream", err)
	}

	return nil
}

func readResultsStream(reader io.Reader) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			event := scanner.Bytes()

			if !json.Valid(event) {
				if !yield(nil, fmt.Errorf("invalid JSON event")) {
					return
				}
				continue
			}

			if !yield(append([]byte(nil), event...), nil) { // copy buffer safely
				return
			}
		}

		// Handle any errors that occurred during scanning
		if err := scanner.Err(); err != nil {
			yield(nil, fmt.Errorf("error reading input: %w", err))
		}
	}
}
